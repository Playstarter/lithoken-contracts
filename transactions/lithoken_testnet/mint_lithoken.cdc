//---------- testnet--------------
import NonFungibleToken from 0x631e88ae7f1d7c20
import LithokenNFT from 0xee10e1f0d87f28c1

//---------- local --------------------
//import NonFungibleToken, LithokenNFT from 0xf8d6e0586b0a20c7

//Mint LithokenNFT token to signer acct

transaction( adresse:Address, fee:UFix64, nbrEdit:UInt64, nomSerie:String, nameNFT:String, nameArtist:String, desc:String, dedi:String, tp:String, lienIPFS:String, nomCollection:String ) {
    let minter: Capability<&LithokenNFT.Minter>
    let receiver: Capability<&{NonFungibleToken.Receiver}>
    //modif
    let name: String 
    let artist: String
    let description: String
    let dedicace: String
    let type: String
    let ipfs: String
    let collection: String
    //let edition: UInt64
    let nameSerie: String
    let nbrEdition: UInt64
    var i : UInt64	

   

    prepare(acct: AuthAccount) {
        if acct.borrow<&LithokenNFT.Collection>(from: LithokenNFT.collectionStoragePath) == nil {
            let collection <- LithokenNFT.createEmptyCollection() //as! @LithokenNFT.Collection
            acct.save(<-collection, to: LithokenNFT.collectionStoragePath)
            acct.link<&{NonFungibleToken.CollectionPublic,NonFungibleToken.Receiver,LithokenNFT.LithokenNFTCollectionPublic}>(LithokenNFT.collectionPublicPath, target: LithokenNFT.collectionStoragePath)
        }

        self.minter = LithokenNFT.minter()
        self.receiver = acct.getCapability<&{NonFungibleToken.Receiver}>(LithokenNFT.collectionPublicPath)
        //modif
        self.name = nameNFT
        self.artist = nameArtist
        self.description = desc
        self.dedicace = dedi
        self.type= tp
        self.ipfs= lienIPFS
        self.collection= nomCollection
        //self.edition=1
        self.nameSerie=nomSerie
        self.nbrEdition=nbrEdit
        self.i = 1

    }

    execute {
        if self.nbrEdition == 1{
            let minter = self.minter.borrow() ?? panic("Could not borrow receiver capability (maybe receiver not configured?)")
                minter.mintTo(creator: self.receiver, name:self.name ,artist:self.artist, description:self.description ,dedicace:self.dedicace ,type:self.type ,ipfs:self.ipfs, collection:self.collection, nomSerie:"", edition:self.nbrEdition ,nbrEdition:self.nbrEdition , royalties: [LithokenNFT.Royalty(Address:adresse, UFix64:fee)])
        } else {
            while self.i <=  self.nbrEdition {
            
                let minter = self.minter.borrow() ?? panic("Could not borrow receiver capability (maybe receiver not configured?)")
                minter.mintTo(creator: self.receiver, name:self.name ,artist:self.artist, description:self.description ,dedicace:self.dedicace ,type:self.type ,ipfs:self.ipfs, collection:self.collection, nomSerie:self.nameSerie, edition:self.i ,nbrEdition:self.nbrEdition , royalties: [LithokenNFT.Royalty(Address:adresse, UFix64:fee)])
                self.i= self.i+1
            }
        }
    }
}