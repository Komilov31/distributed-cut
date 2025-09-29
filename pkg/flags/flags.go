package flags

import (
	"log"

	"github.com/pborman/getopt/v2"
)

type Flags struct {
	FlagF    []int
	FlagD    string
	FlagS    bool
	FileName string
}

func Parse() *Flags {
	flagF := getopt.String('f', "", `select only these fields;  also print any line
                            that contains no delimiter character, unless
                            the -s option is specified`)
	flagD := getopt.String('d', "\t", "use DELIM instead of TAB for field delimiter")
	flagS := getopt.Bool('s', "do not print lines not containing delimiters")
	getopt.Parse()

	flags := Flags{
		FlagD: *flagD,
		FlagS: *flagS,
	}

	if len(*flagF) == 0 {
		log.Fatal("cut: you must specify fields")
	}

	fields, err := parseFlagF(*flagF)
	if err != nil {
		log.Fatal("cut: invalid field value")
	}

	flags.FlagF = fields

	args := getopt.Args()
	if len(args) != 0 {
		flags.FileName = args[0]
	}

	return &flags
}
