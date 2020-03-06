package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ship87/json-csv-converter-golang/helpers"
	"github.com/ship87/json-csv-converter-golang/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var cfg models.Config

func init() {

	cfg.Fill()
}

func main() {

	fmt.Println("Starting server...")
	urlDirectoryDownload := helpers.ConcatStrings([]string{"/", cfg.DirectoryDownload, "/"})
	http.HandleFunc(urlDirectoryDownload, downloadFile)
	http.HandleFunc("/", root)
	port := helpers.ConcatStrings([]string{":", cfg.AppPort})
	http.ListenAndServe(port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			io.WriteString(w, fmt.Sprintf("Media type '%s' not supported\n", contentType))
			return
		}
		handleJson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, fmt.Sprintf("Method '%s' not allowed\n", r.Method))
	}
}

func downloadFile(w http.ResponseWriter, r *http.Request) {

	contentDispositionHeader := helpers.ConcatStrings([]string{"attachment; filename=", r.URL.Path[1:]})
	w.Header().Set("Content-Disposition", contentDispositionHeader)
	w.Header().Set("Content-Type", "text/csv")
	http.ServeFile(w, r, r.URL.Path[1:])
}

func handleJson(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now().Nanosecond()
	var lines []models.Line
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	isErr := checkServerError("Cannot read body", err, w)
	if isErr {
		return
	}

	err = json.Unmarshal(data, &lines)
	isErr = checkServerError("Cannot unmarshal json", err, w)
	if isErr {
		return
	}

	resultJson, isErr := saveCsvFile(lines, w)
	if isErr {
		return
	}

	endTime := time.Now().Nanosecond()
	elapsedTime := float64(endTime - startTime)
	elapsedTimeInMs := strconv.FormatFloat(elapsedTime/1000000, 'f', -1, 64)

	w.Header().Set("X-Elapsed-Time", helpers.ConcatStrings([]string{elapsedTimeInMs, " ms"}))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}

func saveCsvFile(lines []models.Line, w http.ResponseWriter) ([]byte, bool) {

	var result models.Result
	var resultJson []byte

	if _, err := os.Stat(cfg.DirectoryDownload); os.IsNotExist(err) {
		err = os.MkdirAll(cfg.DirectoryDownload, 0755)
		isErr := checkServerError("Cannot create directory", err, w)
		if isErr {
			return resultJson, isErr
		}
	}

	file, err := ioutil.TempFile(cfg.DirectoryDownload, cfg.PrefixFile)
	isErr := checkServerError("Cannot create file", err, w)
	if isErr {
		return resultJson, isErr
	}
	defer file.Close()

	result.DownloadLink = getUrlDownload(file.Name())

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		err := writer.Write(line.Columns)
		isErr = checkServerError("Cannot write to file", err, w)
		if isErr {
			return resultJson, isErr
		}
	}

	resultJson, err = json.Marshal(result)
	isErr = checkServerError("Cannot write to file", err, w)

	return resultJson, isErr
}

func getUrlDownload(path string) string {

	link, err := url.Parse(cfg.AppUrl)
	if err != nil {
		log.Fatal(err)
	}
	link.Path = path

	return link.String()
}

func checkServerError(message string, err error, w http.ResponseWriter) bool {

	if err == nil {
		return false
	}

	log.Println(message, err)
	w.WriteHeader(http.StatusInternalServerError)

	return true
}
