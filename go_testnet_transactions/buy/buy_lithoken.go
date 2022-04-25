package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//buy Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs and AccountArgument if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("Buy Lithoken on flow blockchain (Testnet)")

	//Buy transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/buy_lithoken").
		SignProposeAndPayAs("buyer1"). // Buyer Address
		UInt64Argument(89576587).      //id NFT storefront
		AccountArgument("artist").     //Seller address
		RunPrintEventsFull()

}
