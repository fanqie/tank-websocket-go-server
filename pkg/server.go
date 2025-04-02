package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// NewManager creates a new WebSocket manager
func NewManager() *Manager {
	return &Manager{
		Clients:           make(map[*Client]bool),
		Broadcast:         make(chan []byte),
		BroadcastTopic:    make(chan *TopicResponse),
		Register:          make(chan *Client),
		Unregister:        make(chan *Client),
		Subscribe:         make(chan *Subscription),
		Unsubscribe:       make(chan *Subscription),
		Errors:            make(chan *ErrorEvent, 100),      // Buffered error event channel
		ConnEvents:        make(chan *ConnectionEvent, 100), // Buffered connection event channel
		Topics:            make(map[string]map[*Client]bool),
		shutdown:          make(chan struct{}),
		isRunning:         false,
		enableHeartbeat:   true,
		heartbeatInterval: 5 * time.Second,  // 每5秒发送一次心跳
		heartbeatTimeout:  15 * time.Second, // 15秒没有响应就认为超时
		authEnabled:       false,
		authFunc:          nil,
		debug:             true, // 默认开启调试日志
	}
}

// Start starts the WebSocket manager
func (m *Manager) Start() {
	m.isRunning = true
	for {
		select {
		case <-m.shutdown:
			// Received shutdown signal, clean up resources
			m.mutex.Lock()
			for client := range m.Clients {
				client.conn.Close()
			}
			m.Clients = make(map[*Client]bool)
			m.Topics = make(map[string]map[*Client]bool)
			m.mutex.Unlock()
			m.isRunning = false
			return
		case client := <-m.Register:
			m.mutex.Lock()
			m.Clients[client] = true
			m.mutex.Unlock()

			// Send connection event notification
			m.ConnEvents <- &ConnectionEvent{
				Client:    client,
				EventType: "connect",
				UserID:    client.userID,
				Time:      time.Now(),
			}
		case client := <-m.Unregister:
			m.mutex.Lock()
			if _, ok := m.Clients[client]; ok {
				delete(m.Clients, client)
				close(client.send)
				for topic := range client.topics {
					delete(m.Topics[topic], client)
				}

				// Send disconnection event notification
				m.ConnEvents <- &ConnectionEvent{
					Client:    client,
					EventType: "disconnect",
					UserID:    client.userID,
					Time:      time.Now(),
				}
			}
			m.mutex.Unlock()
		case sub := <-m.Subscribe:
			m.mutex.Lock()
			if _, ok := m.Topics[sub.topic]; !ok {
				m.Topics[sub.topic] = make(map[*Client]bool)
			}
			m.Topics[sub.topic][sub.client] = true
			sub.client.topics[sub.topic] = true
			m.mutex.Unlock()

			// Send subscription event notification
			m.ConnEvents <- &ConnectionEvent{
				Client:    sub.client,
				EventType: "subscribe",
				UserID:    sub.client.userID,
				Topic:     sub.topic,
				Time:      time.Now(),
			}
		case unsub := <-m.Unsubscribe:
			m.mutex.Lock()
			if clients, ok := m.Topics[unsub.topic]; ok {
				delete(clients, unsub.client)
				delete(unsub.client.topics, unsub.topic)

				// Send unsubscription event notification
				m.ConnEvents <- &ConnectionEvent{
					Client:    unsub.client,
					EventType: "unsubscribe",
					UserID:    unsub.client.userID,
					Topic:     unsub.topic,
					Time:      time.Now(),
				}
			}
			m.mutex.Unlock()
		case message := <-m.Broadcast:
			m.mutex.Lock()
			fmt.Printf("-------------%v----------", string(message))
			for client := range m.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
				}
			}
			m.mutex.Unlock()
		case message := <-m.BroadcastTopic:
			m.mutex.Lock()
			if clients, ok := m.Topics[message.Topic]; ok {
				messageBytes, err := json.Marshal(message)
				if err != nil {
					m.Errors <- &ErrorEvent{
						Message: "Message serialization failed",
						Code:    1001,
						Time:    time.Now(),
					}
					log.Println("Message serialization failed:", err)
					return
				}
				for client := range clients {
					select {
					case client.send <- messageBytes:
					default:
						close(client.send)
						delete(clients, client)
					}
				}
			}
			m.mutex.Unlock()
		}
	}
}

// Configure WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all CORS requests
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SetCustomUpgrader allows setting a custom WebSocket upgrader
func SetCustomUpgrader(customUpgrader websocket.Upgrader) {
	upgrader = customUpgrader
}

