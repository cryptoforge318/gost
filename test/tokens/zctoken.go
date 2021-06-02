// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// ZcTokenABI is the input ABI used to generate the binding from.
const ZcTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
var ZcTokenBin = "0x60806040523480156200001157600080fd5b506040516200204d3803806200204d833981810160405281019062000037919062000250565b818181600390805190602001906200005192919062000100565b5080600490805190602001906200006a92919062000100565b50505083600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260078190555033600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050620004d0565b8280546200010e90620003c1565b90600052602060002090601f0160209004810192826200013257600085556200017e565b82601f106200014d57805160ff19168380011785556200017e565b828001600101855582156200017e579182015b828111156200017d57825182559160200191906001019062000160565b5b5090506200018d919062000191565b5090565b5b80821115620001ac57600081600090555060010162000192565b5090565b6000620001c7620001c18462000317565b620002ee565b905082815260208101848484011115620001e057600080fd5b620001ed8482856200038b565b509392505050565b60008151905062000206816200049c565b92915050565b600082601f8301126200021e57600080fd5b815162000230848260208601620001b0565b91505092915050565b6000815190506200024a81620004b6565b92915050565b600080600080608085870312156200026757600080fd5b60006200027787828801620001f5565b94505060206200028a8782880162000239565b935050604085015167ffffffffffffffff811115620002a857600080fd5b620002b6878288016200020c565b925050606085015167ffffffffffffffff811115620002d457600080fd5b620002e2878288016200020c565b91505092959194509250565b6000620002fa6200030d565b9050620003088282620003f7565b919050565b6000604051905090565b600067ffffffffffffffff8211156200033557620003346200045c565b5b62000340826200048b565b9050602081019050919050565b60006200035a8262000361565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015620003ab5780820151818401526020810190506200038e565b83811115620003bb576000848401525b50505050565b60006002820490506001821680620003da57607f821691505b60208210811415620003f157620003f06200042d565b5b50919050565b62000402826200048b565b810181811067ffffffffffffffff821117156200042457620004236200045c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b620004a7816200034d565b8114620004b357600080fd5b50565b620004c18162000381565b8114620004cd57600080fd5b50565b611b6d80620004e06000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636f307dc311610097578063a457c2d711610066578063a457c2d7146102d9578063a9059cbb14610309578063dd62ed3e14610339578063f851a4401461036957610100565b80636f307dc31461023d57806370a082311461025b57806395d89b411461028b5780639dc29fac146102a957610100565b806323b872dd116100d357806323b872dd1461018f578063313ce567146101bf57806339509351146101dd57806340c10f191461020d57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd14610153578063204f83f914610171575b600080fd5b61010d610387565b60405161011a9190611469565b60405180910390f35b61013d60048036038101906101389190611201565b610419565b60405161014a919061144e565b60405180910390f35b61015b610430565b60405161016891906115eb565b60405180910390f35b61017961043a565b60405161018691906115eb565b60405180910390f35b6101a960048036038101906101a491906111b2565b610440565b6040516101b6919061144e565b60405180910390f35b6101c7610533565b6040516101d49190611606565b60405180910390f35b6101f760048036038101906101f29190611201565b61053c565b604051610204919061144e565b60405180910390f35b61022760048036038101906102229190611201565b6105da565b604051610234919061144e565b60405180910390f35b610245610682565b6040516102529190611433565b60405180910390f35b6102756004803603810190610270919061114d565b6106a8565b60405161028291906115eb565b60405180910390f35b6102936106f0565b6040516102a09190611469565b60405180910390f35b6102c360048036038101906102be9190611201565b610782565b6040516102d0919061144e565b60405180910390f35b6102f360048036038101906102ee9190611201565b61082a565b604051610300919061144e565b60405180910390f35b610323600480360381019061031e9190611201565b610910565b604051610330919061144e565b60405180910390f35b610353600480360381019061034e9190611176565b610927565b60405161036091906115eb565b60405180910390f35b6103716109ae565b60405161037e9190611433565b60405180910390f35b6060600380546103969061174f565b80601f01602080910402602001604051908101604052809291908181526020018280546103c29061174f565b801561040f5780601f106103e45761010080835404028352916020019161040f565b820191906000526020600020905b8154815290600101906020018083116103f257829003601f168201915b5050505050905090565b60006104263384846109d4565b6001905092915050565b6000600254905090565b60075481565b600061044d848484610b9f565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610511576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105089061152b565b60405180910390fd5b610527853385846105229190611693565b6109d4565b60019150509392505050565b60006012905090565b60006105d0338484600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105cb919061163d565b6109d4565b6001905092915050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610664906114cb565b60405180910390fd5b6106778484610e13565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546106ff9061174f565b80601f016020809104026020016040519081016040528092919081815260200182805461072b9061174f565b80156107785780601f1061074d57610100808354040283529160200191610778565b820191906000526020600020905b81548152906001019060200180831161075b57829003601f168201915b5050505050905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080c906114cb565b60405180910390fd5b61081f8484610f5b565b600191505092915050565b600080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156108ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e6906115ab565b60405180910390fd5b610905338585846109009190611693565b6109d4565b600191505092915050565b600061091d338484610b9f565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3b9061158b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ab4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aab906114eb565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610b9291906115eb565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c069061156b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c769061148b565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610d05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cfc9061150b565b60405180910390fd5b8181610d119190611693565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610da1919061163d565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e0591906115eb565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7a906115cb565b60405180910390fd5b8060026000828254610e95919061163d565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610eea919061163d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f4f91906115eb565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610fcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc29061154b565b60405180910390fd5b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611051576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611048906114ab565b60405180910390fd5b818161105d9190611693565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546110b19190611693565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161111691906115eb565b60405180910390a3505050565b60008135905061113281611b09565b92915050565b60008135905061114781611b20565b92915050565b60006020828403121561115f57600080fd5b600061116d84828501611123565b91505092915050565b6000806040838503121561118957600080fd5b600061119785828601611123565b92505060206111a885828601611123565b9150509250929050565b6000806000606084860312156111c757600080fd5b60006111d586828701611123565b93505060206111e686828701611123565b92505060406111f786828701611138565b9150509250925092565b6000806040838503121561121457600080fd5b600061122285828601611123565b925050602061123385828601611138565b9150509250929050565b611246816116c7565b82525050565b611255816116d9565b82525050565b600061126682611621565b611270818561162c565b935061128081856020860161171c565b611289816117df565b840191505092915050565b60006112a160238361162c565b91506112ac826117f0565b604082019050919050565b60006112c460228361162c565b91506112cf8261183f565b604082019050919050565b60006112e760148361162c565b91506112f28261188e565b602082019050919050565b600061130a60228361162c565b9150611315826118b7565b604082019050919050565b600061132d60268361162c565b915061133882611906565b604082019050919050565b600061135060288361162c565b915061135b82611955565b604082019050919050565b600061137360218361162c565b915061137e826119a4565b604082019050919050565b600061139660258361162c565b91506113a1826119f3565b604082019050919050565b60006113b960248361162c565b91506113c482611a42565b604082019050919050565b60006113dc60258361162c565b91506113e782611a91565b604082019050919050565b60006113ff601f8361162c565b915061140a82611ae0565b602082019050919050565b61141e81611705565b82525050565b61142d8161170f565b82525050565b6000602082019050611448600083018461123d565b92915050565b6000602082019050611463600083018461124c565b92915050565b60006020820190508181036000830152611483818461125b565b905092915050565b600060208201905081810360008301526114a481611294565b9050919050565b600060208201905081810360008301526114c4816112b7565b9050919050565b600060208201905081810360008301526114e4816112da565b9050919050565b60006020820190508181036000830152611504816112fd565b9050919050565b6000602082019050818103600083015261152481611320565b9050919050565b6000602082019050818103600083015261154481611343565b9050919050565b6000602082019050818103600083015261156481611366565b9050919050565b6000602082019050818103600083015261158481611389565b9050919050565b600060208201905081810360008301526115a4816113ac565b9050919050565b600060208201905081810360008301526115c4816113cf565b9050919050565b600060208201905081810360008301526115e4816113f2565b9050919050565b60006020820190506116006000830184611415565b92915050565b600060208201905061161b6000830184611424565b92915050565b600081519050919050565b600082825260208201905092915050565b600061164882611705565b915061165383611705565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561168857611687611781565b5b828201905092915050565b600061169e82611705565b91506116a983611705565b9250828210156116bc576116bb611781565b5b828203905092915050565b60006116d2826116e5565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561173a57808201518184015260208101905061171f565b83811115611749576000848401525b50505050565b6000600282049050600182168061176757607f821691505b6020821081141561177b5761177a6117b0565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f73656e646572206d7573742062652061646d696e000000000000000000000000600082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b611b12816116c7565b8114611b1d57600080fd5b50565b611b2981611705565b8114611b3457600080fd5b5056fea2646970667358221220a8dd7e555aa7ce435fa3ee6b7a348904d862461b54a0a9f7709fbebfec7c43c764736f6c63430008040033"

