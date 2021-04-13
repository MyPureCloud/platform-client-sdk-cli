package logging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"

	"github.com/spf13/cobra"
)

var loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "Manages the logging for the CLI",
	Long:  `Manages the logging for the CLI`,
}

func Cmdlogging() *cobra.Command {
	loggingCmd.AddCommand(setLogFilePathCmd)
	loggingCmd.AddCommand(enableCmd)
	loggingCmd.AddCommand(disableCmd)
	return loggingCmd
}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables logging",
	Long:  `Enables logging`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		setLogging(cmd, true)
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables logging",
	Long:  `Disables logging`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		setLogging(cmd, false)
	},
}

var setLogFilePathCmd = &cobra.Command{
	Use:   "path [filePath]",
	Short: "Sets the log file path",
	Long:  `Sets the log file path`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		c, err := config.GetConfig(profileName)
		if err != nil {
			logger.Fatal(err)
		}
	
		err = config.UpdateLogFilePath(c, args[0])
		if err != nil {
			logger.Fatal(err)
		}
	},
}

func setLogging(cmd *cobra.Command, loggingEnabled bool) {
	profileName, _ := cmd.Root().Flags().GetString("profile")
	c, err := config.GetConfig(profileName)
	if err != nil {
		logger.Fatal(err)
	}

	err = config.SetLoggingEnabled(c, loggingEnabled)
	if err != nil {
		logger.Fatal(err)
	}
}