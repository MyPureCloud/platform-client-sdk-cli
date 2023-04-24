package proxy

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

func Cmdproxy() *cobra.Command {
        utils.AddFlag(setProxyFileCmd.Flags(), "string", "file", "", "proxy configuration in file")
		setProxyFileCmd.AddCommand(disableCmd)
        return setProxyFileCmd
}

var setProxyFileCmd = &cobra.Command{
        Use:   "proxy",
        Short: "Manages the proxy for the CLI",
        Long:  `Manages the proxy for the CLI`,
        Args:  cobra.NoArgs,

        Run: func(cmd *cobra.Command, args []string) {
                profileName, _ := cmd.Root().Flags().GetString("profile")
                c, err := config.GetConfig(profileName)

                proxyconfig := ResolveConfigurationFlag(cmd)
                if err != nil {
                        logger.Fatal(err)
                }
                err = config.UpdateProxyConfiguration(c, proxyconfig)
                if err != nil {
                        logger.Fatal(err)
                }
        },
}

var disableCmd = &cobra.Command{
		Use:   "disable",
		Short: "disables proxy",
		Long:  `disables proxy`,
		Args:  cobra.NoArgs,

		Run: func(cmd *cobra.Command, args []string) {
				profileName, _ := cmd.Root().Flags().GetString("profile")
				c, err := config.GetConfig(profileName)

				if err != nil {
						logger.Fatal(err)
				}
				err = config.UpdateProxyConfiguration(c, nil)
                if err != nil {
                        logger.Fatal(err)
                }
		},
}

func ResolveConfigurationFlag(cmd *cobra.Command) *(config.ProxyConfiguration) {
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

func convertToJSON(fileName string) *(config.ProxyConfiguration) {
        jsonFile, err := os.Open(fileName)
        if err != nil {
                logger.Fatal(fmt.Sprintf("Unable to open file %s.", fileName), err)
        }

        // defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        fileContent, _ := ioutil.ReadAll(jsonFile)
        proxyconf := config.ProxyConfiguration{}
        if err := json.Unmarshal(fileContent, &proxyconf); err != nil {
                panic(err)
        }
        return &proxyconf
}