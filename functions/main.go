// accept interface but return concrete types
package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	prefix string
}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println(cl.prefix + message)
}

func (cl *ConsoleLogger) SetPrefix(prefix string) {
	cl.prefix = prefix
}

func CallLogMethod(logger Logger) {
	logger.Log(" is a genius")
}

func NewLogger() *ConsoleLogger {
	return &ConsoleLogger{
		prefix: "Subhayan",
	}
}

func main() {
	logger := NewLogger()
	CallLogMethod(logger)
	logger.SetPrefix("Shaayan")
	CallLogMethod(logger)
}
