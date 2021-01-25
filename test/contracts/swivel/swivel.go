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
	Floating   bool
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Expiry     *big.Int
	Nonce      *big.Int
}

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Initiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Release\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CTOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"agreements\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"released\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"release\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"fillFixed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"fillFloating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x60806040523480156200001157600080fd5b50604051620028a4380380620028a483398181016040528101906200003791906200024b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620000aa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a190620002e3565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161462000127578162000129565b305b9050620001ae6040518060400160405280600e81526020017f53776976656c2046696e616e63650000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e300000000000000000000000000000000000000000000000000000008152508684620001be60201b620013461760201c565b6001819055505050505062000388565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b6000815190506200022e8162000354565b92915050565b60008151905062000245816200036e565b92915050565b6000806000606084860312156200026157600080fd5b6000620002718682870162000234565b935050602062000284868287016200021d565b925050604062000297868287016200021d565b9150509250925092565b6000620002b0601f8362000305565b91507f636f6d706f756e6420746f6b656e2061646472657373207265717569726564006000830152602082019050919050565b60006020820190508181036000830152620002fe81620002a1565b9050919050565b600082825260208201905092915050565b600062000323826200032a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200035f8162000316565b81146200036b57600080fd5b50565b62000379816200034a565b81146200038557600080fd5b50565b61250c80620003986000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063592ab2fb11610066578063592ab2fb1461013457806379d5e86714610164578063a3f4df7e14610194578063be3b52ec146101b2578063ffa1ad74146101eb57610093565b8063288cdc91146100985780632ac12622146100c857806343295411146100f857806352a9674b14610116575b600080fd5b6100b260048036038101906100ad919061197a565b610209565b6040516100bf91906121a0565b60405180910390f35b6100e260048036038101906100dd919061197a565b610221565b6040516100ef9190611fc3565b60405180910390f35b610100610241565b60405161010d9190611eac565b60405180910390f35b61011e610265565b60405161012b9190611fde565b60405180910390f35b61014e600480360381019061014991906119df565b61026b565b60405161015b9190611fc3565b60405180910390f35b61017e600480360381019061017991906119df565b610a32565b60405161018b9190611fc3565b60405180910390f35b61019c6111f9565b6040516101a9919061203e565b60405180910390f35b6101cc60048036038101906101c791906119a3565b611232565b6040516101e29a99989796959493929190611ec7565b60405180910390f35b6101f361130d565b604051610200919061203e565b60405180910390f35b60036020528060005260406000206000915090505481565b60026020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6000848260001515600260008460000135815260200190815260200160002060009054906101000a900460ff161515146102da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d1906120a0565b60405180910390fd5b428260e001351015610321576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031890612060565b60405180910390fd5b61033e610338600154610333856113a5565b611459565b8261149a565b73ffffffffffffffffffffffffffffffffffffffff1682602001602081019061036791906118ff565b73ffffffffffffffffffffffffffffffffffffffff16146103bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b490612080565b60405180910390fd5b6003600088600001358152602001908152602001600020548760a001356103e491906122b8565b861115610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041d906120c0565b60405180910390fd5b61042e61179c565b670de0b6b3a76400008860800135670de0b6b3a764000089610450919061225e565b61045a919061222d565b8960a00135610469919061225e565b610473919061222d565b8160e0018181525050868160c0018181525050600088604001602081019061049b91906118ff565b90508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8a60200160208101906104cc91906118ff565b308560e001516040518463ffffffff1660e01b81526004016104f093929190611f63565b602060405180830381600087803b15801561050a57600080fd5b505af115801561051e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105429190611951565b610581576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057890612120565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd33308560c001516040518463ffffffff1660e01b81526004016105c293929190611f63565b602060405180830381600087803b1580156105dc57600080fd5b505af11580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106149190611951565b610653576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064a90612100565b60405180910390fd5b60006106848a604001602081019061066b91906118ff565b8b608001358c60a0013561067f91906121d7565b6115ef565b116106c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb90612140565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508960200160208101906106fd91906118ff565b836000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505033836020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505089604001602081019061077f91906118ff565b836040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060018360600190151590811515815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561080f57600080fd5b505af1158015610823573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108479190611a46565b8360a00181815250508960c00135836101000181815250504283610100015161087091906121d7565b8361012001818152505082600460008c60000135815260200190815260200160002060008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a0820151816003015560c0820151816004015560e0820151816005015561010082015181600601556101208201518160070155905050878a600001357fa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb60405160405180910390a3600195505050505050949350505050565b6000848260001515600260008460000135815260200190815260200160002060009054906101000a900460ff16151514610aa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a98906120a0565b60405180910390fd5b428260e001351015610ae8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adf90612060565b60405180910390fd5b610b05610aff600154610afa856113a5565b611459565b8261149a565b73ffffffffffffffffffffffffffffffffffffffff16826020016020810190610b2e91906118ff565b73ffffffffffffffffffffffffffffffffffffffff1614610b84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7b90612080565b60405180910390fd5b6003600088600001358152602001908152602001600020548760a00135610bab91906122b8565b861115610bed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be4906120c0565b60405180910390fd5b610bf561179c565b670de0b6b3a76400008860a00135670de0b6b3a764000089610c17919061225e565b610c21919061222d565b8960800135610c30919061225e565b610c3a919061222d565b8160c0018181525050868160e00181815250506000886040016020810190610c6291906118ff565b90508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8a6020016020810190610c9391906118ff565b308560c001516040518463ffffffff1660e01b8152600401610cb793929190611f63565b602060405180830381600087803b158015610cd157600080fd5b505af1158015610ce5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d099190611951565b610d48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3f90612120565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd33308560e001516040518463ffffffff1660e01b8152600401610d8993929190611f63565b602060405180830381600087803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ddb9190611951565b610e1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1190612100565b60405180910390fd5b6000610e4b8a6040016020810190610e3291906118ff565b8b60a001358c60800135610e4691906121d7565b6115ef565b11610e8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8290612140565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050896020016020810190610ec491906118ff565b836000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505033836020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050896040016020810190610f4691906118ff565b836040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060008360600190151590811515815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610fd657600080fd5b505af1158015610fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100e9190611a46565b8360a00181815250508960c00135836101000181815250504283610100015161103791906121d7565b8361012001818152505082600460008c60000135815260200190815260200160002060008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a0820151816003015560c0820151816004015560e0820151816005015561010082015181600601556101208201518160070155905050878a600001357fa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb60405160405180910390a3600195505050505050949350505050565b6040518060400160405280600e81526020017f53776976656c2046696e616e636500000000000000000000000000000000000081525081565b6004602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060020160159054906101000a900460ff1690806003015490806004015490806005015490806006015490806007015490508a565b6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007ff01cc16c244dd394ae954a5b2cd48a4f17007f995776783399d1190f45ada62360001b82600001358360200160208101906113e391906118ff565b8460400160208101906113f691906118ff565b8560600160208101906114099190611928565b86608001358760a001358860c001358960e001358a610100013560405160200161143c9a99989796959493929190611df8565b604051602081830303815290604052805190602001209050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0826040013560001c1115611506576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114fd90612180565b60405180910390fd5b601b82600001602081019061151b9190611a6f565b60ff16148061153f5750601c82600001602081019061153a9190611a6f565b60ff16145b61157e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611575906120e0565b60405180910390fd5b6001838360000160208101906115949190611a6f565b84602001358560400135604051600081526020016040526040516115bb9493929190611ff9565b6020604051602081039080840390855afa1580156115dd573d6000803e3d6000fd5b50505060206040510351905092915050565b6000808390508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401611650929190611f9a565b602060405180830381600087803b15801561166a57600080fd5b505af115801561167e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a29190611951565b6116e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d890612160565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a0712d68856040518263ffffffff1660e01b815260040161174091906121a0565b602060405180830381600087803b15801561175a57600080fd5b505af115801561176e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117929190611a46565b9250505092915050565b604051806101400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160001515815260200160001515815260200160008152602001600081526020016000815260200160008152602001600081525090565b60008135905061184481612463565b92915050565b6000813590506118598161247a565b92915050565b60008151905061186e8161247a565b92915050565b60008135905061188381612491565b92915050565b60006060828403121561189b57600080fd5b81905092915050565b600061012082840312156118b757600080fd5b81905092915050565b6000813590506118cf816124a8565b92915050565b6000815190506118e4816124a8565b92915050565b6000813590506118f9816124bf565b92915050565b60006020828403121561191157600080fd5b600061191f84828501611835565b91505092915050565b60006020828403121561193a57600080fd5b60006119488482850161184a565b91505092915050565b60006020828403121561196357600080fd5b60006119718482850161185f565b91505092915050565b60006020828403121561198c57600080fd5b600061199a84828501611874565b91505092915050565b600080604083850312156119b657600080fd5b60006119c485828601611874565b92505060206119d585828601611874565b9150509250929050565b6000806000806101c085870312156119f657600080fd5b6000611a04878288016118a4565b945050610120611a16878288016118c0565b935050610140611a2887828801611874565b925050610160611a3a87828801611889565b91505092959194509250565b600060208284031215611a5857600080fd5b6000611a66848285016118d5565b91505092915050565b600060208284031215611a8157600080fd5b6000611a8f848285016118ea565b91505092915050565b611aa1816122ec565b82525050565b611ab8611ab3826122ec565b61237e565b82525050565b611ac7816122fe565b82525050565b611ade611ad9826122fe565b612390565b82525050565b611aed8161230a565b82525050565b611b04611aff8261230a565b6123a2565b82525050565b6000611b15826121bb565b611b1f81856121c6565b9350611b2f81856020860161234b565b611b3881612438565b840191505092915050565b6000611b506011836121c6565b91507f6f726465722068617320657870697265640000000000000000000000000000006000830152602082019050919050565b6000611b906011836121c6565b91507f696e76616c6964207369676e61747572650000000000000000000000000000006000830152602082019050919050565b6000611bd06018836121c6565b91507f6f7264657220686173206265656e2063616e63656c6c656400000000000000006000830152602082019050919050565b6000611c10601f836121c6565b91507f74616b657220616d6f756e74203e20617661696c61626c6520766f6c756d65006000830152602082019050919050565b6000611c50601b836121c6565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b6000611c90601a836121c6565b91507f7472616e736665722066726f6d2074616b6572206661696c65640000000000006000830152602082019050919050565b6000611cd0601a836121c6565b91507f7472616e736665722066726f6d206d616b6572206661696c65640000000000006000830152602082019050919050565b6000611d106015836121c6565b91507f43546f6b656e206d696e74696e67206661696c656400000000000000000000006000830152602082019050919050565b6000611d50601a836121c6565b91507f756e6465726c79696e6720617070726f76616c206661696c65640000000000006000830152602082019050919050565b6000611d90601b836121c6565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b611dcc81612334565b82525050565b611de3611dde82612334565b6123be565b82525050565b611df28161233e565b82525050565b6000611e04828d611af3565b602082019150611e14828c611af3565b602082019150611e24828b611aa7565b601482019150611e34828a611aa7565b601482019150611e448289611acd565b600182019150611e548288611dd2565b602082019150611e648287611dd2565b602082019150611e748286611dd2565b602082019150611e848285611dd2565b602082019150611e948284611dd2565b6020820191508190509b9a5050505050505050505050565b6000602082019050611ec16000830184611a98565b92915050565b600061014082019050611edd600083018d611a98565b611eea602083018c611a98565b611ef7604083018b611a98565b611f04606083018a611abe565b611f116080830189611abe565b611f1e60a0830188611dc3565b611f2b60c0830187611dc3565b611f3860e0830186611dc3565b611f46610100830185611dc3565b611f54610120830184611dc3565b9b9a5050505050505050505050565b6000606082019050611f786000830186611a98565b611f856020830185611a98565b611f926040830184611dc3565b949350505050565b6000604082019050611faf6000830185611a98565b611fbc6020830184611dc3565b9392505050565b6000602082019050611fd86000830184611abe565b92915050565b6000602082019050611ff36000830184611ae4565b92915050565b600060808201905061200e6000830187611ae4565b61201b6020830186611de9565b6120286040830185611ae4565b6120356060830184611ae4565b95945050505050565b600060208201905081810360008301526120588184611b0a565b905092915050565b6000602082019050818103600083015261207981611b43565b9050919050565b6000602082019050818103600083015261209981611b83565b9050919050565b600060208201905081810360008301526120b981611bc3565b9050919050565b600060208201905081810360008301526120d981611c03565b9050919050565b600060208201905081810360008301526120f981611c43565b9050919050565b6000602082019050818103600083015261211981611c83565b9050919050565b6000602082019050818103600083015261213981611cc3565b9050919050565b6000602082019050818103600083015261215981611d03565b9050919050565b6000602082019050818103600083015261217981611d43565b9050919050565b6000602082019050818103600083015261219981611d83565b9050919050565b60006020820190506121b56000830184611dc3565b92915050565b600081519050919050565b600082825260208201905092915050565b60006121e282612334565b91506121ed83612334565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612222576122216123da565b5b828201905092915050565b600061223882612334565b915061224383612334565b92508261225357612252612409565b5b828204905092915050565b600061226982612334565b915061227483612334565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156122ad576122ac6123da565b5b828202905092915050565b60006122c382612334565b91506122ce83612334565b9250828210156122e1576122e06123da565b5b828203905092915050565b60006122f782612314565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561236957808201518184015260208101905061234e565b83811115612378576000848401525b50505050565b6000612389826123ac565b9050919050565b600061239b826123c8565b9050919050565b6000819050919050565b60006123b782612456565b9050919050565b6000819050919050565b60006123d382612449565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b61246c816122ec565b811461247757600080fd5b50565b612483816122fe565b811461248e57600080fd5b50565b61249a8161230a565b81146124a557600080fd5b50565b6124b181612334565b81146124bc57600080fd5b50565b6124c88161233e565b81146124d357600080fd5b5056fea2646970667358221220161f710756d8a492ba4aabc5efec3374675f196c628bdcc47afd30ef5205180364736f6c63430008000033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend, i *big.Int, c common.Address, v common.Address) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend, i, c, v)
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

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelCaller) CTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "CTOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelSession) CTOKEN() (common.Address, error) {
	return _Swivel.Contract.CTOKEN(&_Swivel.CallOpts)
}

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelCallerSession) CTOKEN() (common.Address, error) {
	return _Swivel.Contract.CTOKEN(&_Swivel.CallOpts)
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

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCaller) Agreements(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "agreements", arg0, arg1)

	outstruct := new(struct {
		Maker      common.Address
		Taker      common.Address
		Underlying common.Address
		Floating   bool
		Released   bool
		Rate       *big.Int
		Principal  *big.Int
		Interest   *big.Int
		Duration   *big.Int
		Release    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maker = out[0].(common.Address)
	outstruct.Taker = out[1].(common.Address)
	outstruct.Underlying = out[2].(common.Address)
	outstruct.Floating = out[3].(bool)
	outstruct.Released = out[4].(bool)
	outstruct.Rate = out[5].(*big.Int)
	outstruct.Principal = out[6].(*big.Int)
	outstruct.Interest = out[7].(*big.Int)
	outstruct.Duration = out[8].(*big.Int)
	outstruct.Release = out[9].(*big.Int)

	return *outstruct, err

}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	return _Swivel.Contract.Agreements(&_Swivel.CallOpts, arg0, arg1)
}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCallerSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	return _Swivel.Contract.Agreements(&_Swivel.CallOpts, arg0, arg1)
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

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactor) FillFixed(opts *bind.TransactOpts, o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "fillFixed", o, a, k, c)
}

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelSession) FillFixed(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, a, k, c)
}

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactorSession) FillFixed(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, a, k, c)
}

