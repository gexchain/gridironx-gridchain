package cli

import (
	"bufio"
	"fmt"

	"github.com/gridironx/gridchain/libs/cosmos-sdk/client/context"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/codec"
	sdk "github.com/gridironx/gridchain/libs/cosmos-sdk/types"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth"
	"github.com/gridironx/gridchain/libs/cosmos-sdk/x/auth/client/utils"
	"github.com/gridironx/gridchain/x/staking/types"
	"github.com/spf13/cobra"
)

// GetCmdEditValidatorCommissionRate gets the edit validator commission rate command
func GetCmdEditValidatorCommissionRate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-validator-commission-rate [commission-rate]",
		Args:  cobra.ExactArgs(1),
		Short: "edit an existing validator commission rate",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(auth.DefaultTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			valAddr := cliCtx.GetFromAddress()

			rate, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return fmt.Errorf("invalid new commission rate: %v", err)
			}

			msg := types.NewMsgEditValidatorCommissionRate(sdk.ValAddress(valAddr), rate)

			// build and sign the transaction, then broadcast to Tendermint
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
