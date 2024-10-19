package main

import (
	index "github.com/miladrezvani/TermDocumentMatrix"
)

func main() {
	hashmap, key := index.TermDocumentMatrix()
	index.Search_engine(hashmap, key)
}