package profiles

import (
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profiles",
	Short: "Manages the profiles for the configuration tool",
	Long:  `Manages the profiles for the configuration tool`,
}

func Cmdprofiles() *cobra.Command {
	profileCmd.AddCommand(currentProfileCmd)
	profileCmd.AddCommand(createProfilesCmd)
	return profileCmd
}
