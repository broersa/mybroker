package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/broersa/lora"
	"github.com/broersa/semtech"
	"github.com/broersa/ttnbroker/bll"
	"github.com/broersa/ttnbroker/bllimpl"
	"github.com/broersa/ttnbroker/dalpsql"
	"github.com/gorilla/mux"

	// Database driver

	_ "github.com/lib/pq"
)

var b bll.Bll

func checkerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("TTNBroker is ALIVE")
	c := os.Getenv("TTNBROKER_DB")
	//s, err := sql.Open("postgres", "postgres://user:password@server/ttn?sslmode=require")
	s, err := sql.Open("postgres", c)
	checkerror(err)
	d := dalpsql.New(s)
	/*d.BeginTransaction()
	var i int64
	i, err = d.AddApplication(&dal.Application{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
	d.CommitTransaction()*/
	b = bllimpl.New(d)
	b.RegisterApplication(&bll.Application{})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/RegisterApplication", RegisterApplication).Methods("POST")
	router.HandleFunc("/HasApplication/{name}", HasApplication).Methods("GET")
	router.HandleFunc("/Message", MessageHandler).Methods("POST")
	//log.Fatal(http.ListenAndServeTLS(":4443", "server.pem", "server.key", router))
	log.Fatal(http.ListenAndServe(":4443", router))
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Andre")
}

// RegisterApplication ...
func RegisterApplication(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
}

// HasApplication ...
func HasApplication(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//appeui := vars["appeui"]
	log.Println("OK")
	fmt.Fprintf(w, "OK")
}

// MessageHandler ...
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	contents, err := ioutil.ReadAll(r.Body)
	checkerror(err)
	var message Message
	err = json.Unmarshal(contents, &message)
	checkerror(err)
	data, err := base64.StdEncoding.DecodeString(message.Package.Data)
	checkerror(err)
	mhdr, err := lora.NewMHDRFromByte(data[0])
	checkerror(err)
	if mhdr.IsJoinRequest() {
		appkey := []byte{0x15, 0x4f, 0x94, 0x7b, 0x41, 0xd0, 0x2f, 0x33, 0x96, 0xf9, 0xaf, 0x6b, 0x4d, 0xb1, 0x0d, 0x5f}
		_, err := lora.NewJoinRequestValidated(appkey, data)
		if err != nil {
			if _, ok := err.(*lora.ErrorMICValidationFailed); ok {
				log.Print(err)
				w.Write([]byte("{error: \"" + err.Error() + "\"}"))
			} else {
				checkerror(err)
			}
		} else {
			joinaccept, err := lora.NewJoinAccept(appkey, 0)
			checkerror(err)
			ja, err := joinaccept.Marshal(appkey)
			checkerror(err)
			responsemessage := &ResponseMessage{
				OriginUDPAddrNetwork: message.OriginUDPAddrNetwork,
				OriginUDPAddrString:  message.OriginUDPAddrString,
				Package: semtech.TXPK{
					Tmst: message.Package.Tmst + 5000000,
					Freq: message.Package.Freq,
					RFCh: message.Package.RFCh,
					Powe: 14,
					Modu: message.Package.Modu,
					DatR: message.Package.DatR,
					CodR: message.Package.CodR,
					IPol: true,
					Size: uint16(len(ja) - 4),
					Data: base64.StdEncoding.EncodeToString(ja)}}
			msg, err := json.Marshal(responsemessage)
			checkerror(err)
			w.Write(msg)
		}
	}
}
