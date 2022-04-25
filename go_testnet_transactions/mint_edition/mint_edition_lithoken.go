package main

import (
	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

// Artist can also mint various editions of his Lithoken. In our website marketplace https://lithoken.com, if artist mint various editions of his art he must choose an editions name.
func main() {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	// mint Lithoken
	flow.TransactionFromFile("lithoken_testnet/mint_lithoken").
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
