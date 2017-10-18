package cli

import (
	"block"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	Bc *block.Blockchain
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (clid *CLI) validateArgs() {
	fmt.Println("this is validateArgs")
}

func (cli *CLI) printUsage() {
	fmt.Println("this is printUsage")
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		aBlock := bci.Next()

		fmt.Printf("Prev. hash: %x\n", aBlock.PrevBlockHash)
		fmt.Printf("Data: %s\n", aBlock.Data)
		fmt.Printf("Hash: %x\n", aBlock.Hash)
		pow := block.NewProofOfWork(aBlock)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(aBlock.PrevBlockHash) == 0 {
			break
		}
	}
}
