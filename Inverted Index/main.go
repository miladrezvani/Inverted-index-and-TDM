package main

import (
	index "github.com/miladrezvani/Inverted_Index"
)


func main() {
	hashmap := index.Inverted_Index()
	index.Search_engine(hashmap)
}