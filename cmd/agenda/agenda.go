package agenda

import (
	"fmt"
	"log"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "Add agenda ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := calendar.NewClient()
		err := client.InsertAgendaID(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Agenda has added with success")
	},
}
