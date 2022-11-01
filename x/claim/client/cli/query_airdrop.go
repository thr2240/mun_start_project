package cli

import (
	"strconv"
	"fmt"
	"strings"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"mun/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdAirdrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "airdrop [airdrop-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the specific airdrop",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for the specific airdrop.

Example:
$ %s query %s airdrop 1
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			airdropId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAirdropRequest{
				AirdropId: airdropId,
			}

			resp, err := queryClient.Airdrop(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
