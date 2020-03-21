package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	indexMap:=make(map[string]int,len(words))
	for _,word:=range words{
		indexMap[word]++
	}

	sortMap(indexMap,k)
	fmt.Println(indexMap)
	return []string{}
}

func sortMap(m map[string]int,k int)  {
	karr:=make([]int,0)
	sarr:=make([]string,0)
	for k,v:=range m{
		karr=append(karr,v)
		sarr=append(sarr,k)
	}
	sort.Ints(karr)
	//karr=karr[len(karr)-k:]
	sort.Strings(sarr)

	fmt.Println(karr)
	fmt.Println(sarr)
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2

	fmt.Println(topKFrequent(words, k))
}
