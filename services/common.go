package commands

import (
	labels "TDBackend/localization"
	"TDBackend/logger"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	repo   = make(map[string]string)
	fNames = []string{"Deals", "FAQs", "Kit", "Partners", "Sessions"}
)

// loadFileToMemory func loads data into memory for better response
func loadFileToMemory() {

	// Get current time
	t := time.Now()

	for _, v := range fNames {
		data, err := readFile(v)
		if err != nil {
			logger.AppLogger(labels.Error, "Error while opening "+v+" file!", time.Since(t).Nanoseconds(), labels.Error+"|"+err.Error())
		}
		repo[v] = data
	}
}

// readFile func reads data from files
func readFile(endPoint string) (string, error) {

	// Get current time
	t := time.Now()

	// Get file name
	fileName := "static/" + endPoint + ".json"

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)

	// if we os.Open returns an error then handle it
	if err != nil {
		logger.AppLogger(labels.Error, "Error while opening "+endPoint+" file!", time.Since(t).Nanoseconds(), labels.Error+"|"+err.Error())
		return "", err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read the data from file
	byteValue, errFile := ioutil.ReadAll(jsonFile)
	if errFile != nil {
		// Set HTTP code to 500
		logger.AppLogger(labels.Error, "Error while reading "+endPoint+" file!", time.Since(t).Nanoseconds(), labels.Error+"|"+errFile.Error())
		return "", errFile
	}

	var data string = string(byteValue)
	logger.AppLogger(labels.Debug, "Data read from "+endPoint+" file!", time.Since(t).Nanoseconds(), labels.Data+"|"+data)

	return data, nil
}

// getData func sends data back
func getData(endPoint string, w http.ResponseWriter, r *http.Request) (int, error) {

	// Get current time
	t := time.Now()

	logger.AppLogger(labels.Debug, "REQUEST HEADERS", time.Since(t).Nanoseconds(), labels.Data+"|"+fmt.Sprintf("%v", r.Header))
	logger.AppLogger(labels.Debug, "REQUEST CONTEXT", time.Since(t).Nanoseconds(), labels.Data+"|"+fmt.Sprintf("%v", r.Context()))
	logger.AppLogger(labels.Debug, "REQUEST", time.Since(t).Nanoseconds(), labels.Data+"|"+fmt.Sprintf("%v", r))

	// Set JSON as a Content-Type and User-Agent for output
	w.Header().Add(labels.HTTPContentTypeKey, labels.HTTPContentTypeValue)
	w.Header().Add(labels.HTTPUSERAGENTKey, labels.HTTPUSERAGENTValue)

	logger.AppLogger(labels.Debug, "Result of GET "+endPoint+" func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusOK))

	// Send data out
	return w.Write([]byte(repo[endPoint]))
}

// getImage func for Serving images
func getImage(w http.ResponseWriter, r *http.Request) (int, error) {

	// Get current time
	t := time.Now()

	// Get file name
	params := mux.Vars(r)
	fileName := "static/img/" + params["name"]

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)

	// if we os.Open returns an error then handle it
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.AppLogger(labels.Error, "Error while opening "+params["name"]+" file!", time.Since(t).Nanoseconds(), labels.Error+"|"+err.Error())
		return 0, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read the data from file
	byteValue, errFile := ioutil.ReadAll(jsonFile)
	if errFile != nil {
		// Set HTTP code to 500
		w.WriteHeader(http.StatusInternalServerError)
		logger.AppLogger(labels.Error, "Error while reading "+params["name"]+" file!", time.Since(t).Nanoseconds(), labels.Error+"|"+errFile.Error())
		return 0, errFile
	}

	// Set JSON as a Content-Type and User-Agent for output
	w.Header().Add(labels.HTTPContentTypeKey, labels.HTTPContentTypeIMGValue)
	w.Header().Add(labels.HTTPUSERAGENTKey, labels.HTTPUSERAGENTValue)

	logger.AppLogger(labels.Debug, "Result of GET func", time.Since(t).Nanoseconds(), labels.Code+"|"+fmt.Sprintf("%v", http.StatusOK))

	// Send data out
	return w.Write(byteValue)
}
