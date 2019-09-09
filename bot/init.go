package bot

import (
    "github.com/andersfylling/disgord"
    "github.com/andersfylling/disgord/std"
    "github.com/spf13/viper"
    "log"
)

// replyPongToPing is a handler that replies pong to ping messages
func replyPongToPing(s disgord.Session, data *disgord.MessageCreate) {
    msg := data.Message
    
    // whenever the message written is "ping", the bot replies "pong"
    if msg.Content == "ping" {
        msg.Reply(s, "pong")
    }
}

func Initialize() {
    log.Println(viper.AllSettings())
    client := disgord.New(&disgord.Config{
        BotToken: viper.GetString("botToken"),
        Logger:   disgord.DefaultLogger(viper.GetBool("verbose")), // debug=false
    })
    defer client.StayConnectedUntilInterrupted()
    
    discordLog, _ := std.NewLogFilter(client)
    filter, _ := std.NewMsgFilter(client)
    filter.SetPrefix(viper.GetString("botPrefix"))
    
    // create a handler and bind it to new message events
    // tip: read the documentation for std.CopyMsgEvt and understand why it is used here.
    client.On(disgord.EvtMessageCreate,
        // middleware
        filter.NotByBot,    // ignore bot messages
        filter.HasPrefix,   // read original
        discordLog.LogMsg,  // log command message
        std.CopyMsgEvt,     // read & copy original
        filter.StripPrefix, // write copy
        // handler
        replyPongToPing) // handles copy
}