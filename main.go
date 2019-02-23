package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}

	trie := NewTrie()

	buf := [1]byte{}
	word := make([]byte, 0, 20)

	for {
		_, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		switch c := buf[0]; {
		case c >= 'a' && c <= 'z':
			word = append(word, c)
		case c >= 'A' && c <= 'Z':
			word = append(word, c+'a'-'A')
		default:
			if len(word) > 0 {
				trie.Add(word)
				word = word[:0]
			}
		}
	}

	queue := NewTopNPriorityQueue()

	// queue implements the Adder interface
	trie.AddAll(queue)

	queue.Print()
}
