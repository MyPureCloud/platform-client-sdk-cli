package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"

	"flag"
	"fmt"
	"gc/cmd/divisions"
	"gc/cmd/groups"
	"gc/cmd/locations"
	"gc/cmd/campaigns"
	"gc/cmd/phones"
	"gc/cmd/queues"
	"gc/cmd/sites"
	"gc/cmd/skills"
	"gc/cmd/stations"
	"gc/cmd/edges"
	"gc/cmd/usage"
	"gc/cmd/users"
	"gc/cmd/notifications"
	"gc/cmd/profiles"
	
	"log"
	"os"
)

var profileName string
var output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gc",
	Short: "gc is a CLI for interacting with Genesys Cloud",
	Long:  `gc is a CLI for interacting with Genesys Cloud`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gc",
	Long:  `All software has versions. This is gc version's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.5")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	flag.CommandLine.Parse(nil)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initViperConfig() {
	homeDir, _ := os.UserHomeDir()
	viper.SetConfigName("config")                       // name of config file (without extension)
	viper.SetConfigType("toml")                         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(fmt.Sprintf("%s/.gc", homeDir)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                            // optionally look for config in the working directory

	viper.SetEnvPrefix("gc")
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize()
	initViperConfig()

	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "DEFAULT", "Name of the profile to use for configuring the cli")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(divisions.Cmddivisions())
	rootCmd.AddCommand(groups.Cmdgroups())
	rootCmd.AddCommand(locations.Cmdlocations())
	rootCmd.AddCommand(campaigns.Cmdcampaigns())
	rootCmd.AddCommand(phones.Cmdphones())
	rootCmd.AddCommand(queues.Cmdqueues())
	rootCmd.AddCommand(sites.Cmdsites())
	rootCmd.AddCommand(skills.Cmdskills())
	rootCmd.AddCommand(stations.Cmdstations())
	rootCmd.AddCommand(edges.Cmdedges())
	rootCmd.AddCommand(usage.Cmdusage())
	rootCmd.AddCommand(users.Cmdusers())
	rootCmd.AddCommand(notifications.Cmdnotifications())
	rootCmd.AddCommand(profiles.Cmdprofiles())
	

	if os.Getenv("GenerateGcDocs") != "" {
		if _, err := os.Stat("/tmp/gc_docs"); os.IsNotExist(err) {
			err = os.Mkdir("/tmp/gc_docs", 0755)
			if err != nil {
				log.Fatal(err)
			}
			err = doc.GenMarkdownTree(rootCmd, "/tmp/gc_docs")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
