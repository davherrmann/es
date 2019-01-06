package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/davherrmann/es/api"
	"github.com/davherrmann/es/api/resolver"
	"github.com/davherrmann/es/base"
	"github.com/davherrmann/es/service"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		log.Printf("command failed: %s", err)
	}
}

func main() {
	bus := base.NewBus()

	s, err := ioutil.ReadFile("schema.gql")
	if err != nil {
		log.Fatal("error reading graphql schema: " + err.Error())
	}

	schema := graphql.MustParseSchema(string(s), &resolver.Root{
		Query: service.New(bus),
	})
	http.Handle("/graphql", CorsMiddleware(&relay.Handler{Schema: schema}))
	http.HandleFunc("/", api.GraphiQL)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

/*
	bus := base.NewBus()

	ctx := context.Background()
	catering := catering.NewService(bus)

	apply := func(command ...base.Command) {
		for _, command := range command {
			err := catering.Apply(ctx, command)
			check(err)
		}
	}

	apply(
		command.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pommes",
		},
		command.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pizza",
		})

	bus.Publish(ctx, event.MoneyTransferred{
		From:   "A",
		To:     "B",
		Amount: 100,
	})

	apply(
		command.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pommes",
		},
		command.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Pizza",
		},
		command.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "",
		})

	bus.Publish(ctx, event.OrderFrozen{
		Date:  time.Now(),
		Place: "X",
	})

	apply(
		command.OrderFood{
			User:  "A",
			Place: "X",
			Date:  time.Now(),
			Food:  "Ich habe nachträglich was geändert...",
		},
		command.OrderFood{
			User:  "B",
			Place: "X",
			Date:  time.Now(),
			Food:  "Ich habe nachträglich noch was geändert...",
		})
*/

/*
	log.Println("Hello World!! ")

	db, err := sql.Open("postgres", "postgres://root@localhost:26257/es?sslmode=disable")

	log.Println("opened connection")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	rows, err := db.Query("SELECT stream, type FROM event")
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
