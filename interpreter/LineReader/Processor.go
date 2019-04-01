package LineReader

import (
	"bufio"
	"github.com/kostyaBro/SICP/interpreter/interfaces"
	"log"
	"os"
	"sort"
	"strings"
)

type processor struct {
	lisp interfaces.ILisp
}

func NewProcessor(lisp interfaces.ILisp) interfaces.IProcessor {
	return &processor{lisp: lisp}
}

func (p *processor) Process() {
	// write string from keyboard
	// ( -> switch ? err
	// wait ) ? err
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

		var spaceIndexes []int
		hookMap := make(map[int]string)

		for index, char := range []byte(out) {
			if strings.Compare(string(char), "(") == 0 || strings.Compare(string(char), ")") == 0 {
				hookMap[index] = string(char)
			}
			if strings.Compare(string(char), " ") == 0 {
				spaceIndexes = append(spaceIndexes, index)
			}
		}

		var hookMapKeys []int
		for index := range hookMap {
			hookMapKeys = append(hookMapKeys, index)
		}
		sort.Ints(hookMapKeys) // todo create fast sort by hand
		sort.Ints(spaceIndexes)
		if hookMapKeys[0] != 0 {
			println("unexpected start")
			continue
		}
		// LOG
		for _, index := range hookMapKeys {
			log.Printf("hookMap[%d] = %s", index, hookMap[index])
		}
		for index, val := range spaceIndexes {
			log.Printf("spaceIndexes[%d] = %d", index, val)
		}

		subTask := out[hookMapKeys[hookMapKeys[0]]+1 : spaceIndexes[0]]

		switch subTask {
		case "+":
		case "-":
		case "*":
		case "/":

		}

		// // find ()
		// var subTask, _ = func () (subTask string, err error) {
		// 	var openIndex int
		// 	for _, index := range hookMapKeys {
		// 		if strings.Compare(hookMap[index], "(") == 0 {
		// 			openIndex = index
		// 		} else if strings.Compare(hookMap[index], ")") == 0 {
		// 			return out[openIndex:index], nil
		// 		}
		// 	}
		// 	return out, errors.New("have not ()")
		// }()
		//
		log.Printf("sub Task = '%s'", subTask)

		//

		// commands := strings.Split(out, " ")
		// if len(commands) == 1 {
		// 	println(commands[0])
		// }
		// for _, command := range commands  {
		//
		// }
		// (define (sum a b) (a + b))
	}
}
