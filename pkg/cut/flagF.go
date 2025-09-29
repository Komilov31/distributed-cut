package cut

import (
	"strings"
)

const (
	doesNotContainDelimeter = 1
)

func (c *Cut) ProcessFlagF(nextLine string) string {
	fields := strings.Split(nextLine, c.flags.FlagD)
	if len(fields) <= doesNotContainDelimeter && c.flags.FlagS {
		return ""
	}

	if len(fields) == doesNotContainDelimeter {
		return fields[0]
	}

	builder := strings.Builder{}
	for _, f := range c.flags.FlagF {
		if f >= 0 && f < len(fields) {
			builder.WriteString(fields[f])
			if f != c.flags.FlagF[len(c.flags.FlagF)-1] {
				builder.WriteString(c.flags.FlagD)
			}
		}
	}

	if len(builder.String()) == 0 {
		return nextLine
	}

	return builder.String()
}
