package LineReader

import (
	"bufio"
	"fmt"
	"github.com/kostyaBro/SICP/interpreter/interfaces"
	"log"
	"os"
	"strings"
)

type processor struct {
	lisp interfaces.ILisp
}

func NewProcessor(lisp interfaces.ILisp) interfaces.IProcessor {
	return &processor{lisp: lisp}
}

func (p *processor) Process() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		out, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Fatal error on reed interpreter/LineReader/Processor.go:24 %s", err.Error())
		}
		out = strings.TrimSuffix(out, "\n")
		if strings.Compare(out, "exit") == 0 {
			return
		}
		if strings.Count(out, "(") < strings.Count(out, ")") {
			println("')' more than '('")
			continue
		} else if strings.Count(out, "(") > strings.Count(out, ")") {
			println("'(' more than ')'")
			continue
		}
		rootNode := new(Node)
		targetNode := rootNode
		for _, char := range []byte(out) {
			switch string(char) {
			case "(":
				targetNode = targetNode.hasOpen()
				continue
			case " ":
				targetNode.hasSpace()
				continue
			case ")":
				targetNode = targetNode.hasClose()
				continue
			default:
				targetNode.hasString(string(char))
				continue
			}
			if targetNode == rootNode {
				log.Printf("breaking on targetNode == rootNode")
				break
			}
		}
		rootNode.Print()
	}
}

// region Node

type Node struct{
	parent   *Node
	value    string
	children []*Node
}

func (node *Node) hasOpen() (control *Node) {
	child := new(Node)
	child.parent = node
	node.children = append(node.children, child)
	return child
}

func (node *Node) hasSpace() {
	return
}

func (node *Node) hasClose() (control *Node) {
	return node.parent
}

func (node *Node) hasString(str string) {
	child := new(Node)
	child.parent = node
	node.children = append(node.children, child)
	child.value = strings.Trim(str, " ")
}

func (node *Node) Print() {
	target := node
	fmt.Printf("%+v\n", target)
	printChild(node, "\t")
}


func printChild(node *Node, indent string) {
	if node.children != nil {
		for _, child := range node.children {
			fmt.Printf("%s%+v\n", indent, child)
			printChild(child, indent + indent)
		}
	}
}

// endregion