package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts
// And you'll need before to have a Lithoken on sale.
//remove Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("Remove marketplace Lithoken on flow blockchain (Testnet)")

	//Remove transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/remove_lithoken_emulator").
		SignProposeAndPayAs("artist").
		UInt64Argument(89638524). //id NFT storefront
		RunPrintEventsFull()

}
