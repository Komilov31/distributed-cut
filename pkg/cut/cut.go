package cut

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Komilov31/distributed-cut/pkg/flags"
	"github.com/wb-go/wbf/zlog"
)

const (
	severNumber = 3
	maxInput    = 5
)

var (
	requests = 0
)

type Cut struct {
	flags *flags.Flags
	file  *os.File
}

func New(flags *flags.Flags) *Cut {
	file := initFile(flags)

	return &Cut{
		flags: flags,
		file:  file,
	}
}

func (c *Cut) ProcessProgram() {
	reader := bufio.NewReader(c.file)
	inputChan := make(chan map[int]string)
	index := 0
	result := make([]string, 0)
	succesfulRequests := 0

	resultChan, gotResult := sendRequestToSevers(inputChan)

	go func() {
		for got := range gotResult {
			if got {
				succesfulRequests++
			}
		}
	}()

	go func() {
		for results := range resultChan {
			for key, value := range results {
				result[key] = value
			}
		}
	}()

	input := make(map[int]string)
	for {
		nextLine, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal("could not read line: ", err)
		}
		result = append(result, "")

		nextLine = strings.TrimSuffix(nextLine, "\n")

		input[index] = nextLine
		index++

		if len(input) == maxInput || err == io.EOF {
			inputChan <- input
			input = make(map[int]string)
		}

		if err == io.EOF {
			break
		}
	}

	if succesfulRequests >= requests/2+1 {
		printResult(result)
		return
	}

	zlog.Logger.Error().Msg("unfotunately severs are not available")
}

func sendRequestToSevers(inputs <-chan map[int]string) (<-chan map[int]string, chan bool) {
	output := make(chan map[int]string)
	gotResult := make(chan bool)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	servers := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	wg := new(sync.WaitGroup)
	for input := range inputs {
		serverIndex := 0

		wg.Add(1)
		go func(address string) {
			defer wg.Done()

			body, err := json.Marshal(input)
			if err != nil {
				zlog.Logger.Error().Msg("could not send request to server: " + err.Error())
				gotResult <- false
				return
			}

			requests++
			var result map[int]string
			req, err := http.NewRequest(http.MethodPost, address, bytes.NewBuffer(body))
			if err != nil {
				zlog.Logger.Error().Msg("could not send request to server: " + err.Error())
				gotResult <- false
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				zlog.Logger.Error().Msg("could not send request to server: " + err.Error())
				gotResult <- false
				return
			}

			body, err = io.ReadAll(resp.Body)
			if err != nil {
				zlog.Logger.Error().Msg("could not read response body from server: " + err.Error())
				gotResult <- false
				return
			}

			if err := json.Unmarshal(body, &result); err != nil {
				zlog.Logger.Error().Msg("could not unmarshal response body from server: " + err.Error())
				gotResult <- false
				return
			}

			serverIndex++
			if serverIndex == severNumber-1 {
				serverIndex = 0
			}

			output <- result
			gotResult <- true
		}(servers[serverIndex])
	}

	go func() {
		wg.Wait()
		close(output)
		close(gotResult)
	}()

	return output, gotResult
}

func printResult(results []string) {
	for _, result := range results {
		if result == "" {
			fmt.Print(result)
		} else {
			fmt.Println(result)
		}
	}
}
