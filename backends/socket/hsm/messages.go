package hsm

import (
	"encoding/hex"
	"errors"
	"fmt"
	breaker "github.com/sony/gobreaker"
	"net"
	"strconv"
	"strings"
	"time"
)

type IConnection interface {
	Open() error
	Close()
	RequestConnection() (net.Conn, error)
}

type IHSMService interface {
	CheckStatus() string
}

func (h *HSM) ExecuteMessage(conn net.Conn, hexString string) (string, error) {
	body, err := h.CircuitBreaker.Execute(func() (interface{}, error) {
		return h.doExecute(conn, hexString)
	})
	return body.(string), err
}

func NewHSM(b IConnection) IHSMService {
	var st breaker.Settings
	st.Name = "HSM"
	st.Timeout = 3
	st.ReadyToTrip = func(counts breaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	cb := breaker.NewCircuitBreaker(st)
	return &HSM{IConnection: b, CircuitBreaker: cb}
}

type HSM struct {
	IConnection
	*breaker.CircuitBreaker
}

func (h *HSM) CheckStatus() string {
	c, _ := h.RequestConnection()
	h.doExecute(c, "")
	return ""
}

func (h *HSM) doExecute(conn net.Conn, hexString string) (string, error) {

	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	timeout, _ := time.ParseDuration("5s")
	err = conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return "", fmt.Errorf("SetReadDeadline failed:%s\n", err.Error())
	}
	_, err = conn.Write(byteArray)
	if err != nil {
		return "", fmt.Errorf("Write to server failed:%s\n", err.Error())
	}

	reply := make([]byte, 512)
	_, err = conn.Read(reply)
	if err != nil {
		return "", fmt.Errorf("Read buffer failed:%s", err.Error())
	}

	var replyHexString string
	replyHexString = hex.EncodeToString(reply)

	var headerLen = 12
	if len(replyHexString) < headerLen {
		return "", errors.New("invalid response message," + replyHexString)
	}
	var bodyLen int
	l64 := hex2int64(replyHexString[11:13])
	strInt64 := strconv.FormatInt(int64(l64), 10)
	bodyLen, err = strconv.Atoi(strInt64)
	if err != nil {
		return "", errors.New("invalid response message," + replyHexString + err.Error())
	}
	var replyActLen = headerLen + bodyLen + 1
	responseHexString := replyHexString[0:replyActLen]
	responseHexString = strings.ToUpper(responseHexString)
	return responseHexString, nil
}

func hex2int64(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}
