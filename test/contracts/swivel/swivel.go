// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swivel

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HashOrder is an auto generated low-level Go binding around an user-defined struct.
type HashOrder struct {
	Key        [32]byte
	Maker      common.Address
	Underlying common.Address
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"name\":\"exit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"initiateZcTokenFillingVaultInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506100d16040518060400160405280600e81526020017f53776976656c2046696e616e63650000000000000000000000000000000000008152506040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525046306100dc60201b610eba1760201c565b60018190555061013b565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b61241b8061014a6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063aba2870111610066578063aba287011461016f578063b3ba20751461019f578063d2144f58146101cf578063f851a440146101ff578063ffa1ad741461021d5761009e565b8063288cdc91146100a35780632ac12622146100d35780633e1608b41461010357806352a9674b14610133578063a3f4df7e14610151575b600080fd5b6100bd60048036038101906100b89190611d86565b61023b565b6040516100ca9190612220565b60405180910390f35b6100ed60048036038101906100e89190611d86565b610253565b6040516100fa9190612047565b60405180910390f35b61011d60048036038101906101189190611daf565b610273565b60405161012a9190612047565b60405180910390f35b61013b61037c565b6040516101489190612062565b60405180910390f35b610159610382565b604051610166919061215e565b60405180910390f35b61018960048036038101906101849190611cb9565b6103bb565b6040516101969190612047565b60405180910390f35b6101b960048036038101906101b49190611ded565b61085c565b6040516101c69190612047565b60405180910390f35b6101e960048036038101906101e49190611cb9565b6109bc565b6040516101f69190612047565b60405180910390f35b610207610e5d565b604051610214919061202c565b60405180910390f35b610225610e81565b604051610232919061215e565b60405180910390f35b60036020528060005260406000206000915090505481565b60026020528060005260406000206000915054906101000a900460ff1681565b600061029261028c60015461028786610f19565b610fdb565b8361101c565b73ffffffffffffffffffffffffffffffffffffffff168360200160208101906102bb9190611c90565b73ffffffffffffffffffffffffffffffffffffffff1614610311576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030890612180565b60405180910390fd5b6001600260008560000135815260200190815260200160002060006101000a81548160ff02191690831515021790555082600001357fe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a60405160405180910390a26001905092915050565b60015481565b6040518060400160405280600e81526020017f53776976656c2046696e616e636500000000000000000000000000000000000081525081565b600080600090505b8787905081101561084d576000151588888381811061040b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050610120020160800160208101906104249190611d5d565b151514156106355760001515888883818110610469577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050610120020160600160208101906104829190611d5d565b1515141561055f576105518888838181106104c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201878784818110610506577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610546577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060600201611171565b61055a57600080fd5b610630565b61062688888381811061059b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905061012002018787848181106105db577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013586868581811061061b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050606002016112d1565b61062f57600080fd5b5b61083a565b60001515888883818110610672577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201606001602081019061068b9190611d5d565b151514156107685761075a8888838181106106cf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050610120020187878481811061070f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050602002013586868581811061074f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060600201611431565b61076357600080fd5b610839565b61082f8888838181106107a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905061012002018787848181106107e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610824577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060600201611591565b61083857600080fd5b5b5b8080610845906122e9565b9150506103c3565b50600190509695505050505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff161515146108cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c2906121c0565b60405180910390fd5b428261010001351015610913576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090a906121a0565b60405180910390fd5b61093061092a60015461092585610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff168260200160208101906109599190611c90565b73ffffffffffffffffffffffffffffffffffffffff16146109af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a690612180565b60405180910390fd5b6001925050509392505050565b600080600090505b87879050811015610e4e5760001515888883818110610a0c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905061012002016080016020810190610a259190611d5d565b15151415610c365760001515888883818110610a6a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905061012002016060016020810190610a839190611d5d565b15151415610b6057610b52888883818110610ac7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201878784818110610b07577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610b47577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050606002016116f1565b610b5b57600080fd5b610c31565b610c27888883818110610b9c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201878784818110610bdc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610c1c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506060020161085c565b610c3057600080fd5b5b610e3b565b60001515888883818110610c73577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905061012002016060016020810190610c8c9190611d5d565b15151415610d6957610d5b888883818110610cd0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201878784818110610d10577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610d50577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060600201611851565b610d6457600080fd5b610e3a565b610e30888883818110610da5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506101200201878784818110610de5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90506020020135868685818110610e25577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9050606002016119b1565b610e3957600080fd5b5b5b8080610e46906122e9565b9150506109c4565b50600190509695505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007f7ddd38ab5ed1c16b61ca90eeb9579e29da1ba821cf42d8cdef8f30a31a6a414660001b8260000135836020016020810190610f579190611c90565b846040016020810190610f6a9190611c90565b856060016020810190610f7d9190611d5d565b866080016020810190610f909190611d5d565b8760a001358860c001358960e001358a6101000135604051602001610fbe9a9998979695949392919061207d565b604051602081830303815290604052805190602001209050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0826040013560001c1115611088576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107f90612200565b60405180910390fd5b601b82600001602081019061109d9190611e3f565b60ff1614806110c15750601c8260000160208101906110bc9190611e3f565b60ff16145b611100576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f7906121e0565b60405180910390fd5b6001838360000160208101906111169190611e3f565b846020013585604001356040516000815260200160405260405161113d9493929190612119565b6020604051602081039080840390855afa15801561115f573d6000803e3d6000fd5b50505060206040510351905092915050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff161515146111e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d7906121c0565b60405180910390fd5b428261010001351015611228576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121f906121a0565b60405180910390fd5b61124561123f60015461123a85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff1682602001602081019061126e9190611c90565b73ffffffffffffffffffffffffffffffffffffffff16146112c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112bb90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff16151514611340576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611337906121c0565b60405180910390fd5b428261010001351015611388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137f906121a0565b60405180910390fd5b6113a561139f60015461139a85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff168260200160208101906113ce9190611c90565b73ffffffffffffffffffffffffffffffffffffffff1614611424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141b90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff161515146114a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611497906121c0565b60405180910390fd5b4282610100013510156114e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114df906121a0565b60405180910390fd5b6115056114ff6001546114fa85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff1682602001602081019061152e9190611c90565b73ffffffffffffffffffffffffffffffffffffffff1614611584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157b90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff16151514611600576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f7906121c0565b60405180910390fd5b428261010001351015611648576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163f906121a0565b60405180910390fd5b61166561165f60015461165a85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff1682602001602081019061168e9190611c90565b73ffffffffffffffffffffffffffffffffffffffff16146116e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116db90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff16151514611760576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611757906121c0565b60405180910390fd5b4282610100013510156117a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179f906121a0565b60405180910390fd5b6117c56117bf6001546117ba85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff168260200160208101906117ee9190611c90565b73ffffffffffffffffffffffffffffffffffffffff1614611844576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183b90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff161515146118c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118b7906121c0565b60405180910390fd5b428261010001351015611908576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ff906121a0565b60405180910390fd5b61192561191f60015461191a85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff1682602001602081019061194e9190611c90565b73ffffffffffffffffffffffffffffffffffffffff16146119a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161199b90612180565b60405180910390fd5b6001925050509392505050565b6000838260001515600260008460000135815260200190815260200160002060009054906101000a900460ff16151514611a20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a17906121c0565b60405180910390fd5b428261010001351015611a68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a5f906121a0565b60405180910390fd5b611a85611a7f600154611a7a85610f19565b610fdb565b8261101c565b73ffffffffffffffffffffffffffffffffffffffff16826020016020810190611aae9190611c90565b73ffffffffffffffffffffffffffffffffffffffff1614611b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611afb90612180565b60405180910390fd5b6001925050509392505050565b600081359050611b2081612372565b92915050565b60008083601f840112611b3857600080fd5b8235905067ffffffffffffffff811115611b5157600080fd5b602083019150836060820283011115611b6957600080fd5b9250929050565b60008083601f840112611b8257600080fd5b8235905067ffffffffffffffff811115611b9b57600080fd5b60208301915083610120820283011115611bb457600080fd5b9250929050565b60008083601f840112611bcd57600080fd5b8235905067ffffffffffffffff811115611be657600080fd5b602083019150836020820283011115611bfe57600080fd5b9250929050565b600081359050611c1481612389565b92915050565b600081359050611c29816123a0565b92915050565b600060608284031215611c4157600080fd5b81905092915050565b60006101208284031215611c5d57600080fd5b81905092915050565b600081359050611c75816123b7565b92915050565b600081359050611c8a816123ce565b92915050565b600060208284031215611ca257600080fd5b6000611cb084828501611b11565b91505092915050565b60008060008060008060608789031215611cd257600080fd5b600087013567ffffffffffffffff811115611cec57600080fd5b611cf889828a01611b70565b9650965050602087013567ffffffffffffffff811115611d1757600080fd5b611d2389828a01611bbb565b9450945050604087013567ffffffffffffffff811115611d4257600080fd5b611d4e89828a01611b26565b92509250509295509295509295565b600060208284031215611d6f57600080fd5b6000611d7d84828501611c05565b91505092915050565b600060208284031215611d9857600080fd5b6000611da684828501611c1a565b91505092915050565b6000806101808385031215611dc357600080fd5b6000611dd185828601611c4a565b925050610120611de385828601611c2f565b9150509250929050565b60008060006101a08486031215611e0357600080fd5b6000611e1186828701611c4a565b935050610120611e2386828701611c66565b925050610140611e3586828701611c2f565b9150509250925092565b600060208284031215611e5157600080fd5b6000611e5f84828501611c7b565b91505092915050565b611e7181612257565b82525050565b611e8081612269565b82525050565b611e8f81612275565b82525050565b6000611ea08261223b565b611eaa8185612246565b9350611eba8185602086016122b6565b611ec381612361565b840191505092915050565b6000611edb601183612246565b91507f696e76616c6964207369676e61747572650000000000000000000000000000006000830152602082019050919050565b6000611f1b600d83612246565b91507f6f726465722065787069726564000000000000000000000000000000000000006000830152602082019050919050565b6000611f5b600f83612246565b91507f6f726465722063616e63656c6c656400000000000000000000000000000000006000830152602082019050919050565b6000611f9b601b83612246565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b6000611fdb601b83612246565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b6120178161229f565b82525050565b612026816122a9565b82525050565b60006020820190506120416000830184611e68565b92915050565b600060208201905061205c6000830184611e77565b92915050565b60006020820190506120776000830184611e86565b92915050565b600061014082019050612093600083018d611e86565b6120a0602083018c611e86565b6120ad604083018b611e68565b6120ba606083018a611e68565b6120c76080830189611e77565b6120d460a0830188611e77565b6120e160c083018761200e565b6120ee60e083018661200e565b6120fc61010083018561200e565b61210a61012083018461200e565b9b9a5050505050505050505050565b600060808201905061212e6000830187611e86565b61213b602083018661201d565b6121486040830185611e86565b6121556060830184611e86565b95945050505050565b600060208201905081810360008301526121788184611e95565b905092915050565b6000602082019050818103600083015261219981611ece565b9050919050565b600060208201905081810360008301526121b981611f0e565b9050919050565b600060208201905081810360008301526121d981611f4e565b9050919050565b600060208201905081810360008301526121f981611f8e565b9050919050565b6000602082019050818103600083015261221981611fce565b9050919050565b6000602082019050612235600083018461200e565b92915050565b600081519050919050565b600082825260208201905092915050565b60006122628261227f565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156122d45780820151818401526020810190506122b9565b838111156122e3576000848401525b50505050565b60006122f48261229f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561232757612326612332565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b61237b81612257565b811461238657600080fd5b50565b61239281612269565b811461239d57600080fd5b50565b6123a981612275565b81146123b457600080fd5b50565b6123c08161229f565b81146123cb57600080fd5b50565b6123d7816122a9565b81146123e257600080fd5b5056fea2646970667358221220f738c3958723a111005632394fc54847323218048d6c048c790f8c170e17c4ca64736f6c63430008000033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// Swivel is an auto generated Go binding around an Ethereum contract.
type Swivel struct {
	SwivelCaller     // Read-only binding to the contract
	SwivelTransactor // Write-only binding to the contract
	SwivelFilterer   // Log filterer for contract events
}

// SwivelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelSession struct {
	Contract     *Swivel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelCallerSession struct {
	Contract *SwivelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwivelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTransactorSession struct {
	Contract     *SwivelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelRaw struct {
	Contract *Swivel // Generic contract binding to access the raw methods on
}

// SwivelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelCallerRaw struct {
	Contract *SwivelCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTransactorRaw struct {
	Contract *SwivelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivel creates a new instance of Swivel, bound to a specific deployed contract.
func NewSwivel(address common.Address, backend bind.ContractBackend) (*Swivel, error) {
	contract, err := bindSwivel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// NewSwivelCaller creates a new read-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelCaller(address common.Address, caller bind.ContractCaller) (*SwivelCaller, error) {
	contract, err := bindSwivel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelCaller{contract: contract}, nil
}

// NewSwivelTransactor creates a new write-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTransactor, error) {
	contract, err := bindSwivel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTransactor{contract: contract}, nil
}

// NewSwivelFilterer creates a new log filterer instance of Swivel, bound to a specific deployed contract.
func NewSwivelFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelFilterer, error) {
	contract, err := bindSwivel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelFilterer{contract: contract}, nil
}

// bindSwivel binds a generic wrapper to an already deployed contract.
func bindSwivel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.SwivelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transact(opts, method, params...)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelCaller) DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelSession) DOMAIN() ([32]byte, error) {
	return _Swivel.Contract.DOMAIN(&_Swivel.CallOpts)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelCallerSession) DOMAIN() ([32]byte, error) {
	return _Swivel.Contract.DOMAIN(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCallerSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCallerSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelSession) Admin() (common.Address, error) {
	return _Swivel.Contract.Admin(&_Swivel.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Swivel *SwivelCallerSession) Admin() (common.Address, error) {
	return _Swivel.Contract.Admin(&_Swivel.CallOpts)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "cancelled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "filled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// Cancel is a paid mutator transaction binding the contract method 0x3e1608b4.
//
// Solidity: function cancel((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactor) Cancel(opts *bind.TransactOpts, o HashOrder, c SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "cancel", o, c)
}

// Cancel is a paid mutator transaction binding the contract method 0x3e1608b4.
//
// Solidity: function cancel((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelSession) Cancel(o HashOrder, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Cancel(&_Swivel.TransactOpts, o, c)
}

// Cancel is a paid mutator transaction binding the contract method 0x3e1608b4.
//
// Solidity: function cancel((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactorSession) Cancel(o HashOrder, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Cancel(&_Swivel.TransactOpts, o, c)
}

