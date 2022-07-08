package marketplacetesting

import (
	"context"
	"math/big"
	test "testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	assertions "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/swivel-finance/gost/test/marketplace"
	"github.com/swivel-finance/gost/test/mocks"
)

type matureMarketSuite struct {
	suite.Suite
	Env           *Env
	Dep           *Dep
	Erc20         *mocks.Erc20Session
	CompoundToken *mocks.CompoundTokenSession
	AaveToken     *mocks.AaveTokenSession
	AavePool      *mocks.AavePoolSession
	MarketPlace   *marketplace.MarketPlaceSession // *Session objects are created by the go bindings
}

func (s *matureMarketSuite) SetupTest() {
	var err error
	assert := assertions.New(s.T())

	s.Env = NewEnv(big.NewInt(ONE_ETH)) // each of the wallets in the env will begin with this balance
	s.Dep, err = Deploy(s.Env)
	assert.Nil(err)

	err = s.Env.Blockchain.AdjustTime(0) // set bc timestamp to 0
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	s.Erc20 = &mocks.Erc20Session{
		Contract: s.Dep.Erc20,
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

	s.AaveToken = &mocks.AaveTokenSession{
		Contract: s.Dep.AaveToken,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	s.AavePool = &mocks.AavePoolSession{
		Contract: s.Dep.AavePool,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// binding owner to both, kind of why it exists - but could be any of the env wallets
	s.MarketPlace = &marketplace.MarketPlaceSession{
		Contract: s.Dep.MarketPlace,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	// the swivel address must be set
	_, err = s.MarketPlace.SetSwivel(s.Dep.SwivelAddress)
	assert.Nil(err)
	s.Env.Blockchain.Commit()
}

func (s *matureMarketSuite) TestMatureFailsWhenPaused() {
	assert := assertions.New(s.T())

	underlying := s.Dep.Erc20Address
	maturity := big.NewInt(123456789)

	tx, err := s.MarketPlace.Pause(true)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MatureMarket(uint8(1), underlying, maturity)

	assert.Nil(tx)
	assert.NotNil(err)
	// assert.Regexp("markets are paused", err.Error())

	// unpause so the other tests don't fail
	tx, err = s.MarketPlace.Pause(false)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()
}

func (s *matureMarketSuite) TestMaturityNotReached() {
	assert := assertions.New(s.T())
	// addresses can be BS in this test as well...
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	ctoken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	tx, err = s.MarketPlace.MatureMarket(uint8(1), underlying, maturity)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("maturity not reached", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func (s *matureMarketSuite) TestMaturityReached() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	ctoken := s.Dep.AaveTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.AaveToken.UNDERLYINGReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.AaveToken.POOLReturns(s.Dep.AavePoolAddress)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(4),
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(4), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultTracker, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	tx, err = vaultTracker.MatureVaultReturns(true)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	rate := big.NewInt(123456789)
	tx, err = s.AavePool.GetReserveNormalizedIncomeReturns(rate)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MatureMarket(uint8(4), underlying, maturity)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// re-fetch the market so the change is reflected...
	market, err = s.MarketPlace.Markets(uint8(4), underlying, maturity)
	assert.Nil(err)
	assert.Equal(rate, market.MaturityRate)

	receipt, err := s.Env.Blockchain.TransactionReceipt(context.Background(), tx.Hash())
	assert.Nil(err)
	assert.NotNil(receipt)

	logs := receipt.Logs
	assert.NotNil(logs)
	assert.Equal(1, len(logs))

	assert.Equal(MATURE_EVENT_SIG, logs[0].Topics[0].Hex())
	assert.Equal(underlying.Hex(), common.HexToAddress(logs[0].Topics[2].Hex()).String())
	assert.Equal(maturity, logs[0].Topics[3].Big())
}

func (s *matureMarketSuite) TestVaultMaturityNotReachedRequireFail() {
	assert := assertions.New(s.T())
	underlying := s.Dep.Erc20Address
	maturity := s.Dep.Maturity
	ctoken := s.Dep.CompoundTokenAddress

	tx, err := s.Erc20.DecimalsReturns(uint8(18))
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.CompoundToken.UnderlyingReturns(underlying)
	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.CreateMarket(
		uint8(1),
		maturity,
		ctoken,
		"awesome market",
		"AM",
	)

	assert.Nil(err)
	assert.NotNil(tx)
	s.Env.Blockchain.Commit()

	// we should be able to fetch the market now...
	market, err := s.MarketPlace.Markets(uint8(1), underlying, maturity)
	assert.Nil(err)
	assert.Equal(market.CTokenAddr, ctoken)

	zcTokenContract, err := mocks.NewZcToken(market.ZcToken, s.Env.Blockchain)
	zcToken := &mocks.ZcTokenSession{
		Contract: zcTokenContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	zcMaturity, err := zcToken.Maturity()
	assert.Equal(maturity, zcMaturity)

	vaultTrackerContract, err := mocks.NewVaultTracker(market.VaultTracker, s.Env.Blockchain)
	vaultTracker := &mocks.VaultTrackerSession{
		Contract: vaultTrackerContract,
		CallOpts: bind.CallOpts{From: s.Env.Owner.Opts.From, Pending: false},
		TransactOpts: bind.TransactOpts{
			From:   s.Env.Owner.Opts.From,
			Signer: s.Env.Owner.Opts.Signer,
		},
	}

	tx, err = vaultTracker.MatureVaultReturns(false)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	// move past the maturity
	err = s.Env.Blockchain.AdjustTime(MATURITY * time.Second)
	assert.Nil(err)
	s.Env.Blockchain.Commit()

	rate := big.NewInt(123456789)
	tx, err = s.CompoundToken.ExchangeRateCurrentReturns(rate)
	assert.Nil(err)
	assert.NotNil(tx)

	s.Env.Blockchain.Commit()

	tx, err = s.MarketPlace.MatureMarket(uint8(1), underlying, maturity)
	assert.NotNil(err)
	// TODO extract the custom error codes?
	// assert.Regexp("mature vault failed", err.Error())
	assert.Nil(tx)

	s.Env.Blockchain.Commit()
}

func TestMatureMarketSuite(t *test.T) {
	suite.Run(t, &matureMarketSuite{})
}
