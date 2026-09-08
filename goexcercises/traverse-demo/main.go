package main

import "strings"

var debug bool
var name string

var subCommands = map[string]bool{
	"run": true,
}

func hasNoOptDefVal(name string) bool {
	switch name {
	case "debug", "d":
		return true
	case "name", "n":
		return false
	default:
		return false
	}
}

func shortHasNoOptDefVal(name string) bool {
	if len(name) == 0 {
		return false
	}
	return hasNoOptDefVal(string(name[0]))
}

func isFlagArg(arg string) bool {
	return (len(arg) >= 3 && arg[0:2] == "--") || (len(arg) >= 2 && arg[0:1] == "-" && arg[1] != '-')
}

func findNext(name string) bool {
	return subCommands[name]
}

func setFlag(key, val string) {
	switch key {
	case "debug", "d":
		debug = val == "true"
	case "name", "n":
		name = val
	}
}

func parseFlags(words []string) {
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "--") {
			key, val, hasEquals := strings.Cut(w[2:], "=")
			if hasEquals {
				setFlag(key, val)
				continue
			}
			if hasNoOptDefVal(key) {
				setFlag(key, "true")
			} else if i+1 < len(words) {
				i++
				setFlag(key, words[i])
			}
			continue
		}
		if len(w) == 2 && w[0] == '-' {
			key := w[1:]
			if hasNoOptDefVal(string(key)) {
				setFlag(string(key), "true")
			} else if i+1 < len(words) {
				i++
				setFlag(string(key), words[i])
			}
		}
	}
}

func traverse(args []string) string {
	flags := []string{}
	inFlag := false
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--") && !strings.Contains(arg, "="):
			inFlag = !hasNoOptDefVal(arg[2:])
			flags = append(flags, arg)
			continue
		case strings.HasPrefix(arg, "-") && !strings.Contains(arg, "=") && len(arg) == 2 && !shortHasNoOptDefVal(arg[1:]):
			inFlag = true
			flags = append(flags, arg)
			continue
		case inFlag:
			inFlag = false
			flags = append(flags, arg)
			continue
		case isFlagArg(arg):
			flags = append(flags, arg)
			continue
		}
		if !findNext(arg) {
			return ""
		}
		parseFlags(flags)
		return arg
	}

	return ""
}

func main() {
	args := []string{"--debug", "false", "run"} // try different inputs

	cmd := traverse(args)
	println("subcommand:", cmd)
	println("debug:", debug)
	println("name:", name)
}
