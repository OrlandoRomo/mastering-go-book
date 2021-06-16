package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	// Global variable of the entire document
	document = js.Global().Get("document")
)

//add adds two numbers as input
func add(this js.Value, i []js.Value) interface{} {
	n1, n2, err := getValues(i)
	if err != nil {
		ShowError()
		return nil
	}
	ShowResult(n1 + n2)
	return n1 + n2
}

//substract substracts two numbers as input
func substract(this js.Value, i []js.Value) interface{} {
	n1, n2, err := getValues(i)
	if err != nil {
		ShowError()
		return nil
	}
	ShowResult(n1 - n2)
	return n1 - n2
}

//multiply multiplies two numbers as input
func multiply(this js.Value, i []js.Value) interface{} {
	n1, n2, err := getValues(i)
	if err != nil {
		ShowError()
		return nil
	}
	ShowResult(n1 * n2)
	return n1 * n2
}

//divide divides two numbers as input
func divide(this js.Value, i []js.Value) interface{} {
	n1, n2, err := getValues(i)
	if err != nil {
		ShowError()
		return nil
	}
	ShowResult(n1 / n2)
	return n1 / n2
}

// getValues returns the values typed by the user and error in case the parsing process fails
func getValues(i []js.Value) (float64, float64, error) {
	num1 := document.Call("getElementById", "num1").Get("value").String()
	num2 := document.Call("getElementById", "num2").Get("value").String()
	n1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, 0, err
	}
	n2, err := strconv.ParseFloat(num2, 64)
	return n1, n2, err
}

// ShowError shows an alert when a user types a non-number value
func ShowError() {
	document.Call("getElementById", "resultDiv").Get("style").Set("visibility", "hidden")
	document.Call("getElementById", "closeError").Get("style").Set("visibility", "visible")
}

// ShowResult shows the result of the performed operation
func ShowResult(result float64) {
	document.Call("getElementById", "resultDiv").Get("style").Set("visibility", "visible")
	document.Call("getElementById", "result").Set("innerHTML", result)
}

// onCloseError closes the alert when the user types a non-number value
func onCloseError(this js.Value, i []js.Value) interface{} {
	document.Call("getElementById", "resultDiv").Get("style").Set("visibility", "hidden")
	document.Call("getElementById", "closeError").Get("style").Set("visibility", "hidden")
	return nil
}

// registerCallbacks handles all the callbacks used in javascript
func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("substract", js.FuncOf(substract))
	js.Global().Set("multiply", js.FuncOf(multiply))
	js.Global().Set("divide", js.FuncOf(divide))
	js.Global().Set("onCloseError", js.FuncOf(onCloseError))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Go Webassembly initialized")
	registerCallbacks()
	<-c
}
