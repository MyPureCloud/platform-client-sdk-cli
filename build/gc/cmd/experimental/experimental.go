package experimental

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/experimental_features"
	"github.com/spf13/cobra"
)

var experimentalCmd = &cobra.Command{
	Use:   "experimental",
	Short: "Manages the experimental features for the CLI",
	Long:  `Manages the experimental features for the CLI`,
}

func Cmdexperimental() *cobra.Command {
	experimentalCmd.AddCommand(enableCmd)
	experimentalCmd.AddCommand(disableCmd)
	experimentalCmd.AddCommand(listExperimentalFeaturesCmd)
	return experimentalCmd
}

var enableCmd = &cobra.Command{
	Use:   "enable [feature]",
	Short: "Enables specified experimental feature",
	Long:  `Enables specified experimental feature`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		setFeature(cmd, args[0], true)
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable [feature]",
	Short: "Disables specified experimental feature",
	Long:  `Disables specified experimental feature`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		setFeature(cmd, args[0], false)
	},
}

var listExperimentalFeaturesCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all experimental features",
	Long:  `Lists all experimental features`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		experimental_features.ListAllFeatures(profileName)
	},
}

func setFeature(cmd *cobra.Command, featureCommand string, enabled bool) {
	profileName, _ := cmd.Root().Flags().GetString("profile")

	if !experimental_features.IsValidCommand(featureCommand) {
		fmt.Printf("The feature '%v' either does not exist, or is no longer experimental.\n", featureCommand)
		return
	}

	err := config.SetExperimentalFeature(profileName, featureCommand, enabled)
	if err != nil {
		fmt.Println(err)
	}
}
