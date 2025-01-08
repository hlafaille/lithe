package lsp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"strings"
)

func Serve() {
	reader := bufio.NewReader(os.Stdin)

	for {
		length := 0
		for {
			header, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Fprintf(os.Stderr, "Header read error: %v\n", err)
				return
			}

			header = strings.TrimSpace(header)
			if header == "" {
				break
			}

			if strings.HasPrefix(header, "Content-Length:") {
				val := strings.TrimPrefix(header, "Content-Length:")
				length, _ = strconv.Atoi(strings.TrimSpace(val))
			}
		}

		body := make([]byte, length)
		if _, err := io.ReadFull(reader, body); err != nil {
			fmt.Fprintf(os.Stderr, "Body read error: %v\n", err)
			return
		}

		var msg = ResponseMessage{
			Id:     "1",
			Result: "Hello",
			Error:  nil,
		}
		var msgJson, _ = msg.MarshalJSON()
		var msgJsonString = string(msgJson)
		var out bytes.Buffer
		out.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(msgJsonString)))
		out.WriteString(msgJsonString)
		_, _ = os.Stdout.Write(out.Bytes())
	}
}
