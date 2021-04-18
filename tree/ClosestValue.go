package tree

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	currentNode := tree
	closest := tree.Value

	for currentNode != nil {
		if absDiff(closest, target) > absDiff(target, currentNode.Value) {
			closest = currentNode.Value
		}
		if target > currentNode.Value {
			currentNode = currentNode.Right
		} else if target < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			break
		}
	}
	return closest
}

func absDiff(first, second int) int {
	if first > second {
		return first - second
	} else {
		return second - first
	}
}

func min(first, second int) int {
	if first > second {
		return second
	}
	return first
}
