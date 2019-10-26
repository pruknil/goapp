package hsm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ianlopshire/go-fixedwidth"
	breaker "github.com/sony/gobreaker"
	"net"
	"strconv"
	"strings"
	"time"
)

type HSM struct {
	IConnection
	*breaker.CircuitBreaker
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

func (h *HSM) ExecuteMessage(conn net.Conn, hexString string) (string, error) {
	body, err := h.CircuitBreaker.Execute(func() (interface{}, error) {
		return h.doExecute(conn, hexString)
	})
	return body.(string), err
}

func (h *HSM) CheckStatus() string {
	conn, _ := h.RequestConnection()
	str, _ := h.ExecuteMessage(conn, "01010000000101")
	inTextByte := []byte(str)
	c := &HSM_FN_01_Response{}
	fixedwidth.Unmarshal(inTextByte, c)
	return fmt.Sprintf("%#v", c)
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
