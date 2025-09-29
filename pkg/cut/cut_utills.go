package cut

import (
	"log"
	"os"

	"github.com/Komilov31/distributed-cut/pkg/flags"
)

func initFile(flags *flags.Flags) *os.File {
	file := os.Stdin

	if flags.FileName != "" {
		var err error
		file, err = os.Open(flags.FileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return file
}
