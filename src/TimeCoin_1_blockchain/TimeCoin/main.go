// TimeCoin project main.go
package main

import (
	"TimeCoin/block"
	"fmt"
)

func main() {
	bc := block.NewBlockChain()

	bc.AddBlock("I am the first transaction...")
	bc.AddBlock("Send a timecoin to you... ")

	for _, blk := range bc.Blocks {
		fmt.Printf("Prev hash: %x\n", blk.PrevBlockHash)
		fmt.Printf("Data: %s\n", blk.Data)
		fmt.Printf("Hash: %x\n", blk.Hash)

		fmt.Println()
	}
}
