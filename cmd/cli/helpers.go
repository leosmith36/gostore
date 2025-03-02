package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgs() (a *args, err error) {
	a = getDefaultArgs()

	list := os.Args[1:]
	for i := 0; i < len(list); i++ {
		name := list[i]

		switch name {
		case "-h":
			fallthrough
		case "--host":
			if a.host, err = safeGetArg(list, "host", i); err != nil {
				return nil, err
			}
			i++
		case "-p":
			fallthrough
		case "--port":
			var val string
			if val, err = safeGetArg(list, "port", i); err != nil {
				return nil, err
			}
			if a.port, err = strconv.Atoi(val); err != nil {
				return nil, fmt.Errorf("ERROR invalid port: %s", val)
			}
			i++
		default:
			return nil, fmt.Errorf("ERROR invalid argument: %s", name)
		}
	}

	return
}

func safeGetArg(list []string, label string, i int) (val string, err error) {
	if i >= len(list)-1 {
		return "", fmt.Errorf("ERROR missing argument for %s", label)
	}

	return list[i+1], nil
}
