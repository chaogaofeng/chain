package chaincmdrunner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/goldnet/chain/pkg/chaincmd"
	"github.com/pkg/errors"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
)

// Start starts the blockchain.
func (r Runner) Start(ctx context.Context, args ...string) error {
	return r.run(
		ctx,
		runOptions{wrappedStdErrMaxLen: 50000},
		r.chainCmd.StartCommand(args...),
	)
}

// Init inits the blockchain.
func (r Runner) Init(ctx context.Context, moniker string, mnemonic string) error {
	if len(mnemonic) == 0 {
		return r.run(ctx, runOptions{}, r.chainCmd.InitCommand(moniker))
	}
	input := &bytes.Buffer{}
	fmt.Fprintln(input, mnemonic)
	return r.run(ctx, runOptions{}, r.chainCmd.RecoverInitCommand(moniker), step.Write(input.Bytes()))
}

// KV holds a key, value pair.
type KV struct {
	key   string
	value string
}

// NewKV returns a new key, value pair.
func NewKV(key, value string) KV {
	return KV{key, value}
}

// ConfigSet updates configurations for a app.
func (r Runner) ConfigSet(ctx context.Context, kvs ...KV) error {
	for _, kv := range kvs {
		if err := r.run(
			ctx,
			runOptions{},
			r.chainCmd.ConfigSetCommand(kv.key, kv.value),
		); err != nil {
			return err
		}
	}
	return nil
}

var gentxRe = regexp.MustCompile(`(?m)"(.+?)"`)

// Gentx generates a genesis tx carrying a self delegation.
func (r Runner) Gentx(
	ctx context.Context,
	validatorName,
	selfDelegation string,
	options ...chaincmd.GentxOption,
) (gentxPath string, err error) {
	b := &bytes.Buffer{}

	if err := r.run(ctx, runOptions{
		stdout: b,
		stderr: b,
		stdin:  os.Stdin,
	}, r.chainCmd.GentxCommand(validatorName, selfDelegation, options...)); err != nil {
		return "", err
	}

	return gentxRe.FindStringSubmatch(b.String())[1], nil
}

// CollectGentxs collects gentxs.
func (r Runner) CollectGentxs(ctx context.Context) error {
	return r.run(ctx, runOptions{}, r.chainCmd.CollectGentxsCommand())
}

// ValidateGenesis validates genesis.
func (r Runner) ValidateGenesis(ctx context.Context) error {
	return r.run(ctx, runOptions{}, r.chainCmd.ValidateGenesisCommand())
}

// UnsafeReset resets the blockchain database.
func (r Runner) UnsafeReset(ctx context.Context) error {
	return r.run(ctx, runOptions{}, r.chainCmd.UnsafeResetCommand())
}

// ShowNodeID shows node id.
func (r Runner) ShowNodeID(ctx context.Context) (nodeID string, err error) {
	b := &bytes.Buffer{}
	err = r.run(ctx, runOptions{stdout: b}, r.chainCmd.ShowNodeIDCommand())
	nodeID = strings.TrimSpace(b.String())
	return
}

// NodeStatus keeps info about node's status.
type NodeStatus struct {
	ChainID string
}

// Status returns the node's status.
func (r Runner) Status(ctx context.Context) (NodeStatus, error) {
	b := newBuffer()

	if err := r.run(ctx, runOptions{stdout: b, stderr: b}, r.chainCmd.StatusCommand()); err != nil {
		return NodeStatus{}, err
	}

	var chainID string

	data, err := b.JSONEnsuredBytes()
	if err != nil {
		return NodeStatus{}, err
	}

	out := struct {
		NodeInfo struct {
			Network string `json:"network"`
		} `json:"NodeInfo"`
	}{}

	if err := json.Unmarshal(data, &out); err != nil {
		return NodeStatus{}, err
	}

	chainID = out.NodeInfo.Network

	return NodeStatus{
		ChainID: chainID,
	}, nil
}

