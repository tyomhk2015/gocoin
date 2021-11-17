# Terms

- <a href="#user-content-blockchain">Blockchain</a>
- <a href="#user-content-append">Append</a>
- <a href="#user-content-decentralization">Decentralization</a>
- <a href="#user-content-previousblockHash">Previous block hash</a>
- <a href="#user-content-blockdata">Block data</a>
- <a href="#user-content-proofofwork">Proof of work</a>
- <a href="#user-content-mining">Mining</a>
- <a href="#user-content-ethereum">Ethereum</a>
- <a href="#user-content-smartcontract">Smart Contract</a>

<hr>

<br id="blockchain">

> Blockchain

Blocks connected in chains, or a chain of blocks.<br>
Many singualar databases joining together in chain.

<br id="blocks">

> Blocks

A database, or data.<br>
A mean of adding data to the blockchain.<br><br>
Adding data is not like adding string data to a variable.<br>
Instead, a block needs to be created with some data in it, and append this block to the desired blockchain.<br><br>
Three parts to look at in a block is <a href="#user-content-blockhash">block hash</a>, <a href="#user-content-previousblockhash">previous block hash</a>, and <a href="#user-content-blockdata">block data</a>.

<br id="append">

> Append

Blockchain is `append-only`, delete or update is not possible.
Adding blocks to chains is the environment of the blockchain.

<br id="decentralization">

> Decentralization

The system where one cannot control the whole database by oneself.<br><br>
People connect to a blockchain will have `same` copy of databases in their computers, nodes. <br>
Therefore, deleting or updating certain data will not be possible, unless more than a half of all computers in the blockchain shutdown.<br><br>
Pretty useful for preventing from manipulated data, such as a house contract, currency, government's censorship etc, you name it.

<br id="blockhash">

> Block hash

Hash: One-way deterministic I/O system.<br><br>
If same input is given, there will always be a same output.<br>
However, with the output, you cannot get the input used in hash system.<br>
Similar to the cryptography, the system of saving passwords of users in a long line of encyptic string.<br><br>
E.g: You can make a steak with raw meat, but you cannot make a raw meat with the steak.

<br id="previousblockhash">

> Previous block hash

Hash code of previous block, an address of previous block.<br>
Required for joining an existing blockchain.<br><br>
The environment of blockchains are appending blocks, previous block do exist definitely in existing blockchains.<br>
By obtaining this previous block hash, and connecting it with your block, you can be a member of the blockchain.<br>
E.g: Becoming a member of 'Club House', the social app.<br><br>
In other word, if the very first block's hash is changed, it will impact the whole blockchain.

<br id="blockdata">

> Block data

The data stored in a block.<br><br>
Crytocurrencies store transaction data.<br>
(Sender, Receiver, Amount, Transaction ID etc)<br><br>
In order to join a blockchain, your block needs to collect all the data related to the blockchain.<br>
(e.g: Bitcoin, history of all Bitcoin transactions.)

<br id="proofofwork">

> Proof of work

<br id="mining">

> Mining

<br id="nonce">

> Nonce

<br id="miningdifficulty">

> Mining Difficulty

<br id="ethereum">

> Ethereum

<br id="smartcontract">

> Smart Contract