package main

import (
	"fmt"
	"os" // osパッケージをインポート
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv" // godotenvパッケージをインポート
)

func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// 環境変数 "DISCORD_BOT_TOKEN" からトークンを取得
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("環境変数 'DISCORD_BOT_TOKEN' が設定されていません。")
		return
	}

	// 取得したトークンを使ってDiscordとのセッションを作成
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("セッションの作成中にエラーが発生しました:", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("接続中にエラーが発生しました:", err)
		return
	}

	fmt.Println("Botが起動しました。CTRL-Cで終了します。")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

// messageCreate関数は変更なし
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "こんにちは" {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" さん、こんにちは！")
	}
}