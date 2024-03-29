package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sup3r-Us3r/golang-calendar-cli/internal/util"
	"github.com/fatih/color"
	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const AGENDA = "YOUR CALENDAR NAME"

var (
	ErrAgendaNotFound = errors.New("agenda not found")
	ErrAddAgenda      = errors.New("error when adding agenda")
	ErrListEventsWeek = errors.New("error list events week")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarID string
}

// NewClient: Creates a new client
func NewClient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("./credentials.json")
	if err != nil {
		log.Fatal("Unable to read credentials JSON\n")
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("Error to create google calendar service: %s\n", err.Error())
	}

	return &Calendar{
		Service: service,
	}
}

// InsertAgendaID: Inserts an existing calendar into the user's calendar list
func (c *Calendar) InsertAgendaID(id string) error {
	calendarListEntry := &gCalendar.CalendarListEntry{
		Id: id,
	}

	_, err := c.Service.CalendarList.Insert(calendarListEntry).Do()
	if err != nil {
		return ErrAddAgenda
	}

	return nil
}

// GetAgendaID: Get calendar ID
func (c *Calendar) GetAgendaID() error {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatal("Error to get agenda list")
	}

	for _, item := range list.Items {
		if item.Summary == AGENDA {
			c.CalendarID = item.Id

			return nil
		}
	}

	return ErrAgendaNotFound
}

type InsertEventData struct {
	Title         string
	Description   string
	Location      string
	DateTimeStart string
	DateTimeEnd   string
}

// InsertEvent: Creates an event
func (c *Calendar) InsertEvent(data InsertEventData) error {
	event := &gCalendar.Event{
		Summary:     data.Title,
		Description: data.Description,
		Location:    data.Location,
		Start: &gCalendar.EventDateTime{
			DateTime: data.DateTimeStart,
			TimeZone: "America/Sao_Paulo",
		},
		End: &gCalendar.EventDateTime{
			DateTime: data.DateTimeEnd,
			TimeZone: "America/Sao_Paulo",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
		// Attendees: []*gCalendar.EventAttendee{
		// 	{Email: ""},
		// },
	}

	event, err := c.Service.Events.Insert(c.CalendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)

	return nil
}

// DeleteEvent: Deletes an event
func (c *Calendar) DeleteEvent(eventId string) error {
	err := c.Service.Events.Delete(c.CalendarID, eventId).Do()
	if err != nil {
		log.Fatalf("Unable to delete event. %v\n", err)
	}

	fmt.Println("Event deleted!")

	return nil
}

// ListTodayEvents: Gets all events of the day
func (c *Calendar) ListTodayEvents() error {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)
	events, err := c.Service.Events.List(c.CalendarID).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrListEventsWeek
	}

	for _, event := range events.Items {
		formatOutput(event)
	}

	return nil
}

// ListWeekEvents: Gets all events of the week
func (c *Calendar) ListWeekEvents() error {
	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)
	events, err := c.Service.Events.List(c.CalendarID).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrListEventsWeek
	}

	for _, event := range events.Items {
		formatOutput(event)
	}

	return nil
}

// formatOutput: Formats the event to be displayed on the user's terminal in a more pleasant way
func formatOutput(event *gCalendar.Event) {
	summary := "-"
	description := "-"
	startDate := "-"

	if event.Summary != "" {
		summary = event.Summary
	}

	if event.Description != "" {
		description = event.Description
	}

	if event.Start != nil {
		startDate = event.Start.Date
	}

	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("ID: %s\n", blue(event.Id))
	fmt.Printf("Summary: %s\n", blue(summary))
	fmt.Printf("Description: %s\n", blue(description))
	fmt.Printf("Status: %s\n", blue(event.Status))
	fmt.Printf("Start Date: %s\n", blue(startDate))
	fmt.Println(util.RepeatString("-", len(event.Id)+3))
}
