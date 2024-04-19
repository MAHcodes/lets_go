package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	_ "github.com/gopxl/beep"
	"github.com/robfig/cron"
)

type Prayers struct {
	Emsak    string `json:"emsak"`
	Fajer    string `json:"fajer"`
	Shorook  string `json:"shorook"`
	Dohor    string `json:"dohor"`
	Aser     string `json:"aser"`
	Moghreb  string `json:"moghreb"`
	Ishaa    string `json:"ishaa"`
	Midnight string `json:"midnight"`
}

type Command string

const (
	HelpCmd     Command = "help"
	VersionCmd  Command = "version"
	CompleteCmd Command = "complete"
	AlarmCmd    Command = "alarm"
	NextCmd     Command = "next"
	AllCmd      Command = "all"
	EmsakCmd    Command = "emsak"
	FajerCmd    Command = "fajer"
	ShorookCmd  Command = "shorook"
	DohorCmd    Command = "dohor"
	AserCmd     Command = "aser"
	MoghrebCmd  Command = "moghreb"
	IshaaCmd    Command = "ishaa"
	MidnightCmd Command = "midnight"
)

var commands = []Command{HelpCmd, VersionCmd, CompleteCmd, AlarmCmd, NextCmd, AllCmd, EmsakCmd, FajerCmd, ShorookCmd, DohorCmd, AserCmd, MoghrebCmd, IshaaCmd, MidnightCmd}

const Endpoint = "https://almanar.com.lb/ajax/prayer-times.php"

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchPrayer() (*Prayers, error) {
	resp, err := http.Get(Endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prayer Prayers
	json.Unmarshal(body, &prayer)
	return &prayer, nil
}

func getTime(cmd Command) string {
	prayer, err := fetchPrayer()
	handle(err)

	switch cmd {
	case EmsakCmd:
		return prayer.Emsak
	case FajerCmd:
		return prayer.Fajer
	case ShorookCmd:
		return prayer.Shorook
	case DohorCmd:
		return prayer.Dohor
	case AserCmd:
		return prayer.Aser
	case MoghrebCmd:
		return prayer.Moghreb
	case IshaaCmd:
		return prayer.Ishaa
	case MidnightCmd:
		return prayer.Midnight
	default:
		return "Invalid Command"
	}
}

func handleCmd(cmd Command) string {
	switch cmd {
	case EmsakCmd, FajerCmd, ShorookCmd, DohorCmd, AserCmd, MoghrebCmd, IshaaCmd, MidnightCmd:
		return getTime(cmd)
	case NextCmd:
		return next().String()
	case AlarmCmd:
		return alarmNext()
	case AllCmd:
		return all()
	case VersionCmd:
		return version()
	case CompleteCmd:
		return complete()
	default:
		return help()
	}
}

func help() string {
	message := "Usage:\n\twhen <command>\nCommands:"
	for _, command := range commands {
		message += fmt.Sprintf("\n\t%s", command)
	}
	return message
}

func parseTime(prayerTime string) time.Time {
	t, err := time.Parse("15:04", prayerTime)
	handle(err)

	year, month, day := time.Now().Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 0, 0, time.Local)
}

type Prayer struct {
	Name string
	When string
}

func (p Prayer) String() string {
	return fmt.Sprintf(" %-9s: %s", p.Name, p.When)
}

func (p Prayer) isEmpty() bool {
	return p == Prayer{}
}

func next() Prayer {
	prayer, err := fetchPrayer()
	handle(err)
	currentTime := time.Now()
	nextPrayerDuration := 86400
	v := reflect.ValueOf(*prayer)
	t := v.Type()

	var nextPrayer Prayer

	for i := 0; i < v.NumField(); i++ {
		prayName := t.Field(i).Name
		prayTime := v.Field(i).String()
		timeDiff := int(parseTime(prayTime).Sub(currentTime).Seconds())

		if timeDiff >= 0 && timeDiff <= nextPrayerDuration {
			nextPrayerDuration = timeDiff
			nextPrayer = Prayer{
				Name: prayName,
				When: prayTime,
			}
		}

		if i == v.NumField()-1 && nextPrayer.isEmpty() {
			nextPrayer = Prayer{
				Name: prayName,
				When: prayTime,
			}
		}
	}

	return nextPrayer
}

func all() (s string) {
	prayer, err := fetchPrayer()
	handle(err)
	v := reflect.ValueOf(*prayer)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		prayName := t.Field(i).Name
		prayTime := v.Field(i).String()
		p := Prayer{
			Name: prayName,
			When: prayTime,
		}
		s += fmt.Sprintf("%s\n", p.String())
	}
	return s
}

func version() string {
	return "v0.0.1"
}

func playAzan() {
	azanPath := "~/Music/azan.mp3"
	_, err := os.Open(azanPath)
	handle(err)
  // TODO:
}

func alarmNext() string {
	c := cron.New()
	prayer, err := fetchPrayer()
	handle(err)

  fajer := strings.Split(prayer.Fajer, ":")
  dohor := strings.Split(prayer.Dohor, ":")
  ishaa := strings.Split(prayer.Ishaa, ":")

  c.AddFunc(fmt.Sprintf("%s %s * * * *", fajer[0], fajer[1]), playAzan)
  fmt.Printf("[NEW ALARM] %s %s * * * *\n", fajer[0], fajer[1])

  c.AddFunc(fmt.Sprintf("%s %s * * * *", dohor[0], dohor[1]), playAzan)
  fmt.Printf("[NEW ALARM] %s %s * * * *\n", dohor[0], dohor[1])

  c.AddFunc(fmt.Sprintf("%s %s * * * *", ishaa[0], ishaa[1]), playAzan)
  fmt.Printf("[NEW ALARM] %s %s * * * *\n", ishaa[0], ishaa[1])

  c.AddFunc("0 0 * * * *", func() {
    fmt.Println("[RESTETTING]")
    c.Stop()
    alarmNext()
  })

	c.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	c.Stop()

	return "\n[ALARMS STOPPED]\n"
}

func complete() string {
	var cmpl []string
	cmpl = append(cmpl, "help", "version", "next", "all", "emsak", "fajer", "shorook", "dohor", "aser", "moghreb", "ishaa", "midnight")
	return fmt.Sprintf("#compdef %s\n\n_arguments -s \\\n%s\n\n",
		"when", strings.Join(cmpl, " "))
}

func main() {
	cmd := "help"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	out := handleCmd(Command(cmd))
	fmt.Println(out)
}
