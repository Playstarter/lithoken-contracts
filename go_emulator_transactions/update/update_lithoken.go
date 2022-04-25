package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) (and mint Lithoken and sell Lithoken) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts
//Update Lithoken : You need to change id NFT storefront (and SignProposeAndPayAs if necessary)
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("Update Lithoken on flow blockchain (Testnet)")

	//Update price on marketplace
	flow.TransactionFromFile("lithoken_emulator/update_lithoken_emulator").
		SignProposeAndPayAs("artist"). // Seller address
		UInt64Argument(89638398).      //id NFT storefront
		UFix64Argument("10.0").        //New price
		RunPrintEventsFull()

}
