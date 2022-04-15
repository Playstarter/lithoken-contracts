# Lithoken smart contracts

This repository contains the smart contracts and transactions for the [Flow](https://www.docs.onflow.org) blockchain
used by Lithoken marketplace.

The smart contracts are written in [Cadence](https://docs.onflow.org/cadence).

## Addresses

| Contract     | Mainnet              | Testnet              |
|--------------|----------------------|----------------------|
| LithokenFee   | `                 ` | `0xee10e1f0d87f28c1` |
| RedevNFT  | `                     ` | `0xee10e1f0d87f28c1` |
| LithokenNFT   | `                 ` | `0xee10e1f0d87f28c1` |
| LithokenOrder | `                 ` | `0xee10e1f0d87f28c1` |

## Smart contracts

`LithokenFee`: This is simple fee manager that holds the rates and addresses for fees.

`RedevNFT`: This is contract interface that adds royalties to NFT. You can implement this `RedevNFT` in your
contract (along with [`NonFungibleToken`](https://github.com/onflow/flow-nft)) and your royalties will be taken when
trading on [Lithoken](https://lithoken.com/).

`LithokenNFT`: The Lithoken NFT contract that implements the [Flow NFT standard](https://github.com/onflow/flow-nft)
which is equivalent to ERC-721 or ERC-1155 on Ethereum.

`LithokenOrder`: This marketplace contract is the wrapper for the
standard [NFTStorefront](https://github.com/onflow/nft-storefront)
for handling market orders.

## Directory structure

The directories here are organized info contracts and transactions. Contracts contain the code that are deployed to
Flow.

- `contracts/`: Where the Lithoken related smart contracts live.
- `contracts/core/`: This contains flow core contracts.
- `transactions/`: This directory contains all the transactions and scripts that are associated with the Lithoken smart
  contracts.
- `transactions/storefront`: Storefront actions for supported NFT's.
