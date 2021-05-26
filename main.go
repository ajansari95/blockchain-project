package main

import (
	"awesomeProject/blockchain"
	"fmt"
	"strconv"
)



func main() {
	chain:= blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("S Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _,block :=range chain.Blocks{
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow :=blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}


}
