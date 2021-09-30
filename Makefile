# Generates a application binary interface (ABI) from a solidity smart contract
.PHONY: gen-abi
gen-abi:
	@solc --abi ./contracts/Store.sol -o build/
	@echo "🦄 generated abi bindings from solidity smart contract"

# Generate Golang bindings from our ABI file.
# This makes it possible to interact with the smart contract with type-safe
# Golang code.
.PHONY: gen-abi-bindings
gen-abi-bindings:
	@abigen --abi ./build/Store.abi \
 			--pkg main \
  			--type Store \
   			--out contract_bindings.go
	@echo "✨ generated go bindings from abi file"


# Generates bindings directly from the solidity file (skipping the abi file)
# Equivalent to gen-abi-bindings
.PHONY: gen-sol-bindings
gen-sol-bindings:
	@abigen --sol contracts/Store.sol \
		--pkg main \
	   	--type Store \
	    --out contract_bindings.go
	@echo "🌀 generated go bindings from solidity smart contract"

# Generates the bytecode/binary file from the smart contract
.PHONY: gen-bytecode
gen-bytecode:
	@solc --bin ./contracts/Store.sol -o build/
	@echo "💖 generated smart contract bytecode from solidity smart contract"


# Generates Golang bindings to deploy our smart contract
# with the smart contract, abi file and binary provided
.PHONY: gen-deploy
gen-deploy: gen-abi gen-abi-bindings gen-bytecode
	@abigen --abi build/Store.abi \
		--pkg main \
	   	--type Store \
	   	--bin build/Store.bin \
	    --out contract_bindings.go
	@echo "🌀 generated go bindings from solidity smart contract"

.PHONY: clean
clean:
	@rm -rf contract_bindings.go build/
	@echo "Cleaned contract builds"
	@ls