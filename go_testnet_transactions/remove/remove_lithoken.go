package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

// And you'll need before to have a Lithoken on sale.
//remove Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("Remove marketplace Lithoken on flow blockchain (Testnet)")

	//Remove transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/remove_lithoken").
		SignProposeAndPayAs("artist"). // Owner Lithoken
		UInt64Argument(89638524).      //id NFT storefront
		RunPrintEventsFull()

}
