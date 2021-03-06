package RuleParser

import "github.com/notpaulmartin/mdParser/Parse"

func RecursiveApply(tree []Parse.ParseTree, applyable Parse.Applyable) []Parse.ParseTree {
	// Base case: empty tree
	if len(tree) == 0 {
		return tree
	}

	// Base case: unit tree and reached unexpanded node
	if len(tree) == 1 && tree[0].TagName == Parse.RawTag {
		success, parse := applyable.Apply(tree[0].Content) // Try to expand node
		if !success {
			// If can't expand, leave as is
			return tree
		}
		return parse
	}

	// Base case: unit tree and reached leaf
	if len(tree) == 1 && len(tree[0].Children) == 0 {
		return tree
	}

	// Unit tree and at expandable node  ->  Go deeper
	if len(tree) == 1 {
		tree[0].Children = RecursiveApply(tree[0].Children, applyable)
		return tree
	}

	// Not a unit tree  ->  Recurse over all root nodes
	// TODO: do in parallel using a worker pool
	var newTree []Parse.ParseTree
	for _, node := range tree {
		newTree = append(newTree, RecursiveApply(Parse.UnitTree(node), applyable)...)
	}
	return newTree
}
