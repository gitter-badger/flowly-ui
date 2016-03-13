package main

import (
  "log"
  "github.com/gorilla/websocket"
)

type Connection struct {
  socket *websocket.Conn
  send   chan []byte
}

func (conn *Connection) writeLoop() {
  defer conn.socket.Close()

  for message := range conn.send {
    // write the message to the socket
    if err := conn.socket.WriteMessage(websocket.TextMessage, message); err != nil {
      log.Print("Could not write message: ", err)
      break
    }
  }
}

/**
 * Reads all the messages of a websocket until it is closed.
 */
func (conn *Connection) readLoop() {
  for {
    if _, _, err := conn.socket.NextReader(); err != nil {
      log.Println("Error reading websocket: ", err)
      conn.socket.Close()
      break
    }
  }
}
