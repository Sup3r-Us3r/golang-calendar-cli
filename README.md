# Golang Calendar CLI

This application is a CLI that was written in Golang using Cobra, it communicates with your Google calendar, allowing you to create, display and delete events via the command line.

## Technologies

- [x] [Golang](https://go.dev)
- [x] [Cobra](https://github.com/spf13/cobra)
- [x] [Google Calendar API](https://developers.google.com/calendar/api/guides/overview)

<details>
<summary>GCP Setup</summary>

Follow the steps below to configure your application on GCP, as we need to create a project in it and activate the Calendar API.

Go to Google Cloud Platform here: [link](https://console.cloud.google.com)

### Step 1

Create a new project

![gcp-step-1](.github/docs/gcp/1.png)

### Step 2

Define project name

![gcp-step-2](.github/docs/gcp/2.png)

### Step 3

Access Google Calendar API page

![gcp-step-3](.github/docs/gcp/3.png)

### Step 4

Enable Google Calendar API

![gcp-step-4](.github/docs/gcp/4.png)

### Step 5

Access credentials page

![gcp-step-5](.github/docs/gcp/5.png)

### Step 6

Access service account option

![gcp-step-6](.github/docs/gcp/6.png)

### Step 7

Service account details

![gcp-step-7](.github/docs/gcp/7.png)

### Step 8

Grant this service account access to project

![gcp-step-8](.github/docs/gcp/8.png)

### Step 9

Click on your created service account

![gcp-step-9](.github/docs/gcp/9.png)

### Step 10

Create new key

![gcp-step-10](.github/docs/gcp/10.png)

### Step 11

Download the JSON file

![gcp-step-11](.github/docs/gcp/11.png)

> Save it in the project root with the name `credentials.json`

</details>

<details>
<summary>Google Calendar Setup</summary>

Follow the steps below to configure your application on Google Calendar.

Go to Google Calendar here: [link](https://calendar.google.com)

### Step 1

Access settings page

![google-calendar-step-1](.github/docs/google-calendar/1.png)

### Step 2

Create a new calendar

![google-calendar-step-2](.github/docs/google-calendar/2.png)
![google-calendar-step-2-1](.github/docs/google-calendar/3.png)

### Step 3

Share with specific people or groups

![google-calendar-step-3](.github/docs/google-calendar/4.png)

> Copy the `client_email` from the credentials file previously downloaded from GCP and add this email in this step.

### Step 4

Copy calendar ID

![google-calendar-step-4](.github/docs/google-calendar/5.png)

> You need to copy the calendar id, as it will be used in the application execution step below.

</details>

<details>
<summary>Execute application</summary>

## Setup app

Make sure the `credentials.json` file is in the root of the project, as instructed in the GCP setup step.

Change the calendar name in [internal/calendar/calendar.go](./internal/calendar/calendar.go) line 17:

```go
const AGENDA = "YOUR CALENDAR NAME"
```

Install dependencies:

```sh
$ go mod download
```

## Execute app

Build application:

```sh
$ go build
```

Create integration:

```sh
$ ./golang-calendar-cli agenda ID
```

> Is the ID copied in the Google Calendar Setup step.

> This action is only necessary once.

### Commands

Create a new event

```sh
$ ./golang-calendar-cli events insert --title "New Event” --description "New Event Description” --location "My home” --dateTimeStart "2024-03-29T09:00:00-03:00" --dateTimeEnd "2024-03-29T17:00:00-03:00"
```

> For more details run: `./golang-calendar-cli events insert -h`

Show events by day

```sh
$ ./golang-calendar-cli events day
```

Show events by week

```sh
$ ./golang-calendar-cli events week
```

Delete event by ID

```sh
$ ./golang-calendar-cli events delete ID
```

> You can get the event ID by viewing events by day or week

</details>
