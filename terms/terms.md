# Terms

- <a href="#user-content-blockchain">Blockchain</a>
- <a href="#user-content-append">Append</a>
- <a href="#user-content-decentralization">Decentralization</a>
- <a href="#user-content-previousblockhash">Previous block hash</a>
- <a href="#user-content-blockdata">Block data</a>
- <a href="#user-content-proofofwork">Proof of work</a>
- <a href="#user-content-mining">Mining</a>
- <a href="#user-content-nonce">Nonce</a>
- <a href="#user-content-miningdifficulty">Mining Difficulty</a>
- <a href="#user-content-ethereum">Ethereum</a>
- <a href="#user-content-smartcontract">Smart Contract</a>
- <a href="#user-content-nft">NFT</a>
- <a href="#user-content-nft">Wallet</a>
- <a href="#user-content-defi">DeFi</a>
- <a href="#user-content-dex">DEX</a>
- <a href="#user-content-stablecoin">Stable Coin</a>
- <a href="#user-content-cons">Cons of crypto currencies</a>
- <a href="#user-content-1stgeneration">The first generation blockchain</a>
- <a href="#user-content-2ndgeneration">The second generation blockchain</a>
- <a href="#user-content-3rdgeneration">The third generation blockchain</a>
- <a href="#user-content-3rdgeneration">The scalability trilemma</a>
- <a href="#user-content-51attack">51% Attack</a>
- <a href="#user-content-confirmation">Confirmation</a>
- <a href="#user-content-consensus">Consensus</a>
- <a href="#user-content-hashrate">Hash rate</a>
- <a href="#user-content-hardfork">Hard fork</a>
- <a href="#user-content-proofofstake">Proof of stake</a>
- <a href="#user-content-dpos">Delegated proof of stake.</a>

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
Therefore, deleting or updating certain data will not be possible, unless someone takes more than a half of all computering power in the blockchain.<br><br>
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

A system of verifying if a new data (block) is reliable or truthy for joining the chain.
<br>
In other words, consensus algorithm.
<br><br>
`Miner` is related to the 'proof of work'.<br>
Verifying is done by asking some questions to miners. If miners get correct answers, miners are eligible for appending a new block to the chain.<br>
The questions are known to all nodes on the blockchain.<br><br>

`Verifying blocks are relatively easier than answering questions.`

<br id="mining">

> Mining / Miner

The worker that verifies, create block, and appends block to the chain.<br><br>
E.g: A bank teller, who verifies records of people's balance, creates transaction, proceeds loan etc. Similar to interface.<br><br>
If miners successfuly helps transaction, they get commission or transfer fee, like Bitcoin.<br><br>
Miners recieves commission for verifying transaction, and Bitcoin(`coinbase transaction`) for solving hash questions and adding a block to the chain.<br><br>
Miners cannot change any data in a block except `nonce`.

<br id="nonce">

> Nonce

A number used only once.<br><br>

Number of trials for getting correct answer, the `hash`, for the question that miner has gotten. This is during the process of verifying the new block coming into the chain.

<br id="miningdifficulty">

> Mining Difficulty

Indicator of the difficulty of obtaining coins, `bits` are used for this.<br>
It's like having a lot of `if` conditions to be satisfied for getting the coins as reward.<br>
If blocks are added to the chain fast, the difficulty gets harder.<br>

<br id="ethereum">

> Ethereum

One of the types of smart contract.

<br id="smartcontract">

> Smart Contract

A system that people interact with each other by having a automated code or entity as an interface.<br><br>
In developer's perspective, the code you've written gets uploaded a server that is decentralized. This means your code is likely to be verified by the blockchain's security and be used by many people/nodes on the blockchain.

Pros ✔️

* No middle stages are needed.

Practical assumption:<br>
* When buying a piece of land or a building, real estate agency will not be needed because the codes in the blockchain will verify the transaction between the buyer and the seller.
* When changing real currency, you won't need a bank anymore.
* In the court, the cases can be judged without any personal bias.

Cons ❌

* Smart contract only works under their own network.
* Changing inputs in smart contracts are not allowed. However, `Oracle` gets in the middle of the smart contract and allows to change the input.

<br id="nft">

> NFT: Non-Fungible Token

A token that cannot be exchangable with other things.<br>
Think NTF as the only one coin that exists in the world.<br>
Pros ✔️: Guarantees the originality / ownership of certain token, or goods.

<br id="wallet"> 

> Wallet 

