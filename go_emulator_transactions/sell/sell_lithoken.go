package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts
//Sell Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("Sell Lithoken on flow blockchain (Testnet)")

	//Sell transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/sell_lithoken_emulator").
		SignProposeAndPayAs("artist"). //<<<<< Put Seller
		UInt64Argument(114).           //<<<<< Put id NFT
		UFix64Argument("20.00").       //Flow price NFT
		RunPrintEventsFull()

}
