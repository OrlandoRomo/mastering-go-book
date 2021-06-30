/*Create a key-value store with the following features.
1. Add new element
2. Delete an existing element from the key-value store on a key
3. Lookin up the value specifying a key in the stor
4. Changing the value of an existing key
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Print  = "PRINT"
	Stop   = "Stop"
	Add    = "ADD"
	LookUp = "LOOKUP"
	Delete = "DELETE"
	Update = "UPDATE"
)

var STORE_LOG = "/tmp/store_log.log"

func NewLogger() (*log.Logger, *os.File) {
	file, err := os.OpenFile(STORE_LOG, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "key-store: ", log.LstdFlags)
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
	return logger, file
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	logger, file := NewLogger()
	defer file.Close()
	s := NewStore()
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}
		switch tokens[0] {
		case Print:
			s.Print()
		case Stop:
			return
		case Delete:
			if !s.Delete(tokens[1]) {
				logger.Printf("Delete operation failed %v\n", tokens[1])
				fmt.Println("Delete operation failed")
			}
		case Add:
			element := &Element{
				tokens[2],
				tokens[3],
				tokens[4],
			}
			if !s.Add(tokens[1], element) {
				logger.Printf("Add operation failed %v\n", tokens[1])
				fmt.Println("Add operation failed")
			}
		case LookUp:
			element := s.store[tokens[1]]
			if element == nil {
				logger.Printf("%v key does not exist\n", tokens[1])
				fmt.Printf("%v key does not exist\n", tokens[1])
				break
			}
			logger.Printf("%v key-element found %v\n", tokens[1], element)
			fmt.Printf("%v key-element found %v\n", tokens[1], element)
		case Update:
			element := &Element{
				tokens[2],
				tokens[3],
				tokens[4],
			}
			if !s.Update(tokens[1], element) {
				logger.Printf("Update operation failed %v\n", tokens[1])
				fmt.Println("Update operation failed")

			}
		default:
			logger.Printf("%v is not a valid.\n", tokens[1])
			fmt.Println("Unknown command - please try again!")
		}
	}
}
