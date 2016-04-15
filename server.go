package main

import (
	"log"
	//"fmt"
	//"reflect"
	"strconv"
	//"strings"
	"encoding/json"
	"gopkg.in/go-playground/validator.v8"
	//"github.com/House3272/suggest/dataStructures/sliceDeStrings"
	"github.com/House3272/suggest/dataStructures/trieDeStrings"
	"flag"
	"os"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var validate *validator.Validate

var gPath string = os.Getenv("GOPATH")  //can't be a constant?
const cBase, cMin, cMax uint64 = 20, 10, 50


var muhTrie *trieDeStrings.ZTrie
//var inMemorySlice *[]string



func main() {

	muhTrie = trieDeStrings.NewTrie()
	//inMemorySlice = sliceDeStrings.NewSlice()

	// determine which file to load
	pathToList := flag.String("f", gPath+"/suggest/wikiTitles/semi-all", "path of file to load")
	flag.Parse()
	f, err := os.Open(*pathToList)
	if err != nil {
		panic(err)
	}
	// analyze opened file
	go muhTrie.LoadTrie(f)
	//go sliceDeStrings.MakeSlice(f,inMemorySlice)
	defer f.Close()


	config := &validator.Config{TagName: "valid"}
	validate = validator.New(config)

	router := httprouter.New()
	router.GET("/go/suggest/:qury", resHandler)
	router.GET("/go/suggest/:qury/:cnt", resHandler)

	log.Fatal(http.ListenAndServe(":4321", router))

}


func resHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//if Trie not ready
	if muhTrie.GetCount() < 99999 {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("The Dataset Has Not Been Loaded Yet"))
		return
	}

	//get params of interest
	qString := ps.ByName("qury")
	qCount := ps.ByName("cnt")


	//get matches
	var amount uint64

	//range not specified
	cntCheck := validate.Field(qCount, "required,numeric")
	if cntCheck != nil {
		amount = cBase
	} else {
	//range specified
		qRange, cntParse := strconv.ParseUint(qCount,10,64)
		if cntParse != nil {
		//parse error
			amount = cBase
		} else if qRange <= cMin {
		//range is less than minimum
			amount = cMin
		} else if qRange >= cMax {
		//range is more than maxium
			amount = cMax
		} else {
		//range set to request value
			amount = qRange
		}
	}

	matches := muhTrie.PrefixSearch(qString, amount)
	// if no results
	if len(*matches) < 1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Matching Title..."))
		return
	}

	resultArray, _ := json.Marshal(*matches)
	w.Write(resultArray)
	return
}










// func myFunc(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	fmt.Println(req)
// 	w.Write([]byte("Hello Dude, path / "))
// }

	//fmt.Println(reflect.TypeOf(*f))


	// aLine, err := r.ReadString(10) // 0x0A separator = newline
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(aLine)