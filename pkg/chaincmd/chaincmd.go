package chaincmd

import (
	"fmt"

	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
)

const (
	commandStart             = "start"
	commandInit              = "init"
	commandKeys              = "keys"
	commandGenesis           = "genesis"
	commandAddPermssion      = "add-permission"
	commandAddIdentify       = "add-identity"
	commandAddToken          = "add-token"
	commandAddGenesisAccount = "add-genesis-account"
	commandGentx             = "gentx"
	commandCollectGentxs     = "collect-gentxs"
	commandValidateGenesis   = "validate-genesis"
	commandShowNodeID        = "show-node-id"
	commandStatus            = "status"
	commandTx                = "tx"
	commandQuery             = "query"
	commandUnsafeReset       = "unsafe-reset-all"
	commandExport            = "export"
	commandConfig            = "config"

	optionHome                             = "--home"
	optionNode                             = "--node"
	optionKeyringBackend                   = "--keyring-backend"
	optionChainID                          = "--chain-id"
	optionOutput                           = "--output"
	optionRecover                          = "--recover"
	optionAddress                          = "--address"
	optionAlgo                             = "--algo"
	optionAmount                           = "--amount"
	optionGas                              = "--gas"
	optionGasAdjustment                    = "--gas-adjustment"
	optionGasPrices                        = "--gas-prices"
	optionValidatorMoniker                 = "--moniker"
	optionValidatorCommissionRate          = "--commission-rate"
	optionValidatorCommissionMaxRate       = "--commission-max-rate"
	optionValidatorCommissionMaxChangeRate = "--commission-max-change-rate"
	optionValidatorMinSelfDelegation       = "--min-self-delegation"
	optionValidatorDetails                 = "--details"
	optionValidatorIdentity                = "--identity"
	optionValidatorWebsite                 = "--website"
	optionValidatorSecurityContact         = "--security-contact"
	optionYes                              = "--yes"
	optionCoinType                         = "--coin-type"
	optionVestingAmount                    = "--vesting-amount"
	optionVestingEndTime                   = "--vesting-end-time"
	optionBroadcastMode                    = "--broadcast-mode"

	constTendermint = "tendermint"
	constJSON       = "json"
)

type KeyringBackend string

const (
	KeyringBackendUnspecified KeyringBackend = ""
	KeyringBackendOS          KeyringBackend = "os"
	KeyringBackendFile        KeyringBackend = "file"
	KeyringBackendPass        KeyringBackend = "pass"
	KeyringBackendTest        KeyringBackend = "test"
	KeyringBackendKwallet     KeyringBackend = "kwallet"
)

type ChainCmd struct {
	appCmd          string
	chainID         string
	homeDir         string
	keyringBackend  KeyringBackend
	keyringPassword string
	nodeAddress     string
	gasPrices       string
	gas             string
	gasAdjustment   string

	isAutoChainIDDetectionEnabled bool
}

// Option configures ChainCmd.
type Option func(*ChainCmd)

// WithHome replaces the default home used by the chain
func WithHome(home string) Option {
	return func(c *ChainCmd) {
		c.homeDir = home
	}
}

// WithChainID provides a specific chain ID for the commands that accept this option
func WithChainID(chainID string) Option {
	return func(c *ChainCmd) {
		c.chainID = chainID
	}
}

// WithGasPrice replaces the default gas price for the commands that accept this option
func WithGasPrices(gasPrices string) Option {
	return func(c *ChainCmd) {
		c.gasPrices = gasPrices
	}
}

// WithGas replaces the default gas for the commands that accept this option
func WithGas(gas string) Option {
	return func(c *ChainCmd) {
		c.gas = gas
	}
}

// WithGasPrice replaces the default gas adjustment for the commands that accept this option
func WithGasAdjustment(gasAdjustment string) Option {
	return func(c *ChainCmd) {
		c.gasAdjustment = gasAdjustment
	}
}

// WithAutoChainIDDetection finds out the chain id by communicating with the node running.
func WithAutoChainIDDetection() Option {
	return func(c *ChainCmd) {
		c.isAutoChainIDDetectionEnabled = true
	}
}

// WithKeyringBackend provides a specific keyring backend for the commands that accept this option
func WithKeyringBackend(keyringBackend KeyringBackend) Option {
	return func(c *ChainCmd) {
		c.keyringBackend = keyringBackend
	}
}

// WithKeyringPassword provides a password to unlock keyring
func WithKeyringPassword(password string) Option {
	return func(c *ChainCmd) {
		c.keyringPassword = password
	}
}

// WithNodeAddress sets the node address for the commands that needs to make an
// API request to the node that has a different node address other than the default one.
func WithNodeAddress(addr string) Option {
	return func(c *ChainCmd) {
		c.nodeAddress = addr
	}
}

