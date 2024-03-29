package events

import (
	"github.com/spf13/cobra"
)

func init() {
	EventsCmd.AddCommand(
		EventsInsertCmd,
		EventsDeleteCmd,
		EventsTodayCmd,
		EventsWeekCmd,
	)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Manage your events",
	Long:  "Insert new events, delete event or check all your events by day or week",
}
