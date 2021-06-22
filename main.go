package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	frequency   int
	setNickname bool
	token       string
	network string
)

func init() {

	flag.IntVar(&frequency, "frequency", 5, "seconds between gas price cycles")
	flag.BoolVar(&setNickname, "setNickname", false, "wether to set nickname of bot")
	flag.StringVar(&token, "token", "", "discord bot token")
	flag.StringVar(&network, "network", "", "ethereum, binance-smart-chain, or polygon")

	flag.Parse()
}

func main() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

	// Set activity as desc
	if setNickname {
		err = dg.UpdateListeningStatus(fmt.Sprintf("%s gas price", network))
		if err != nil {
			fmt.Printf("Unable to set activity: \n", err)
		} else {
			fmt.Println("Set activity")
		}
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		setNickname = false
	}

	changeFrequency := time.Duration(frequency) * time.Second
	var nickname string

	// watch gas price
	for {

		// get gas prices
		gasPrices, err := GetGasPrices(network)
		if err != nil {
			fmt.Printf("Error getting rates: %s\n", err)
			time.Sleep(changeFrequency)
			continue
		}

		nickname = fmt.Sprintf("Standard: %dgwei", gasPrices.Standard)

		// change nickname
		if setNickname {

			for _, g := range guilds {

				err = dg.GuildMemberNickname(g.ID, "@me", nickname)
				if err != nil {
					fmt.Printf("Error updating nickname: %s\n", err)
					continue
				} else {
					fmt.Printf("Set nickname in %s: %s\n", g.Name, nickname)
				}
			}
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)

		nickname = fmt.Sprintf("Fast: %dgwei", gasPrices.Fast)

		// change nickname
		if setNickname {

			for _, g := range guilds {

				err = dg.GuildMemberNickname(g.ID, "@me", nickname)
				if err != nil {
					fmt.Printf("Error updating nickname: %s\n", err)
					continue
				} else {
					fmt.Printf("Set nickname in %s: %s\n", g.Name, nickname)
				}
			}
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)

		nickname = fmt.Sprintf("Instant: %dgwei", gasPrices.Instant)

		// change nickname
		if setNickname {

			for _, g := range guilds {

				err = dg.GuildMemberNickname(g.ID, "@me", nickname)
				if err != nil {
					fmt.Printf("Error updating nickname: %s\n", err)
					continue
				} else {
					fmt.Printf("Set nickname in %s: %s\n", g.Name, nickname)
				}
			}
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)
	}

}
