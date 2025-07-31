package chat

import (
	"strings"
	"testing"
	"time"

	"github.com/yyovil/tandem/internal/message"
)

func TestFormatTimestampDiffSeconds(t *testing.T) {
	tests := []struct {
		name     string
		start    int64
		end      int64
		expected string
	}{
		{
			name:     "less than 1 second",
			start:    1000,
			end:      1500,
			expected: "0.5s",
		},
		{
			name:     "exactly 1 second",
			start:    1000,
			end:      2000,
			expected: "1.0s",
		},
		{
			name:     "multiple seconds",
			start:    1000,
			end:      3500,
			expected: "2.5s",
		},
		{
			name:     "more than 60 seconds",
			start:    1000,
			end:      61000,
			expected: "60.0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatTimestampDiffSeconds(tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("formatTimestampDiffSeconds(%d, %d) = %s, want %s", tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestGetCurrentTimeFormatted(t *testing.T) {
	result := getCurrentTimeFormatted()

	// Check that the result has the correct format (HH:MM:SS)
	if len(result) != 8 {
		t.Errorf("getCurrentTimeFormatted() length = %d, want 8", len(result))
	}

	// Check that it has colons in the right places
	if result[2] != ':' || result[5] != ':' {
		t.Errorf("getCurrentTimeFormatted() = %s, want format HH:MM:SS", result)
	}

	// Parse it to ensure it's a valid time
	_, err := time.Parse("15:04:05", result)
	if err != nil {
		t.Errorf("getCurrentTimeFormatted() returned invalid time format: %s, error: %v", result, err)
	}
}

func TestRenderAssistantMessage_TimeDisplay(t *testing.T) {
	// Create a test message with finish data
	now := time.Now().Unix()
	startTime := now - 2500 // 2.5 seconds ago

	msg := message.Message{
		ID:        "test-id",
		CreatedAt: startTime,
		Parts: []message.ContentPart{
			message.TextContent{Text: "Hello, this is a test message"},
			message.Finish{
				Reason: message.FinishReasonEndTurn,
				Time:   now,
			},
		},
	}

	// Render the message
	uiMessages := renderAssistantMessage(msg, []message.Message{}, false, 80, 0)

	// There should be one UI message
	if len(uiMessages) != 1 {
		t.Errorf("Expected 1 UI message, got %d", len(uiMessages))
		return
	}

	content := uiMessages[0].content

	// Check that the content contains both current time format and duration in seconds
	// The format should be "HH:MM:SS (X.Xs)"
	if !strings.Contains(content, ":") {
		t.Errorf("Message content should contain time with colons, got: %s", content)
	}

	if !strings.Contains(content, ".") {
		t.Errorf("Message content should contain duration with decimal seconds, got: %s", content)
	}

	if !strings.Contains(content, "s") {
		t.Errorf("Message content should contain 's' for seconds, got: %s", content)
	}

	if !strings.Contains(content, "(") || !strings.Contains(content, ")") {
		t.Errorf("Message content should contain duration in parentheses, got: %s", content)
	}
}

func TestRenderAssistantMessage_NonFinishedMessage(t *testing.T) {
	// Create a test message without finish data
	msg := message.Message{
		ID:        "test-id",
		CreatedAt: time.Now().Unix(),
		Parts: []message.ContentPart{
			message.TextContent{Text: "Hello, this is an unfinished message"},
		},
	}

	// Render the message
	uiMessages := renderAssistantMessage(msg, []message.Message{}, false, 80, 0)

	// There should be one UI message
	if len(uiMessages) != 1 {
		t.Errorf("Expected 1 UI message, got %d", len(uiMessages))
		return
	}

	content := uiMessages[0].content

	// Check that the content does not contain time format for unfinished messages
	// Should not contain ":" from time format or "(" from duration
	lines := strings.Split(content, "\n")
	lastLine := lines[len(lines)-1]

	// The time format should not appear in an unfinished message
	if strings.Contains(lastLine, ":") && strings.Contains(lastLine, "(") {
		t.Errorf("Unfinished message should not contain time display, got: %s", lastLine)
	}
}
