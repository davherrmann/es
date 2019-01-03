package main

import (
	"context"
	"log"
	"time"

	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/commands"
	"github.com/davherrmann/es/events"
	"github.com/davherrmann/es/services/catering"
	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		log.Printf("command failed: %s", err)
	}
}

func main() {
	bus := base.NewBus()

	ctx := context.Background()
	catering := catering.NewService(bus)

	apply := func(commands ...base.Command) {
		for _, command := range commands {
			err := catering.Apply(ctx, command)
			check(err)
		}
	}

	apply(
		commands.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pommes",
		},
		commands.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pizza",
		})

	bus.Publish(ctx, events.MoneyTransferred{
		From:   "A",
		To:     "B",
		Amount: 100,
	})

	apply(
		commands.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pommes",
		},
		commands.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pizza",
		},
		commands.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "",
		})

	bus.Publish(ctx, events.OrderFrozen{
		Date:  time.Now(),
		Place: "X",
	})

	apply(
		commands.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Ich habe nachträglich was geändert...",
		},
		commands.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Ich habe nachträglich noch was geändert...",
		})
}

/*
	log.Println("Hello World!! ")

	db, err := sql.Open("postgres", "postgres://root@localhost:26257/es?sslmode=disable")

	log.Println("opened connection")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	rows, err := db.Query("SELECT stream, type FROM events")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stream, eventType string

		if err := rows.Scan(&stream, &eventType); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s %s\n", stream, eventType)
	}

	log.Println("done")
*/
