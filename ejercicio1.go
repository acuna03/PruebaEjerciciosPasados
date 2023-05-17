package main

import "fmt"

func obtenermultiplosdetres(letters []string) []string {
	var result []string

	for i, letter := range letters {
		if (i+1)%3 == 0 {
			fmt.Println("Posicion:", i+1, "Letter:", letter)
			result = append(result, letter)
		}
	}

	return result
}

func main() {
	alfabeto := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	fmt.Println("Alfabeto:", alfabeto)

	multiplosdetres := obtenermultiplosdetres(alfabeto)

	fmt.Println("obtenermultiplosdetres:", multiplosdetres)
}