// EnableAuth enables authentication using the provided function
func (m *Manager) EnableAuth(authFunc func(r *http.Request) bool) {
	m.authEnabled = true
	m.authFunc = authFunc
}

// DisableAuth disables authentication
func (m *Manager) DisableAuth() {
	m.authEnabled = false
	m.authFunc = nil
}

// HandleConnection handles WebSocket request
func (m *Manager) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Check authentication if enabled
	if m.authEnabled && m.authFunc != nil {
		if !m.authFunc(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			m.Errors <- &ErrorEvent{
				Message: "Authentication failed",
				Code:    1007,
				Time:    time.Now(),
			}
			return
		}
	}

	// Upgrade HTTP connection to WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m.Errors <- &ErrorEvent{
			Message: "Connection upgrade failed",
			Code:    1002,
			Time:    time.Now(),
		}
		log.Println("Connection upgrade failed:", err)
		return
	}

	// 生成唯一的客户端ID
	clientID := r.URL.Query().Get("user_id")
	if clientID == "" {
		clientID = fmt.Sprintf("client_%d", time.Now().UnixNano())
	}

	// Create new client
	client := &Client{
		manager: m,
		conn:    conn,
		send:    make(chan []byte, 256),
		userID:  clientID,
		topics:  make(map[string]bool),
	}

	// 发送欢迎消息
	welcomeMsg := fmt.Sprintf("Hello, %s! Welcome to WebSocket server.", clientID)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(welcomeMsg)); err != nil {
		m.debugLog("Failed to send welcome message: %v", err)
		conn.Close()
		return
	}
	m.debugLog("Sent welcome message to client %s", clientID)

	// Register new client
	client.manager.Register <- client

	// Start heartbeat if enabled
	client.startHeartbeat()

	// Start goroutines for read/write operations
	go client.readPump()
	go client.writePump()
}

// Client read message
func (c *Client) readPump() {
	defer func() {
		c.manager.Unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.manager.Errors <- &ErrorEvent{
					Client:  c,
					Message: fmt.Sprintf("Read message error: %v", err),
					Code:    1003,
					Time:    time.Now(),
				}
				c.manager.debugLog("Client %s: Read error: %v", c.userID, err)
			}
			break
		}

		msgStr := string(message)
		c.manager.debugLog("Client %s: Received message: %s", c.userID, msgStr)

		// 尝试解析JSON消息
		var msgMap map[string]interface{}
		if err := json.Unmarshal(message, &msgMap); err == nil {
			// 如果是心跳响应消息，忽略
			if msgType, ok := msgMap["type"].(string); ok && msgType == "heartbeat" {
				continue
			}
		}

		// Handle subscription message
		if strings.HasPrefix(msgStr, "sub:") {
			topic := msgStr[4:]
			c.manager.debugLog("Client %s: Subscribing to topic: %s", c.userID, topic)
			c.manager.Subscribe <- &Subscription{client: c, topic: topic}
		} else if strings.HasPrefix(msgStr, "unsub:") {
			topic := msgStr[6:]
			c.manager.debugLog("Client %s: Unsubscribing from topic: %s", c.userID, topic)
			c.manager.Unsubscribe <- &Subscription{client: c, topic: topic}
		} else {
			// 广播消息给其他客户端
			c.manager.debugLog("Client %s: Broadcasting message to other clients: %s", c.userID, msgStr)
			c.manager.mutex.Lock()
			for client := range c.manager.Clients {
				// 不发送给消息发送者自己
				if client != c {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(c.manager.Clients, client)
					}
				}
			}
			c.manager.mutex.Unlock()
		}
	}
}

// Client write message
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		message, ok := <-c.send
		if !ok {
			// Channel is closed
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			c.manager.Errors <- &ErrorEvent{
				Client:  c,
				Message: fmt.Sprintf("Write message error: %v", err),
				Code:    1004,
				Time:    time.Now(),
			}
			c.manager.debugLog("Client %s: Write error: %v", c.userID, err)
			return
		}
		c.manager.debugLog("Client %s: Sent message: %s", c.userID, string(message))
	}
}

// BroadcastMessage broadcasts a message to all connected clients except the excluded one
func (m *Manager) BroadcastMessage(message []byte, excludeClient *Client) {
	excludeID := "none"
	if excludeClient != nil {
		excludeID = excludeClient.userID
	}
	m.debugLog("Broadcasting message to all clients (except %s): %s", excludeID, string(message))

	m.mutex.Lock()
	for client := range m.Clients {
		if excludeClient == nil || client != excludeClient {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(m.Clients, client)
			}
		}
	}
	m.mutex.Unlock()
}

