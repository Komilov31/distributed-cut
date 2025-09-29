package flags

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseFlagF(flagF string) ([]int, error) {
	fields := []int{}

	fFields := strings.Split(flagF, ",")
	for _, field := range fFields {
		f, err := strconv.Atoi(field)
		if err != nil {
			intervals := strings.Split(field, "-")
			if len(intervals) != 2 {
				return nil, fmt.Errorf("wrong interval")
			}

			start, err := strconv.Atoi(intervals[0])
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(intervals[1])
			if err != nil {
				return nil, err
			}

			for i := start - 1; i < end; i++ {
				fields = append(fields, i)
			}
		}
		if f != 0 {
			fields = append(fields, f-1)
		}
	}
	slices.Sort(fields)

	return fields, nil
}