// New creates a new ChainCmd to launch command with the chain app
func New(appCmd string, options ...Option) ChainCmd {
	chainCmd := ChainCmd{
		appCmd: appCmd,
	}

	for _, applyOption := range options {
		applyOption(&chainCmd)
	}

	return chainCmd
}

// Copy makes a copy of ChainCmd by overwriting its options with given options.
func (c ChainCmd) Copy(options ...Option) ChainCmd {
	for _, applyOption := range options {
		applyOption(&c)
	}

	return c
}

func (c ChainCmd) Home() string {
	return c.homeDir
}

// StartCommand returns the command to start the daemon of the chain
func (c ChainCmd) StartCommand(options ...string) step.Option {
	command := append([]string{
		commandStart,
	}, options...)
	return c.daemonCommand(command)
}

// InitCommand returns the command to initialize the chain
func (c ChainCmd) InitCommand(moniker string) step.Option {
	command := []string{
		commandInit,
		moniker,
	}
	command = c.attachChainID(command)
	return c.daemonCommand(command)
}

// RecoverInitCommand returns the command to initialize the chain
func (c ChainCmd) RecoverInitCommand(moniker string) step.Option {
	command := []string{
		commandInit,
		moniker,
		optionRecover,
	}
	command = c.attachChainID(command)
	return c.daemonCommand(command)
}

// AddKeyCommand returns the command to add a new key in the chain keyring
func (c ChainCmd) AddKeyCommand(accountName, coinType string, algo string) step.Option {
	command := []string{
		commandKeys,
		"add",
		accountName,
		optionOutput,
		constJSON,
	}
	if coinType != "" {
		command = append(command, optionCoinType, coinType)
	}
	if algo != "" {
		command = append(command, optionAlgo, algo)
	}
	command = c.attachKeyringBackend(command)

	return c.cliCommand(command)
}

// RecoverKeyCommand returns the command to recover a key into the chain keyring from a mnemonic
func (c ChainCmd) RecoverKeyCommand(accountName, coinType string, algo string) step.Option {
	command := []string{
		commandKeys,
		"add",
		accountName,
		optionRecover,
		optionOutput,
		constJSON,
	}
	if coinType != "" {
		command = append(command, optionCoinType, coinType)
	}
	if algo != "" {
		command = append(command, optionAlgo, algo)
	}
	command = c.attachKeyringBackend(command)

	return c.cliCommand(command)
}

// ImportKeyCommand returns the command to import a key into the chain keyring from a key file
func (c ChainCmd) ImportKeyCommand(accountName, keyFile string) step.Option {
	command := []string{
		commandKeys,
		"import",
		accountName,
		keyFile,
	}
	command = c.attachKeyringBackend(command)

	return c.cliCommand(command)
}

// ShowKeyAddressCommand returns the command to print the address of a key in the chain keyring
func (c ChainCmd) ShowKeyAddressCommand(accountName string) step.Option {
	command := []string{
		commandKeys,
		"show",
		accountName,
		optionAddress,
	}
	command = c.attachKeyringBackend(command)

	return c.cliCommand(command)
}

// ListKeysCommand returns the command to print the list of a keys in the chain keyring
func (c ChainCmd) ListKeysCommand() step.Option {
	command := []string{
		commandKeys,
		"list",
		optionOutput,
		constJSON,
	}
	command = c.attachKeyringBackend(command)

	return c.cliCommand(command)
}

// AddGenesisAccountRolesCommand returns the command to add a new account in the genesis file of the chain
func (c ChainCmd) AddGenesisAccountRolesCommand(address string, roles []string) step.Option {
	command := []string{
		commandGenesis,
		commandAddPermssion,
		address,
	}
	command = append(command, roles...)

	return c.daemonCommand(command)
}

// AddGenesisAccountCommand returns the command to add a new account in the genesis file of the chain
func (c ChainCmd) AddGenesisAccountCommand(address, coins string) step.Option {
	command := []string{
		commandAddGenesisAccount,
		address,
		coins,
	}

	return c.daemonCommand(command)
}

// AddVestingAccountCommand returns the command to add a delayed vesting account in the genesis file of the chain
func (c ChainCmd) AddVestingAccountCommand(address, originalCoins, vestingCoins string, vestingEndTime int64) step.Option {
	command := []string{
		commandAddGenesisAccount,
		address,
		originalCoins,
		optionVestingAmount,
		vestingCoins,
		optionVestingEndTime,
		fmt.Sprintf("%d", vestingEndTime),
	}

	return c.daemonCommand(command)
}

// GentxOption for the GentxCommand
type GentxOption func([]string) []string

