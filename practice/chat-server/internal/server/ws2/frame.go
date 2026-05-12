package websocketlite

import (
	"encoding/binary"
	"io"
	"net"
)

func readFrame(conn net.Conn) ([]byte, error) {

	header := make([]byte, 2)

	_, err := io.ReadFull(conn, header)
	if err != nil {
		return nil, err
	}

	payloadLen := int(header[1] & 0x7F)

	maskKey := make([]byte, 4)

	_, err = io.ReadFull(conn, maskKey)
	if err != nil {
		return nil, err
	}

	payload := make([]byte, payloadLen)

	_, err = io.ReadFull(conn, payload)
	if err != nil {
		return nil, err
	}

	// unmask
	for i := 0; i < payloadLen; i++ {
		payload[i] ^= maskKey[i%4]
	}

	return payload, nil
}

func writeFrame(conn net.Conn, payload []byte) error {

	frame := []byte{}

	// FIN + TEXT FRAME
	frame = append(frame, 0x81)

	payloadLen := len(payload)

	frame = append(frame, byte(payloadLen))

	frame = append(frame, payload...)

	_, err := conn.Write(frame)

	return err
}
