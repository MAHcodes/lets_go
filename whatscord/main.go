package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var (
	wp                 *whatsmeow.Client
	dg                 *discordgo.Session
	DISCORD_BOT_ID     string
	DISCORD_CHANNEL_ID string
	DISCORD_TOKEN      string
)

func HandleMessageEvent(v *events.Message) {
	isGroup := v.Info.IsGroup
	sender := v.Info.Sender.User
	senderName := v.Info.PushName
	conv := v.Message.GetConversation()
	msg := fmt.Sprintf("%s - %s\n%s", sender, senderName, conv)
	if !isGroup {
		if v.Info.Type == "text" {
			dg.ChannelMessageSend(DISCORD_CHANNEL_ID, msg)
		}
	}
}

func SendMessage(dist string, msg *string) {
	wp.SendMessage(context.Background(),
		types.JID{User: dist, Server: "s.whatsapp.net"},
		&proto.Message{Conversation: msg})
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		HandleMessageEvent(v)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ReferencedMessage != nil {
		referencedMessageContent := m.ReferencedMessage.Content
		lines := strings.Split(referencedMessageContent, "\n")
		firstLine := lines[0]
		sender := strings.Split(firstLine, " ")[0]
		SendMessage(sender, &m.Content)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	DISCORD_BOT_ID = os.Getenv("DISCORD_BOT_ID")
	DISCORD_CHANNEL_ID = os.Getenv("DISCORD_CHANNEL_ID")

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	wp = whatsmeow.NewClient(deviceStore, clientLog)
	wp.AddEventHandler(eventHandler)

	// Create a new Discord session using the provided bot token.
	dg, err = discordgo.New("Bot " + DISCORD_TOKEN) // TODO: use env for this
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	if wp.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := wp.GetQRChannel(context.Background())
		err = wp.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = wp.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	wp.Disconnect()
	dg.Close()
}
