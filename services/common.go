package commands

import (
	labels "TDBackend/localization"
	"TDBackend/logger"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// getFileData func for RESTful API
func readFileData(endPoint string, w http.ResponseWriter, r *http.Request) {

	// Get current time
	t := time.Now()

	// Get file name
	fileName := "static/" + endPoint + ".json"

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)

	// if we os.Open returns an error then handle it
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.AppLogger(labels.Error, "Result of GET "+endPoint+"!", time.Since(t).Nanoseconds(), labels.Error+"|"+err.Error())
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read the data from file
	byteValue, errFile := ioutil.ReadAll(jsonFile)
	if errFile != nil {
		// Set HTTP code to 500
		w.WriteHeader(http.StatusInternalServerError)
		logger.AppLogger(labels.Error, "Result of GET Partners", time.Since(t).Nanoseconds(), labels.Error+"|"+errFile.Error())
	}

	// Set JSON as a Content-Type and User-Agent for output
	w.Header().Add(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.Header().Add(labels.HTTPUSERAGENTKey, labels.HTTPUSERAGENTValue)

	// Send data out
	w.Write(byteValue)

	// Set HTTP code to 200
	w.WriteHeader(http.StatusOK)
	logger.AppLogger(labels.Debug, "Result of GET func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusOK))
}