// BroadcastTopicMessage broadcasts a message to all subscribers of a specific topic
func (m *Manager) BroadcastTopicMessage(topic string, data string) {
	m.debugLog("Broadcasting message to topic %s: %s", topic, data)
	m.BroadcastTopic <- &TopicResponse{
		Topic: topic,
		Data:  data,
	}
}

// SetHTTPServer sets HTTP server reference for shutdown
func (m *Manager) SetHTTPServer(server *http.Server) {
	m.httpServer = server
}

// Shutdown gracefully shuts down the WebSocket manager and HTTP server
func (m *Manager) Shutdown(ctx context.Context) error {
	if !m.isRunning {
		return fmt.Errorf("server is not running")
	}

	// Notify all clients of imminent shutdown
	closeMessage := []byte("Server is shutting down")
	m.BroadcastMessage(closeMessage, nil)

	// Send shutdown signal
	m.shutdown <- struct{}{}

	// Shut down HTTP server
	if m.httpServer != nil {
		return m.httpServer.Shutdown(ctx)
	}

	return nil
}

// GetClientCount gets the number of currently connected clients
func (m *Manager) GetClientCount() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return len(m.Clients)
}

// GetTopicSubscriberCount gets the number of subscribers for a specific topic
func (m *Manager) GetTopicSubscriberCount(topic string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if clients, ok := m.Topics[topic]; ok {
		return len(clients)
	}
	return 0
}

// GetAllTopics gets all available topics
func (m *Manager) GetAllTopics() []string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	topics := make([]string, 0, len(m.Topics))
	for topic := range m.Topics {
		topics = append(topics, topic)
	}
	return topics
}

// IsRunning checks if the server is running
func (m *Manager) IsRunning() bool {
	return m.isRunning
}

// CloseClient closes the connection to a specific client
func (m *Manager) CloseClient(userID string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for client := range m.Clients {
		if client.userID == userID {
			client.conn.Close()
			return true
		}
	}
	return false
}

// EnableHeartbeat enables the heartbeat mechanism with specified interval
func (m *Manager) EnableHeartbeat(interval time.Duration) {
	m.enableHeartbeat = true
	m.heartbeatInterval = interval
}

// DisableHeartbeat disables the heartbeat mechanism
func (m *Manager) DisableHeartbeat() {
	m.enableHeartbeat = false
}

// startHeartbeat starts the heartbeat mechanism for a client
func (c *Client) startHeartbeat() {
	if !c.manager.enableHeartbeat {
		c.manager.debugLog("Client %s: Heartbeat is disabled", c.userID)
		return
	}

	c.manager.debugLog("Client %s: Starting heartbeat with interval %v",
		c.userID, c.manager.heartbeatInterval)

	// 启动心跳发送goroutine
	go func() {
		// 等待一小段时间再开始发送心跳，避免连接建立后立即发送
		time.Sleep(1 * time.Second)

		ticker := time.NewTicker(c.manager.heartbeatInterval)
		defer ticker.Stop()

		consecutiveFailures := 0 // 连续发送失败次数
		maxFailures := 3         // 最大允许的连续失败次数

		for {
			select {
			case <-ticker.C:

				// 尝试发送心跳消息
				if err := c.conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
					consecutiveFailures++
					c.manager.debugLog("Client %s: Failed to send heartbeat (%d/%d failures): %v",
						c.userID, consecutiveFailures, maxFailures, err)

					if consecutiveFailures >= maxFailures {
						c.manager.debugLog("Client %s: Too many consecutive heartbeat failures, closing connection", c.userID)
						c.manager.Errors <- &ErrorEvent{
							Client:  c,
							Message: fmt.Sprintf("Heartbeat failed %d times consecutively", consecutiveFailures),
							Code:    1005,
							Time:    time.Now(),
						}
						c.conn.Close()
						return
					}
				} else {
					// 发送成功，重置失败计数
					if consecutiveFailures > 0 {
						c.manager.debugLog("Client %s: Heartbeat restored after %d failures", c.userID, consecutiveFailures)
					}
					consecutiveFailures = 0
				}
			}
		}
	}()
}

// EnableDebug enables debug logging
func (m *Manager) EnableDebug() {
	m.debug = true
}

// DisableDebug disables debug logging
func (m *Manager) DisableDebug() {
	m.debug = false
}

// debugLog prints debug message if debug is enabled
func (m *Manager) debugLog(format string, v ...interface{}) {
	if m.debug {
		log.Printf(format, v...)
	}
}
