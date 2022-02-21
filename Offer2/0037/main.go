package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	queue := []*TreeNode{root}
	res := []string{}
	var node *TreeNode
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	return "[" + strings.Join(res, ",") + "]"
}

func deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	res := strings.Split(data[1:len(data)-1], ",")
	val, _ := strconv.Atoi(res[0])
	root := &TreeNode{val, nil, nil}
	queue := []*TreeNode{root}
	i := 1
	var node *TreeNode
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		if res[i] != "null" {
			val, _ = strconv.Atoi(res[i])
			node.Left = &TreeNode{val, nil, nil}
			queue = append(queue, node.Left)
		}
		i += 1
		if res[i] != "null" {
			val, _ = strconv.Atoi(res[i])
			node.Right = &TreeNode{val, nil, nil}
			queue = append(queue, node.Right)
		}
		i += 1
	}
	return root
}
