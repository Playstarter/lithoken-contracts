//---------- emulator--------------
import RedevNFT from 0xf8d6e0586b0a20c7
import LithokenNFT from 0xf8d6e0586b0a20c7
import LithokenOrder from 0xf8d6e0586b0a20c7
import FlowToken from 0x0ae53cb6e3f42a79
import FungibleToken from 0xee82856bf20e2aa6
import NonFungibleToken from 0xf8d6e0586b0a20c7
import NFTStorefront from 0xf8d6e0586b0a20c7

// Cancels order with [orderId], then open new order with same LithokenNFT token for FlowToken [price]

transaction(orderId: UInt64) {
    let nftProvider: Capability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>
    let storefront: &NFTStorefront.Storefront
    let listing: &NFTStorefront.Listing{NFTStorefront.ListingPublic}
    let orderAddress: Address

    prepare(acct: AuthAccount) {
        let nftProviderPath = /private/LithokenNFTProviderForNFTStorefront
        if !acct.getCapability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath)!.check() {
            acct.link<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath, target: LithokenNFT.collectionStoragePath)
        }

        self.nftProvider = acct.getCapability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath)!
        assert(self.nftProvider.borrow() != nil, message: "Missing or mis-typed nft collection provider")

        self.storefront = acct.borrow<&NFTStorefront.Storefront>(from: NFTStorefront.StorefrontStoragePath)
            ?? panic("Missing or mis-typed NFTStorefront Storefront")

        self.listing = self.storefront.borrowListing(listingResourceID: orderId)
            ?? panic("No Offer with that ID in Storefront")

        self.orderAddress = acct.address
    }

    execute {
        LithokenOrder.removeOrder(
            storefront: self.storefront,
            orderId: orderId,
            orderAddress: self.orderAddress,
            listing: self.listing,
        )
    }
}