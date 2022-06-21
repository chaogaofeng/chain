package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/goldnet/chain/pkg/cacmd"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/tempfile"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
)

// https://github.com/vulnsystem/OpenssLabs
// https://github.com/bitpay/bitpay-go
// openssl ec -in root_key.pem -noout -text
// openssl x509 -in root_cert.pem -noout -text
func caCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "ca",
		Short: "tools for ca certificate",
	}

	c.AddCommand(GenRootCommand())
	c.AddCommand(GenRootCertCommand())
	c.AddCommand(IssueCertCommand())
	c.AddCommand(GenKeyCommand())

	return c
}

var (
	FlagRootKey  = "root_key"
	FlagRootCert = "root_cert"
	FlagKey      = "key"
	FlagCer      = "cer"
	FlagCert     = "cert"
	FlagDays     = "days"
	FlagSubject  = "subject"
	FlagType     = "type"
	FlagOutFile  = "out_file"
)

// GenRootCommand returns a command that sets the root cert.
func GenRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-root",
		Short: "generate self-sign root key & certificate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)
			subject, _ := cmd.Flags().GetString(FlagSubject)
			days, _ := cmd.Flags().GetInt(FlagDays)
			rootKey, _ := cmd.Flags().GetString(FlagRootKey)
			rootCert, _ := cmd.Flags().GetString(FlagRootCert)

			return cacmd.GenRootCert(algo, rootKey, rootCert, subject, strconv.Itoa(days))
		},
	}

	cmd.Flags().String(FlagRootKey, "root_key.pem", "file of certificate's private key")
	cmd.Flags().String(FlagRootCert, "root_cert.pem", "file of certificate")
	cmd.Flags().String(FlagSubject, "/C=CN/ST=root/L=root/O=root/OU=root/CN=root", "certificate subject")
	cmd.Flags().Int(FlagDays, 3650, "certificate expire days")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Sm2Type), "certificate algorithm")
	return cmd
}

// GenRootCertCommand returns a command that sets the root cert.
func GenRootCertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-root-cert [root_key]",
		Short: "generate self-sign root certificate from the root key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)
			subject, _ := cmd.Flags().GetString(FlagSubject)
			days, _ := cmd.Flags().GetInt(FlagDays)
			rootCert, _ := cmd.Flags().GetString(FlagRootCert)

			return cacmd.GenSelfSignCert(algo, args[0], rootCert, subject, strconv.Itoa(days))
		},
	}

	cmd.Flags().String(FlagRootCert, "root_cert.pem", "file of certificate")
	cmd.Flags().String(FlagSubject, "/C=CN/ST=root/L=root/O=root/OU=root/CN=root", "certificate subject")
	cmd.Flags().Int(FlagDays, 3650, "certificate expire days")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Sm2Type), "certificate algorithm")
	return cmd
}

// GenRootCertCommand returns a command that sets the root cert.
func IssueCertCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-cert [key]",
		Short: "issue certificate from the parent certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)
			rootKey, _ := cmd.Flags().GetString(FlagRootKey)
			rootCert, _ := cmd.Flags().GetString(FlagRootCert)
			subject, _ := cmd.Flags().GetString(FlagSubject)
			cer, _ := cmd.Flags().GetString(FlagCer)
			cert, _ := cmd.Flags().GetString(FlagCert)
			days, _ := cmd.Flags().GetInt(FlagDays)

			if err := cacmd.GenCertRequest(algo, args[0], cer, subject); err != nil {
				return err
			}
			return cacmd.IssueCert(algo, cer, rootCert, rootKey, cert, strconv.Itoa(days))
		},
	}

	cmd.Flags().String(FlagRootKey, "root_key.pem", "root key file path")
	cmd.Flags().String(FlagRootCert, "root_cert.pem", "root cert file path")
	cmd.Flags().String(FlagCer, "cer.pem", "cer file path")
	cmd.Flags().String(FlagCert, "cert.pem", "cert file path")
	cmd.Flags().String(FlagSubject, "/C=CN/ST=test/L=test/O=test/OU=test/CN=test", "cert subject")
	cmd.Flags().Int(FlagDays, 3650, "expire days")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Sm2Type), "Key algorithm to generate certificate for")

	return cmd
}

// GenKeyCommand returns a command that generates the key from priv_validator_key.json or node_key.json
// Will used to generate CA request.
func GenKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-key",
		Short:   "generate ca private key from the private key",
		Args:    cobra.NoArgs,
		PreRunE: preCheckCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			keyType, _ := cmd.Flags().GetString(FlagType)
			outFile, _ := cmd.Flags().GetString(FlagOutFile)

			var privKey crypto.PrivKey
			if keyType := strings.TrimSpace(keyType); keyType == "node" {
				nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
				if err != nil {
					return err
				}
				privKey = nodeKey.PrivKey
			} else if keyType == "validator" {
				filePv := privval.LoadFilePV(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())
				privKey = filePv.Key.PrivKey
			} else if keyType == "account" {
				if clientCtx.Keyring == nil {
					return errors.New("keyring must be set")
				}
				priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
				if err != nil {
					return err
				}
				privKey, err = cacmd.ToTmPrivKeyInterface(priv)
				if err != nil {
					return err
				}
			}

			key, err := cacmd.Genkey(privKey)
			if err != nil {
				return err
			}
			return tempfile.WriteFileAtomic(outFile, key, 0600)
		},
	}

	cmd.Flags().String(flags.FlagKeyringBackend, "test", "Select keyring's backend (os|file|kwallet|pass|test|memory)")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(FlagType, "account", "key type (node|validator|account)")
	cmd.Flags().String(FlagOutFile, "key.pem", "ca private key file path")

	return cmd
}

func preCheckCmd(cmd *cobra.Command, _ []string) error {
	flags := cmd.Flags()
	if flags.Changed(FlagType) {
		keyType, _ := flags.GetString(FlagType)
		if keyType != "node" && keyType != "validator" && keyType != "account" {
			return fmt.Errorf("key type must be node or validator or account")
		}
	}
	return nil
}
