package autopagination

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
)

var autoPaginationCmd = &cobra.Command{
	Use:   "autopagination",
	Short: "autopagination",
	Long:  `automatic pagination`,
}

var enableAutoPaginationCmd = &cobra.Command{
	Use:   "enable",
	Short: "Permanently enable autopagination.",
	Long:  `Permanently enable autopagination.`,
	Args:  utils.DetermineArgs([]string{}),

	Run: func(cmd *cobra.Command, args []string) {
		const enabled = true
		setAutoPagination(cmd, enabled)
	},
}

var disableAutoPaginationCmd = &cobra.Command{
	Use:   "disable",
	Short: "Permanently disable autopagination.",
	Long:  `Permanently disable autopagination.`,
	Args:  utils.DetermineArgs([]string{}),

	Run: func(cmd *cobra.Command, args []string) {
		const disabled = false
		setAutoPagination(cmd, disabled)
	},
}

var checkAutoPaginationCmd = &cobra.Command{
	Use:   "status",
	Short: "Check autopagination status in config.",
	Long:  "Check autopagination status in config.",
	Args:  utils.DetermineArgs([]string{}),

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		status, err := config.GetAutoPaginationEnabled(profileName)
		if err != nil {
			logger.Fatal(err)
		}
		if status {
			fmt.Print("Auto-pagination is currently enabled.\n")
		} else {
			fmt.Print("Auto-pagination is currently disabled.\n")
		}
	},
}

func setAutoPagination(cmd *cobra.Command, autoPaginationEnabled bool) {
	profileName, _ := cmd.Root().Flags().GetString("profile")
	c, err := config.GetConfig(profileName)
	if err != nil {
		logger.Fatal(err)
	}
	err = config.SetAutoPaginationEnabled(c, autoPaginationEnabled)
	if err != nil {
		logger.Fatal(err)
	}
}

func Cmdautopagination() *cobra.Command {
	autoPaginationCmd.AddCommand(enableAutoPaginationCmd)
	autoPaginationCmd.AddCommand(disableAutoPaginationCmd)
	autoPaginationCmd.AddCommand(checkAutoPaginationCmd)
	return autoPaginationCmd
}
