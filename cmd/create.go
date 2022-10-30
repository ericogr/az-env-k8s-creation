package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new environment",
	Long:  `Use this command to create a new AzureDevOps Environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: must also specify a resource like kubernetes")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("project", "p", "", "[required] AzureDevOps project name with organization (ex: myorg/myproject)")
	err := createCmd.MarkPersistentFlagRequired("project")
	if err != nil {
		fmt.Println(err.Error())
	}

	createCmd.PersistentFlags().StringP("name", "n", "", "[required] AzureDevOps environment name")
	err = createCmd.MarkPersistentFlagRequired("name")
	if err != nil {
		fmt.Println(err.Error())
	}

	createCmd.PersistentFlags().StringP("service-connection", "c", "", "[required] AzureDevOps service connection name")
	err = createCmd.MarkPersistentFlagRequired("service-connection")
	if err != nil {
		fmt.Println(err.Error())
	}
}
