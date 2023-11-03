package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//firstCLIApp()
	mainHandler()
}

func firstCLIApp() {
	fmt.Println("What would you like me to scream?")
	myInput := bufio.NewReader(os.Stdin)
	userInput, _ := myInput.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	userInput = strings.ToUpper(userInput)
	fmt.Println(userInput + "!")
}

func mainHandler() {
	http.HandleFunc("/", ShowingMenu)
	http.ListenAndServe(":8080", nil)
}

func ShowingMenu(res http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("./menu.txt")

	io.Copy(res, file)
}
