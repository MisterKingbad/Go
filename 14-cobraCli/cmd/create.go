/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"taxgo/14-cobraCli/internal/database"

	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples`,
		RunE: runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _:= cmd.Flags().GetString("name")
		description, _:= cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategory(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name","n","","Name of the category")
	createCmd.Flags().StringP("description","d","","Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
