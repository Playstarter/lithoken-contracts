//---------- testnet--------------
//import NonFungibleToken from 0x631e88ae7f1d7c20
//import LithokenNFT from 0x6ed96c1bdf81bd6e

// ----- Emulator -------------

import NonFungibleToken, LithokenNFT from 0xf8d6e0586b0a20c7
// Setup storage for LithokenNFT on signer account
//
transaction {
    prepare(acct: AuthAccount) {
        if acct.borrow<&LithokenNFT.Collection>(from: LithokenNFT.collectionStoragePath) == nil {
            let collection <- LithokenNFT.createEmptyCollection() //as! @LithokenNFT.Collection
            acct.save(<-collection, to: LithokenNFT.collectionStoragePath)
            acct.link<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver,LithokenNFT.LithokenNFTCollectionPublic}>(LithokenNFT.collectionPublicPath, target: LithokenNFT.collectionStoragePath)
            // create a public capability for the collection
           // acct.link<&LithokenNFT.Collection{NonFungibleToken.CollectionPublic, LithokenNFT.LithokenNFTCollectionPublic}>(LithokenNFT.collectionPublicPath, target: LithokenNFT.collectionStoragePath)
        }
    }
}