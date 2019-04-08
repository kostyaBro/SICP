package LineReader

import (
	"bufio"
	"fmt"
	"github.com/kostyaBro/SICP/interpreter/interfaces"
	"log"
	"os"
	"reflect"
	"strconv"
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
		// tree
		rootNode := new(Node)
		targetNode := rootNode
		var str string
		for _, char := range []byte(out) {
			switch string(char) {
			case "(":
				if len(str) != 0 {
					targetNode.hasString(str)
					str = ""
				}
				targetNode = targetNode.hasOpen()
				continue
			case " ":
				if len(str) != 0 {
					targetNode.hasString(str)
					str = ""
				}
				targetNode.hasSpace()
				continue
			case ")":
				if len(str) != 0 {
					targetNode.hasString(str)
					str = ""
				}
				targetNode = targetNode.hasClose()
				continue
			case "+", "-", "*", "/":
				if len(str) != 0 {
					targetNode.hasString(str)
					str = ""
				}
				targetNode.hasString(string(char))
				continue
			default:
				str += string(char)
				continue
			}
			if targetNode == rootNode {
				log.Printf("breaking on targetNode == rootNode")
				break
			}
		}
		rootNode.Print()
		// calculate
		o := calculate(p, rootNode)
		log.Println("type of output", reflect.ValueOf(o).Kind())
		log.Println("output: ", o)
	}
}

func calculate(p *processor, root *Node) (output interface{}) {
	if len(root.value) == 0 {
		for _, child := range root.children {
			calculate(p, child)
			if len(root.value) != 0 {
				break
			}
		}
	} else {
		switch root.value {
		case "+", "-", "*", "/":
			var args []interface{}
			for i, parentsChildren := range root.parent.children {
				if i == 0 {
					continue
				}
				if len(parentsChildren.value) == 0 {
					args = append(args, calculate(p, parentsChildren))
				}
				if i, err := strconv.Atoi(parentsChildren.value); err != nil {
					log.Printf("unexpected value on func '+' %s err %s", parentsChildren.value, err.Error())
					return -1
				} else {
					args = append(args, i)
				}
			}
			var err error
			switch root.value {
			case "+":
				output, err = p.lisp.Add(args...)
				if err != nil {
					log.Printf("exception %s", err.Error())
					return -2
				}
			case "-":
				output, err = p.lisp.Rem(args)
				if err != nil {
					log.Printf("exception %s", err.Error())
					return -2
				}
			case "*":
				output, err = p.lisp.Mul(args)
				if err != nil {
					log.Printf("exception %s", err.Error())
					return -2
				}
			case "/":
				output, err = p.lisp.Div(args)
				if err != nil {
					log.Printf("exception %s", err.Error())
					return -2
				}
			}
			if reflect.ValueOf(output).Kind() > reflect.Int && reflect.ValueOf(output).Kind() < reflect.Uint64 {
				root.parent.value = strconv.FormatInt(reflect.ValueOf(output).Int(), 10)
			} else {
				root.parent.value = reflect.ValueOf(output).String()
			}
			return root.parent.value
		default:
			log.Printf("unexpected symbol %s", root.value)
			return -3
		}
	}
	return root.value
}

// region Node

type Node struct {
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
			printChild(child, indent+indent)
		}
	}
}

func (node *Node) Calculate() {

}

// endregion
