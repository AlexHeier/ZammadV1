package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

// Husk at å følge linux med / ikke \
var CSVPath string = "C:/Users/Alexa/Desktop/bedkombedpressinvsender/test.csv"

// Navnet til den som skal stå som eier. Må være en av de som allerede eksisterer.
var EmailOwner string = "Alexander Engebrigtsen Heier"

// Om boten skal sende mailen eller ikke
var ToSend bool = false

// Company represents a single entry in your CSV file
type Company struct {
	Bedrift string
	Emails  string
	CC      []string
}

var emailTitle string = "Login inviterer til bedriftspresentasjon"

func main() {
	time.Sleep(1 * time.Second)
	companies, err := importCSV()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Example output to verify the import
	for _, company := range companies {
		sendMail(company)
	}
}

func importCSV() ([]Company, error) {
	// Open the CSV file
	file, err := os.Open(CSVPath) // Replace with your CSV file path
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Optionally set the reader to trim leading and trailing spaces
	reader.TrimLeadingSpace = true

	// Read all the data from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	// Array to store the parsed data
	var companies []Company

	// Loop through the CSV records and parse each entry
	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		CCList := strings.Split(record[2], " ")

		// Create a Company struct and populate it
		company := Company{
			Bedrift: record[0],
			Emails:  record[1],
			CC:      CCList,
		}

		// Add the company to the list
		companies = append(companies, company)
	}

	return companies, nil
}

// moveMouseByPercentage moves the mouse to a location based on percentage of the screen width and height
func moveMouseByPercentage(xPercent, yPercent float64) {
	// Get screen size in pixels
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// Convert percentages (0-100) to decimals (0.0-1.0) and calculate pixel positions
	x := int((xPercent / 100) * float64(screenWidth))
	y := int((yPercent / 100) * float64(screenHeight))

	// Move the mouse instantly to the calculated position
	robotgo.Move(x, y)
}

func replaceInputFieldContent(content string) {
	// Simulate clicking on the text field (assuming you have already moved to it)
	robotgo.Click()

	// Wait briefly to ensure the select action is registered
	time.Sleep(50 * time.Millisecond) // Reduced sleep duration

	// Simulate deleting the selected text (Backspace or Delete key)
	robotgo.KeyTap("delete") // Use "delete" key to clear the field

	// Wait briefly before typing the new content
	time.Sleep(50 * time.Millisecond) // Reduced sleep duration

	// Type the new content (the emailTitle)
	robotgo.TypeStr(content)

	// Optional: Add a new line after typing if needed
	robotgo.KeyTap("enter")
}

func replaceInputFieldCC(persons []string) {
	// Simulate clicking on the text field (assuming you have already moved to it)
	robotgo.Click()

	// Wait briefly to ensure the select action is registered
	time.Sleep(50 * time.Millisecond) // Reduced sleep duration

	// Simulate deleting the selected text (Backspace or Delete key)
	robotgo.KeyTap("delete") // Use "delete" key to clear the field

	// Wait briefly before typing the new content
	time.Sleep(50 * time.Millisecond) // Reduced sleep duration

	for _, person := range persons {
		robotgo.TypeStr(person)
		time.Sleep(300 * time.Millisecond)
		robotgo.KeyTap("tab")
	}

	// Optional: Add a new line after typing if needed
	robotgo.KeyTap("enter")
}

func replaceInputFieldContentMultiLine(companyName string) {
	// Simulate clicking on the text field (assuming you have already moved to it)
	robotgo.Click()

	// Wait briefly before typing the new content
	time.Sleep(100 * time.Millisecond) // Reduced sleep duration
	var emailContent = []string{
		"Hei,",
		"",
		fmt.Sprintf("Vi i Login ønsker å begynne planlegging av bedriftspresentasjoner for neste semester. Dette er en super måte hvor dere kan vise hva %s holder på med og få studentene interessert til å søke hos dere! Vi ønsker å høre om dette er noe dere er interessert i å avholde hos oss. Mer informasjon om selve gjennomføring, dato og pris kommer dersom dere er interessert :)", companyName),
		"",
		"Hører fra dere :)",
	}

	// Iterate through each line in the content slice
	for _, line := range emailContent {
		// Type the current line
		robotgo.TypeStr(line)
		// Simulate pressing "Enter" to create a new line
		robotgo.KeyTap("enter")
	}
}

func sendMail(company Company) {

	moveMouseByPercentage(9.6875, 94.375)
	robotgo.Click("left", false)
	robotgo.MouseSleep = 3000
	moveMouseByPercentage(32.8125, 25.0697)
	replaceInputFieldContent(emailTitle)
	moveMouseByPercentage(32.8125, 30.4861)
	replaceInputFieldContent(company.Emails)
	moveMouseByPercentage(19.5, 42.5)
	robotgo.Click()
	moveMouseByPercentage(32.8125, 36.4583)
	replaceInputFieldCC(company.CC)

	moveMouseByPercentage(32.8125, 52.8472)
	robotgo.Click("left", false)
	robotgo.MouseSleep = 300
	moveMouseByPercentage(32.8125, 61.0417)
	robotgo.Click("left", false)

	moveMouseByPercentage(19.5, 42.5)
	robotgo.Click()
	robotgo.MouseSleep = 200

	moveMouseByPercentage(53.6328, 67.1528)
	robotgo.MouseSleep = 1000
	robotgo.Click("left", false)
	robotgo.Click("esc")
	robotgo.MouseSleep = 1000
	robotgo.TypeStr(EmailOwner)
	robotgo.KeyTap("enter")

	robotgo.MouseSleep = 100

	moveMouseByPercentage(32.8125, 42.5)
	replaceInputFieldContentMultiLine(company.Bedrift)

	moveMouseByPercentage(19.5, 42.5)
	robotgo.Scroll(0, -1000)

	robotgo.MouseSleep = 500
	moveMouseByPercentage(65.7422, 91.4)

	if ToSend {
		robotgo.Click()
	}
}
