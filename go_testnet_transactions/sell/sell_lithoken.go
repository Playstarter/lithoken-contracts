package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//Sell Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("Sell Lithoken on flow blockchain (Testnet)")

	//Sell transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/sell_lithoken").
		SignProposeAndPayAs("artist"). // Seller
		UInt64Argument(114).           //id NFT
		UFix64Argument("20.00").       //Flow price NFT
		RunPrintEventsFull()

}
