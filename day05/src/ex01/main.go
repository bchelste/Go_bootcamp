package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type TreeNode struct {
    HasToy	bool
    Left	*TreeNode
    Right	*TreeNode
}

func areToysBalanced(root *TreeNode) bool {
	if (root == nil) {
		return false
	}
	left := root.Left
	right := root.Right

	leftToys := countToys(left)
	rightToys := countToys(right)

	return (leftToys == rightToys)
}

func countToys(root *TreeNode) int {
	var result int = 0

	if (root == nil) {
		return result
	}

	var storage []*TreeNode
	storage = append(storage, root)
	var front *TreeNode = nil

	for len(storage) != 0 {
		size := len(storage)
		for ; size > 0; size-- {
			front = storage[0]
			storage = storage[1:]
			if (front.HasToy == true) {
				result += 1
			}
			if (front.Left != nil) {
				storage = append(storage, front.Left)
			}
			if (front.Right != nil) {
				storage = append(storage, front.Right)
			}
				
		}
	}
	return result
}

func printSpace(nbr int) {
	for i := 0; i < nbr; i++ {
		fmt.Print(" ")
	}
}

func addToLine(line *string, item *TreeNode) {
	if (item == nil) {
		*line += "x"
	} else {
		if item.HasToy {
			*line += "1"
		} else {
			*line += "0"
		}
	}
}

func treeHeight(tree *TreeNode) (int) {
	var result int = 0

	if (tree == nil) {
		return result
	}

	var storage []*TreeNode
	storage = append(storage, tree)
	var front *TreeNode = nil

	for len(storage) != 0 {
		size := len(storage)
		for ; size > 0; size-- {
			front = storage[0]
			storage = storage[1:]
			if (front.Left != nil) {
				storage = append(storage, front.Left)
			}
			if (front.Right != nil) {
				storage = append(storage, front.Right)
			}
				
		}
		result += 1
	}
	return result
}

func getLineFromTree(tree *TreeNode) string {
	if (tree == nil) {
		return ""
	}

	var line string = ""
	var next string = ""
	var prev string = ""

	var storage []*TreeNode
	storage = append(storage, tree)
	var front *TreeNode = nil
	addToLine(&line, tree)

	for len(storage) != 0 {
		size := len(storage)
		next = ""
		i := 0
		for ; size > 0; size-- {

			for ; i < len(prev); i++ {
				if prev[i] == 'x' {
					next += "xx"
				} else {
					break
				}
			}

			front = storage[0]
			storage = storage[1:]
			if (front.Left != nil) {
				storage = append(storage, front.Left)
				addToLine(&next, front.Left)
			} else {
				addToLine(&next, nil)
			}
			if (front.Right != nil) {
				storage = append(storage, front.Right)
				addToLine(&next, front.Right)
			} else {
				addToLine(&next, nil)
			}
			i += 1
				
		}

		line += next
		prev = next
	}

	return line
}

func outputTree(tree *TreeNode) {
	height := treeHeight(tree)
	line := getLineFromTree(tree)
	var space int = int(math.Pow(2, float64(height)))

	for i := 0; i < height; i += 1 {
		space = space / 2
		nbr := int(math.Pow(2, float64(i)))
		for j := 0; j < nbr; j += 1 {
			printSpace(space - 1)
			fmt.Printf("%c", line[nbr - 1 + j])
			printSpace(space)
		}
		fmt.Print("\n")
	}
}

