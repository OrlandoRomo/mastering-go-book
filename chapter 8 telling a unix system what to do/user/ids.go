package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User ID", os.Getuid())

	var u *user.User
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print("Group IDs: ")
	groupIds, err := u.GroupIds()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, id := range groupIds {
		fmt.Print(id, " ")
	}
	fmt.Println()
}
