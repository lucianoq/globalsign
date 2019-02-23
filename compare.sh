#!/usr/bin/env bash
set -x

#echo "Solution:"
cat mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20

#echo "Result:"
go run main.go top_n_priority_queue.go trie.go

#echo "Differences:"

diff <(cat mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20) <(go run main.go top_n_priority_queue.go trie.go)