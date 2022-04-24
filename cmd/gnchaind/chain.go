package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"

	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/tendermint/starport/starport/pkg/xurl"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/fatih/color"
	"github.com/glodnet/chain/pkg/chaincmd"
	chaincmdrunner "github.com/glodnet/chain/pkg/chaincmd/runner"
	"github.com/imdario/mergo"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/tendermint/starport/starport/pkg/confile"
)

const (
	moniker = "mynode"
)

// returns a command that groups sub commands related to compiling, serving blockchains and so on.
func chainCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "chain [./chain.yaml]",
		Short:   "Initialize files for a blockchain",
		Long:    `chain will populate with ecessary files (private validator, genesis, config, etc.) in home directories.`,
		Aliases: []string{"c"},
		Args:    cobra.RangeArgs(0, 1),
		RunE:    chainInitHandler,
	}

	cmd.Flags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().String(flags.FlagKeyringBackend, "test", "Select keyring's backend (os|file|kwallet|pass|test|memory)")
	cmd.Flags().String(flags.FlagGasPrices, "", "Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)")
	cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to Tendermint RPC interface for this chain")
	cmd.Flags().BoolP(cli.FlagOverwrite, "o", false, "overwrite directory for config and data")
	return cmd
}

func chainInitHandler(cmd *cobra.Command, args []string) error {
	homeDir, _ := cmd.Flags().GetString(flags.FlagHome)
	chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
	backend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
	gasPrices, _ := cmd.Flags().GetString(flags.FlagGasPrices)
	nodeAddress, _ := cmd.Flags().GetString(flags.FlagNode)
	overwrite, _ := cmd.Flags().GetBool(cli.FlagOverwrite)
	if overwrite {
		os.RemoveAll(homeDir)
	}

	path := "./chain.yml"
	if len(args) > 0 {
		path = args[0]
	}

	var conf Chain
	if err := parseConfig(path, &conf, &Chain{}); err != nil {
		return err
	}
	if conf.Genesis != nil {
		if id, ok := conf.Genesis["chain_id"]; ok {
			chainID = id.(string)
		}
	}

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

	ctx := cmd.Context()
	if err := initChain(ctx, runner, &conf); err != nil {
		return err
	}

	// overwrite configuration changes from config.yml to
	// over app's sdk configs.
	if err := configure(homeDir, &conf); err != nil {
		return err
	}

	fmt.Printf("🗃  Initialized. Checkout your chain's home (data) directory: %s\n", infoColor(homeDir))
	return nil
}

var infoColor = color.New(color.FgYellow).SprintFunc()

// initChain initializes the chain accounts and creates validator gentxs
func initChain(ctx context.Context, runner chaincmdrunner.Runner, conf *Chain) error {
	generatedAccounts := map[string]chaincmdrunner.Account{}
	accounts := map[string]Account{}
	// add accounts from config into keys
	for _, account := range conf.Accounts {
		// If the account doesn't provide an address, we create one
		if account.Address == "" {
			generatedAccount, err := runner.AddAccount(ctx, account.Name, account.Mnemonic, account.CoinType, account.Algo)
			if err != nil {
				return err
			}
			generatedAccounts[account.Name] = generatedAccount
			accounts[account.Name] = account
		}
	}

	mnemonic := ""
	if len(conf.Validators) > 0 {
		if account, ok := generatedAccounts[conf.Validators[0].Name]; ok {
			mnemonic = account.Mnemonic
		}
	}
	if err := runner.Init(ctx, moniker, mnemonic); err != nil {
		return err
	}

	// add accounts from config into genesis
	for _, account := range conf.Accounts {
		accountAddress := account.Address
		if account.Address == "" {
			generatedAccount := generatedAccounts[account.Name]
			accountAddress = generatedAccount.Address
			fmt.Printf(
				"🙂 Created account %q with address %q with mnemonic: %q\n",
				generatedAccount.Name,
				generatedAccount.Address,
				generatedAccount.Mnemonic,
			)
		} else {
			fmt.Printf(
				"🙂 Imported an account %q with address: %q\n",
				account.Name,
				account.Address,
			)
		}
		coins := strings.Join(account.Coins, ",")
		if err := runner.AddGenesisAccount(ctx, accountAddress, coins); err != nil {
			return err
		}
		if len(account.Roles) > 0 {
			if err := runner.AddGenesisAccountRoles(ctx, accountAddress, account.Roles); err != nil {
				return err
			}
		}
	}

	homeDir := runner.Cmd().Home()
	gentxsDir := filepath.Join(".", "gentxs")
	defer os.RemoveAll(gentxsDir)
	for _, v := range conf.Validators {
		nodeDir := filepath.Join(gentxsDir, fmt.Sprintf("node_%s", v.Name))
		trunner, _ := chaincmdrunner.New(ctx, runner.Cmd().Copy(chaincmd.WithHome(nodeDir)))

		mnemonic := ""
		if account, ok := generatedAccounts[v.Name]; ok {
			mnemonic = account.Mnemonic
		}
		if err := trunner.Init(ctx, moniker, mnemonic); err != nil {
			return err
		}
		if account, ok := accounts[v.Name]; ok {
			_, err := trunner.AddAccount(ctx, account.Name, generatedAccounts[v.Name].Mnemonic, account.CoinType, account.Algo)
			if err != nil {
				return err
			}
		}
		if err := copy.Copy(filepath.Join(homeDir, "config", "genesis.json"), filepath.Join(nodeDir, "config", "genesis.json")); err != nil {
			return err
		}

		_, err := trunner.Gentx(
			ctx,
			v.Name,
			v.StakingAmount,
			chaincmd.GentxWithMoniker(v.Moniker),
			chaincmd.GentxWithCommissionRate(v.CommissionRate),
			chaincmd.GentxWithCommissionMaxRate(v.CommissionMaxRate),
			chaincmd.GentxWithCommissionMaxChangeRate(v.CommissionMaxChangeRate),
			chaincmd.GentxWithMinSelfDelegation(v.MinSelfDelegation),
			chaincmd.GentxWithDetails(v.Details),
			chaincmd.GentxWithIdentity(v.Identity),
			chaincmd.GentxWithWebsite(v.Website),
			chaincmd.GentxWithSecurityContact(v.SecurityContact),
		)
		if err != nil {
			return err
		}
		fmt.Printf(
			"🙂 Created an validator %q with staking amount: %q\n",
			v.Name,
			v.StakingAmount,
		)
		copy.Copy(filepath.Join(nodeDir, "config", "gentx"), filepath.Join(homeDir, "config", "gentx"))
	}

	return runner.CollectGentxs(ctx)
}

