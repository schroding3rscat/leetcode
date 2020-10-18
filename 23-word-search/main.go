package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{byte('A'), byte('B'), byte('C'), byte('E')},
		{byte('S'), byte('F'), byte('C'), byte('S')},
		{byte('A'), byte('D'), byte('E'), byte('E')},
	}
	word := "ABCCED"
	ex := exist(board, word)
	fmt.Println(ex)
	fmt.Printf("\n")
	board = [][]byte{
		{byte('a'), byte('b')},
		{byte('c'), byte('d')},
	}
	word = "acdb"
	ex = exist(board, word)
	fmt.Println(ex)
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return len(word) == 0
	}

	wordBytes := []byte(word)
	li := len(board)
	lj := len(board[0])

	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			if board[i][j] == wordBytes[0] {
				visited := make([][]bool, li)
				for k := 0; k < li; k++ {
					visited[k] = make([]bool, lj)
				}
				if dfs(board, i, j, wordBytes, 0, visited) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, wordBytes []byte, depth int, visited [][]bool) bool {
	fmt.Printf("i = %d, j = %d: %s, visited: %v\n", i, j, string(board[i][j]), visited[i][j])
	if visited[i][j] {
		return false
	}
	visited[i][j] = true
	if wordBytes[depth] == board[i][j] {
		if len(wordBytes)-1 == depth {
			return true
		}

		if i-1 >= 0 {
			if dfs(board, i-1, j, wordBytes, depth+1, visited) {
				return true
			}
		}
		if j+1 < len(board[0]) {
			if dfs(board, i, j+1, wordBytes, depth+1, visited) {
				return true
			}
		}
		if i+1 < len(board) {
			if dfs(board, i+1, j, wordBytes, depth+1, visited) {
				return true
			}
		}
		if j-1 >= 0 {
			if dfs(board, i, j-1, wordBytes, depth+1, visited) {
				return true
			}
		}
	}
	visited[i][j] = false
	return false
}
