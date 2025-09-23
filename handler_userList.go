package main

import (
	"context"
	"fmt"
)

func handlerUserList(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return err
	}
	current := s.cfg.CurrentUserName // or wherever you read it

	for _, u := range users {
		name := u.Name
		if name == current {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}
