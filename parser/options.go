package parser

import (
	"github.com/gomarkdown/markdown/ast"
)

// ParserOptions is a collection of supplementary parameters tweaking the behavior of various parts of the parser.
type ParserOptions struct {
	ParserHook    BlockFunc
	ReadIncludeFn ReadIncludeFunc
}

// BlockFunc allows to registration of a parser function. If successful it
// returns an ast.Node, a buffer that should be parsed as a block and the the number of bytes consumed.
type BlockFunc func(data []byte) (ast.Node, []byte, int)

// ReadIncludeFunc should read the file under path and returns the read bytes,
// from will be set to the name of the current file being parsed. Initially
// this will be empty. address is the optional address specifier of which lines
// of the file to return. If this function is not set no data will be read.
type ReadIncludeFunc func(from, path string, address []byte) []byte
