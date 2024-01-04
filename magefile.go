//go:build mage
//+build mage

package main

import (
	"github.com/magefile/mage/sh"
	
	"os"
	"os/exec"
	"path/filepath"
)

var BINARY_NAME = "logistics-routes"
var BIN = "bin"
var CODE_FOLDERS = "internal"

// Construir el programa principal
func Build() error {
	internalPath := "./" + CODE_FOLDERS
	files, err := filepath.Glob(filepath.Join(internalPath, "*.go"))
	if err != nil {
		return err
	}
	
	args := append([]string{"build"}, files...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

// Instalación de las dependencias
func Install() {
	println("Instalando las dependencias...")
	sh.Run("go", "mod", "tidy")
	println("Instalación de dependencias finalizada")
}

// Ejecutar el programa
func Run() {
	println("Ejecutando el programa %s...\n", BINARY_NAME)
	// runCommand(filepath.Join(BIN, BINARY_NAME))
}

// Limpiar el proyecto
func Clean() {
	println("Limpiando los binarios...")
	sh.Run("go", "clean", BIN)
}

// Comprobar la sintaxis
func Check() {
	internalPath := "./" + CODE_FOLDERS
	println("Comprobando sintaxis del proyecto...")
	sh.Run("gofmt", "-l", internalPath)
	println("Check finalizado")
}

// Ejecutar pruebas
func Test() {
	internalPath := "./" + CODE_FOLDERS
	println("Ejecutando pruebas...")
	execCmd("go", "test", "-v", internalPath)
	println("Pruebas finalizadas")
}

func execCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		println("Error executing command: %v\n", err)
		os.Exit(1)
	}
}