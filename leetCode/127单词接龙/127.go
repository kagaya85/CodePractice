package main

import "fmt"

func main() {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordID := map[string]int{}
	graph := [][]int{}

	// hashmap 里添加单词和对应id
	addWord := func(word string) int {
		id, has := wordID[word]
		if !has {
			id = len(wordID)
			wordID[word] = id
			graph = append(graph, []int{})
		}
		return id
	}

	// 建图
	addEdge := func(word string) int {
		id1 := addWord(word)
		charList := []byte(word)
		for i, c := range charList {
			charList[i] = '*'
			id2 := addWord(string(charList))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			charList[i] = c
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}

	beginID := addEdge(beginWord)
	endID, has := wordID[endWord]
	if !has {
		return 0
	}

	distBegin := make([]int, len(wordID))
	for i := range distBegin {
		distBegin[i] = -1
	}
	distBegin[beginID] = 0

	distEnd := make([]int, len(wordID))
	for i := range distEnd {
		distEnd[i] = -1
	}
	distEnd[endID] = 0

	beginQueue := []int{beginID}
	endQeueue := []int{endID}

	for len(beginQueue) > 0 && len(endQeueue) > 0 {
		q := beginQueue
		beginQueue = nil
		for _, from := range q {
			if distEnd[from] > 0 {
				return (distBegin[from]+distEnd[from])/2 + 1
			} else {
				for _, to := range graph[from] {
					if distBegin[to] < 0 {
						distBegin[to] = distBegin[from] + 1
						beginQueue = append(beginQueue, to)
					}
				}
			}
		}
		q = endQeueue
		endQeueue = nil
		for _, from := range q {
			if distBegin[from] > 0 {
				return (distEnd[from]+distBegin[from])/2 + 1
			} else {
				for _, to := range graph[from] {
					if distEnd[to] < 0 {
						distEnd[to] = distEnd[from] + 1
						endQeueue = append(endQeueue, to)
					}
				}
			}
		}
	}

	return 0
}
