package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	asciiart "asciiart/functionFiles"
)

const (
	notFound            = http.StatusNotFound
	internalServerError = http.StatusInternalServerError
	methodNotAllowed    = http.StatusMethodNotAllowed
	badRequest          = http.StatusBadRequest
)

var (
	feedback       string
	currentAsciiArt string // This holds the current ASCII art to be exported
)

// Handle errors by redirecting to an error page
func handleError(writer http.ResponseWriter, statusCode int, message string) {
	// Construct the URL for the error page with query parameters
	target := fmt.Sprintf("/error?code=%d&message=%s", statusCode, url.QueryEscape(message))
	http.Redirect(writer, &http.Request{URL: &url.URL{Path: target}}, target, http.StatusSeeOther)
}

// Handle GET requests to the root path
func Request(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/" {
		handleError(writer, notFound, "Page not found")
		feedback = "This page does not exist"
		return
	}
	if reader.Method != http.MethodGet {
		handleError(writer, methodNotAllowed, "Method not allowed")
		return
	}
	tmpl := GetTemplate()
	err := tmpl.Execute(writer, Data{Success: false})
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		fmt.Printf("Error executing template: %s\n", err)
	}
	fmt.Println("GET / - 200 OK") // Log success in the terminal
}

// Handle POST requests to generate ASCII art
func Post(writer http.ResponseWriter, reader *http.Request) {
	if reader.Method != http.MethodPost {
		handleError(writer, methodNotAllowed, "Method not allowed")
		return
	}

	userInput := reader.FormValue("text")
	banner := reader.FormValue("banner")
	characterMap, err := asciiart.CreateMap(banner)
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		feedback = fmt.Sprintf("Error loading %s banner file", banner)
		fmt.Printf("Error creating map: %s\n", err)
		return
	}

	result := asciiart.DisplayAsciiArt(characterMap, userInput)
	if result == "" {
		handleError(writer, badRequest, "Bad Request")
		feedback = "The Input must NOT contain non-ascii characters"
		fmt.Println("Character not found")
		return
	}

	tmpl := GetTemplate()
	err = tmpl.Execute(writer, Data{Success: true, UserInput: userInput, Result: result})
	if err != nil {
		handleError(writer, internalServerError, "Internal Server Error")
		fmt.Printf("Error executing template: %s\n", err)
	}
	fmt.Println("POST /ascii-art - 200 OK") // Log success with input data

	// Store the result in session or global state for export
	// In practice, use a session to store this data to be used in ExportHandler
	currentAsciiArt = result
}

// Handle errors and render the error page
func ErrorHandler(writer http.ResponseWriter, reader *http.Request) {
	statusCodeStr := reader.URL.Query().Get("code")
	statusCode, err := strconv.Atoi(statusCodeStr)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	message := reader.URL.Query().Get("message")

	err = errorTmpl.Execute(writer, Data{ErrorMessage: message, StatusCode: statusCode, Feedback: feedback})
	if err != nil {
		http.Error(writer, "Error rendering error page", http.StatusInternalServerError)
		fmt.Printf("Error executing error template: %s\n", err)
	}
}

// ExportHandler allows users to download ASCII art as a text file
func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleError(w, methodNotAllowed, "Method not allowed")
		return
	}

	// Retrieve ASCII art from stored session or variable
	asciiArt := currentAsciiArt // This variable should hold the ASCII art to export

	if asciiArt == "" {
		handleError(w, badRequest, "No ASCII art to export")
		return
	}

	// Set HTTP headers for file download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(asciiArt)))

	// Write ASCII art to the response
	_, err := w.Write([]byte(asciiArt))
	if err != nil {
		http.Error(w, "Unable to write file", http.StatusInternalServerError)
		return
	}
	fmt.Println("Export /ascii-art - 200 OK") // Log success in the terminal
}
