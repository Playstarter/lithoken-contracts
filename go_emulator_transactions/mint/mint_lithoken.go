package main

import (
	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts ==> make deploy_contracts
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	// mint Lithoken
	flow.TransactionFromFile("lithoken_emulator/mint_lithoken_emulator").
		SignProposeAndPayAs("artist").
		AccountArgument("artist").
		UFix64Argument("0.10").                                                                                      //royalty creator
		UInt64Argument(1).                                                                                           //number edition
		StringArgument("").                                                                                          //name of serie ("" if number edition !> 1)
		StringArgument("My beautiful Lithoken").                                                                     //name of lithoken (NFT)
		StringArgument("MyArtistName").                                                                              //artist name
		StringArgument("This is a description of my art Lithoken").                                                  //description
		StringArgument("This is a dedication if I want it").                                                         //dedication
		StringArgument("Type of file").                                                                              //type of file
		StringArgument("https://ipfs.pixura.io/ipfs/QmUyARmq5RUJk5zt7KUeaMLYB8SQbKHp3Gdqy5WSxRtPNa/SeaofRoses.jpg"). //link IPFS
		StringArgument("Painting").                                                                                  //name of a classification
		RunPrintEventsFull()

}
