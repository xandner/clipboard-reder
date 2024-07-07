package main

import "clip/pkg"

func main() {
	run()
}

func run(){
	newPkg:=pkg.NewProcess()
	newPkg.Init()
}