package termdocumentmatrix

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

func applyOp(hashmap map[string]map[string]int,data map[string][]int, keys int ,a, b string, op string) []int {
	switch op {
	case "&":
		return intersection(hashmap,data, keys,a,b)
	case "|":
		return union(hashmap,data,keys,a,b)
	case "~|":
		return union_not(hashmap,data,keys,a,b,false)
	case "|~":
		return union_not(hashmap,data,keys,a,b,true)
	case "~&":
		return intersection_not(hashmap,data,keys,a,b,false)
	case "&~":
		return intersection_not(hashmap,data,keys,a,b,true)
	default:
		panic("Invalid operator")
	}
}

func search(tokens string, hashmap map[string]map[string]int, keys int) (map[string][]int, error) {
	var tokensArray []string
	data := make(map[string][]int)

	tokenizer := StringTokenizer.NewStringTokenizer(strings.NewReader(tokens), " ", false /*includeDelimiters */)
	for tokenizer.HasMoreTokens() {
		tokensArray = append(tokensArray, tokenizer.NextToken())
	}
	var values []string
	var ops []string
	var calculate []int

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
					calculate = applyOp(hashmap,data,keys,val1,val2[1:] , op + "~")
				}else if []rune(val1)[0] == '~' {
					calculate = applyOp(hashmap,data,keys,val1[1:],val2 , "~" + op)
				}else {
					calculate = applyOp(hashmap,data,keys,val1,val2 , op)
				}

				data[tokens] = calculate
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
					calculate = applyOp(hashmap,data,keys,val1,val2[1:] , op + "~")
				}else if []rune(val1)[0] == '~' {
					calculate = applyOp(hashmap,data,keys,val1[1:],val2 , "~" + op)
				}else {
					calculate = applyOp(hashmap,data,keys,val1,val2 , op)
				}

				data[tokens] = calculate
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
			calculate = applyOp(hashmap,data,keys,val1,val2[1:] , op + "~")
		}else if []rune(val1)[0] == '~' {
			calculate = applyOp(hashmap,data,keys,val1[1:],val2 , "~" + op)
		}else {
			calculate = applyOp(hashmap,data,keys,val1,val2 , op)
		}

		data[tokens] = calculate
		values = append(values, tokens)
	}

	if len(values) != 1 {
		return nil, errors.New("invalid expression")
	}

	return data, nil
}

func Search_engine(hashmap map[string]map[string]int, keys int) {
	fmt.Print("search : ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
    if err != nil {
		log.Fatal(err)
    }
	token := strings.TrimSuffix(line, "\n")

	// token := "خداحافظ | ( ( خوب & سلام ) & درب )"
	flag := true
	i:= 0
	fmt.Print("result -> ")
	for i <= keys {
		if hashmap[strconv.Itoa(i)][token] == 1 {
			fmt.Print(i)
			fmt.Print(", ")
			flag = false
		}
		i++
	}
	if flag {
		data, _ := search(token,hashmap, keys)
		fmt.Println(data[token])
	}
}