.PHONY: compile_solidity_mock_erc compile_go_mock_erc compile_mock_erc
.PHONY: compile_solidity_mock_yield compile_go_mock_yield compile_mock_yield
.PHONY: compile_solidity_mock_pendle compile_go_mock_pendle compile_mock_pendle
.PHONY: compile_solidity_mock_pendle_token compile_go_mock_pendle_token compile_mock_pendle_token
.PHONY: compile_solidity_mock_tempus compile_go_mock_tempus compile_mock_tempus
.PHONY: compile_solidity_mock_sense compile_go_mock_sense compile_mock_sense
.PHONY: compile_solidity_mock_apwine_token compile_go_mock_apwine_token compile_mock_apwine_token
.PHONY: compile_solidity_mock_apwine compile_go_mock_apwine compile_mock_apwine
.PHONY: compile_solidity_mock_sense_token compile_go_mock_sense_token compile_mock_sense_token
.PHONY: compile_solidity_mock_element_token compile_go_mock_element_token compile_mock_element_token
.PHONY: compile_solidity_mock_illuminate compile_go_mock_illuminate compile_mock_illuminate
.PHONY: compile_mocks

# TODO under? api-specific sol?

# .PHONY: compile_solidity_zct compile_go_zct compile_zct
# .PHONY: compile_tokens

.PHONY: compile_solidity_illuminate_test compile_go_illuminate_test compile_illuminate_test
.PHONY: compile_solidity_lender_test compile_go_lender_test compile_lender_test
.PHONY: compile_solidity_redeemer_test compile_go_redeemer_test compile_redeemer_test
.PHONY: compile_test

.PHONY: clean_test_abi clean_test_bin clean_test_go
.PHONY: clean_test

# .PHONY: clean_build_sol clean_build_abi clean_build_bin clean_build_go
# .PHONY: clean_build

# .PHONY: copy_to_build

.PHONY: all

# Mocks
compile_solidity_mock_erc:
	@echo "compiling Mock ERC20 solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Erc20.sol

compile_go_mock_erc:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Erc20.abi --bin ./test/mocks/Erc20.bin -pkg mocks -type Erc20 -out ./test/mocks/erc20.go 

compile_mock_erc: compile_solidity_mock_erc compile_go_mock_erc

compile_solidity_mock_yield:
	@echo "compiling Mock Yield solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Yield.sol

compile_go_mock_yield:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Yield.abi --bin ./test/mocks/Yield.bin -pkg mocks -type Yield -out ./test/mocks/yield.go 

compile_mock_yield: compile_solidity_mock_yield compile_go_mock_yield

compile_solidity_mock_pendle_token:
	@echo "compiling Mock Pendle Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/PendleToken.sol

compile_go_mock_pendle_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/PendleToken.abi --bin ./test/mocks/PendleToken.bin -pkg mocks -type PendleToken -out ./test/mocks/pendletoken.go 

compile_mock_pendle_token: compile_solidity_mock_pendle_token compile_go_mock_pendle_token

compile_solidity_mock_pendle:
	@echo "compiling Mock Pendle solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Pendle.sol

compile_go_mock_pendle:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Pendle.abi --bin ./test/mocks/Pendle.bin -pkg mocks -type Pendle -out ./test/mocks/pendle.go 

compile_mock_pendle: compile_solidity_mock_pendle compile_go_mock_pendle

compile_solidity_mock_element_token:
	@echo "compiling Mock ElementToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/ElementToken.sol

compile_go_mock_element_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/ElementToken.abi --bin ./test/mocks/ElementToken.bin -pkg mocks -type ElementToken -out ./test/mocks/elementtoken.go 

compile_mock_element_token: compile_solidity_mock_element_token compile_go_mock_element_token

compile_solidity_mock_element:
	@echo "compiling Mock Element solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Element.sol

compile_go_mock_element:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Element.abi --bin ./test/mocks/Element.bin -pkg mocks -type Element -out ./test/mocks/element.go 

compile_mock_element: compile_solidity_mock_element compile_go_mock_element

compile_solidity_mock_swivel:
	@echo "compiling Mock Swivel solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Swivel.sol

compile_go_mock_swivel:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Swivel.abi --bin ./test/mocks/Swivel.bin -pkg mocks -type Swivel -out ./test/mocks/swivel.go 

compile_mock_swivel: compile_solidity_mock_swivel compile_go_mock_swivel

compile_solidity_mock_illuminate:
	@echo "compiling Mock Illuminate solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Illuminate.sol

compile_go_mock_illuminate:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Illuminate.abi --bin ./test/mocks/Illuminate.bin -pkg mocks -type Illuminate -out ./test/mocks/illuminate.go 

compile_mock_illuminate: compile_solidity_mock_illuminate compile_go_mock_illuminate

compile_solidity_mock_zc_token:
	@echo "compiling Mock ZcToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/ZcToken.sol

compile_go_mock_zc_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/ZcToken.abi --bin ./test/mocks/ZcToken.bin -pkg mocks -type ZcToken -out ./test/mocks/zctoken.go 

compile_mock_zc_token: compile_solidity_mock_zc_token compile_go_mock_zc_token

compile_solidity_mock_tempus:
	@echo "compiling Mock Tempus solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Tempus.sol

compile_go_mock_tempus:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Tempus.abi --bin ./test/mocks/Tempus.bin -pkg mocks -type Tempus -out ./test/mocks/tempus.go 

