package tree_sitter_ron_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-ron"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ron.Language())
	if language == nil {
		t.Errorf("Error loading Ron grammar")
	}
}
