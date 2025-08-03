package logging

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/log"
)

func TestCharmbraceletLogIntegration(t *testing.T) {
	// Clear any existing messages
	defaultLogData.messages = []LogMessage{}
	
	// Setup charmbracelet/log to write to our custom writer
	writer := NewWriter()
	logger := log.NewWithOptions(writer, log.Options{
		Level:           log.DebugLevel,
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	
	// Set as default
	log.SetDefault(logger)
	
	// Test logging functions
	Info("Test info message", "key1", "value1", "key2", "value2")
	Warn("Test warning message", "error", "some error")
	Error("Test error message", "stack", "stack trace here")
	Debug("Test debug message", "source", "/path/to/file.go:123")
	
	// Give a small delay to ensure messages are processed
	time.Sleep(10 * time.Millisecond)
	
	// Get captured messages
	messages := List()
	
	if len(messages) < 4 {
		t.Fatalf("Expected at least 4 messages, got %d", len(messages))
	}
	
	// Check that we have the expected message types
	levels := make(map[string]bool)
	for _, msg := range messages {
		levels[msg.Level] = true
		
		// Verify message is not empty
		if msg.Message == "" {
			t.Errorf("Message should not be empty for level %s", msg.Level)
		}
		
		// Verify timestamp is recent
		if time.Since(msg.Time) > time.Minute {
			t.Errorf("Timestamp seems too old: %v", msg.Time)
		}
	}
	
	// Check we got all expected levels
	expectedLevels := []string{"info", "warn", "error", "debug"}
	for _, level := range expectedLevels {
		if !levels[level] {
			t.Errorf("Expected to find level %s in messages", level)
		}
	}
	
	// Check that attributes were parsed correctly
	found := false
	for _, msg := range messages {
		if strings.Contains(msg.Message, "Test info message") {
			// Check for attributes
			keyFound := false
			for _, attr := range msg.Attributes {
				if attr.Key == "key1" && attr.Value == "value1" {
					keyFound = true
					break
				}
			}
			if !keyFound {
				t.Error("Expected to find key1=value1 attribute in info message")
			}
			found = true
			break
		}
	}
	if !found {
		t.Error("Could not find test info message")
	}
}

func TestPersistentLogging(t *testing.T) {
	// Clear any existing messages
	defaultLogData.messages = []LogMessage{}
	
	// Setup charmbracelet/log to write to our custom writer
	writer := NewWriter()
	logger := log.NewWithOptions(writer, log.Options{
		Level:           log.DebugLevel,
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	
	// Set as default
	log.SetDefault(logger)
	
	// Test persistent logging functions
	InfoPersist("Test persistent info", "important", "data")
	WarnPersist("Test persistent warning")
	ErrorPersist("Test persistent error")
	
	// Give a small delay to ensure messages are processed
	time.Sleep(10 * time.Millisecond)
	
	// Get captured messages
	messages := List()
	
	if len(messages) < 3 {
		t.Fatalf("Expected at least 3 messages, got %d", len(messages))
	}
	
	// Check that persistent flag is set
	persistentCount := 0
	for _, msg := range messages {
		if msg.Persist {
			persistentCount++
		}
	}
	
	if persistentCount != 3 {
		t.Errorf("Expected 3 persistent messages, got %d", persistentCount)
	}
}