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
	@mkdir bindings
	@abigen --abi ./build/Store.abi \
 			--pkg bindings \
  			--type Store \
   			--out bindings/bindings.go
	@echo "✨ generated go bindings from abi file"


# Generates bindings directly from the solidity file (skipping the abi file)
# Equivalent to gen-abi-bindings
.PHONY: gen-sol-bindings
gen-sol-bindings:
	@mkdir bindings
	@abigen --sol contracts/Store.sol \
		--pkg bindings \
	   	--type Store \
	    --out bindings/bindings.go
	@echo "🌀 generated go bindings from solidity smart contract"

# Generates the bytecode/binary file from the smart contract
.PHONY: gen-bytecode
gen-bytecode:
	@solc --bin ./contracts/Store.sol -o build/
	@echo "💖 generated smart contract bytecode from solidity smart contract"


# Cleans existing contract builds and generates Golang bindings to deploy our smart contract
# with the smart contract, abi file and binary provided
.PHONY: gen-deploy
gen-deploy: clean gen-abi gen-abi-bindings gen-bytecode
	@abigen --abi build/Store.abi \
		--pkg bindings \
	   	--type Store \
	   	--bin build/Store.bin \
	    --out bindings/bindings.go
	@echo "🌀 generated go bindings from solidity smart contract"

.PHONY: clean
clean:
	@rm -rf bindings/ build/
	@echo "Cleaned contract builds"
	@ls