func configure(homePath string, conf *Chain) error {
	if err := appTOML(homePath, conf); err != nil {
		return err
	}
	if err := clientTOML(homePath); err != nil {
		return err
	}
	if err := configTOML(homePath, conf); err != nil {
		return err
	}

	genesisPath := filepath.Join(homePath, "config/genesis.json")
	appTOMLPath := filepath.Join(homePath, "config/app.toml")
	configTOMLPath := filepath.Join(homePath, "config/config.toml")
	clientTOMLPath := filepath.Join(homePath, "config/client.toml")

	appconfigs := []struct {
		ec      confile.EncodingCreator
		path    string
		changes map[string]interface{}
	}{
		{confile.DefaultJSONEncodingCreator, genesisPath, conf.Genesis},
		{confile.DefaultTOMLEncodingCreator, appTOMLPath, conf.Init.App},
		{confile.DefaultTOMLEncodingCreator, clientTOMLPath, conf.Init.Client},
		{confile.DefaultTOMLEncodingCreator, configTOMLPath, conf.Init.Config},
	}

	for _, ac := range appconfigs {
		cf := confile.New(ac.ec, ac.path)
		var conf map[string]interface{}
		if err := cf.Load(&conf); err != nil {
			return err
		}
		if err := mergo.Merge(&conf, ac.changes, mergo.WithOverride); err != nil {
			return err
		}
		if err := cf.Save(conf); err != nil {
			return err
		}
	}
	return nil
}

func appTOML(homePath string, conf *Chain) error {
	// TODO find a better way in order to not delete comments in the toml.yml
	path := filepath.Join(homePath, "config/app.toml")
	config, err := toml.LoadFile(path)
	if err != nil {
		return err
	}
	config.Set("api.enable", true)
	config.Set("api.enabled-unsafe-cors", true)
	config.Set("rpc.cors_allowed_origins", []string{"*"})
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = config.WriteTo(file)
	return err
}

func configTOML(homePath string, conf *Chain) error {
	// TODO find a better way in order to not delete comments in the toml.yml
	path := filepath.Join(homePath, "config/config.toml")
	config, err := toml.LoadFile(path)
	if err != nil {
		return err
	}
	config.Set("rpc.cors_allowed_origins", []string{"*"})
	config.Set("consensus.timeout_commit", "1s")
	config.Set("consensus.timeout_propose", "1s")
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = config.WriteTo(file)
	return err
}

func clientTOML(homePath string) error {
	path := filepath.Join(homePath, "config/client.toml")
	config, err := toml.LoadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	config.Set("keyring-backend", "test")
	config.Set("broadcast-mode", "block")
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = config.WriteTo(file)
	return err
}

type Account struct {
	Name     string   `yaml:"name"`
	Coins    []string `yaml:"coins,omitempty"`
	Mnemonic string   `yaml:"mnemonic,omitempty"`
	Address  string   `yaml:"address,omitempty"`
	CoinType string   `yaml:"coin_type,omitempty"`
	Algo     string   `yaml:"algo,omitempty"`
	Roles    []string `yaml:"roles,omitempty"`
}

type Validator struct {
	Name                    string `json:"name"`
	Moniker                 string `json:"moniker"`
	StakingAmount           string `json:"sef_delegation"`
	CommissionRate          string `json:"commission_rate"`
	CommissionMaxRate       string `json:"commission_max_rate"`
	CommissionMaxChangeRate string `json:"commission_max_change_rate"`
	MinSelfDelegation       string `json:"min_self_delegation"`
	Details                 string `json:"details"`
	Identity                string `json:"identity"`
	Website                 string `json:"website"`
	SecurityContact         string `json:"security_contact"`
}

// Init overwrites sdk configurations with given values.
type Init struct {
	// App overwrites appd's config/app.toml configs.
	App map[string]interface{} `yaml:"app"`

	// Client overwrites appd's config/client.toml configs.
	Client map[string]interface{} `yaml:"client"`

	// Config overwrites appd's config/config.toml configs.
	Config map[string]interface{} `yaml:"config"`

	// Home overwrites default home directory used for the app
	Home string `yaml:"home"`

	// KeyringBackend is the default keyring backend to use for blockchain initialization
	KeyringBackend string `yaml:"keyring-backend"`
}

type Chain struct {
	Accounts   []Account              `yaml:"accounts"`
	Validators []Validator            `yaml:"validators"`
	Genesis    map[string]interface{} `yaml:"genesis"`
	Init       Init                   `yaml:"init"`
}
