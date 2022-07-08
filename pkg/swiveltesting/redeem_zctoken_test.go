package swiveltesting

import (
	"math/big"
	test "testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/mocks"
	"github.com/swivel-finance/gost/test/swivel"
)

type redeemZcTokenSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	EulerToken    *mocks.EulerTokenSession
	CompoundToken *mocks.CompoundTokenSession
	MarketPlace   *mocks.MarketPlaceSession
	Swivel        *swivel.SwivelSession
}

func (s *redeemZcTokenSuite) SetupTest() {
	var err error

	s.Env = NewEnv(big.NewInt(ONE_ETH))
	s.Dep, err = Deploy(s.Env)

	if err != nil {
		panic(err)
	}

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.EulerToken = &mocks.EulerTokenSession{
		Contract: s.Dep.EulerToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.CompoundToken = &mocks.CompoundTokenSession{
		Contract: s.Dep.CompoundToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.MarketPlace = &mocks.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.Swivel = &swivel.SwivelSession{
		Contract: s.Dep.Swivel,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}
}

func (s *redeemZcTokenSuite) TestRedeemZcToken() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity

	// stub the underlying to return true or the Safe lib will revert
	tx, err := s.Erc20.TransferReturns(true)
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CTokenAddressReturns(s.Dep.EulerTokenAddress)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// stub the marketplace mock.redeemZcToken
	redeemed := big.NewInt(12345)
	tx, err = s.MarketPlace.RedeemZcTokenReturns(redeemed)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456)
	tx, err = s.Swivel.RedeemZcToken(uint8(5), underlying, maturity, amount)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// underlying tranfer should have been called with an amount redeemed
	transferred, err := s.Erc20.TransferCalled(s.Env.Owner.Opts.From)
	assert.Nil(err)
	assert.Equal(redeemed, transferred)

	withdrawn, err := s.EulerToken.WithdrawCalled(big.NewInt(0))
	assert.NotNil(withdrawn)
	assert.Nil(err)
	assert.Equal(withdrawn, redeemed)
}

func (s *redeemZcTokenSuite) TestCompoundRedeemZcTokenRedeemUnderlyingFails() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity

	tx, err := s.CompoundToken.RedeemUnderlyingReturns(big.NewInt(1))
	assert.NotNil(tx)
	assert.Nil(err)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CTokenAddressReturns(s.Dep.CompoundTokenAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// stub the marketplace mock.redeemZcToken
	tx, err = s.MarketPlace.RedeemZcTokenReturns(big.NewInt(12345))
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	amount := big.NewInt(123456)
	tx, err = s.Swivel.RedeemZcToken(uint8(1), underlying, maturity, amount)
	assert.NotNil(err)
	// TODO as these are now "execution reverted" we'll skip until we identify how to get the args
	// assert.Regexp("withdraw failed", err.Error())
	assert.Nil(tx)
}

func TestRedeemZcTokenSuite(t *test.T) {
	suite.Run(t, &redeemZcTokenSuite{})
}
