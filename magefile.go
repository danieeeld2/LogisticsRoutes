package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var BINARY_NAME = "logistics-routes"
var BIN = "./bin"
var CODE_FOLDERS = "internal/*"

// Construir el programa principal
func Build() {
	fmt.Println("Construyendo el proyecto...")
	runCommand("go", "build")
}

// Instalación de las dependencias
func InstallDeps() {
	fmt.Println("Instalando las dependencias...")
	runCommand("go", "mod", "tidy")
}

// Ejecutar el programa
func Run() {
	fmt.Printf("Ejecutando el programa %s...\n", BINARY_NAME)
	runCommand(filepath.Join(BIN, BINARY_NAME))
}

// Limpiar el proyecto
func Clean() {
	fmt.Println("Limpiando los binarios...")
	os.RemoveAll(filepath.Join(BIN, BINARY_NAME))
	runCommand("go", "clean", "./...")
}

// Comprobar la sintaxis
func Check() {
	fmt.Println("Comprobando sintaxis del proyecto...")
	runCommand("gofmt", "-l", CODE_FOLDERS)
}