// GentxWithMoniker provides moniker option for the gentx command
func GentxWithMoniker(moniker string) GentxOption {
	return func(command []string) []string {
		if len(moniker) > 0 {
			return append(command, optionValidatorMoniker, moniker)
		}
		return command
	}
}

// GentxWithCommissionRate provides commission rate option for the gentx command
func GentxWithCommissionRate(commissionRate string) GentxOption {
	return func(command []string) []string {
		if len(commissionRate) > 0 {
			return append(command, optionValidatorCommissionRate, commissionRate)
		}
		return command
	}
}

// GentxWithCommissionMaxRate provides commission max rate option for the gentx command
func GentxWithCommissionMaxRate(commissionMaxRate string) GentxOption {
	return func(command []string) []string {
		if len(commissionMaxRate) > 0 {
			return append(command, optionValidatorCommissionMaxRate, commissionMaxRate)
		}
		return command
	}
}

// GentxWithCommissionMaxChangeRate provides commission max change rate option for the gentx command
func GentxWithCommissionMaxChangeRate(commissionMaxChangeRate string) GentxOption {
	return func(command []string) []string {
		if len(commissionMaxChangeRate) > 0 {
			return append(command, optionValidatorCommissionMaxChangeRate, commissionMaxChangeRate)
		}
		return command
	}
}

// GentxWithMinSelfDelegation provides minimum self delegation option for the gentx command
func GentxWithMinSelfDelegation(minSelfDelegation string) GentxOption {
	return func(command []string) []string {
		if len(minSelfDelegation) > 0 {
			return append(command, optionValidatorMinSelfDelegation, minSelfDelegation)
		}
		return command
	}
}

// GentxWithDetails provides validator details option for the gentx command
func GentxWithDetails(details string) GentxOption {
	return func(command []string) []string {
		if len(details) > 0 {
			return append(command, optionValidatorDetails, details)
		}
		return command
	}
}

// GentxWithIdentity provides validator identity option for the gentx command
func GentxWithIdentity(identity string) GentxOption {
	return func(command []string) []string {
		if len(identity) > 0 {
			return append(command, optionValidatorIdentity, identity)
		}
		return command
	}
}

// GentxWithWebsite provides validator website option for the gentx command
func GentxWithWebsite(website string) GentxOption {
	return func(command []string) []string {
		if len(website) > 0 {
			return append(command, optionValidatorWebsite, website)
		}
		return command
	}
}

// GentxWithSecurityContact provides validator security contact option for the gentx command
func GentxWithSecurityContact(securityContact string) GentxOption {
	return func(command []string) []string {
		if len(securityContact) > 0 {
			return append(command, optionValidatorSecurityContact, securityContact)
		}
		return command
	}
}

func (c ChainCmd) IsAutoChainIDDetectionEnabled() bool {
	return c.isAutoChainIDDetectionEnabled
}

// GentxCommand returns the command to generate a gentx for the chain
func (c ChainCmd) GentxCommand(
	validatorName string,
	selfDelegation string,
	options ...GentxOption,
) step.Option {
	command := []string{
		commandGentx,
		validatorName,
		selfDelegation,
	}

	// Apply the options provided by the user
	for _, applyOption := range options {
		command = applyOption(command)
	}

	command = c.attachChainID(command)

	command = c.attachKeyringBackend(command)

	command = c.attachGasPrice(command)
	command = c.attachGas(command)
	command = c.attachGasAdjustment(command)

	return c.daemonCommand(command)
}

// CollectGentxsCommand returns the command to gather the gentxs in /gentx dir into the genesis file of the chain
func (c ChainCmd) CollectGentxsCommand() step.Option {
	command := []string{
		commandCollectGentxs,
	}
	return c.daemonCommand(command)
}

// ValidateGenesisCommand returns the command to check the validity of the chain genesis
func (c ChainCmd) ValidateGenesisCommand() step.Option {
	command := []string{
		commandValidateGenesis,
	}
	return c.daemonCommand(command)
}

// ShowNodeIDCommand returns the command to print the node ID of the node for the chain
func (c ChainCmd) ShowNodeIDCommand() step.Option {
	command := []string{
		constTendermint,
		commandShowNodeID,
	}
	return c.daemonCommand(command)
}

// UnsafeResetCommand returns the command to reset the blockchain database
func (c ChainCmd) UnsafeResetCommand() step.Option {
	command := []string{
		commandUnsafeReset,
	}
	return c.daemonCommand(command)
}

// ExportCommand returns the command to export the state of the blockchain into a genesis file
func (c ChainCmd) ExportCommand() step.Option {
	command := []string{
		commandExport,
	}
	return c.daemonCommand(command)
}

