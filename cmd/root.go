package cmd

import (
	"fmt"
	"os"

	"github.com/Sup3r-Us3r/golang-calendar-cli/cmd/agenda"
	"github.com/Sup3r-Us3r/golang-calendar-cli/cmd/events"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "calendar",
		Short: "Your calendar CLI",
	}

	rootCmd.AddCommand(agenda.AgendaCmd)
	rootCmd.AddCommand(events.EventsCmd)

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
