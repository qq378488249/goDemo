package main

import "goDemo/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("1")
	bc.SendData("2")

	bc.Print()
}
