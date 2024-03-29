package events

import (
	"log"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var eventId string

func init() {
	EventsDeleteCmd.Flags().StringVar(&eventId, "eventId", "", "Event ID")
}

var EventsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete event by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := calendar.NewClient()
		err := client.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}

		client.DeleteEvent(eventId)
	},
}
