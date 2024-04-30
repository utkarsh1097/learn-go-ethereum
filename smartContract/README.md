# Create Go contract file

We can generate the ABI from a solidity source file.

`solc --abi Store.sol -o build`

It'll write it to a file called ./build/Store.abi

Now let's convert the ABI to a Go file that we can import. This new file will contain all the available methods the we can use to interact with the smart contract from our Go application.

`abigen --abi=./build/Store.abi --pkg=store --out=Store.go`

In order to deploy a smart contract from Go, we also need to compile the solidity smart contract to EVM bytecode. The EVM bytecode is what will be sent in the data field of the transaction. The bin file is required for generating the deploy methods on the Go contract file.

`solc --bin Store.sol -o build`

Now we compile the Go contract file which will include the deploy methods because we includes the bin file.

`abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go`



