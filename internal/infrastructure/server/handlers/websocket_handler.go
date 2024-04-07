package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/kafka/producer"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	// conn *websocket.Conn
	uc *usecase.ChatUseCase
}

func NewWebsocketHandler(uc *usecase.ChatUseCase) *WebSocketConnection {
	return &WebSocketConnection{uc}
}

var OnlineUsers = make(map[string]*websocket.Conn)

var upgrade = websocket.Upgrader{
	HandshakeTimeout: 10 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

func (ws *WebSocketConnection) EstablishConnection(ctx *gin.Context) {
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusInternalServerError,
			Error:   err.Error(),
			Message: "Failed to establish websocket connection",
		})
		return
	}

	user := ctx.GetStringMap("user")
	userId, ok := user["userId"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   e.ErrKeyNotFound.Error(),
			Message: "Failed to get userId from token",
		})
		return
	}

	defer delete(OnlineUsers, userId.(string))
	defer conn.Close()
	OnlineUsers[userId.(string)] = conn

	for {
		fmt.Println("==loop start ", OnlineUsers)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		} else {
			ws.SendMessage(ctx, msg, userId.(string))
		}
	}
}

func (ws *WebSocketConnection) SendMessage(ctx context.Context, msg []byte, userId string) {

	message := new(req.SendChatReq)
	message.CreatedAt = helper.CurrentIsoDateTimeString()
	message.SenderID = userId

	err := json.Unmarshal(msg, message)
	if err != nil {
		conn, ok := OnlineUsers[userId]
		if ok {
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}
	}
	msgBytes, _ := json.Marshal(message)

	conn, ok := OnlineUsers[message.RecipientID]
	if ok {
		err := conn.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			delete(OnlineUsers, message.RecipientID)
		} else {
			message.ReceivedAt = helper.CurrentIsoDateTimeString()
		}
	}else{
		producer.NewProducer("chat", userId, msg)
	}
	resp := ws.uc.SaveMessage(ctx, message)
	if resp.Error != nil {
		log.Println("Error Storing Msg:", err.Error())
	}
}
