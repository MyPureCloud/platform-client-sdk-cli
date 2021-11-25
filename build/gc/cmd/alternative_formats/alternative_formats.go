package alternative_formats

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/spf13/cobra"
	"strings"
)

var validFormats = [...]string{"JSON", "YAML"}

func isValidDataFormat(formatName string) bool {
	for _, v := range validFormats {
		if strings.EqualFold(v, formatName) {
			return true
		}
	}

	fmt.Printf("The data format you entered is either invalid or not supported. The supported formats are: ")
	for i, v := range validFormats {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s", v)
	}
	fmt.Print("\n")
	return false
}

var alternativeFormatsCmd = &cobra.Command{
	Use:   "alternativeformats",
	Short: "Used to specify the desired input and output formats",
	Long:  `Used to specify the desired input and output formats`,
}

var getInputFormatCmd = &cobra.Command{
	Use:   "getinput",
	Short: "Gets the configured input data format.",
	Long:  `Gets the configured input data format.`,

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		inputFormat, _ := config.GetInputFormat(profileName)
		if inputFormat != "" {
			fmt.Printf("Input format: %s\n", strings.ToUpper(inputFormat))
		} else {
			fmt.Println("No configuration set. (Defaulting to JSON)")
		}
	},
}

var getOutputFormatCmd = &cobra.Command{
	Use:   "getoutput",
	Short: "Gets the configured output data format.",
	Long:  `Gets the configured output data format.`,

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		outputFormat, _ := config.GetOutputFormat(profileName)
		if outputFormat != "" {
			fmt.Printf("Output format: %s\n", strings.ToUpper(outputFormat))
		} else {
			fmt.Println("No configuration set. (Defaulting to JSON)")
		}
	},
}

var setOutputFormatCmd = &cobra.Command{
	Use:   "setoutput [format]",
	Short: "Pins the specified output format in the configuration file.",
	Long:  `Pins the specified output format in the configuration file.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if isValidDataFormat(args[0]) {
			profileName, _ := cmd.Root().Flags().GetString("profile")
			err := config.SetOutputFormat(profileName, args[0])
			if err != nil {
				logger.Fatal(err)
			}
		}
	},
}

var setInputFormatCmd = &cobra.Command{
	Use:   "setinput [format]",
	Short: "Pins the specified input format in the configuration file.",
	Long:  `Pins the specified input format in the configuration file.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if isValidDataFormat(args[0]) {
			profileName, _ := cmd.Root().Flags().GetString("profile")
			err := config.SetInputFormat(profileName, args[0])
			if err != nil {
				logger.Fatal(err)
			}
		}
	},
}

func Cmdalternative_formats() *cobra.Command {
	alternativeFormatsCmd.AddCommand(setInputFormatCmd)
	alternativeFormatsCmd.AddCommand(setOutputFormatCmd)
	alternativeFormatsCmd.AddCommand(getInputFormatCmd)
	alternativeFormatsCmd.AddCommand(getOutputFormatCmd)
	return alternativeFormatsCmd
}
