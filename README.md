# golang-blockchain
Simple Blockchain implementation in go


## Usage

### Get Balance
Retrieve the balance for a specific address.
```sh
getbalance -address ADDRESS
```

### Create Blockchain
Create a new blockchain and send a reward to the specified address.
```sh
createblockchain -address ADDRESS
```

### Print Chain
Print all the blocks in the blockchain.
```sh
printchain
```

### Send Coins
Send a specified amount of coins from one address to another.
```sh
send -from FROM -to TO -amount AMOUNT
```

### Create Wallet
Create a new wallet.
```sh
createwallet
```

### List Addresses
List all addresses stored in the wallet file.
```sh
listaddresses
```

### Reindex Unspent Transactions
Rebuild the UTXO set.
```sh
reindexutxo
```
