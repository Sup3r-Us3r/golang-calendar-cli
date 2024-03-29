package events

import (
	"log"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "List all today events",
	Run: func(cmd *cobra.Command, args []string) {
		client := calendar.NewClient()
		err := client.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}

		client.ListTodayEvents()
	},
}