// DeployZcToken deploys a new Ethereum contract, binding an instance of ZcToken to it.
func DeployZcToken(auth *bind.TransactOpts, backend bind.ContractBackend, u common.Address, m *big.Int, n string, s string) (common.Address, *types.Transaction, *ZcToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZcTokenBin), backend, u, m, n, s)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// ZcToken is an auto generated Go binding around an Ethereum contract.
type ZcToken struct {
	ZcTokenCaller     // Read-only binding to the contract
	ZcTokenTransactor // Write-only binding to the contract
	ZcTokenFilterer   // Log filterer for contract events
}

// ZcTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZcTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZcTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZcTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZcTokenSession struct {
	Contract     *ZcToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZcTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZcTokenCallerSession struct {
	Contract *ZcTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZcTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZcTokenTransactorSession struct {
	Contract     *ZcTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZcTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZcTokenRaw struct {
	Contract *ZcToken // Generic contract binding to access the raw methods on
}

// ZcTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZcTokenCallerRaw struct {
	Contract *ZcTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZcTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZcTokenTransactorRaw struct {
	Contract *ZcTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZcToken creates a new instance of ZcToken, bound to a specific deployed contract.
func NewZcToken(address common.Address, backend bind.ContractBackend) (*ZcToken, error) {
	contract, err := bindZcToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// NewZcTokenCaller creates a new read-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenCaller(address common.Address, caller bind.ContractCaller) (*ZcTokenCaller, error) {
	contract, err := bindZcToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenCaller{contract: contract}, nil
}

// NewZcTokenTransactor creates a new write-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZcTokenTransactor, error) {
	contract, err := bindZcToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransactor{contract: contract}, nil
}

// NewZcTokenFilterer creates a new log filterer instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZcTokenFilterer, error) {
	contract, err := bindZcToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZcTokenFilterer{contract: contract}, nil
}

// bindZcToken binds a generic wrapper to an already deployed contract.
func bindZcToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.ZcTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenSession) Admin() (common.Address, error) {
	return _ZcToken.Contract.Admin(&_ZcToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Admin() (common.Address, error) {
	return _ZcToken.Contract.Admin(&_ZcToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCallerSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenSession) TotalSupply() (*big.Int, error) {
	return _ZcToken.Contract.TotalSupply(&_ZcToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ZcToken.Contract.TotalSupply(&_ZcToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Burn(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "burn", f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Mint(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mint", t, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Mint(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, t, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Mint(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, sender, recipient, amount)
}

// ZcTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZcToken contract.
type ZcTokenApprovalIterator struct {
	Event *ZcTokenApproval // Event containing the contract specifics and raw log

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
func (it *ZcTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZcTokenApproval)
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
		it.Event = new(ZcTokenApproval)
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
func (it *ZcTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZcTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZcTokenApproval represents a Approval event raised by the ZcToken contract.
type ZcTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZcTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZcToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZcTokenApprovalIterator{contract: _ZcToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZcTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZcToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZcTokenApproval)
				if err := _ZcToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) ParseApproval(log types.Log) (*ZcTokenApproval, error) {
	event := new(ZcTokenApproval)
	if err := _ZcToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZcTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZcToken contract.
type ZcTokenTransferIterator struct {
	Event *ZcTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ZcTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZcTokenTransfer)
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
		it.Event = new(ZcTokenTransfer)
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
func (it *ZcTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZcTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZcTokenTransfer represents a Transfer event raised by the ZcToken contract.
type ZcTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZcTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransferIterator{contract: _ZcToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZcTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZcTokenTransfer)
				if err := _ZcToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) ParseTransfer(log types.Log) (*ZcTokenTransfer, error) {
	event := new(ZcTokenTransfer)
	if err := _ZcToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
