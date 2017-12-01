package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/go-nats"
)

type server struct {
	nc *nats.Conn
}

func (s server) baseRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basic NATS based microservice example v0.0.1")
}

func (s server) createTask(w http.ResponseWriter, r *http.Request) {
	requestAt := time.Now()
	response, err := s.nc.Request("tasks", []byte("help please"), 5*time.Second)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}
	duration := time.Since(requestAt)

	fmt.Fprintf(w, "Task scheduled in %+v\nResponse: %v\n", duration, string(response.Data))
}

func (s server) healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	fmt.Println("starting")
	go natsSender()
	go natsListener()
	// waits for a key to be pressed
	fmt.Scanln()
}

func natsSender() {
	var s server
	var err error
	//uri := os.Getenv("NATS_URI")
	uri := "nats://172.18.0.2:4222"

	for i := 0; i < 5; i++ {
		nc, err := nats.Connect(uri)
		if err == nil {
			s.nc = nc
			break
		}

		fmt.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}

	fmt.Println("Connected to NATS at:", s.nc.ConnectedUrl())
	http.HandleFunc("/", s.baseRoot)
	http.HandleFunc("/createTask", s.createTask)
	http.HandleFunc("/healthz", s.healthz)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func natsListener() {

	//uri := os.Getenv("NATS_URI")
	uri := "nats://172.18.0.2:4222"
	var err error
	var nc *nats.Conn

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(uri)
		if err == nil {
			break
		}

		fmt.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	fmt.Println("Connected to NATS at:", nc.ConnectedUrl())
	nc.Subscribe("tasks", func(m *nats.Msg) {
		fmt.Println("Got task request on:", m.Subject)
		nc.Publish(m.Reply, []byte("Done!"))
	})

	fmt.Println("Worker subscribed to 'tasks' for processing requests...")
	fmt.Println("Server listening on port 8181...")

	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "OK")
}