Consist of a `public key` (a lock 🔒) and a `private key` (key to open the lock 🔑).<br>
Create with asymmetric encryption.
<br>
(Not symmetric, where things can correspond to each other.)
<br><br>
The `private key` is the most important part of the blockchain transaction.<br>
Careless management of the private key may make you lose all the coins you possess.
<br><br>
To transact or trade some coin or stuff on the blockchain, the sender packs some consensus contract and his/her own public key together. <br>
Then the sender locks the whole thing with receiver's public key.<br>
When the receiver gets the package, he/she opens it with his/her own private key.

<br id="defi">

> DeFi: Decentralized Finanace

Banks without centralized system and real human tellers.<br><br>
No centeralized system: Barely gets account locked or banned.<br>
No real human: No commission fee or no manipulation during transaction.


<br id="dex">

> DEX: Decentralized Exchange

There is no middle man b/ buyer and seller.<br>
Exchange process is done by smart contracts, or code, not by people.


<br id="stablecoin">

> Stable Coin

One crypto currency's value is equal or nearly same as a dollar.<br>
This is can be used for smart contracts.

<br id="cons">

> Cons of crypto currencies

* Too much electrical power is required to gain / transfer the currency from one wallet to the another one, which is not eco-friendly. The power usage, during stage of the `proof of work`, is intense.

* Government has power to prohibit the crypto currency markets, like Binance, Upbit, Coinone etc.

* Anonymous makes the crypto currency used for illegal activities, such as drugs.

* When trying to buy some goods with your crypto currency, you have to reveal all the transaction history and be verified before actually paying with the currency, such as buying a latte in Starbucks.

<br id="1stgeneration">

> The first generation blockchain

Use `proof of work` as security.<br>
E.g: The goal of `Bitcoin` was to create a money that no one could hack.<br>
The `proof of work` requires a lot of computational powers, which causes some ecological problems, such as electrical power.<br>
Does not have programable features, only transactions are possible.

<br id="2ndgeneration">

> The 2nd generation blockchain

Use `smart contracts`. E.g: Ethereum.<br>
The goal of Ethereum is to create world-shared-computer.<br>

1st & 2nd generation blockchain 

<br id="3rdgeneration">

> The 3nd generation blockchain

The blockchain that satisfies the three factors of `the scalability trilemma`, scalability, decentralization, and security.<br>
E.g: Ethereum 2.0

<br id="trilemma">

> The scalability trilemma

<img src="https://i.imgur.com/TAjfVev.png" alt="The scalability trilemma" />

Introduces the three factors that crypto currencies should have.

<br id="51attack">

> 51% Attack

Convincing or taking control of the blockchain network by occupying equal to or more than 51% computing power of the whole blockchain.
<br>
It is hardly possible to accomplish such attack because to aquire 51% of the whole blockchain computing power, you need land, equipments, maintenance etc, which the cost will be astronomical.
<br>
Not only financial problem, but 'confirmation' and 'consensus' makes 51% attack realistically impossible.

<br id="confirmation">

> Confirmation

After a block is added by a winner computer, the loser computers have to verify if the block is legit.

<br id="consensus">

> Consensus

The process of agreeing on something.
<br>
The blockchain is having consensus via confirmation between nodes on the blockchain network.

<br id="hashrate">

> Hash rate

Computing power of computers on the blockchain network.
<br>
The speed of finding a new block, or finding correct hash for solving problems to obtain a new block.

<br id="hardforck">

> Hard fork

A mean of ignoring irregualar nodes. This usually happens when miners disagree. Hard fork is also a mean of updating the blockchain.

<br id="proofofstake">

> Proof of stake.

The amounts of reward you will get from the blockchain network will depend on how many coins you hold, period of holding those coins, and commitment on the network.
<br><br>
Locking your coins to become a miner or a verifyer is mandatory on the blockchain network that uses PoS.
It is kind of accountability or insurance of becoming a member of the blockchain network.
If a member of the network tries to take malicious action, that member will get banned from the network, and the member's coins will be locked forever.
<br><br>
PoS is the key feature that secures the 3rd generation crypto currencies, and cosumes much less electricity than PoW.
<br><br>
Cons: Riches become richer, poors cannot have a chance to become a member of the network. 

<br id="dpos">

> Delegated proof of stake.

Delegate a trustworthy person for validating blocks on the blockchain network by voting with the coins the voters have.
<br><br>
Voters delegate someone to run computers, and when the delegatee gets reward from the network, the reward will be shared with delegators. This is similar to stock option and running a business.
<br><br>
If the delegatee is not doing its job, the voters can take their coins back and vote to other candidates who wants to be a delegatee.
<br><br>
DPoS or PoS does look like democratic system.

