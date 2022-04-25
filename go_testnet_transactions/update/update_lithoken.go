package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to mint Lithoken and sell Lithoken (Lithoken must be on sale)
//Update Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("Update Lithoken on flow blockchain (Testnet)")

	//Sell transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/update_lithoken").
		SignProposeAndPayAs("artist"). // Seller address
		UInt64Argument(89638398).      //id NFT storefront
		UFix64Argument("10.0").        //New price
		RunPrintEventsFull()

}
