package service

import "time"

type ReqHeader struct {
	FuncNm       string `json:"funcNm"`
	RqUID        string `json:"rqUID"`
	RqDt         string `json:"rqDt"`
	RqAppID      string `json:"rqAppId"`
	UserLangPref string `json:"userLangPref"`
}

type ResHeader struct {
	FuncNm     string     `json:"funcNm"`
	RqUID      string     `json:"rqUID"`
	RsAppID    string     `json:"rsAppId"`
	RsUID      string     `json:"rsUID"`
	RsDt       time.Time  `json:"rsDt"`
	StatusCode string     `json:"statusCode"`
	ErrorVect  *ErrorVect `json:"errorVect,omitempty"`
}

type ErrorVect struct {
	Error []Error `json:"error"`
}

type Error struct {
	ErrorAppID    string `json:"errorAppId"`
	ErrorAppAbbrv string `json:"errorAppAbbrv"`
	ErrorCode     string `json:"errorCode"`
	ErrorDesc     string `json:"errorDesc"`
	ErrorSeverity string `json:"errorSeverity"`
}

type ReqMsg struct {
	Header ReqHeader   `json:"Header"`
	Body   interface{} `json:"Body"`
}

type ResMsg struct {
	Header ResHeader   `json:"Header"`
	Body   interface{} `json:"Body,omitempty"`
}
