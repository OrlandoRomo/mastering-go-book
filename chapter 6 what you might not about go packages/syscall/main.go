package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// pid proccess of this golang program
	pid, _, _ := syscall.Syscall(syscall.SYS_GETPPID, 0, 0, 0)
	fmt.Println("My pid is: ", pid)
	// uid of the user who's running this process
	uid, _, _ := syscall.Syscall(syscall.SYS_GETUID, 0, 0, 0)
	fmt.Println("My uid is: ", uid)

	message := []byte("Hello running system calls\n")
	syscall.Write(1, message)

	fmt.Println("Using syscall.Exec()")
	cmd := "/bin/ls"
	env := os.Environ()
	syscall.Exec(cmd, []string{"ls", "-a"}, env)

}
