package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abadojack/whatlanggo"
	"io/ioutil"
	"net/http"
)

func detectLang(w http.ResponseWriter, r *http.Request) {
	var (
		requestMap map[string]string
		detectResult whatlanggo.Info
	)
	resultMap := map[string]string{}

	if r.Method != http.MethodPost {
		http.Error(w, "POST method expected", http.StatusNotFound)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &requestMap)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	options := whatlanggo.Options{
		Whitelist: map[whatlanggo.Lang]bool{
			whatlanggo.Eng: true,
			whatlanggo.Por: true,
			whatlanggo.Ukr: true,
			whatlanggo.Pol: true,
			whatlanggo.Fra: true,
			whatlanggo.Spa: true,
			whatlanggo.Rus: true,
		},
	}
	for k, v := range requestMap {
		detectResult = whatlanggo.DetectWithOptions(v, options)
		resultMap[k] = detectResult.Lang.Iso6391()
	}
	res, _ := json.Marshal(resultMap)
	_, _ = w.Write([]byte(res))
}
func main() {
	var (
		httpServerPort int
	)
	flag.IntVar(&httpServerPort, "port", 8080, "")
	flag.Parse()
	http.HandleFunc("/detect-lang", detectLang)
	http.ListenAndServe(fmt.Sprintf(":%d", httpServerPort), nil)
}
