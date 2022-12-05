package handler

import (
	"fmt"
)

type ErrorState uint

const (
	INTERNAL = iota
	UNKNOWN
	RUNTIME
)

var ErrorType = []string{
	"INTERNAL",
	"UNKNOWN",
	"RUNTIME",
}

type CustomError struct {
	ErrorState
	Message string
}

func (dt ErrorState) name() string {
	return ErrorType[dt]
}

func (dt ErrorState) ordinal() int {
	return int(dt)
}

func (dt ErrorState) values() *[]string {
	return &ErrorType
}

func (err *CustomError) Error() {
	fmt.Println("Error: TYPE:", err.name(), "MESSAGE:", err.Message)
}

func HandleError(errType ErrorState, err ...error) {
	for _, e := range err {
		if e != nil {
			custom := CustomError{ErrorState: errType, Message: fmt.Sprint(e)}
			custom.Error()
		}
	}
}
