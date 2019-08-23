package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

type jokeValue struct {
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

type joke struct {
	Type  string    `json:"type"`
	Value jokeValue `json:"value"`
}

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}
var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")

	log.Print("starting joker web server, listening on port :5000")

	http.ListenAndServe(":5000", r)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	var personJSON string
	var p person
	var err error

	// TODO: When certain foreign language unicode characters are unmarshalled into person.Name
	// and person.Surname (from the raw JSON returned by the http request to uinames.com/api),
	// and then that person is used to build the http request to api.icndb.com, the request
	// will fail with a bad request error. Refactor joker to work without specifying a region in
	// the uinames.com/api request.
	if personJSON, err = makeHTTPRequest("http://uinames.com/api?region=england"); err != nil {
		log.Fatalln(fmt.Sprintf("a server side error occurred: %v", err))
	}

	// TODO: If the client runs joker many times in rapid succession, uinames.com/api will flag
	// the (session?, IP?) of the joker web server as a bot and will start sending access denied
	// error html instead of valid JSON. Currently the only fix appears to be to restart the joker
	// web server. Find a solution to this problem that doesn't require a joker web server restart.
	if err = json.Unmarshal([]byte(personJSON), &p); err != nil {
		fmt.Fprintf(w, "uh, houston, we've had a problem: the response from uinames.com/api was not JSON\n")
		log.Printf("response from uinames.com/api was not JSON: %v ", string(personJSON))
		return
	}

	var sb strings.Builder

	sb.WriteString("http://api.icndb.com/jokes/random?firstName=")
	sb.WriteString(p.Name)
	sb.WriteString("&lastName=")
	sb.WriteString(p.Surname)
	sb.WriteString("&limitTo=[nerdy]")

	var jokeJSON string
	var j joke

	if jokeJSON, err = makeHTTPRequest(sb.String()); err != nil {
		log.Fatalln(fmt.Sprintf("a server side error occurred: %v", err))
	}

	if err = json.Unmarshal([]byte(jokeJSON), &j); err != nil {
		fmt.Fprintf(w, "uh, houston, we've had a problem: the response from api.icndb.com/jokes was not JSON\n")
		log.Printf("response from api.icndb.com/jokes was not JSON: %v ", string(jokeJSON))
		return
	}

	fmt.Fprintf(w, j.Value.Joke+"\n")

}

func makeHTTPRequest(url string) (string, error) {

	resp, err := netClient.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return (string(body)), nil
}
