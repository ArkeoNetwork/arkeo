package arkeocli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/common/cosmos"
	"github.com/arkeonetwork/arkeo/x/arkeo/types"
)

func newOpenContractCmd() *cobra.Command {
	openContractCmd := &cobra.Command{
		Use:   "open-contract",
		Short: "open a contract",
		Args:  cobra.ExactArgs(0),
		RunE:  runOpenContractCmd,
	}

	flags.AddTxFlagsToCmd(openContractCmd)
	openContractCmd.Flags().String("provider-pubkey", "", "provider pubkey")
	openContractCmd.Flags().String("service", "", "provider service name")
	openContractCmd.Flags().String("client-pubkey", "", "client pubkey")
	openContractCmd.Flags().String("delegate-pubkey", "", "delegate pubkey")
	openContractCmd.Flags().Bool("no-delegate", false, "delegate pubkey")
	openContractCmd.Flags().String("contract-type", "", "contract type (subscription or pay-as-you-go)")
	openContractCmd.Flags().Int64("deposit", 0, "deposit amount")
	openContractCmd.Flags().Int64("duration", 0, "contract duration")
	openContractCmd.Flags().Int64("rate", 0, "contract rate")
	openContractCmd.Flags().Int64("qpm", 0, "queries per minute")
	openContractCmd.Flags().Int64("settlement-duration", 0, "contract settlement duration")
	openContractCmd.Flags().String("contract-authorization", "strict", "contract authorization (strict or open)")
	return openContractCmd
}

func runOpenContractCmd(cmd *cobra.Command, args []string) (err error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	key, err := ensureKeys(cmd)
	if err != nil {
		return err
	}
	addr, err := key.GetAddress()
	if err != nil {
		return
	}
	clientCtx = clientCtx.WithFromName(key.Name).WithFromAddress(addr)
	if err = client.SetCmdClientContext(cmd, clientCtx); err != nil {
		return
	}

	argProviderPubkey, _ := cmd.Flags().GetString("provider-pubkey")
	if argProviderPubkey == "" {
		argProviderPubkey, err = promptForArg(cmd, "Specify provider pubkey: ")
		if err != nil {
			return
		}
	}

	argService, _ := cmd.Flags().GetString("service")
	if argService == "" {
		argService, err = promptForArg(cmd, "Specify service (e.g. gaia-mainnet-rpc-archive, btc-mainnet-fullnode, etc): ")
		if err != nil {
			return err
		}
	}

	argClientPubkey, _ := cmd.Flags().GetString("client-pubkey")
	if argClientPubkey == "" {
		argClientPubkey, err = toPubkey(cmd, addr)
		if err != nil {
			return err
		}
	}

	var argDelegatePubkey string
	argNoDelegate, _ := cmd.Flags().GetBool("no-delegate")
	if !argNoDelegate {
		argDelegatePubkey, _ = cmd.Flags().GetString("delegate-pubkey")
		if argDelegatePubkey == "" {
			argDelegatePubkey, err = promptForArg(cmd, "Specify delegate pubkey (leave blank to use client key): ")
			if err != nil {
				return err
			}
		}
	}

	argContractType, _ := cmd.Flags().GetString("contract-type")
	if argContractType == "" {
		argContractType, err = promptForArg(cmd, "Specify contract type (subscription or pay-as-you-go): ")
		if err != nil {
			return err
		}
	}
	argContractType = strings.ToUpper(strings.ReplaceAll(argContractType, "-", "_"))
	if _, ok := types.ContractType_value[argContractType]; !ok {
		return fmt.Errorf("invalid contract type: %s", argContractType)
	}
	contractType := types.ContractType(types.ContractType_value[argContractType])

	argDuration, _ := cmd.Flags().GetInt64("duration")
	if argDuration == 0 {
		duration, err := promptForArg(cmd, "Specify contract duration (in blocks): ")
		if err != nil {
			return err
		}
		argDuration, err = strconv.ParseInt(duration, 10, 64)
		if err != nil {
			return err
		}
	}

	argRate, _ := cmd.Flags().GetString("rate")
	if len(argRate) == 0 {
		argRate, err = promptForArg(cmd, "Specify rate (must match provider): ")
		if err != nil {
			return err
		}
	}
	rate, err := cosmos.ParseCoin(argRate)
	if err != nil {
		return err
	}

	argDeposit, _ := cmd.Flags().GetInt64("deposit")
	if argDeposit == 0 {
		deposit, err := promptForArg(cmd, "Specify deposit amount (product of rate and duration): ")
		if err != nil {
			return err
		}
		argDeposit, err = strconv.ParseInt(deposit, 10, 64)
		if err != nil {
			return err
		}
	}

	argSettlementDuration, _ := cmd.Flags().GetInt64("settlement-duration")
	if argSettlementDuration == 0 {
		settlementDuration, err := promptForArg(cmd, "Specify settlement duration (in blocks): ")
		if err != nil {
			return err
		}
		argSettlementDuration, err = strconv.ParseInt(settlementDuration, 10, 64)
		if err != nil {
			return err
		}
	}

	argQPM, _ := cmd.Flags().GetInt64("qpm")
	if argQPM == 0 {
		qpm, err := promptForArg(cmd, "Specify queries per minute: ")
		if err != nil {
			return err
		}
		argSettlementDuration, err = strconv.ParseInt(qpm, 10, 64)
		if err != nil {
			return err
		}
	}

	argContractAuth, _ := cmd.Flags().GetString("contract-authorization")
	if argContractAuth == "" {
		argContractAuth, err = promptForArg(cmd, "Specify contract authorization (strict or open): ")
		if err != nil {
			return err
		}
	}
	argContractAuth = strings.ToUpper(strings.ReplaceAll(argContractAuth, "-", "_"))
	if _, ok := types.ContractAuthorization_value[argContractAuth]; !ok {
		return fmt.Errorf("invalid contract authorization: %s", argContractAuth)
	}
	contractAuth := types.ContractAuthorization(types.ContractAuthorization_value[argContractAuth])

	pubkey, err := common.NewPubKey(argProviderPubkey)
	if err != nil {
		return err
	}
	deposit := cosmos.NewInt(argDeposit)
	msg := types.NewMsgOpenContract(
		clientCtx.GetFromAddress(),
		pubkey,
		argService,
		common.PubKey(argClientPubkey),
		common.PubKey(argDelegatePubkey),
		contractType,
		argDuration,
		argSettlementDuration,
		rate,
		deposit,
		contractAuth,
		argQPM,
	)
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}
