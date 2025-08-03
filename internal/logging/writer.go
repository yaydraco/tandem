package logging

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/yaydraco/tandem/internal/pubsub"
)

const (
	persistKeyArg  = "$_persist"
	PersistTimeArg = "$_persist_time"
)

type LogData struct {
	messages []LogMessage
	*pubsub.Broker[LogMessage]
	lock sync.Mutex
}

func (l *LogData) Add(msg LogMessage) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.messages = append(l.messages, msg)
	l.Publish(pubsub.CreatedEvent, msg)
}

func (l *LogData) List() []LogMessage {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.messages
}

var defaultLogData = &LogData{
	messages: make([]LogMessage, 0),
	Broker:   pubsub.NewBroker[LogMessage](),
}

type writer struct{}

func (w *writer) Write(p []byte) (int, error) {
	// Parse charmbracelet/log format: "2025/08/03 12:30:05 INFO <file:line> message key=value key2=value2"
	line := strings.TrimSpace(string(p))
	if line == "" {
		return len(p), nil
	}

	// Regex to parse the charmbracelet/log format
	// Format: "2025/08/03 12:30:05 LEVEL <file:line> message key=value"
	logRegex := regexp.MustCompile(`^(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) ([A-Z]{4}) (?:<([^>]+)> )?(.+)`)
	matches := logRegex.FindStringSubmatch(line)
	if len(matches) < 4 {
		// If we can't parse the line, treat it as a plain message
		msg := LogMessage{
			ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
			Time:    time.Now(),
			Level:   "info",
			Message: line,
		}
		defaultLogData.Add(msg)
		return len(p), nil
	}

	// Parse timestamp
	timestamp, err := time.Parse("2006/01/02 15:04:05", matches[1])
	if err != nil {
		timestamp = time.Now()
	}

	// Extract level
	level := strings.ToLower(matches[2])
	if level == "erro" {
		level = "error"
	} else if level == "debu" {
		level = "debug"
	}

	// Extract caller info if present
	caller := ""
	if len(matches) > 3 && matches[3] != "" {
		caller = matches[3]
	}

	// Extract message and attributes
	messageAndAttrs := matches[4]
	
	// Parse attributes - look for key=value or key="quoted value" patterns
	// Updated regex to handle special characters like $_persist
	attrRegex := regexp.MustCompile(`([\w$_-]+)=("([^"]*)"|(\S+))`)
	attrMatches := attrRegex.FindAllStringSubmatch(messageAndAttrs, -1)
	
	// Remove attributes from message
	message := messageAndAttrs
	for _, match := range attrMatches {
		message = strings.Replace(message, match[0], "", 1)
	}
	message = strings.TrimSpace(message)

	msg := LogMessage{
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
		Time:    timestamp,
		Level:   level,
		Message: message,
	}

	// Add caller as attribute if present
	if caller != "" {
		msg.Attributes = append(msg.Attributes, Attr{
			Key:   "source",
			Value: caller,
		})
	}

	// Process attributes
	for _, match := range attrMatches {
		key := match[1]
		var value string
		if match[3] != "" {
			// Quoted value
			value = match[3]
		} else {
			// Unquoted value
			value = match[4]
		}

		if key == persistKeyArg {
			msg.Persist = value == "true"
		} else if key == PersistTimeArg {
			if parsed, err := time.ParseDuration(value); err == nil {
				msg.PersistTime = parsed
			}
		} else {
			msg.Attributes = append(msg.Attributes, Attr{
				Key:   key,
				Value: value,
			})
		}
	}

	defaultLogData.Add(msg)
	return len(p), nil
}

func NewWriter() *writer {
	w := &writer{}
	return w
}

func Subscribe(ctx context.Context) <-chan pubsub.Event[LogMessage] {
	return defaultLogData.Subscribe(ctx)
}

func List() []LogMessage {
	return defaultLogData.List()
}