// FillFloating is a paid mutator transaction binding the contract method 0x592ab2fb.
//
// Solidity: function fillFloating((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactor) FillFloating(opts *bind.TransactOpts, o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "fillFloating", o, a, k, c)
}

// FillFloating is a paid mutator transaction binding the contract method 0x592ab2fb.
//
// Solidity: function fillFloating((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelSession) FillFloating(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFloating(&_Swivel.TransactOpts, o, a, k, c)
}

// FillFloating is a paid mutator transaction binding the contract method 0x592ab2fb.
//
// Solidity: function fillFloating((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactorSession) FillFloating(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFloating(&_Swivel.TransactOpts, o, a, k, c)
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

// SwivelInitiateIterator is returned from FilterInitiate and is used to iterate over the raw logs and unpacked data for Initiate events raised by the Swivel contract.
type SwivelInitiateIterator struct {
	Event *SwivelInitiate // Event containing the contract specifics and raw log

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
func (it *SwivelInitiateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelInitiate)
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
		it.Event = new(SwivelInitiate)
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
func (it *SwivelInitiateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelInitiateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelInitiate represents a Initiate event raised by the Swivel contract.
type SwivelInitiate struct {
	OrderKey     [32]byte
	AgreementKey [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiate is a free log retrieval operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) FilterInitiate(opts *bind.FilterOpts, orderKey [][32]byte, agreementKey [][32]byte) (*SwivelInitiateIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Initiate", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelInitiateIterator{contract: _Swivel.contract, event: "Initiate", logs: logs, sub: sub}, nil
}

// WatchInitiate is a free log subscription operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) WatchInitiate(opts *bind.WatchOpts, sink chan<- *SwivelInitiate, orderKey [][32]byte, agreementKey [][32]byte) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Initiate", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelInitiate)
				if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
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

// ParseInitiate is a log parse operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) ParseInitiate(log types.Log) (*SwivelInitiate, error) {
	event := new(SwivelInitiate)
	if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelReleaseIterator is returned from FilterRelease and is used to iterate over the raw logs and unpacked data for Release events raised by the Swivel contract.
type SwivelReleaseIterator struct {
	Event *SwivelRelease // Event containing the contract specifics and raw log

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
func (it *SwivelReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelRelease)
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
		it.Event = new(SwivelRelease)
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
func (it *SwivelReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelRelease represents a Release event raised by the Swivel contract.
type SwivelRelease struct {
	OrderKey     [32]byte
	AgreementKey [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelease is a free log retrieval operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) FilterRelease(opts *bind.FilterOpts, orderKey [][32]byte, agreementKey [][32]byte) (*SwivelReleaseIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Release", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelReleaseIterator{contract: _Swivel.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *SwivelRelease, orderKey [][32]byte, agreementKey [][32]byte) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Release", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelRelease)
				if err := _Swivel.contract.UnpackLog(event, "Release", log); err != nil {
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

// ParseRelease is a log parse operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) ParseRelease(log types.Log) (*SwivelRelease, error) {
	event := new(SwivelRelease)
	if err := _Swivel.contract.UnpackLog(event, "Release", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
