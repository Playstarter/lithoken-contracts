# Lithoken smart contracts

This repository contains the smart contracts and transactions for the [Flow](https://www.docs.onflow.org) blockchain
used by https://lithoken.com

The smart contracts are written in [Cadence](https://docs.onflow.org/cadence).

## Addresses

| Contract     | Mainnet             | Testnet              |
|--------------|---------------------|----------------------|
| LithokenFee  | `                 ` | `0xee10e1f0d87f28c1` |
| RedevNFT     | `                 ` | `0xee10e1f0d87f28c1` |
| LithokenNFT  | `                 ` | `0xee10e1f0d87f28c1` |
| LithokenOrder| `                 ` | `0xee10e1f0d87f28c1` |

## Smart contracts

`LithokenFee`: This is simple fee manager that holds the rates and addresses for fees.

`RedevNFT`: This is contract interface that adds royalties to NFT. You can implement this `RedevNFT` in your
contract (along with [`NonFungibleToken`](https://github.com/onflow/flow-nft)) and your royalties will be taken when
trading on [Lithoken](https://lithoken.com/).

`LithokenNFT`: The Lithoken NFT contract that implements the [Flow NFT standard](https://github.com/onflow/flow-nft).

`LithokenOrder`: This marketplace contract is the wrapper for the
standard [NFTStorefront](https://github.com/onflow/nft-storefront)
for handling market orders.

## Directory structure

- `contracts/`: Where the Lithoken related smart contracts live.
- `contracts/core/`: This contains flow core contracts.
- `demo_emulator/`: This contains the full demo automated (on "Emulator") of the all transactions available on the Lithoken marketplace.
- `demo_testnet/`: This contains the full demo automated (on "Tesnet") of the all transactions available on the Lithoken marketplace.
- `go_emulator_transactions/`: This contains all transactions one by one, used in "demo_emulator".
- `go_testnet_transactions/`: This contains all transactions one by one, used in "testnet_emulator".
- `transactions/`: This directory contains all the transactions that are associated with the Lithoken smart
  contracts.
- `transactions/flowtoken`: Transaction for deposit Flowtoken into an account (use on emulator)
- `transactions/lithoken_emulator`: All transactions written in "Cadence" available on the marketplace Lithoken (used for the Go tests on the Emulator).
- `transactions/lithoken_testnet`: All transactions written in "Cadence" available on the marketplace Lithoken (used for the Go tests on the Testnet).


Follow the guide below to set it up and test locally or in the testnet, in the terminal.

## Prerequisites

1. Ensure Go is [installed on your machine](https://golang.org/dl/) `recommended version 1.17^`
2. [Install the Flow CLI](https://docs.onflow.org/docs/cli) 
3. Run `$ git clone https://github.com/Playstarter/lithoken-contracts.git` in a terminal window
4. Change to the project directory `cd lithoken-contracts`

## How to run the demo marketplace

On the emulator :
  Start two terminals. Both from the root directory.
    
    In the first terminal (ensure flow emulator is not already running):
  - `make emulator`
   In the second terminal :
  - `make marketplace_demo_emulator`

On the Testnet :
 - `make marketplace_demo_testnet`

## What happends in the demo (Emulator)

1. deploy all the contracts
2. setup 2 accounts "artist" and a "buyer1" who each receive 1000 FlowToken
2. a Lithoken (NFT) is minted by the artist account
3. the sale of this Lithoken is made by "artist"
4. purchase of this Lithoken by "buyer1"
5. a new sale of this Lithoken is made by "buyer1"
6. update of the selling price of Lithoken on the marketplace
7. remove Lithoken from the sale 

## What happends in the demo (Testnet)
 
1. a Lithoken (NFT) is minted by the artist account
2. the sale of this Lithoken is made by "artist"
3. purchase of this Lithoken by "buyer1"
4. a new sale of this Lithoken is made by "buyer1"
5. update of the selling price of Lithoken on the marketplace
6. remove Lithoken from the sale by "buyer1" (who is the owner)