// Exit is a paid mutator transaction binding the contract method 0xaba28701.
//
// Solidity: function exit((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactor) Exit(opts *bind.TransactOpts, o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "exit", o, a, c)
}

// Exit is a paid mutator transaction binding the contract method 0xaba28701.
//
// Solidity: function exit((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelSession) Exit(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Exit(&_Swivel.TransactOpts, o, a, c)
}

// Exit is a paid mutator transaction binding the contract method 0xaba28701.
//
// Solidity: function exit((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactorSession) Exit(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Exit(&_Swivel.TransactOpts, o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactor) Initiate(opts *bind.TransactOpts, o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiate", o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelSession) Initiate(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, c)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] c) returns(bool)
func (_Swivel *SwivelTransactorSession) Initiate(o []HashOrder, a []*big.Int, c []SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, c)
}

// InitiateZcTokenFillingVaultInitiate is a paid mutator transaction binding the contract method 0xb3ba2075.
//
// Solidity: function initiateZcTokenFillingVaultInitiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, uint256 a, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactor) InitiateZcTokenFillingVaultInitiate(opts *bind.TransactOpts, o HashOrder, a *big.Int, c SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiateZcTokenFillingVaultInitiate", o, a, c)
}

// InitiateZcTokenFillingVaultInitiate is a paid mutator transaction binding the contract method 0xb3ba2075.
//
// Solidity: function initiateZcTokenFillingVaultInitiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, uint256 a, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelSession) InitiateZcTokenFillingVaultInitiate(o HashOrder, a *big.Int, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateZcTokenFillingVaultInitiate(&_Swivel.TransactOpts, o, a, c)
}

// InitiateZcTokenFillingVaultInitiate is a paid mutator transaction binding the contract method 0xb3ba2075.
//
// Solidity: function initiateZcTokenFillingVaultInitiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o, uint256 a, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactorSession) InitiateZcTokenFillingVaultInitiate(o HashOrder, a *big.Int, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateZcTokenFillingVaultInitiate(&_Swivel.TransactOpts, o, a, c)
}

// SwivelCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the Swivel contract.
type SwivelCancelIterator struct {
	Event *SwivelCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwivelCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwivelCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwivelCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelCancel represents a Cancel event raised by the Swivel contract.
type SwivelCancel struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) FilterCancel(opts *bind.FilterOpts, key [][32]byte) (*SwivelCancelIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelCancelIterator{contract: _Swivel.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *SwivelCancel, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelCancel)
				if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancel is a log parse operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) ParseCancel(log types.Log) (*SwivelCancel, error) {
	event := new(SwivelCancel)
	if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
