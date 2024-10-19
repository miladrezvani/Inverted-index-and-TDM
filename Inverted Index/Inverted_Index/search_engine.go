package Inverted_Index

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	StringTokenizer "github.com/wolfeidau/stringtokenizer"
)

func precedence(op string) int {
	if op == "&" || op == "|" {
		return 1
	}
	if op == "~|" || op == "|~" || op == "~&" || op == "&~" {
		return 2
	}
	return 0
}

func applyOp(a, b []string, op string) []string {
	switch op {
	case "&":
		return intersection(a,b)
	case "|":
		return union(a,b)
	case "~|":
		return union_not(a,b,false)
	case "|~":
		return union_not(a,b,true)
	case "~&":
		return intersection_not(a,b,false)
	case "&~":
		return intersection_not(a,b,true)
	default:
		panic("Invalid operator")
	}
}


func search(tokens string, hashmap map[string][]string) ([]string, error) {
	var tokensArray []string 
	
	tokenizer := StringTokenizer.NewStringTokenizer(strings.NewReader(tokens), " ", false)
	for tokenizer.HasMoreTokens() {
		tokensArray = append(tokensArray, tokenizer.NextToken())
	}
	var values []string
	var ops []string
	var calculate []string
	
	for i := 0; i < len(tokensArray); i++ {
		if tokensArray[i] == "(" {
			ops = append(ops, tokensArray[i])
		}else if tokensArray[i] == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				if []rune(val2)[0] == '~' {
					calculate = applyOp(hashmap[val1],hashmap[val2[1:]] , op + "~")
				}else if []rune(val1)[0] == '~' {
					calculate = applyOp(hashmap[val1[1:]],hashmap[val2], "~" + op)
				}else {
					calculate = applyOp(hashmap[val1],hashmap[val2] , op)
				}
				
				hashmap[tokens] = calculate[:]
				values = append(values, tokens)
			}
			if len(ops) == 0 || ops[len(ops)-1] != "(" {
				return nil, errors.New("unbalanced parentheses")
			}
			ops = ops[:len(ops)-1]
		}else if tokensArray[i] == "&" || tokensArray[i] == "|" {
			for len(ops) > 0 && precedence(ops[len(ops)-1]) >= precedence(tokensArray[i]) {
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]

				if []rune(val2)[0] == '~' {
					calculate = applyOp(hashmap[val1],hashmap[val2[1:]] , op + "~")
				}else if []rune(val1)[0] == '~' {
					calculate = applyOp(hashmap[val1[1:]],hashmap[val2], "~" + op)
				}else {
					calculate = applyOp(hashmap[val1],hashmap[val2] , op)
				}

				hashmap[tokens] = calculate[:]
				values = append(values, tokens)
			}
			ops = append(ops, tokensArray[i])
		}else {
			values = append(values, tokensArray[i])
		}
	}
	for len(ops) > 0 {
		val2 := values[len(values)-1]
		values = values[:len(values)-1]
		val1 := values[len(values)-1]
		values = values[:len(values)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		if []rune(val2)[0] == '~' {
			calculate = applyOp(hashmap[val1],hashmap[val2[1:]] , op + "~")
		}else if []rune(val1)[0] == '~' {
			calculate = applyOp(hashmap[val1[1:]],hashmap[val2], "~" + op)
		}else {
			calculate = applyOp(hashmap[val1],hashmap[val2] , op)
		}

		hashmap[tokens] = calculate[:]
		values = append(values, tokens)
	}

	if len(values) != 1 {
		return nil, errors.New("invalid expression")
	}

	return values, nil
}

func Search_engine(hashmap map[string][]string) {
	fmt.Print("search : ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
    if err != nil {
		log.Fatal(err)
    }
	token := strings.TrimSuffix(line, "\n")

		// token := "( ( سلام & خوب ) & درب ) | خداحافظ"
	if len(hashmap[token]) == 0 {
		search(token,hashmap)
	}
	fmt.Println("result ->", hashmap[token])
}