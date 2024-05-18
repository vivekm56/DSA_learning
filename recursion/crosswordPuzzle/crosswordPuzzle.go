package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printMatrix(matrix []string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func checkHorizontal(x, y int, matrix []string, currentWord string) bool {
	n := len(currentWord)
	for i := 0; i < n; i++ {
		if y+i >= len(matrix[x]) || (matrix[x][y+i] != '-' && matrix[x][y+i] != currentWord[i]) {
			return false
		}
	}
	return true
}

func placeHorizontal(x, y int, matrix []string, currentWord string) []string {
	newMatrix := make([]string, len(matrix))
	copy(newMatrix, matrix)
	for i := 0; i < len(currentWord); i++ {
		newMatrix[x] = newMatrix[x][:y+i] + string(currentWord[i]) + newMatrix[x][y+i+1:]
	}
	return newMatrix
}

func checkVertical(x, y int, matrix []string, currentWord string) bool {
	n := len(currentWord)
	for i := 0; i < n; i++ {
		if x+i >= len(matrix) || (matrix[x+i][y] != '-' && matrix[x+i][y] != currentWord[i]) {
			return false
		}
	}
	return true
}

func placeVertical(x, y int, matrix []string, currentWord string) []string {
	newMatrix := make([]string, len(matrix))
	copy(newMatrix, matrix)
	for i := 0; i < len(currentWord); i++ {
		newMatrix[x+i] = newMatrix[x+i][:y] + string(currentWord[i]) + newMatrix[x+i][y+1:]
	}
	return newMatrix
}

func solvePuzzle(words []string, matrix []string, index int) bool {
	if index == len(words) {
		printMatrix(matrix)
		return true
	}

	currentWord := words[index]
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if checkHorizontal(i, j, matrix, currentWord) {
				newMatrix := placeHorizontal(i, j, matrix, currentWord)
				if solvePuzzle(words, newMatrix, index+1) {
					return true
				}
			}
			if checkVertical(i, j, matrix, currentWord) {
				newMatrix := placeVertical(i, j, matrix, currentWord)
				if solvePuzzle(words, newMatrix, index+1) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read the grid
	matrix := make([]string, 10)
	for i := 0; i < 10; i++ {
		line, _ := reader.ReadString('\n')
		matrix[i] = strings.TrimSpace(line)
	}

	// Read the words
	line, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(line), ";")

	// Solve the puzzle
	solvePuzzle(words, matrix, 0)
}
