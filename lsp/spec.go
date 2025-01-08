package lsp

import "encoding/json"

const PARSE_ERROR = -32700
const INVALID_REQUEST = -32600
const METHOD_NOT_FOUND = -32601
const INVALID_PARAMS = -32602
const INTERNAL_ERROR = -32603

/** https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#headerPart */
type HeaderPart struct {
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
}

/** https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#contentPart */
type ContentPart struct {
	HeaderPart     HeaderPart
	RequestMessage RequestMessage
}

/** https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage */
type RequestMessage struct {
	Id     string         `json:"id"`
	Method string         `json:"method"`
	Params *[]interface{} `json:"params"`
}

/** https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage */
type ResponseError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    LSPAny `json:"data"`
}

/** https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage */
type ResponseMessage struct {
	Id     string         `json:"id"`
	Result LSPAny         `json:"result"`
	Error  *ResponseError `json:"error"`
}

func (s ResponseMessage) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      s.Id,
		"result":  s.Result,
		"error":   s.Error,
	}
	return json.Marshal(out)
}
