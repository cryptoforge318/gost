package testing

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/contracts/swivel"
)

type fillFixedTestSuite struct {
	suite.Suite
	Env    *Env
	Dep    *Dep
	Swivel *swivel.SwivelSession // *Session objects are created by the go bindings
}

func (s *fillFixedTestSuite) SetupSuite() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *fillFixedTestSuite) TestFoo() {
	assert := assert.New(s.T())
	assert.Equal(true, true)
}

func TestSwivelSuite(t *test.T) {
	suite.Run(t, &fillFixedTestSuite{})
}
