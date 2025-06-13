package is04

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handle_index() http.HandlerFunc {
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
			w.Write([]byte(fmt.Sprintf(`{"error": %s}`, strconv.Quote(err.Error()))))
		}
		fmt.Fprintf(w, `%s`, res)
	}
}

func handle_self() http.HandlerFunc{

}
func handle_sources() http.HandlerFunc {

}
func handle_flows() http.HandlerFunc {

}
func handle_devices() http.HandlerFunc {

}
func handle_senders() http.HandlerFunc {

}
func handle_receivers() http.HandlerFunc {

}

func RegisterRoutes(server *http.ServeMux) {
	server.HandleFunc("/", handle_index())
	server.HandleFunc("/self", handle_self())
	server.HandleFunc("/sources", handle_sources())
	server.HandleFunc("/flows", handle_flows())
	server.HandleFunc("/devices", handle_devices())
	server.HandleFunc("/senders", handle_senders())
	server.HandleFunc("/receivers", handle_receivers())
}
