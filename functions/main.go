package main
import "fmt"

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	// Add fields as needed
	userData map[string]string
}

func (s SimpleDataStore) UserNameForId(userId string) (string, bool) {
	// Simulate a simple data store lookup
	name, ok := s.userData[userId]
	return name, ok
}

// NewSimpleDataStore initializes a SimpleDataStore with some dummy data
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"123": "Alice",
			"456": "Bob",
		},
	}
}

// UserNameForId retrieves the user name for a given user ID
// Simple interface which can be used to mock the data store
type DataStore interface {
	UserNameForId(userId string) (string, bool)
}


// Logger interface for logging messages
//  My logging functionality or library needs to implement this interface
type Logger interface {
	Log(message string)
}

// LoggerAdapter meets the Logger interface
// and wraps a function that takes a string
type LoggerAdapter func(message string)
func (l LoggerAdapter) Log(message string) {
	l(message)
}

type SimpleLogic struct {
	dataStore DataStore
	logger    Logger
}

func (s SimpleLogic) SayHello(userId string) (string, error) {
	s.l.Log("SayHello called with userId: " + userId)
	name, ok := s.dataStore.UserNameForId(userId)
	if !ok {
		return "", fmt.Errorf("user not found")
	}
	return "Hello " + name, nil
}

// factory function to create a new SimpleLogic instance
func NewSimpleLogic(dataStore DataStore, logger Logger) SimpleLogic {
	return SimpleLogic{
		dataStore: dataStore,
		logger:    logger,
	}
}

// i just need to say hello to the user
type Logic interface {
	SayHello(userId string) (string, error)
}


func main() {
	// Initialize the data store
	dataStore := NewSimpleDataStore()

	// Initialize the logger
	logger := LoggerAdapter(LogOutput)

	// Example usage
	userId := "123"
	if name, ok := dataStore.UserNameForId(userId); ok {
		logger.Log("User found: " + name)
	} else {
		logger.Log("User not found")
	}
}