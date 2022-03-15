package validate

import "github.com/gorilla/websocket"

// IsWSError 判断是否是ws错误
func IsWSError(ws *websocket.Conn, err error) bool {
	if err != nil {
		if _, ok := err.(*websocket.CloseError); ok {
			return true
		}
		return false
	}
	return false
}