compile_mock_tempus: compile_solidity_mock_tempus compile_go_mock_tempus

compile_solidity_mock_sense:
	@echo "compiling Mock Sense solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/Sense.sol

compile_go_mock_sense:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/Sense.abi --bin ./test/mocks/Sense.bin -pkg mocks -type Sense -out ./test/mocks/sense.go 

compile_mock_sense: compile_solidity_mock_sense compile_go_mock_sense

compile_solidity_mock_sense_token:
	@echo "compiling Mock SenseToken solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/SenseToken.sol

compile_go_mock_sense_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/SenseToken.abi --bin ./test/mocks/SenseToken.bin -pkg mocks -type SenseToken -out ./test/mocks/sensetoken.go 

compile_mock_sense_token: compile_solidity_mock_sense_token compile_go_mock_sense_token

compile_solidity_mock_apwine_token:
	@echo "compiling Mock APWine Token solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/APWineToken.sol

compile_go_mock_apwine_token:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/APWineToken.abi --bin ./test/mocks/APWineToken.bin -pkg mocks -type APWineToken -out ./test/mocks/apwinetoken.go 

compile_mock_apwine_token: compile_solidity_mock_apwine_token compile_go_mock_apwine_token

compile_solidity_mock_apwine:
	@echo "compiling Mock APWine solidity source into abi and bin files"
	solc -o ./test/mocks --abi --bin --overwrite ./test/mocks/APWine.sol

compile_go_mock_apwine:
	@echo "compiling abi and bin files to golang"
	abigen --abi ./test/mocks/APWine.abi --bin ./test/mocks/APWine.bin -pkg mocks -type APWine -out ./test/mocks/apwine.go 

compile_mock_apwine: compile_solidity_mock_apwine compile_go_mock_apwine

compile_mocks: compile_mock_erc compile_mock_yield compile_mock_pendle compile_mock_pendle_token compile_mock_element_token compile_mock_element compile_mock_illuminate compile_mock_zc_token compile_mock_swivel compile_mock_tempus compile_mock_sense compile_mock_sense_token compile_go_mock_apwine_token compile_mock_apwine compile_mock_apwine_token


# Real Tokens
# compile_solidity_zct:
# 	@echo "compiling ZCT solidity source into abi and bin files"
# 	solc -o ./test/tokens --abi --bin --overwrite ./test/tokens/ZcToken.sol

# compile_go_zct:
# 	@echo "compiling abi and bin files to golang"
# 	abigen --abi ./test/tokens/ZcToken.abi --bin ./test/tokens/ZcToken.bin -pkg tokens -type ZcToken -out ./test/tokens/zctoken.go 

# compile_zct: compile_solidity_zct compile_go_zct

# compile_tokens: compile_zct

# Fakes

# Contracts
compile_solidity_illuminate_test:
	@echo "compiling Illuminate solidity source into abi and bin files"
	solc -o ./test/illuminate --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/illuminate/Illuminate.sol

compile_go_illuminate_test:
	@echo "compiling Illuminate abi and bin files to golang"
	abigen --abi ./test/illuminate/Illuminate.abi --bin ./test/illuminate/Illuminate.bin -pkg illuminate -type Illuminate -out ./test/illuminate/illuminate.go

compile_illuminate_test: compile_solidity_illuminate_test compile_go_illuminate_test

compile_solidity_lender_test:
	@echo "compiling Lender solidity source into abi and bin files"
	solc -o ./test/lender --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/lender/Lender.sol

compile_go_lender_test:
	@echo "compiling Lender abi and bin files to golang"
	abigen --abi ./test/lender/Lender.abi --bin ./test/lender/Lender.bin -pkg lender -type Lender -out ./test/lender/lender.go

compile_lender_test: compile_solidity_lender_test compile_go_lender_test

compile_solidity_redeemer_test:
	@echo "compiling Redeemer solidity source into abi and bin files"
	solc -o ./test/redeemer --optimize --optimize-runs=15000 --abi --bin --overwrite ./test/redeemer/Redeemer.sol

compile_go_redeemer_test:
	@echo "compiling Redeemer abi and bin files to golang"
	abigen --abi ./test/redeemer/Redeemer.abi --bin ./test/redeemer/Redeemer.bin -pkg redeemer -type Redeemer -out ./test/redeemer/redeemer.go

compile_redeemer_test: compile_solidity_redeemer_test compile_go_redeemer_test
compile_test: compile_illuminate_test compile_lender_test compile_redeemer_test

# Cleaning
clean_test_abi:
	@echo "removing abi files from test/ dirs"
	rm test/**/*.abi

clean_test_bin:
	@echo "removing bin files from test/ dirs"
	rm test/**/*.bin

clean_test_go:
	@echo "removing generated go bindings from test/ dirs"
	rm test/**/*.go

clean_test: clean_test_abi clean_test_bin clean_test_go

# clean_build_sol:
# 	@echo "removing sol files from build/ dirs"
# 	rm build/**/*.sol
# 
# clean_build_abi:
# 	@echo "removing abi files from build/ dirs"
# 	rm build/**/*.abi
# 
# clean_build_bin:
# 	@echo "removing bin files from build/ dirs"
# 	rm build/**/*.bin
# 
# clean_build_go:
# 	@echo "removing go files from build/ dirs"
# 	rm build/**/*.go
# 
# clean_build: clean_build_sol clean_build_abi clean_build_bin clean_build_go

all: clean_test compile_test
