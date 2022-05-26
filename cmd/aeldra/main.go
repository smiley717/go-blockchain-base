package main

import (
	"aeldra/fs"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"
const flagIP = "ip"
const flagPort = "port"

func main() {
	var aeldraCmd = &cobra.Command{
		Use:   "aeldra",
		Short: "Aeldra CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	aeldraCmd.AddCommand(migrateCmd())
	aeldraCmd.AddCommand(versionCmd)
	aeldraCmd.AddCommand(runCmd())
	aeldraCmd.AddCommand(balancesCmd())

	err := aeldraCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