// BankSend sends amount from fromAccount to toAccount.
func (r Runner) BankSend(ctx context.Context, fromAccount, toAccount, amount string) (string, error) {
	b := newBuffer()
	opt := []step.Option{
		r.chainCmd.BankSendCommand(fromAccount, toAccount, amount),
	}

	if r.chainCmd.KeyringPassword() != "" {
		input := &bytes.Buffer{}
		fmt.Fprintln(input, r.chainCmd.KeyringPassword())
		fmt.Fprintln(input, r.chainCmd.KeyringPassword())
		fmt.Fprintln(input, r.chainCmd.KeyringPassword())
		opt = append(opt, step.Write(input.Bytes()))
	}

	if err := r.run(ctx, runOptions{stdout: b}, opt...); err != nil {
		if strings.Contains(err.Error(), "key not found") || // stargate
			strings.Contains(err.Error(), "unknown address") || // launchpad
			strings.Contains(b.String(), "item could not be found") { // launchpad
			return "", errors.New("account doesn't have any balances")
		}

		return "", err
	}

	txResult, err := decodeTxResult(b)
	if err != nil {
		return "", err
	}

	if txResult.Code > 0 {
		return "", fmt.Errorf("cannot send tokens (SDK code %d): %s", txResult.Code, txResult.RawLog)
	}

	return txResult.TxHash, nil
}

// WaitTx waits until a tx is successfully added to a block and can be queried
func (r Runner) WaitTx(ctx context.Context, txHash string, retryDelay time.Duration, maxRetry int) error {
	retry := 0

	// retry querying the request
	checkTx := func() error {
		b := newBuffer()
		if err := r.run(ctx, runOptions{stdout: b}, r.chainCmd.QueryTxCommand(txHash)); err != nil {
			// filter not found error and check for max retry
			if !strings.Contains(err.Error(), "not found") {
				return backoff.Permanent(err)
			}
			retry++
			if retry == maxRetry {
				return backoff.Permanent(fmt.Errorf("can't retrieve tx %s", txHash))
			}
			return err
		}

		// parse tx and check code
		txResult, err := decodeTxResult(b)
		if err != nil {
			return backoff.Permanent(err)
		}
		if txResult.Code != 0 {
			return backoff.Permanent(fmt.Errorf("tx %s failed: %s", txHash, txResult.RawLog))
		}

		return nil
	}
	return backoff.Retry(checkTx, backoff.WithContext(backoff.NewConstantBackOff(retryDelay), ctx))
}

// Export exports the state of the chain into the specified file
func (r Runner) Export(ctx context.Context, exportedFile string) error {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	if err := r.run(ctx, runOptions{stdout: stdout, stderr: stderr}, r.chainCmd.ExportCommand()); err != nil {
		return err
	}

	// Exported genesis is written on stderr from Cosmos-SDK v0.44.0
	var exportedState []byte
	if stdout.Len() > 0 {
		exportedState = stdout.Bytes()
	} else {
		exportedState = stderr.Bytes()
	}

	// Save the new state
	return os.WriteFile(exportedFile, exportedState, 0644)
}

// EventSelector is used to query events.
type EventSelector struct {
	typ   string
	attr  string
	value string
}

// NewEventSelector creates a new event selector.
func NewEventSelector(typ, addr, value string) EventSelector {
	return EventSelector{typ, addr, value}
}

// Event represents a TX event.
type Event struct {
	Type       string
	Attributes []EventAttribute
	Time       time.Time
}

// EventAttribute holds event's attributes.
type EventAttribute struct {
	Key   string
	Value string
}

// QueryTxEvents queries tx events by event selectors.
func (r Runner) QueryTxEvents(
	ctx context.Context,
	selector EventSelector,
	moreSelectors ...EventSelector,
) ([]Event, error) {
	// prepare the slector.
	var list []string

	eventsSelectors := append([]EventSelector{selector}, moreSelectors...)

	for _, event := range eventsSelectors {
		list = append(list, fmt.Sprintf("%s.%s=%s", event.typ, event.attr, event.value))
	}

	query := strings.Join(list, "&")

	// execute the commnd and parse the output.
	b := newBuffer()

	if err := r.run(ctx, runOptions{stdout: b}, r.chainCmd.QueryTxEventsCommand(query)); err != nil {
		return nil, err
	}

	out := struct {
		Txs []struct {
			Logs []struct {
				Events []struct {
					Type  string `json:"type"`
					Attrs []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"attributes"`
				} `json:"events"`
			} `json:"logs"`
			TimeStamp string `json:"timestamp"`
		} `json:"txs"`
	}{}

	data, err := b.JSONEnsuredBytes()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	var events []Event

	for _, tx := range out.Txs {
		for _, log := range tx.Logs {
			for _, e := range log.Events {
				var attrs []EventAttribute
				for _, attr := range e.Attrs {
					attrs = append(attrs, EventAttribute{
						Key:   attr.Key,
						Value: attr.Value,
					})
				}

				txTime, err := time.Parse(time.RFC3339, tx.TimeStamp)
				if err != nil {
					return nil, err
				}

				events = append(events, Event{
					Type:       e.Type,
					Attributes: attrs,
					Time:       txTime,
				})
			}
		}
	}

	return events, nil
}
