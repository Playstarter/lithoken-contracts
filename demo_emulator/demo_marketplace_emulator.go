package main

import (
	"fmt"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

func mint_lithoken(accountName string) {

	//Use Lithoken contracts on Flow Emulator
	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("\n \n ******* Mint Lithoken on flow blockchain (Emulator) *******\n \n ")

	//Transaction to mint a Lithoken (NFT)
	flow.TransactionFromFile("lithoken_emulator/mint_lithoken_emulator").
		SignProposeAndPayAs(accountName).
		AccountArgument(accountName).
		UFix64Argument("0.10").                                                                                      //royalty creator
		UInt64Argument(1).                                                                                           //number edition
		StringArgument("").                                                                                          //name of editions ("" if number edition !> 1)
		StringArgument("My beautiful Lithoken").                                                                     //name of lithoken (NFT)
		StringArgument("MyArtistName").                                                                              //artist name
		StringArgument("This is a description of my art Lithoken").                                                  //description
		StringArgument("This is a dedication if I want it").                                                         //dedication
		StringArgument("Type of file").                                                                              //type of file
		StringArgument("https://ipfs.pixura.io/ipfs/QmUyARmq5RUJk5zt7KUeaMLYB8SQbKHp3Gdqy5WSxRtPNa/SeaofRoses.jpg"). //(example) link IPFS
		StringArgument("Painting").                                                                                  //name of a classification
		RunPrintEventsFull()
}

// ----- Function for the sale of Lithoken on marketplace -----//
func sell_lithoken(accountName string, idNFT uint64, price string) {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("\n \n******* Sell Lithoken on flow blockchain (Emulator) *******\n \n ")

	//Sell transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/sell_lithoken_emulator").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idNFT). //id NFT
		UFix64Argument(price). //Flow price NFT
		RunPrintEventsFull()

}

// ----- Function for update price Lithoken on marketplace -----//
func update_lithoken(accountName string, idStoreFront uint64, newPrice string) {
	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("\n \n ******* Update Lithoken on flow blockchain (Emulator) *******\n \n ")

	//Update price transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/update_lithoken_emulator").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idStoreFront). //id NFT storefront
		UFix64Argument(newPrice).     //New price
		RunPrintEventsFull()
}

// ----- Function for remove Lithoken on marketplace -----//
func remove_lithoken(accountName string, idNFTstoreFront uint64) {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("\n \n ******* Remove marketplace Lithoken on flow blockchain (Emulator) *******\n \n ")

	//Transaction to remove Lithoken on marketplace
	flow.TransactionFromFile("lithoken_emulator/remove_lithoken_emulator").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idNFTstoreFront). //id NFT storefront
		RunPrintEventsFull()

}

// ----- Function to buy Lithoken on marketplace -----//
func buy_lithoken(buyer string, seller string, idNFTstoreFront uint64) {

	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")

	fmt.Println("\n \n ******* Buy Lithoken on flow blockchain (Emulator) *******\n \n ")

	//Buy transaction on marketplace
	flow.TransactionFromFile("lithoken_emulator/buy_lithoken_emulator").
		SignProposeAndPayAs(buyer).
		UInt64Argument(idNFTstoreFront). //id NFT storefront
		AccountArgument(seller).         //Seller address
		RunPrintEventsFull()

}

// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// 	This main() function is a demo will show you all transactions available on marketplace Lithoken :
// 	mint Lithoken, buy Lithoken, sell Lithoken, update price Lithoken, and remove Lithoken.
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func main() {

	//------------------------------------------------------------//
	//initialization contracts and accounts : emulator blockchain //
	//------------------------------------------------------------//
	flow := gwtf.NewGoWithTheFlowForNetwork("emulator")
	//Deploy contracts
	flow.InitializeContracts()
	//Deposit FlowToken on "artist" account and "buyer1" account
	fmt.Println("\n \n ******* Deposit FlowToken on 'artist' account and 'buyer1' account (Emulator) *******\n \n ")
	flow.TransactionFromFile("flowtoken/mint_token").SignProposeAndPayAsService().AccountArgument("artist").UFix64Argument("1000.0").RunPrintEventsFull()
	flow.TransactionFromFile("flowtoken/mint_token").SignProposeAndPayAsService().AccountArgument("buyer1").UFix64Argument("1000.0").RunPrintEventsFull()

	//------------------------------------------------------------//
	//			Mint Lithoken on artist account               	  //
	//------------------------------------------------------------//
	mint_lithoken("artist")

	//------------------------------------------------------------//
	//			Sell Lithoken of artist account               	  //
	//------------------------------------------------------------//
	sell_lithoken("artist", 0, "10.00")

	//------------------------------------------------------------//
	//			Buy Lithoken (artist ==> buyer1)              	  //
	//------------------------------------------------------------//
	buy_lithoken("buyer1", "artist", 39)

	//------------------------------------------------------------//
	//			Re-sell Lithoken by the buyer1              	  //
	//------------------------------------------------------------//
	sell_lithoken("buyer1", 0, "10.00")

	//------------------------------------------------------------//
	//			Update price buyer1's Lithoken on sale            //
	//------------------------------------------------------------//
	update_lithoken("buyer1", 47, "50.00")

	//------------------------------------------------------------//
	//			Remove buyer1's Lithoken from sale marketplace    //
	//------------------------------------------------------------//
	remove_lithoken("buyer1", 48)
}
