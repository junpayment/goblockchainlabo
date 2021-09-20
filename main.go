package main

import (
	"fmt"
	"strconv"

	"github.com/junpayment/goblockchainlabo/lib"
)

func main()  {
	var blocks [10]*lib.Block
	for h:=0; h < 10; h++ {
		hashPrev := func() string {
			if h == 0 {
				return lib.FirstHash
			}
			return blocks[h-1].HashOwn
		}()
		blocks[h] = lib.NewBlock(strconv.Itoa(h), hashPrev)
	}
	for i, v := range blocks {
		fmt.Printf("%d: %v \n", i, v)
	}
}
