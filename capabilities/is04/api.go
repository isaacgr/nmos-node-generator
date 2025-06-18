package is04

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type IS04Server struct {
	server     *http.Server
	signalChan chan os.Signal
	Node       *Node
}

func (s IS04Server) registerRoutes(server *http.ServeMux) {
	server.HandleFunc("/", s.handleIndex())
	server.HandleFunc("/self", s.handleSelf())
	server.HandleFunc("/devices", s.handleDevices())
}

func (s IS04Server) registerSignals(server *http.Server) {
	go func() {
		s.signalChan = make(chan os.Signal, 1)
		signal.Notify(s.signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-s.signalChan

		err := server.Close()
		if err != nil {
			log.Fatalf("Unable to close HTTP server. Error [%s]", err)
		}
	}()
}

func (s IS04Server) Addr() string {
	return fmt.Sprintf(
		"%s:%d",
		s.Node.Api.Endpoints[0].Host,
		s.Node.Api.Endpoints[0].Port,
	)
}

func (s IS04Server) Start() error {
	mux := http.NewServeMux()
	s.server = &http.Server{
		Addr:    s.Addr(),
		Handler: mux,
	}

	s.registerRoutes(mux)
	s.registerSignals(s.server)

	err := s.server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s IS04Server) Stop() error {
	return nil
}

func (s IS04Server) handleIndex() http.HandlerFunc {
	endpoints := []string{
		"self/",
		"sources/",
		"flows/",
		"devices/",
		"senders/",
		"receivers/",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(endpoints)
		if err != nil {
			fmt.Printf(`Failed to serve base content. Err [%s]`, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(
				fmt.Appendf(
					res,
					fmt.Sprintf(`{"error": %s}`, strconv.Quote(err.Error())),
				),
			)
		}
		fmt.Fprintf(w, `%s`, res)
	}
}

func (s IS04Server) handleSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(s.Node)
		if err != nil {
			fmt.Printf(`Failed to serve /self content. Err [%s]`, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(
				fmt.Appendf(
					res,
					fmt.Sprintf(`{"error": %s}`, strconv.Quote(err.Error())),
				),
			)
		}
		fmt.Fprintf(w, `%s`, res)

	}
}

func (s IS04Server) handleDevices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(s.Node.devices)
		if err != nil {
			fmt.Printf(`Failed to serve /devices content. Err [%s]`, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(
				fmt.Appendf(
					res,
					fmt.Sprintf(`{"error": %s}`, strconv.Quote(err.Error())),
				),
			)
		}
		fmt.Fprintf(w, `%s`, res)

	}
}
