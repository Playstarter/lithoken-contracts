package main

import (
	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

//NB! before you will need to launch emulator and create account(s) ==> in root directory ==> make emulator ==> (and in another terminal )==> make artist_and_buyer_accounts ==> make deploy_contracts
// Artist can also mint various editions of his Lithoken. In our website marketplace https://lithoken.com, if artist mint various editions of his art he must choose an editions name.
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	// mint Lithoken
	flow.TransactionFromFile("lithoken_emulator/mint_lithoken_emulator").
		SignProposeAndPayAs("artist").
		AccountArgument("artist").
		UFix64Argument("0.10").                                                                                      //royalty creator
		UInt64Argument(5).                                                                                           //number edition
		StringArgument("My editions name").                                                                          //name of editions ("" if number edition !> 1)
		StringArgument("My beautiful Lithoken").                                                                     //name of lithoken (NFT)
		StringArgument("MyArtistName").                                                                              //artist name
		StringArgument("This is a description of my art Lithoken").                                                  //description
		StringArgument("This is a dedication if I want it").                                                         //dedication
		StringArgument("Type of file").                                                                              //type of file
		StringArgument("https://ipfs.pixura.io/ipfs/QmUyARmq5RUJk5zt7KUeaMLYB8SQbKHp3Gdqy5WSxRtPNa/SeaofRoses.jpg"). //link IPFS
		StringArgument("Painting").                                                                                  //name of a classification
		RunPrintEventsFull()

}
