package termdocumentmatrix

import (
	"fmt"
	"log"
	"strings"

	StringTokenizer "github.com/wolfeidau/stringtokenizer"
	"github.com/xuri/excelize/v2"
)

func TermDocumentMatrix() (map[string]map[string]int, int) {
	hashmap := make(map[string]map[string]int)
    key := 0
	f, err := excelize.OpenFile("./TermDocumentMatrix/comment.xlsx")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

	rows, err := f.GetRows("result")
    if err != nil {
        log.Fatal(err)
    }
    for _, row := range rows {
        tokenizer := StringTokenizer.NewStringTokenizer(strings.NewReader(row[2]), " ", false)
		temp := make(map[string]int)
        for tokenizer.HasMoreTokens() {
            token := tokenizer.NextToken()
            if lemmatization[token] != "" {
				token = lemmatization[token]
			}
            if !persian[token] {
				temp[token] = 1
            }
    	}
		hashmap[row[3]] = temp
        key++
	}

    return hashmap, key
}