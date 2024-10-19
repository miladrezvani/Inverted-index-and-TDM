package Inverted_Index

import (
	"fmt"
	"log"
	"strings"

	StringTokenizer "github.com/wolfeidau/stringtokenizer"
	"github.com/xuri/excelize/v2"
)

func Inverted_Index() map[string][]string {
	hashmap := make(map[string][]string)
	f, err := excelize.OpenFile("./Inverted_Index/comment.xlsx")
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
		tokenizer := StringTokenizer.NewStringTokenizer(strings.NewReader(row[2]), " ", false /*includeDelimiters */)
		for tokenizer.HasMoreTokens() {
			token := tokenizer.NextToken()
			if lemmatization[token] != "" {
				token = lemmatization[token]
			}
			if !persian[token] {
				hashmap[token] = append(hashmap[token], row[3])
			}
			
		}
	}
	return hashmap
}
