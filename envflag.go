package envflag

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

func StringVar(p *string, name, value, usage string) {
	if s := os.Getenv(strings.ToUpper(name)); s != "" {
		// set it to environment value if available
		*p = s
	} else {
		// fall back to default
		*p = value
	}

	// pass it on to flag
	flag.StringVar(p, name, *p, usage)
}

func IntVar(p *int, name string, value int, usage string) {
	if s := os.Getenv(strings.ToUpper(name)); s != "" {
		// set it to environment value if available
		if i, err := strconv.ParseInt(s, 10, 63); err != nil {
			*p = int(i)
		}
	} else {
		// fall back to default
		*p = value
	}

	// pass it on to flag
	flag.IntVar(p, name, *p, usage)
}

func BoolVar(p *bool, name string, value bool, usage string) {
	if s := os.Getenv(strings.ToUpper(name)); s != "" {
		// set it to environment value if available
		if b, err := strconv.ParseBool(s); err != nil {
			*p = b
		}
	} else {
		// fall back to default
		*p = value
	}

	// pass it on to flag
	flag.BoolVar(p, name, *p, usage)
}
