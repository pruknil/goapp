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
	conn, _ := h.RequestConnection()
	str, _ := h.doExecute(conn, "01010000000101")
	inTextByte := []byte(str)
	c := &HSM_FN_01_Response{}
	fixedwidth.Unmarshal(inTextByte, c)

	return fmt.Sprintf("%#v", c)
}

type HSM_FN_01_Response struct {
	ResponseHeader    string `fixed:"1,10"`
	ResponseLen       string `fixed:"11,12"`
	Fn                string `fixed:"13,14"`
	Rc                string `fixed:"15,16"`
	RAMStatus         string `fixed:"17,18"`
	ROMStatus         string `fixed:"19,20"`
	DESStatus         string `fixed:"21,22"`
	HostPortStatus    string `fixed:"23,24"`
	BatteryStatus     string `fixed:"25,26"`
	AESStatus         string `fixed:"27,28"`
	HardDiskStatus    string `fixed:"29,30"`
	RSAAccelerator    string `fixed:"31,32"`
	PerformanceLevel  string `fixed:"33,34"`
	ResetCount        string `fixed:"35,38"`
	CallsInLastMinute string `fixed:"39,46"`
	CallsInLast10Mins string `fixed:"47,54"`
	SoftwareIDLength  string `fixed:"55,56"`
	SoftwareID        string `fixed:"57,64"`
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
