package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	/*
		p1 := person{
			First: "Jenny",
		}

		p2 := person{
			First: "James",
		}

		xp := []person{p1, p2}

		bs, err := json.Marshal(xp)
		if err != nil {
			log.Panic(err)
		}
		//  fmt.Println(string(bs)) // till marshal you need to include this line also

		//UNMARSHAL FROM JSON TO GO
		fmt.Println("PRINT JSON", string(bs))

		xp2 := []person{}
		err = json.Unmarshal(bs, &xp2)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("back into a Go data structure", xp2) */

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "jerry",
	}
	err := json.NewEncoder(w).Encode(p1) // NewEncoder returns a new encoder that writes to w.Encode writes the JSON encoding of v to the stream, followed by a newline character.
	if err != nil {
		log.Println("encoded bad data", err)
	}

}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person                              // we need to have person to intending to receive
	err := json.NewDecoder(r.Body).Decode(&p1) // newdecoder that reads the request body as json and decode that intop1
	if err != nil {
		log.Println("decoded bad data", err) // log is the timestamps at what time it happends
	}
	log.Println("person:", p1)

}
