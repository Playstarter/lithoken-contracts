package main

import (
	"fmt"
	"strconv"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
)

func mint_lithoken(accountName string) uint64 {

	var idLithoken uint64

	//Use Lithoken contracts on Flow Testnet
	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("\n \n ******* Mint Lithoken on flow blockchain (Testnet) *******\n \n ")

	//Transaction to mint a Lithoken (NFT)
	flow.TransactionFromFile("lithoken_testnet/mint_lithoken").
		SignProposeAndPayAs(accountName).
		AccountArgument(accountName).
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

	eventsFetcher := flow.EventFetcher().
		Last(1000).
		Event("A.ee10e1f0d87f28c1.LithokenNFT.Mint").
		EventIgnoringFields("A.ee10e1f0d87f28c1.LithokenNFT.Mint", []string{"creator", "royalties"})

	events, err := eventsFetcher.Run()
	if err != nil {
		panic(err)
	} else {
		nbrEvent := len(events)
		i := 1
		for _, event := range events {
			if i == nbrEvent {
				id, err := strconv.ParseUint(string(event.Fields["id"].(string)), 10, 64)
				if err != nil {
					panic(err)
				} else {
					idLithoken = id
				}
			}
			i++
		}
	}
	return idLithoken
}

// ----- Function for the sale of Lithoken on marketplace -----//
func sell_lithoken(accountName string, idNFT uint64, price string) uint64 {

	var idLithokenStrorefront uint64
	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("\n \n******* Sell Lithoken on flow blockchain (Testnet) *******\n \n ")

	//Sell transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/sell_lithoken").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idNFT). //id NFT
		UFix64Argument(price). //Flow price NFT
		RunPrintEventsFull()

	eventsFetcher := flow.EventFetcher().
		Last(1000).
		Event("A.ee10e1f0d87f28c1.NFTStorefront.ListingAvailable").
		EventIgnoringFields("A.ee10e1f0d87f28c1.NFTStorefront.ListingAvailable", []string{"ftVaultType", "nftID", "nftType", "storefrontAddress"})

	events, err := eventsFetcher.Run()
	if err != nil {
		panic(err)
	} else {
		nbrEvent := len(events)
		i := 1
		for _, event := range events {
			if i == nbrEvent {
				id, err := strconv.ParseUint(string(event.Fields["listingResourceID"].(string)), 10, 64)
				if err != nil {
					panic(err)
				} else {
					idLithokenStrorefront = id
				}
			}
			i++
		}
	}

	return idLithokenStrorefront
}

// ----- Function for update price Lithoken on marketplace -----//
func update_lithoken(accountName string, idStoreFront uint64, newPrice string) uint64 {

	var idLithokenStrorefront uint64
	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("\n \n ******* Update Lithoken on flow blockchain (Testnet) *******\n \n ")

	//Update price transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/update_lithoken").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idStoreFront). //id NFT storefront
		UFix64Argument(newPrice).     //New price
		RunPrintEventsFull()

	eventsFetcher := flow.EventFetcher().
		Last(1000).
		Event("A.ee10e1f0d87f28c1.NFTStorefront.ListingAvailable").
		EventIgnoringFields("A.ee10e1f0d87f28c1.NFTStorefront.ListingAvailable", []string{"ftVaultType", "nftID", "nftType", "storefrontAddress"})

	events, err := eventsFetcher.Run()
	if err != nil {
		panic(err)
	} else {
		nbrEvent := len(events)
		i := 1
		for _, event := range events {
			if i == nbrEvent {
				id, err := strconv.ParseUint(string(event.Fields["listingResourceID"].(string)), 10, 64)
				if err != nil {
					panic(err)
				} else {
					idLithokenStrorefront = id
				}
			}
			i++
		}
	}

	return idLithokenStrorefront

}

// ----- Function for remove Lithoken on marketplace -----//
func remove_lithoken(accountName string, idNFTstoreFront uint64) {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("\n \n ******* Remove marketplace Lithoken on flow blockchain (Testnet) *******\n \n ")

	//Transaction to remove Lithoken on marketplace
	flow.TransactionFromFile("lithoken_testnet/remove_lithoken").
		SignProposeAndPayAs(accountName).
		UInt64Argument(idNFTstoreFront). //id NFT storefront
		RunPrintEventsFull()

}

// ----- Function to buy Lithoken on marketplace -----//
func buy_lithoken(buyer string, seller string, idNFTstoreFront uint64) {

	flow := gwtf.NewGoWithTheFlowForNetwork("testnet")

	fmt.Println("\n \n ******* Buy Lithoken on flow blockchain (Testnet) *******\n \n ")

	//Buy transaction on marketplace
	flow.TransactionFromFile("lithoken_testnet/buy_lithoken").
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

	//Mint a Lithoken
	var lithokenId uint64 = mint_lithoken("artist")

	//Sell Lithoken on marketplace
	var lithokenIdStoreFront uint64 = sell_lithoken("artist", lithokenId, "20.00")

	// Buy Lithoken (artist ==> buyer1)
	buy_lithoken("buyer1", "artist", lithokenIdStoreFront)

	//Re-sell Lithoken by the buyer1
	lithokenIdStoreFront = sell_lithoken("buyer1", lithokenId, "10.00")

	// Update price Lithoken's buyer1 on sale
	lithokenIdStoreFront = update_lithoken("buyer1", lithokenIdStoreFront, "50.00")

	//Remove Lithoken's buyer1
	remove_lithoken("buyer1", lithokenIdStoreFront)

}
