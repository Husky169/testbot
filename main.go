package main

import (
	"BoostTool/Core/Bot"
	"BoostTool/Core/Discord"
	b "BoostTool/Core/Keyauth"
	"BoostTool/Core/Utils"
	"errors"
	"os"
	"time"

	title "github.com/lxi1400/GoTitle"
)

var config, _ = Utils.LoadConfig()

// * KeyAuth Application Details *//
var name = "boost bot"
var ownerid = "bbV2t9c90P"
var version = "4.1"

func CheckConfig() error {
	proxyStat, _ := os.Stat("./Data/proxies.txt")
	if proxyStat.Size() == 0 {
		return errors.New("Proxies.txt File is Empty, Please add Proxies!")
	}

	if config.License == "" {
		return errors.New("No License Provided.")
	}

	if config.CapService == "" || config.CapKey == "" {
		return errors.New("No Captcha Service or Captcha Key Provided.")
	}

	if config.DiscordSettings.Token == "" {
		return errors.New("No Discord Bot Token Provided.")
	}

	if config.DiscordSettings.GuildID == "" {
		return errors.New("No Guild ID Provided.")
	}

	if len(config.DiscordSettings.Owners) == 0 {
		return errors.New("No Owner ID(s) Provided.")
	}

	if config.DiscordSettings.EmbedColor == "" {
		return errors.New("No Embed Color Provided.")
	}

	if config.DiscordSettings.LogsChannel == "" {
		return errors.New("No Logs Channel Provided.")
	}

	return nil
}

func main() {

	Utils.ClearScreen()
	Utils.LogInfo("Initializing Checks (This may take a moment)", "", "")
	err := CheckConfig()
	if err != nil {
		Utils.LogError(err.Error(), "", "")
		time.Sleep(time.Second * 5)
		os.Exit(0)
	}
	time.Sleep(time.Second * 3)
	Utils.LogInfo("Logging in.", "", "")
	time.Sleep(time.Second * 3)

	go func() {
		for {

			b.Api(name, ownerid, version) // Important to set up the API Details
			b.Init()
			b.License(config.License)
			time.Sleep(time.Minute * 30)
		}
	}()

	title.SetTitle("Boost Bot V3 | discord.gg/boostsociety | t.me/boostsociety | myvouch.es/bloodonhershirt")
	Utils.ClearScreen()
	Utils.PrintASCII()

	if config.CustomPersonalization.Onliner {
		Utils.LogInfo("Token Onliner: Enabled", "", "")
		go Discord.Websocket()
	} else {
		Utils.LogInfo("Token Onliner: Disabled", "", "")
	}

	go Bot.Automation()
	Bot.StartBot()
}
