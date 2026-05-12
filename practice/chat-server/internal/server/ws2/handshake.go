package ws2

import (
  "crypto/sha1"
  "encoding/base64"
)

const magicGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func createWebSocketAcceptKey(key string) string {
	hash := sha1.Sum([]byte(key + magicGUID))
	return base64.StdEncoding.EncodeToString(hash[:])
}