func creatRandomTree(height int) (*TreeNode) {

	if (height == 0) {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	levels := height - 1

	var storage []*TreeNode
	var p *TreeNode = new(TreeNode)
	storage = append(storage, p)
	var front *TreeNode = nil
	var rnd int
	
	for len(storage) != 0 {
		size := len(storage)
		for ; size > 0; size-- {
			front = storage[0]
			storage = storage[1:]
			rnd = rand.Intn(2)
			if (rnd == 0) {
				front.HasToy = false
			} else {
				front.HasToy = true
			}
			if (levels > 0) {
				rnd = rand.Intn(2)
				if (rnd == 0) && (levels < 2) && (front != nil) {
					front.Left = nil
				} else {
					front.Left = new(TreeNode)
				}
				rnd = rand.Intn(2)
				if (rnd == 0) && (levels < 2) && (front != nil) {
					front.Right = nil
				} else {
					front.Right = new(TreeNode)
				}
			}
			if (front.Left != nil) {
				storage = append(storage, front.Left)
			}
			if (front.Right != nil) {
				storage = append(storage, front.Right)
			}
			
		}
		levels--
	}

	return p
}

func incrCntr(cntr *int, src string, delt int) bool {
	if (*cntr + delt < len(src)) {
		*cntr += delt
		return true
	} else {
		return false
	}
}

func addToResult(result *[]bool, tmp string, cntr int) {
	if ((cntr & 1) != 0) {
		for i := len(tmp) - 1; i >= 0; i-- {
			if (tmp[i] == '0') {
				*result = append(*result, false)
			} else if (tmp[i] == '1') {
				*result = append(*result, true)
			}
		}	
	} else {
		for i := 0; i < len(tmp); i++ {
			if (tmp[i] == '0') {
				*result = append(*result, false)
			} else if (tmp[i] == '1') {
				*result = append(*result, true)
			}
		}
	}
}

func unrollGarland(root *TreeNode) []bool {
	if (root == nil) {
		return []bool{}
	}

	var result = []bool{}

	var next string = ""
	

	var storage []*TreeNode
	storage = append(storage, root)
	var front *TreeNode = nil
	
	var cntr int = 1

	for len(storage) != 0 {
		size := len(storage)
		next = ""
		for ; size > 0; size-- {
			front = storage[0]
			addToLine(&next, front)
			storage = storage[1:]
			if (front.Left != nil) {
				storage = append(storage, front.Left)
			} 
			if (front.Right != nil) {
				storage = append(storage, front.Right)
			}
		}
		addToResult(&result, next, cntr)
		cntr++;
	}

	return result
} 

func main() {

	var height int64 = 0

	flag.Int64Var(&height, "R", 0, "count lines")
	flag.Parse()
	
	if (height != 0) {
		fmt.Printf("The height of random tree is: %d\n", height)
		randomTree := creatRandomTree(int(height))
		fmt.Println("-----")
		outputTree(randomTree)
		fmt.Println("-----")

		fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(randomTree))

		return
	}

	fmt.Println("EX00 examples for: \"areToysBalanced\"")
	fmt.Println("-----------------------------")

	exampleTree1 := &TreeNode{false, 
		&TreeNode{false,
			&TreeNode{false, nil, nil},
			&TreeNode{true, nil, nil}},
		&TreeNode{true, nil, nil}}
	outputTree(exampleTree1)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleTree1))
	fmt.Println("-----------------------------")

	exampleTree2 := &TreeNode{true, 
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	outputTree(exampleTree2)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleTree2))
	fmt.Println("-----------------------------")

	exampleTree3 := &TreeNode{true, 
		&TreeNode{true, nil,nil},
		&TreeNode{false, nil, nil}}
	outputTree(exampleTree3)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleTree3))
	fmt.Println("-----------------------------")

	exampleTree4 := &TreeNode{false, 
		&TreeNode{true,
			nil,
			&TreeNode{true, nil, nil}},
		&TreeNode{false,
			nil,
			&TreeNode{true, nil, nil}}}
	outputTree(exampleTree4)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleTree4))
	fmt.Println("-----------------------------")

	exampleSingleRoot := &TreeNode{false,nil,nil}
	outputTree(exampleSingleRoot)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleSingleRoot))
	fmt.Println("-----------------------------")

	var exampleNilRoot *TreeNode = nil
	outputTree(exampleNilRoot)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleNilRoot))
	fmt.Println("-----------------------------")
	
	exampleTreeBig := &TreeNode{true,
		&TreeNode{false,
			&TreeNode{true,
				&TreeNode{false,
					&TreeNode{true,nil,nil},
					&TreeNode{true,nil,nil}},
				&TreeNode{true,
					nil,
					&TreeNode{false,nil,nil}}},
			&TreeNode{true,
				&TreeNode{true,
					nil,
					nil},
				&TreeNode{true,
					nil,
					nil}}},
		&TreeNode{true,
			&TreeNode{false,
				&TreeNode{false,
					&TreeNode{true,nil,nil},
					&TreeNode{true,nil,nil}},
				nil},
			&TreeNode{true,
				&TreeNode{true, 
					nil, 
					nil},
				&TreeNode{true,
					&TreeNode{true,nil,nil},
					&TreeNode{false,nil,nil}}}}}

	outputTree(exampleTreeBig)
	fmt.Printf("unrollGarland result: ")
	fmt.Println(unrollGarland(exampleTreeBig))
	fmt.Println("-----------------------------")
}