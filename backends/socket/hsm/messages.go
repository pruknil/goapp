package hsm

import (
	"strconv"
	"strings"
)

type HSMStatusResponse struct {
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

func hex2int64(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}
