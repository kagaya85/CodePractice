package main

import (
    "fmt"
)

func main() {
    board := [][]byte{
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'},
    }
    
    word := "ABCCED"
    ans := exist(board, word)
    
    fmt.Println(ans)
}

func exist(board [][]byte, word string) bool {
    width := len(board[0])
    height := len(board)
    
    var visited []bool = make([]bool, width * height)
    
    var check func(x, y, z int) bool
    check = func(x, y, z int) bool {
        visited[y * width + x] = true
        if board[y][x] != word[z] || z >= len(word) {
            visited[y * width + x] = false
            return false
        } 
        
        if z == len(word) - 1 {
            return true
        }

        if x > 0 && !visited[y * width + x - 1] {
            if check(x-1, y, z+1) {
                return true
            }
        }
        if x < width - 1 && !visited[y * width + x + 1] {
            if check(x+1, y, z+1) {
                return true
            }
        }
        if y > 0 && !visited[(y - 1) * width + x] {
            if check(x, y-1, z+1) {
                return true
            }
        }
        if y < height - 1 && !visited[(y + 1) * width + x]  {
            if check(x, y+1, z+1) {
                return true
            }
        }
        
        visited[y * width + x] = false
        return false
    }
    
    
    for i, _ := range board {
        for j, _ := range board[i] {
            if check(j, i, 0) {
                return true
            }            
        }
    }
    
    return false
}
