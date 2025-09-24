package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Predator792002/WebGator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error:")
		os.Exit(1)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   feedName,
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("ID: %s\nName: %s\nURL: %s\nUserID: %s\nCreatedAt: %s\nUpdatedAt: %s\n",
		feed.ID, feed.Name, feed.Url, feed.UserID, feed.CreatedAt, feed.UpdatedAt)

	return nil
}
