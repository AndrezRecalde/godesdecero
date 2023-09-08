package main

import (
	"fmt"
	/* "godesdecero/variables" */
	"runtime"
)

func main() {
	//variables.RestoVariables()
	/* estado, texto := variables.ConvierteaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */

	os := runtime.GOOS
	if os == "Linux." || os == "OS X." {
		fmt.Println("No estas wn Windows")
	} else {
		fmt.Println("Estas en Windows: ", os)
	}

	switch os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Println("Este SO es: ", os)
	}
}
