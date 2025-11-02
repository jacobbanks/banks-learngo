package main

import ( f "fmt")

var packageScope = 10 

func scopeTest() {
	var packageScope = 5
	f.Printf("packageScope redeclared in block %d \n", packageScope)
}
