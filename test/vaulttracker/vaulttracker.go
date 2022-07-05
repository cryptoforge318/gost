// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulttracker

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VaultTrackerMetaData contains all meta data concerning the VaultTracker contract.
var VaultTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002aa138038062002aa1833981810160405281019062000038919062000683565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508360ff166101008160ff16815250508260e081815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1681525050604051806060016040528060008152602001600081526020016200011b86856200018760201b620016b01760201c565b8152506000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050505050505062000820565b600060016005811115620001a0576200019f620006f5565b5b60ff168360ff160362000227578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200021f919062000724565b90506200059a565b600360058111156200023e576200023d620006f5565b5b60ff168360ff1603620002c5578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000297573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002bd919062000724565b90506200059a565b60046005811115620002dc57620002db620006f5565b5b60ff168360ff1603620004585760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200033a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000360919062000756565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620003c7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003ed919062000756565b6040518263ffffffff1660e01b81526004016200040b919062000799565b602060405180830381865afa15801562000429573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200044f919062000724565b9150506200059a565b6005808111156200046e576200046d620006f5565b5b60ff168360ff16036200050d578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401620004c1919062000803565b602060405180830381865afa158015620004df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000505919062000724565b90506200059a565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b815260040162000553919062000803565b602060405180830381865afa15801562000571573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000597919062000724565b90505b92915050565b600080fd5b600060ff82169050919050565b620005bd81620005a5565b8114620005c957600080fd5b50565b600081519050620005dd81620005b2565b92915050565b6000819050919050565b620005f881620005e3565b81146200060457600080fd5b50565b6000815190506200061881620005ed565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200064b826200061e565b9050919050565b6200065d816200063e565b81146200066957600080fd5b50565b6000815190506200067d8162000652565b92915050565b60008060008060808587031215620006a0576200069f620005a0565b5b6000620006b087828801620005cc565b9450506020620006c38782880162000607565b9350506040620006d6878288016200066c565b9250506060620006e9878288016200066c565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000602082840312156200073d576200073c620005a0565b5b60006200074d8482850162000607565b91505092915050565b6000602082840312156200076f576200076e620005a0565b5b60006200077f848285016200066c565b91505092915050565b62000793816200063e565b82525050565b6000602082019050620007b0600083018462000788565b92915050565b6000819050919050565b6000819050919050565b6000620007eb620007e5620007df84620007b6565b620007c0565b620005e3565b9050919050565b620007fd81620007ca565b82525050565b60006020820190506200081a6000830184620007f2565b92915050565b60805160a05160c05160e051610100516121be620008e360003960008181610573015281816109a001528181610c9d01528181610f6f01528181611029015261142e01526000610b310152600081816103280152818161138a015261157f015260008181610594015281816109c101528181610cbe0152818161104a0152818161144f015261166a0152600081816103540152818161089801528181610b5701528181610ecf01528181610f950152818161128b015261168e01526121be6000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806364ae3c9d1161008c578063a622ee7c11610066578063a622ee7c14610288578063b326258d146102ba578063b7dd3483146102ea578063f851a44014610308576100ea565b806364ae3c9d1461020a5780638ce744261461023a578063a01cfffb14610258576100ea565b806319caf46c116100c857806319caf46c1461015b578063204f83f91461018b578063613a28d1146101a95780636392a51f146101d9576100ea565b8063012b264a146100ef57806311554c431461010d578063177946731461012b575b600080fd5b6100f7610326565b6040516101049190611adb565b60405180910390f35b61011561034a565b6040516101229190611b0f565b60405180910390f35b61014560048036038101906101409190611b87565b610350565b6040516101529190611bf5565b60405180910390f35b61017560048036038101906101709190611c10565b610894565b6040516101829190611b0f565b60405180910390f35b610193610b2f565b6040516101a09190611b0f565b60405180910390f35b6101c360048036038101906101be9190611c3d565b610b53565b6040516101d09190611bf5565b60405180910390f35b6101f360048036038101906101ee9190611c10565b610e49565b604051610201929190611c7d565b60405180910390f35b610224600480360381019061021f9190611ca6565b610ecb565b6040516102319190611bf5565b60405180910390f35b610242610f6d565b60405161024f9190611cef565b60405180910390f35b610272600480360381019061026d9190611c3d565b610f91565b60405161027f9190611bf5565b60405180910390f35b6102a2600480360381019061029d9190611c10565b61125d565b6040516102b193929190611d0a565b60405180910390f35b6102d460048036038101906102cf9190611c3d565b611287565b6040516102e19190611bf5565b60405180910390f35b6102f2611668565b6040516102ff9190611adb565b60405180910390f35b61031061168c565b60405161031d9190611adb565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611d9e565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361044f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044690611e0a565b60405180910390fd5b60008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050848260000151101561056c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056390611e76565b60405180910390fd5b60006105b87f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006116b0565b9050600080600154111561060b576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006001546105f09190611ec5565b6105fa9190611f4e565b6106049190611f7f565b905061064a565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000846106339190611ec5565b61063d9190611f4e565b6106479190611f7f565b90505b60006a52b7d2dcc80cd2e40000008560000151836106689190611ec5565b6106729190611f4e565b905080856020018181516106869190611fb3565b91508181525050878560000181815161069f9190611f7f565b9150818152505082856040018181525050846000808c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060008460000151111561080e576000600154111561076d576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006001546107529190611ec5565b61075c9190611f4e565b6107669190611f7f565b91506107ac565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000856107959190611ec5565b61079f9190611f4e565b6107a99190611f7f565b91505b60006a52b7d2dcc80cd2e40000008560000151846107ca9190611ec5565b6107d49190611f4e565b905080856020018181516107e89190611fb3565b9150818152505088856000018181516108019190611fb3565b9150818152505050610819565b878460000181815250505b82846040018181525050836000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600196505050505050509392505050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610925576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091c90611d9e565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008160200151905060006109e57f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006116b0565b90506000806001541115610a38576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000600154610a1d9190611ec5565b610a279190611f4e565b610a319190611f7f565b9050610a77565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e400000084610a609190611ec5565b610a6a9190611f4e565b610a749190611f7f565b90505b60006a52b7d2dcc80cd2e4000000856000015183610a959190611ec5565b610a9f9190611f4e565b9050828560400181815250506000856020018181525050846000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050508084610b229190611fb3565b9650505050505050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdb90611d9e565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508381600001511015610c96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8d90612055565b60405180910390fd5b6000610ce27f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006116b0565b90506000806001541115610d35576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000600154610d1a9190611ec5565b610d249190611f4e565b610d2e9190611f7f565b9050610d74565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e400000084610d5d9190611ec5565b610d679190611f4e565b610d719190611f7f565b90505b60006a52b7d2dcc80cd2e4000000846000015183610d929190611ec5565b610d9c9190611f4e565b90508084602001818151610db09190611fb3565b915081815250508684600001818151610dc99190611f7f565b9150818152505082846040018181525050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b60008060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015181602001519250925050915091565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5390611d9e565b60405180910390fd5b826001819055506001915050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611022576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101990611d9e565b60405180910390fd5b600061106e7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006116b0565b905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816000015111156111db576000806001541115611139576a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e400000060015461111e9190611ec5565b6111289190611f4e565b6111329190611f7f565b9050611178565b6a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e4000000856111619190611ec5565b61116b9190611f4e565b6111759190611f7f565b90505b60006a52b7d2dcc80cd2e40000008360000151836111969190611ec5565b6111a09190611f4e565b905080836020018181516111b49190611fb3565b9150818152505086836000018181516111cd9190611fb3565b9150818152505050506111e6565b848160000181815250505b81816040018181525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506001935050505092915050565b60006020528060005260406000206000915090508060000154908060010154908060020154905083565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611318576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130f90611d9e565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905084826000018181516114209190611f7f565b9150818152505060006114737f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006116b0565b905060008183604001511461156057600060015411156114d2576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e40000006001546114b79190611ec5565b6114c19190611f4e565b6114cb9190611f7f565b9050611511565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000846114fa9190611ec5565b6115049190611f4e565b61150e9190611f7f565b90505b60006a52b7d2dcc80cd2e400000084600001518361152f9190611ec5565b6115399190611f4e565b9050808460200181815161154d9190611fb3565b9150818152505082846040018181525050505b86836000018181516115729190611fb3565b91508181525050826000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000600160058111156116c6576116c5612075565b5b60ff168360ff1603611748578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561171d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061174191906120b9565b9050611a94565b6003600581111561175c5761175b612075565b5b60ff168360ff16036117de578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d791906120b9565b9050611a94565b600460058111156117f2576117f1612075565b5b60ff168360ff16036119615760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561184e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061187291906120fb565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118fc91906120fb565b6040518263ffffffff1660e01b81526004016119189190611adb565b602060405180830381865afa158015611935573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061195991906120b9565b915050611a94565b60058081111561197457611973612075565b5b60ff168360ff1603611a0c578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b81526004016119c4919061216d565b602060405180830381865afa1580156119e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a0591906120b9565b9050611a94565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401611a50919061216d565b602060405180830381865afa158015611a6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9191906120b9565b90505b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611ac582611a9a565b9050919050565b611ad581611aba565b82525050565b6000602082019050611af06000830184611acc565b92915050565b6000819050919050565b611b0981611af6565b82525050565b6000602082019050611b246000830184611b00565b92915050565b600080fd5b611b3881611aba565b8114611b4357600080fd5b50565b600081359050611b5581611b2f565b92915050565b611b6481611af6565b8114611b6f57600080fd5b50565b600081359050611b8181611b5b565b92915050565b600080600060608486031215611ba057611b9f611b2a565b5b6000611bae86828701611b46565b9350506020611bbf86828701611b46565b9250506040611bd086828701611b72565b9150509250925092565b60008115159050919050565b611bef81611bda565b82525050565b6000602082019050611c0a6000830184611be6565b92915050565b600060208284031215611c2657611c25611b2a565b5b6000611c3484828501611b46565b91505092915050565b60008060408385031215611c5457611c53611b2a565b5b6000611c6285828601611b46565b9250506020611c7385828601611b72565b9150509250929050565b6000604082019050611c926000830185611b00565b611c9f6020830184611b00565b9392505050565b600060208284031215611cbc57611cbb611b2a565b5b6000611cca84828501611b72565b91505092915050565b600060ff82169050919050565b611ce981611cd3565b82525050565b6000602082019050611d046000830184611ce0565b92915050565b6000606082019050611d1f6000830186611b00565b611d2c6020830185611b00565b611d396040830184611b00565b949350505050565b600082825260208201905092915050565b7f73656e646572206d75737420626520617574686f72697a656400000000000000600082015250565b6000611d88601983611d41565b9150611d9382611d52565b602082019050919050565b60006020820190508181036000830152611db781611d7b565b9050919050565b7f63616e6e6f74207472616e73666572206e6f74696f6e616c20746f2073656c66600082015250565b6000611df4602083611d41565b9150611dff82611dbe565b602082019050919050565b60006020820190508181036000830152611e2381611de7565b9050919050565b7f616d6f756e74206578636565647320617661696c61626c652062616c616e6365600082015250565b6000611e60602083611d41565b9150611e6b82611e2a565b602082019050919050565b60006020820190508181036000830152611e8f81611e53565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611ed082611af6565b9150611edb83611af6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611f1457611f13611e96565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611f5982611af6565b9150611f6483611af6565b925082611f7457611f73611f1f565b5b828204905092915050565b6000611f8a82611af6565b9150611f9583611af6565b925082821015611fa857611fa7611e96565b5b828203905092915050565b6000611fbe82611af6565b9150611fc983611af6565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611ffe57611ffd611e96565b5b828201905092915050565b7f616d6f756e742065786365656473207661756c742062616c616e636500000000600082015250565b600061203f601c83611d41565b915061204a82612009565b602082019050919050565b6000602082019050818103600083015261206e81612032565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000815190506120b381611b5b565b92915050565b6000602082840312156120cf576120ce611b2a565b5b60006120dd848285016120a4565b91505092915050565b6000815190506120f581611b2f565b92915050565b60006020828403121561211157612110611b2a565b5b600061211f848285016120e6565b91505092915050565b6000819050919050565b6000819050919050565b600061215761215261214d84612128565b612132565b611af6565b9050919050565b6121678161213c565b82525050565b6000602082019050612182600083018461215e565b9291505056fea26469706673582212201f49fa4e020092363b794e54de3cc327f7a1edc4c802a5b5f19f96a4f4095d5d64736f6c634300080d0033",
}

// VaultTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultTrackerMetaData.ABI instead.
var VaultTrackerABI = VaultTrackerMetaData.ABI

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultTrackerMetaData.Bin instead.
var VaultTrackerBin = VaultTrackerMetaData.Bin

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, p uint8, m *big.Int, c common.Address, s common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := VaultTrackerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultTrackerBin), backend, p, m, c, s)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// VaultTracker is an auto generated Go binding around an Ethereum contract.
type VaultTracker struct {
	VaultTrackerCaller     // Read-only binding to the contract
	VaultTrackerTransactor // Write-only binding to the contract
	VaultTrackerFilterer   // Log filterer for contract events
}

// VaultTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultTrackerSession struct {
	Contract     *VaultTracker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultTrackerCallerSession struct {
	Contract *VaultTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTrackerTransactorSession struct {
	Contract     *VaultTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultTrackerRaw struct {
	Contract *VaultTracker // Generic contract binding to access the raw methods on
}

// VaultTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultTrackerCallerRaw struct {
	Contract *VaultTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTrackerTransactorRaw struct {
	Contract *VaultTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultTracker creates a new instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTracker(address common.Address, backend bind.ContractBackend) (*VaultTracker, error) {
	contract, err := bindVaultTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// NewVaultTrackerCaller creates a new read-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerCaller(address common.Address, caller bind.ContractCaller) (*VaultTrackerCaller, error) {
	contract, err := bindVaultTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerCaller{contract: contract}, nil
}

// NewVaultTrackerTransactor creates a new write-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTrackerTransactor, error) {
	contract, err := bindVaultTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerTransactor{contract: contract}, nil
}

// NewVaultTrackerFilterer creates a new log filterer instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultTrackerFilterer, error) {
	contract, err := bindVaultTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerFilterer{contract: contract}, nil
}

// bindVaultTracker binds a generic wrapper to an already deployed contract.
func bindVaultTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.VaultTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerSession) Admin() (common.Address, error) {
	return _VaultTracker.Contract.Admin(&_VaultTracker.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) Admin() (common.Address, error) {
	return _VaultTracker.Contract.Admin(&_VaultTracker.CallOpts)
}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCaller) BalancesOf(opts *bind.CallOpts, o common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "balancesOf", o)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerSession) BalancesOf(o common.Address) (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.BalancesOf(&_VaultTracker.CallOpts, o)
}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCallerSession) BalancesOf(o common.Address) (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.BalancesOf(&_VaultTracker.CallOpts, o)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCaller) CTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "cTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) MaturityRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "maturityRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerSession) MaturityRate() (*big.Int, error) {
	return _VaultTracker.Contract.MaturityRate(&_VaultTracker.CallOpts)
}

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) MaturityRate() (*big.Int, error) {
	return _VaultTracker.Contract.MaturityRate(&_VaultTracker.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerCaller) Protocol(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerSession) Protocol() (uint8, error) {
	return _VaultTracker.Contract.Protocol(&_VaultTracker.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerCallerSession) Protocol() (uint8, error) {
	return _VaultTracker.Contract.Protocol(&_VaultTracker.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCaller) Swivel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "swivel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerCaller) Vaults(opts *bind.CallOpts, arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Notional     *big.Int
		Redeemable   *big.Int
		ExchangeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Notional = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Redeemable = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExchangeRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerSession) Vaults(arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	return _VaultTracker.Contract.Vaults(&_VaultTracker.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerCallerSession) Vaults(arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	return _VaultTracker.Contract.Vaults(&_VaultTracker.CallOpts, arg0)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) AddNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "addNotional", o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) MatureVault(opts *bind.TransactOpts, c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "matureVault", c)
}

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerSession) MatureVault(c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts, c)
}

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) MatureVault(c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts, c)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactor) RedeemInterest(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "redeemInterest", o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactorSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) RemoveNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "removeNotional", o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFee(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFee", f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFrom", f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}
