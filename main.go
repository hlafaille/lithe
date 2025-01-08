package main

import (
	"fmt"

	"kerosenelabs.com/lithe/lsp"
)

func main() {
	// setup
	lsp.InstallHandler(lsp.RequestHandler{
		Method: "initialize",
		Do: func() (error, lsp.ResponseMessage) {
			fmt.Println("Got initialize")
			return nil, lsp.ResponseMessage{
				Id: 1,,
			}
		},
	})

	// we're ready
	lsp.Serve()
}
