package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	root := &Node{0, nil, nil}
	tempHead := head
	tempRoot := root

	map1 := make(map[*Node]int)
	map2 := make(map[int]*Node)
	i := 0
	for tempHead != nil {
		tempRoot.Next = &Node{tempHead.Val, nil, nil}

		map1[tempHead] = i
		map2[i] = tempRoot.Next

		tempRoot = tempRoot.Next
		tempHead = tempHead.Next
		i++
	}

	tempHead = head
	tempRoot = root
	var tempRandom *Node = nil

	for tempHead != nil {
		tempRandom = tempHead.Random
		if tempRandom != nil {
			if index, ok := map1[tempRandom]; ok {
				tempRoot.Next.Random = map2[index]
			}
		}

		tempRoot = tempRoot.Next
		tempHead = tempHead.Next
	}

	return root.Next
}
