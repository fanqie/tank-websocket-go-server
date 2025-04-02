package pkg

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type TopicResponse struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// ErrorEvent represents an error event in the WebSocket service
type ErrorEvent struct {
	Client  *Client   `json:"-"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
	Time    time.Time `json:"time"`
}

// ConnectionEvent represents a connection status change event
type ConnectionEvent struct {
	Client    *Client   `json:"-"`
	EventType string    `json:"event_type"` // "connect", "disconnect", "subscribe", "unsubscribe"
	UserID    string    `json:"user_id"`
	Topic     string    `json:"topic,omitempty"`
	Time      time.Time `json:"time"`
}

// Manager manages all WebSocket connections
type Manager struct {
	Clients        map[*Client]bool
	Broadcast      chan []byte
	BroadcastTopic chan *TopicResponse
	Register       chan *Client
	Unregister     chan *Client
	Subscribe      chan *Subscription
	Unsubscribe    chan *Subscription
	Errors         chan *ErrorEvent      // Error event channel
	ConnEvents     chan *ConnectionEvent // Connection event channel
	mutex          sync.Mutex
	Topics         map[string]map[*Client]bool
	httpServer     *http.Server  // For closing HTTP server
	shutdown       chan struct{} // Channel for shutdown notification
	isRunning      bool          // Server running status

	// Heartbeat configuration
	enableHeartbeat   bool
	heartbeatInterval time.Duration
	heartbeatTimeout  time.Duration

	// Authentication
	authEnabled bool
	authFunc    func(r *http.Request) bool

	// Debug configuration
	debug bool // 是否启用调试日志
}

// Client represents a WebSocket connection
type Client struct {
	manager *Manager
	conn    *websocket.Conn
	send    chan []byte
	userID  string
	topics  map[string]bool
}

// Subscription represents a topic subscription by a client
type Subscription struct {
	client *Client
	topic  string
}
