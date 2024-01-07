package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

// RenderTemplate renders a template file with the provided data.
//
// This function takes a file path to a template file and a map of data
// to be applied to the template. It parses the template file, executes it with
// the provided data, and returns the resulting string as a byte slice.
//
// If the template file is not found, the function returns an error with a
// relevant message. Any parsing or execution errors encountered during the
// template processing are also returned as errors.
//
// Example:
//
//	templateFile := "path/to/file.tmpl"
//	data := map[string]interface{}{"key": "value"}
//	result, err := RenderTemplate(templateFile, data)
//	if err != nil {
//	    log.Println("Error:", err)
//	    return
//	}
//	fmt.Println("Rendered String:", string(result))
//
// Parameters:
//   - templateFile: The file path to the template file.
//   - data:         A map of data to be applied to the template.
//
// Returns:
//   - ([]byte, error): The rendered string as a byte slice and any encountered errors.
func RenderTemplate(templateFile string, data map[string]interface{}) ([]byte, error) {
	_, err := os.Stat(templateFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("template file not found: %v", err)
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %v", err)
	}

	var response bytes.Buffer

	err = tmpl.Execute(&response, data)
	if err != nil {
		return nil, fmt.Errorf("error executing template: %v", err)
	}

	return response.Bytes(), nil
}

// GenerateRandomID generates a random integer ID.
//
// This function utilizes the math/rand package to generate a pseudo-random integer ID.
// It initializes the random number generator with the current Unix timestamp, ensuring
// a different seed for each invocation. The generated ID is a positive integer less
// than 100000000 (10^8).
//
// Example:
//
//	randomID := GenerateRandomID()
//	fmt.Println("Random ID:", randomID)
//
// Returns:
//   - int: A randomly generated integer ID.
func GenerateRandomID() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(100000000)
}

// GenerateTimestamp generates a timestamp string in the format "2006-01-02T15:04:05.999Z".
//
// This function utilizes the time package to obtain the current timestamp and formats it
// as a string using the specified layout. The layout "2006-01-02T15:04:05.999Z" represents
// the ISO 8601 format with milliseconds and the 'Z' indicating UTC.
//
// Example:
//
//	timestamp := GenerateTimestamp()
//	fmt.Println("Timestamp:", timestamp)
//
// Returns:
//   - string: A timestamp string formatted as "2006-01-02T15:04:05.999Z".
func GenerateTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z")
}

// GenerateRandomName generates a random word as a string.
//
// This function fetches a random word from the "https://random-word-api.vercel.app" API,
// decodes the JSON response, and returns the first word from the array. In case of any
// errors during the HTTP request or JSON decoding, the function logs the error and
// terminates the program.
//
// Example:
//
//	randomName := GenerateRandomName()
//	fmt.Println("Random Name:", randomName)
//
// Returns:
//   - string: A randomly generated word.
func GenerateRandomWord() string {
	resp, err := http.Get("https://random-word-api.vercel.app/api?words=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var randomWord []string
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&randomWord)
	if err != nil {
		log.Fatal(err)
	}

	return randomWord[0]
}
