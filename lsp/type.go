package lsp

type LSPAny interface{}

type LSPObject map[string]LSPAny
type LSPArray []LSPAny
