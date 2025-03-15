package main

import (
    "fmt"
    "calculate/shared/utils"  // Importation du package utils
)

func main() {
    fmt.Println("Hello from app1!")
    utils.PrintHello()  // Appel à la fonction PrintHello du package utils
}
