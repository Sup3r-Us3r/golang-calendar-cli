package events

import (
	"log"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "List all week events",
	Run: func(cmd *cobra.Command, args []string) {
		client := calendar.NewClient()
		err := client.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}

		client.ListWeekEvents()
	},
}