// BankSendCommand returns the command for transferring tokens.
func (c ChainCmd) BankSendCommand(fromAddress, toAddress, amount string) step.Option {
	command := []string{
		commandTx,
	}

	command = append(command,
		"bank",
		"send",
		fromAddress,
		toAddress,
		amount,
		optionYes,
	)

	command = c.attachChainID(command)
	command = c.attachKeyringBackend(command)
	command = c.attachNode(command)
	command = c.attachGasPrice(command)
	command = c.attachGas(command)
	command = c.attachGasAdjustment(command)

	return c.cliCommand(command)
}

// QueryTxCommand returns the command to query tx
func (c ChainCmd) QueryTxCommand(txHash string) step.Option {
	command := []string{
		commandQuery,
		"tx",
		txHash,
	}

	command = c.attachNode(command)
	return c.cliCommand(command)
}

// QueryTxEventsCommand returns the command to query events.
func (c ChainCmd) QueryTxEventsCommand(query string) step.Option {
	command := []string{
		commandQuery,
		"txs",
		"--events",
		query,
		"--page", "1",
		"--limit", "1000",
	}

	command = c.attachNode(command)
	return c.cliCommand(command)
}

// ConfigSetCommand
func (c ChainCmd) ConfigSetCommand(name string, value string) step.Option {
	command := []string{
		commandConfig,
		name,
		value,
	}

	return c.cliCommand(command)
}

// StatusCommand returns the command that fetches node's status.
func (c ChainCmd) StatusCommand() step.Option {
	command := []string{
		commandStatus,
	}

	command = c.attachNode(command)
	return c.cliCommand(command)
}

// KeyringBackend returns the underlying keyring backend.
func (c ChainCmd) KeyringBackend() KeyringBackend {
	return c.keyringBackend
}

// KeyringPassword returns the underlying keyring password.
func (c ChainCmd) KeyringPassword() string {
	return c.keyringPassword
}

// attachChainID appends the chain ID flag to the provided command
func (c ChainCmd) attachChainID(command []string) []string {
	if c.chainID != "" {
		command = append(command, []string{optionChainID, c.chainID}...)
	}
	return command
}

// attachGasPrice appends the gas price flag to the provided command
func (c ChainCmd) attachGasPrice(command []string) []string {
	if c.gasPrices != "" {
		command = append(command, []string{optionGasPrices, c.gasPrices}...)
	}
	return command
}

// attachGas appends the gas price flag to the provided command
func (c ChainCmd) attachGas(command []string) []string {
	if c.gas != "" {
		command = append(command, []string{optionGas, c.gas}...)
	}
	return command
}

// attachGasAdjustment appends the gas price flag to the provided command
func (c ChainCmd) attachGasAdjustment(command []string) []string {
	if c.gasAdjustment != "" {
		command = append(command, []string{optionGasAdjustment, c.gasAdjustment}...)
	}
	return command
}

// attachKeyringBackend appends the keyring backend flag to the provided command
func (c ChainCmd) attachKeyringBackend(command []string) []string {
	if c.keyringBackend != "" {
		command = append(command, []string{optionKeyringBackend, string(c.keyringBackend)}...)
	}
	return command
}

// attachHome appends the home flag to the provided command
func (c ChainCmd) attachHome(command []string) []string {
	if c.homeDir != "" {
		command = append(command, []string{optionHome, c.homeDir}...)
	}
	return command
}

// attachNode appends the node flag to the provided command
func (c ChainCmd) attachNode(command []string) []string {
	if c.nodeAddress != "" {
		command = append(command, []string{optionNode, c.nodeAddress}...)
	}
	return command
}

// daemonCommand returns the daemon command from the provided command
func (c ChainCmd) daemonCommand(command []string, options ...string) step.Option {
	return step.Exec(c.appCmd, c.attachHome(command)...)
}

// cliCommand returns the cli command from the provided command
// cli is the daemon for Stargate
func (c ChainCmd) cliCommand(command []string) step.Option {
	return step.Exec(c.appCmd, c.attachHome(command)...)
}

// KeyringBackendFromString returns the keyring backend from its string
func KeyringBackendFromString(kb string) (KeyringBackend, error) {
	existingKeyringBackend := map[KeyringBackend]bool{
		KeyringBackendUnspecified: true,
		KeyringBackendOS:          true,
		KeyringBackendFile:        true,
		KeyringBackendPass:        true,
		KeyringBackendTest:        true,
		KeyringBackendKwallet:     true,
	}

	if _, ok := existingKeyringBackend[KeyringBackend(kb)]; ok {
		return KeyringBackend(kb), nil
	}
	return KeyringBackendUnspecified, fmt.Errorf("unrecognized keyring backend: %s", kb)
}
