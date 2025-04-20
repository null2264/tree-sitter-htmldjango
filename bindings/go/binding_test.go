package tree_sitter_htmljinja_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_htmljinja "github.com/null2264/tree-sitter-htmldjango/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_htmljinja.Language())
	if language == nil {
		t.Errorf("Error loading Htmljinja grammar")
	}
}
