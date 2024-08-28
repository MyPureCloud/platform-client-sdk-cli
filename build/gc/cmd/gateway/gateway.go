package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func Cmdgateway() *cobra.Command {
	utils.AddFlag(setGatewayFileCmd.Flags(), "string", "file", "", "gateway configuration in file")
	setGatewayFileCmd.AddCommand(disableCmd)
	return setGatewayFileCmd
}

var setGatewayFileCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Manages the gateway for the CLI",
	Long:  `Manages the gateway for the CLI`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		c, err := config.GetConfig(profileName)

		gateWayConfig := ResolveConfigurationFlag(cmd)
		if err != nil {
			logger.Fatal(err)
		}
		err = config.UpdateGateWayConfiguration(c, gateWayConfig)
		if err != nil {
			logger.Fatal(err)
		}
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "disables gateway",
	Long:  `disables gateway`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Root().Flags().GetString("profile")
		c, err := config.GetConfig(profileName)

		if err != nil {
			logger.Fatal(err)
		}
		err = config.UpdateGateWayConfiguration(c, nil)
		if err != nil {
			logger.Fatal(err)
		}
	},
}

func ResolveConfigurationFlag(cmd *cobra.Command) *config.GateWayConfiguration {
	fileName, err := cmd.Flags().GetString("file")
	if err != nil {
		logger.Fatal(err)
	}

	if fileName != "" {
		return convertToJSON(fileName)
	}

	for _, command := range cmd.Commands() {
		files, _ := command.Flags().GetString("file")
		if files != "" {
			return convertToJSON(files)
		}
	}
	return nil
}

func convertToJSON(fileName string) *(config.GateWayConfiguration) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to open file %s.", fileName), err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	fileContent, _ := ioutil.ReadAll(jsonFile)
	gconf := config.GateWayConfiguration{}
	if err := json.Unmarshal(fileContent, &gconf); err != nil {
		panic(err)
	}
	return &gconf
}
