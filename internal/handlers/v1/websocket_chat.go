package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (customize as needed)
	},
}

type Message struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

func (h *Handler) handleWebSocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "WebSocket upgrade failed"})
		return
	}
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		response := processMessage(msg, h, c)

		err = conn.WriteJSON(response)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func processMessage(msg Message, h *Handler, c *gin.Context) Message {
	var response Message
	response.Action = "response"
	switch msg.Action {
	case "find_product":
		product, err := h.services.product.GetSimilarProduct(c, msg.Payload)
		if err != nil {
			response.Action = "error"
			response.Payload = err.Error()
		} else {
			response.Payload = fmt.Sprintf("Возможно, вы ищите это: <a href=/products/%d>%s</a>", product.ID, product.Name)
		}
		fmt.Println(product.ID, product.Name)
		return response
	case "add_order":
		email := c.Request.Context().Value("email").(string)
		err := h.services.order.AddOrder(c, email, msg.Payload)
		if err != nil {
			response.Action = "error"
			response.Payload = fmt.Sprintf("Ошибка оформления заказа: %s", err.Error())
			return response
		}
		response.Payload = "Заказ оформлен! Посмотреть заказы <a href=\"/render-auth/order-history\">здесь</a>"
		return response
	case "check_history":
		response.Payload = "История заказов находится <a href=\"/render-auth/order-history\">здесь</a>"
		return response
	default:
		response.Action = "error"
		response.Payload = "Unknown action"
		return response
	}
}
