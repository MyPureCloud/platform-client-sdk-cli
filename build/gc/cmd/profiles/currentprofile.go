package profiles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"

	"github.com/spf13/cobra"
)

var currentProfileCmd = &cobra.Command{
	Use:   "get [profileName]",
	Short: "Retrieves profile by name. The default profile is returned if no name is provided",
	Long:  "Retrieves profile by name. The default profile is returned if no name is provided",
	Args:  cobra.RangeArgs(0, 1),

	Run: func(cmd *cobra.Command, args []string) {
		var profileName string
		if len(args) == 0 {
			profileName, _ = cmd.Root().Flags().GetString("profile")
		} else {
			profileName = args[0]
		}

		config, err := config.GetConfig(profileName)

		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(config.String())

	},
}
