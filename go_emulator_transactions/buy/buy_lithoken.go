package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts
//buy Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs and AccountArgument if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("Buy Lithoken on flow blockchain (Testnet)")

	//Buy Lithoken transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/buy_lithoken_emulator").
		SignProposeAndPayAs("buyer1"). //<<<<< Put Buyer address
		UInt64Argument(89576587).      //<<<<< Put id NFT storefront
		AccountArgument("artist").     //<<<<< Put Seller address
		RunPrintEventsFull()

}
