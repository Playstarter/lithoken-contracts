//---------- emulator --------------
import RedevNFT from 0xf8d6e0586b0a20c7
import LithokenNFT from 0xf8d6e0586b0a20c7
import LithokenOrder from 0xf8d6e0586b0a20c7
import FlowToken from 0x0ae53cb6e3f42a79
import FungibleToken from 0xee82856bf20e2aa6
import NFTStorefront from 0xf8d6e0586b0a20c7
import NonFungibleToken from 0xf8d6e0586b0a20c7

// Sell LithokenNFT token for FlowToken with NFTStorefront
//
transaction(tokenId: UInt64, price: UFix64) {
    let nftProvider: Capability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>
    let storefront: &NFTStorefront.Storefront

    prepare(acct: AuthAccount) {
        let nftProviderPath = /private/LithokenNFTProviderForNFTStorefront
        if !acct.getCapability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath)!.check() {
            acct.link<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath, target: LithokenNFT.collectionStoragePath)
        }

        self.nftProvider = acct.getCapability<&{NonFungibleToken.Provider,NonFungibleToken.CollectionPublic,RedevNFT.CollectionPublic}>(nftProviderPath)!
        assert(self.nftProvider.borrow() != nil, message: "Missing or mis-typed nft collection provider")

        if acct.borrow<&NFTStorefront.Storefront>(from: NFTStorefront.StorefrontStoragePath) == nil {
            let storefront <- NFTStorefront.createStorefront() as! @NFTStorefront.Storefront
            acct.save(<-storefront, to: NFTStorefront.StorefrontStoragePath)
            acct.link<&NFTStorefront.Storefront{NFTStorefront.StorefrontPublic}>(NFTStorefront.StorefrontPublicPath, target: NFTStorefront.StorefrontStoragePath)
        }
        self.storefront = acct.borrow<&NFTStorefront.Storefront>(from: NFTStorefront.StorefrontStoragePath)
            ?? panic("Missing or mis-typed NFTStorefront Storefront")
    }

    execute {
        let royalties: [LithokenOrder.PaymentPart] = []
        let extraCuts: [LithokenOrder.PaymentPart] = []
        
        for royalty in self.nftProvider.borrow()!.getRoyalties(id: tokenId) {
            royalties.append(LithokenOrder.PaymentPart(address: royalty.address, rate: royalty.fee))
        }
        
        
        LithokenOrder.addOrder(
            storefront: self.storefront,
            nftProvider: self.nftProvider,
            nftType: Type<@LithokenNFT.NFT>(),
            nftId: tokenId,
            vaultPath: /public/flowTokenReceiver,
            vaultType: Type<@FlowToken.Vault>(),
            price: price,
            extraCuts: extraCuts,
            royalties: royalties
        )
    }
}