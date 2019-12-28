package strs

import (
	"bytes"
	"math"
	"strconv"
)

var result [][]byte

type node struct {
	value int
	down  *node
	right *node
}

func FindMaxPath(A [][]int) string {
	row, col := 0, 0
	root := newNode(A[0][0])
	createTree(A, row, col, root)
	findPaths(root)
	return strconv.FormatInt(int64(findMax()), 10)
}

func findPaths(current *node) {
	push(current.value)

	// move down
	if current.down != nil {
		findPaths(current.down)
	}

	// move right
	if current.right != nil {
		findPaths(current.right)
	}

	// we reached to leaf
	// save current path
	// back track to previous node
	if current.right == nil && current.down == nil {
		savePath()
	}
	pop()
}

func createTree(a [][]int, row, col int, parent *node) {
	right, down := -1, -1

	if col < len(a[0])-1 {
		right = a[row][col+1]
	}

	if row < len(a)-1 {
		down = a[row+1][col]
	}

	if right >= down && right != -1 {
		parent.right = newNode(right)
		createTree(a, row, col+1, parent.right)
	}
	if down >= right && down != -1 {
		parent.down = newNode(down)
		createTree(a, row+1, col, parent.down)
	}
}

func newNode(v int) *node {
	return &node{
		value: v,
		down:  nil,
		right: nil,
	}
}

// super minimal stack
var stack []int

func push(i int) {
	stack = append(stack, i)
}

func pop() int {
	var top int = -1
	if len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return top
}

// save current path from the root to leaf
func savePath() {
	var path bytes.Buffer
	for _, v := range stack {
		path.WriteByte(byte(v))
	}

	result = append(result, path.Bytes())
}
func findMax() int {
	var max int
	for _, v := range result {
		var current int
		for i, vv := range v {
			current += int(vv) * int(math.Pow(float64(10), float64(len(v)-i-1)))
		}
		if current > max {
			max = current
		}
	}
	return max
}
