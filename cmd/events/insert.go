package events

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/calendar"
	"github.com/spf13/cobra"
)

var title, description, location, dateTimeStart, dateTimeEnd string

func init() {
	EventsInsertCmd.Flags().StringVar(&title, "title", "New event", "Event title name")
	EventsInsertCmd.Flags().StringVar(&description, "description", "", "Description about your event")
	EventsInsertCmd.Flags().StringVar(&location, "location", "", "Address where your event will take place")
	EventsInsertCmd.Flags().StringVar(&dateTimeStart, "dateTimeStart", getDate().DateTimeStart, "Start date and time of your event")
	EventsInsertCmd.Flags().StringVar(&dateTimeEnd, "dateTimeEnd", getDate().DateTimeEnd, "End date and time of your event")
}

var EventsInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new event",
	Run: func(cmd *cobra.Command, args []string) {
		client := calendar.NewClient()
		err := client.GetAgendaID()
		if err != nil {
			log.Fatal(err)
		}

		client.InsertEvent(calendar.InsertEventData{
			Title:         title,
			Description:   description,
			Location:      location,
			DateTimeStart: dateTimeStart,
			DateTimeEnd:   dateTimeEnd,
		})
	},
}

type dateTimeFormat struct {
	DateTimeStart string
	DateTimeEnd   string
}

func getDate() dateTimeFormat {
	currentYear, currentMonth, currentDay := time.Now().Date()
	monthFormatted := currentMonth.String()
	dayFormatted := strconv.Itoa(currentDay)

	if currentMonth < 10 {
		monthFormatted = "0" + strconv.Itoa(int(currentMonth))
	}

	if currentDay < 10 {
		dayFormatted = "0" + strconv.Itoa(currentDay)
	}

	dates := dateTimeFormat{
		DateTimeStart: fmt.Sprintf(
			"%v-%v-%vT00:00:00-03:00", currentYear, monthFormatted, dayFormatted,
		),
		DateTimeEnd: fmt.Sprintf(
			"%v-%v-%vT23:59:59-03:00", currentYear, monthFormatted, dayFormatted,
		),
	}

	return dates
}
