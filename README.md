## Context

Almost all NFT smart contracts on Ethereum are deployed following
the [ERC721 standard](https://eips.ethereum.org/EIPS/eip-721). A common
standard allows interactions with NFTs to be simple and consistent
across many different services and applications. One commonly used
external function is

```
/// @notice Find the owner of an NFT
/// @dev NFTs assigned to zero address are considered invalid, and queries
///  about them do throw.
/// @param _tokenId The identifier for an NFT
/// @return The address of the owner of the NFT
function ownerOf(uint256 _tokenId) external view returns (address);
```

The `ownerOf` function is particularly useful as it allows a caller to figure out
the Ethereum account that owns a specific NFT. However, developers are often
surprised to find that the standard ERC721 interface does **not** implement a
function such as

```
function tokensOf(address owner) external view returns (uint256[]);
```

that allows an application to query all owned tokens for a specific Ethereum account.
This obviously makes it more complicated to implement a simple feature such as showing
an Ethereum user a list of all of their owned tokens for a given NFT.

Enumerable functions such as the one above are omitted from the standard as it has
significant implications on the gas cost efficiency of the contract.
[Gas](https://ethereum.org/en/developers/docs/gas/) is a unit that measures the
computational power required to execute blockchain operations. In order for a
transaction to be processed by the network, the gas cost of the transaction is
converted into ETH, and then paid by the sender. Functions that only read blockchain
state without modifying it (for example the `ownerOf` call), can be executed without
the sender paying for it. However, these read-only calls are also measured in Gas
and Ethereum nodes will generally also limit the amount of gas that can be spent
on executing read-only calls within a certain timeframe.

## Assessment

With the above constraints in mind, many organizations provide a service to make
querying an account's owned tokens easier and more efficiently. For example,
OpenSea provides an [API for querying a list of owned assets for a specific account](https://docs.opensea.io/reference/getting-assets). This is at times referred to as an asset
directory.

The goal of this assessment is to implement a simple asset directory service,
with a REST API that allows querying for owned assets for a given user.

### Requirements

- Implement an asset directory to track the following contract addresses

  - `0x61B91a780945971b07ba3898A8E0Dc8201dB46b3`
  - `0xd89148d8dEFc7E8942cA4b16DdBE0E2f6485a4c8`
  - `0x9D7f9672060EED641ebc1b22443132eDf4967D91`

  The above addresses represent three different ERC721 NFT contracts deployed
  on the Goerli Test Network ([Test Networks](https://ethereum.org/en/developers/docs/development-networks/#public-beacon-testchains)).

- Implement the route `/assets?account=address`, which returns a list of owned
  Token Ids for each of those three contracts in the following JSON structure

  ```
  curl localhost:3333/assets?account=0x448c8e9e1816300Dd052e77D2A44c990A2807D15

  {
    "assets": [
      { token: "0x61B91a780945971b07ba3898A8E0Dc8201dB46b3", ids: [20, 33] },
      { token: "0xd89148d8dEFc7E8942cA4b16DdBE0E2f6485a4c8", ids: [400, 490] },
      { token: "0x9D7f9672060EED641ebc1b22443132eDf4967D91", ids: [86, 98] }
    ]
  }
  ```

  Note: Token ids are represented by the `uint256` type, but assume that the
  token id won't exceed 10000 (thus it is safe to return just integers in the response).

- Repeatedly invoking this endpoint should not incur a gas cost on the network that
  is equal to or greater than the cost as if the client were to just query the
  blockchain network directly.

- The API should correctly track changes in ownership of tokens. If token
  `100` was transferred from `0xAlice` to `0xBob`, then calling the endpoint
  with `0xBob` as the param should show ownership of token `100`.

### Submission

Fork this repository and provide a link to the fork once this task is complete.
The code should compile. Provide instructions if necessary. Testing this
submission will be performed by compiling and executing the binary, then
performing the above curl command

```
go build -o main

./main $

curl localhost:3333/assets?account=0x448c8e9e1816300Dd052e77D2A44c990A2807D15
```

### Notes

- Here are a set of recommended readings to help get started on this task
  - https://goethereumbook.org/en/
  - https://ethereum.org/en/developers/docs/
- This repository is quite bare, though does include an ABI for the ERC721
  interface (`IERC721.abi`). It's highly recommended to use the [abigen](https://geth.ethereum.org/docs/dapp/native-bindings) tool to
  generate go bindings to the contract.
- There is more than one way to implement this API, though we advise
  utilizing event logs https://goethereumbook.org/en/event-read/, particularly
  the ERC721 transfer event to do so.

  ```
  /// @dev This emits when ownership of any NFT changes by any mechanism.
  ///  This event emits when NFTs are created (`from` == 0) and destroyed
  ///  (`to` == 0). Exception: during contract creation, any number of NFTs
  ///  may be created and assigned without emitting Transfer. At the time of
  ///  any transfer, the approved address for that NFT (if any) is reset to none.
  event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
  ```

  For this assessment, ignore the exception `Exception: during contract creation, any number of NFTs may be created and assigned without emitting Transfer`, and safely assume all
  transfers will have a corresponding transfer event.

- Interacting with the Goerli Testnet requires an endpoint to connect to. This is
  provided as a const in `main.go`.

## Assessment Criteria

This task is fairly trivial for a developer that has experience in the relevant
tech stack (golang, go-ethereum, web3, etc), and could be completed within an hour.
This test is mostly assessing your ability (and willingness) to quickly learn
and grasp web3 concepts.
