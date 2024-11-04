package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Actor struct {
	ID int 					`json:"id"`
	Login string 			`json:"login"`
	DisplayLogin string 	`json:"display_login"`
	URL string 				`json:"url"`
	AvatarURL string 		`json:"avatar_url"`
}

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
}

type Commit struct {
    SHA     string `json:"sha"`
    Message string `json:"message"`
}

type Payload struct {
    PushID  int64    `json:"push_id"`
    Ref     string   `json:"ref"`
    Commits []Commit `json:"commits"`
}

type Event struct {
    ID        string    `json:"id"`
    Type      string    `json:"type"`
    Actor     Actor     `json:"actor"`
    Repo      Repo      `json:"repo"`
    Payload   Payload   `json:"payload"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
	var rootCmd = &cobra.Command{
		Use: "github-user-activity",
		Short: "Tracks the user activity of the user",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]

			var events []Event

			var url = "https://api.github.com/users/" + username + "/events"
			resp, err := http.Get(url)

			if err != nil {
				log.Fatal("Failed to fetch data:", err)
			}
			defer resp.Body.Close()

            if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
                log.Fatal(err)
            }
            for _, event := range events {
                fmt.Printf("Event ID: %s\n", event.ID)
                fmt.Printf("Type: %s\n", event.Type)
                fmt.Printf("Actor: %s\n", event.Actor.Login)
                fmt.Printf("Repo: %s\n", event.Repo.Name)
                if len(event.Payload.Commits) > 0 {
                    fmt.Printf("Latest Commit: %s\n", event.Payload.Commits[0].Message)
                }
                fmt.Printf("Created At: %s\n", event.CreatedAt.Format(time.RFC822))
                fmt.Println("----------------------------------------")
            }
		},
	}

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}