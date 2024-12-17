package main

import "github.com/lemavisaitov/applied-informatics_3/internal/controller"

func main() {
	appController := controller.New()
	appController.Run()
}
