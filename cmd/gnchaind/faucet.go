package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glodnet/chain/pkg/chaincmd"
	chaincmdrunner "github.com/glodnet/chain/pkg/chaincmd/runner"
	"github.com/glodnet/chain/pkg/cosmosfaucet"
	"github.com/goccy/go-yaml"
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
	"github.com/tendermint/starport/starport/pkg/xhttp"
	"github.com/tendermint/starport/starport/pkg/xurl"
)

func faucetServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "faucet [./faucet.yml]",
		Short:   "Start a faucet server",
		Long:    "The faucet service sends tokens to addresses.",
		Aliases: []string{"f"},
		Args:    cobra.RangeArgs(0, 1),
		RunE:    faucetServeHandler,
	}

	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().String(flags.FlagKeyringBackend, "test", "Select keyring's backend (os|file|kwallet|pass|test|memory)")
	cmd.Flags().String(flags.FlagGasPrices, "", "Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)")
	cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to Tendermint RPC interface for this chain")
	return cmd
}

func faucetServeHandler(cmd *cobra.Command, args []string) error {
	homeDir, _ := cmd.Flags().GetString(flags.FlagHome)
	chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
	backend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
	gasPrices, _ := cmd.Flags().GetString(flags.FlagGasPrices)
	nodeAddress, _ := cmd.Flags().GetString(flags.FlagNode)
	chainCommandOptions := []chaincmd.Option{
		chaincmd.WithHome(homeDir),
		chaincmd.WithChainID(chainID),
		chaincmd.WithGasPrices(gasPrices),
		chaincmd.WithNodeAddress(xurl.TCP(nodeAddress)),
		chaincmd.WithKeyringBackend(chaincmd.KeyringBackend(backend)),
	}
	cc := chaincmd.New(binary(), chainCommandOptions...)

	ccrOptions := make([]chaincmdrunner.Option, 0)
	//ccrOptions = append(ccrOptions,
	//	chaincmdrunner.Stdout(os.Stdout),
	//	chaincmdrunner.Stderr(os.Stderr),
	//)
	runner, err := chaincmdrunner.New(cmd.Context(), cc, ccrOptions...)
	if err != nil {
		return err
	}

	path := "./faucet.yml"
	if len(args) > 0 {
		path = args[0]
	}

	var conf Faucet
	if err := parseConfig(path, &conf, &Faucet{
		Host:     "0.0.0.0:5400",
		API:      "http://localhost:1317",
		Algo:     "sm2",
		CoinType: "118",
	}); err != nil {
		return err
	}
	if len(conf.Mnemonic) > 0 {

	} else if _, err := runner.ShowAccount(cmd.Context(), conf.Name); err != nil {
		return err
	}

	faucetOptions := []cosmosfaucet.Option{
		cosmosfaucet.ChainID(chainID),
		cosmosfaucet.Account(conf.Name, conf.Mnemonic, conf.CoinType, conf.Algo),
		cosmosfaucet.OpenAPI(xurl.HTTP(conf.API)),
	}
	// parse coins to pass to the faucet as coins.
	for _, coin := range conf.Coins {
		parsedCoin, err := sdk.ParseCoinNormalized(coin)
		if err != nil {
			return fmt.Errorf("%s: %s", err, coin)
		}

		var amountMax uint64
		// find out the max amount for this coin.
		for _, coinMax := range conf.CoinsMax {
			parsedMax, err := sdk.ParseCoinNormalized(coinMax)
			if err != nil {
				return fmt.Errorf("%s: %s", err, coin)
			}
			if parsedMax.Denom == parsedCoin.Denom {
				amountMax = parsedMax.Amount.Uint64()
				break
			}
		}
		faucetOptions = append(faucetOptions, cosmosfaucet.Coin(parsedCoin.Amount.Uint64(), amountMax, parsedCoin.Denom))
	}

	if conf.RateLimitWindow != "" {
		rateLimitWindow, err := time.ParseDuration(conf.RateLimitWindow)
		if err != nil {
			return fmt.Errorf("%s: %s", err, conf.RateLimitWindow)
		}
		faucetOptions = append(faucetOptions, cosmosfaucet.RefreshWindow(rateLimitWindow))
	}

	ctx := cmd.Context()

	faucet, err := cosmosfaucet.New(ctx, runner, faucetOptions...)
	if err != nil {
		return err
	}

	fmt.Printf("🌍 Token faucet: %s\n", infoColor(xurl.HTTP(conf.Host)))

	return xhttp.Serve(ctx, &http.Server{
		Addr:    conf.Host,
		Handler: faucet,
	})
}

func parseConfig(path string, conf interface{}, defaultConf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(conf); err != nil {
		return err
	}

	if err := mergo.Merge(conf, defaultConf); err != nil {
		return err
	}
	return nil
}

type Faucet struct {
	// Name is faucet account's name.
	Name string `yaml:"name"`
	// Mnemonic used to generate an account. This field is ignored if name already exists.
	Mnemonic string `yaml:"mnemonic,omitempty"`
	// CoinType number for HD derivation to generate an account.
	CoinType string `yaml:"coin_type,omitempty"`
	// Algo signing algorithm used to generate an account.
	Algo string `yaml:"algo,omitempty"`
	// Coins holds type of coin denoms and amounts to distribute.
	Coins []string `yaml:"coins"`
	// CoinsMax holds of chain denoms and their max amounts that can be transferred to single user.
	CoinsMax []string `yaml:"coins_max"`
	// LimitRefreshTime sets the timeframe at the end of which the limit will be refreshed
	RateLimitWindow string `yaml:"rate_limit_window,omitempty"`
	// Host is the host of the faucet server
	Host string `yaml:"host"`
	// API is the api of the blockchain node
	API string `yaml:"api"`
}

func binary() string {
	executableName, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return executableName
}
