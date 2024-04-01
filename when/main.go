package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"
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

var commands = []Command{HelpCmd, VersionCmd, NextCmd, AllCmd, EmsakCmd, FajerCmd, ShorookCmd, DohorCmd, AserCmd, MoghrebCmd, IshaaCmd, MidnightCmd}

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
		next()
		return ""
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
	return t
}

type Prayer struct {
	Name string
	When string
}

func (p Prayer) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.When)
}

func next() string {
	prayer, err := fetchPrayer()
	handle(err)

	currentTime := time.Now()

	nextPrayerDuration := 86400

	v := reflect.ValueOf(*prayer)
	t := v.Type()

	var nextPrayer Prayer

	for i := 0; i < v.NumField(); i++ {
		fieldTime := v.Field(i).String()
		timeDiff := currentTime.Sub(parseTime(fieldTime))
    fmt.Println("time diff: ", timeDiff.Seconds())
    fmt.Println(int(timeDiff.Seconds()), nextPrayerDuration)
		if int(timeDiff.Seconds()) <= nextPrayerDuration {
			nextPrayerDuration = int(timeDiff)
			nextPrayer = Prayer{
				Name: t.Field(i).Name,
				When: fieldTime,
			}
		}
	}

	return nextPrayer.String()
}

func main() {
	cmd := "help"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	out := handleCmd(Command(cmd))
	fmt.Println(out)
}
