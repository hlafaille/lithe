package lsp

import (
	"errors"
	"fmt"
)

/** RequestHandler represents a registered handler for a method */
type RequestHandler struct {
	Method string
	Do     func() (error, ResponseMessage)
}

var handlers []RequestHandler = make([]RequestHandler, 0)

/** InstallHandler installs a request handler */
func InstallHandler(handler RequestHandler) {
	handlers = append(handlers, handler)
}

/** HandleMessage handles a given message */
func HandleMessage(message RequestMessage) (error, ResponseMessage) {
	for _, handler := range handlers {
		if handler.Method == message.Method {
			return handler.Do()
		}
	}
	return errors.New(fmt.Sprintf("no registered handlers with method '%s'", message.Method)), ResponseMessage{}
}
