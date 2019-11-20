package http

/*type ReqHeader struct {
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
}*/

type AQIRes []struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Aqi         int    `json:"aqi"`
	Type        string `json:"type"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
}
type KPeopleReq struct {
	EMPID  string `json:"EMP_ID"`
	COTPCD string `json:"CO_TP_CD"`
}

type KPeopleRes struct {
	EMPID           string      `json:"EMP_ID,omitempty"`
	THTTL           string      `json:"TH_TTL,omitempty"`
	THFRSTNM        string      `json:"TH_FRST_NM,omitempty"`
	THSURNM         string      `json:"TH_SUR_NM,omitempty"`
	ENTTL           string      `json:"EN_TTL,omitempty"`
	ENFRSTNM        string      `json:"EN_FRST_NM,omitempty"`
	ENSURNM         string      `json:"EN_SUR_NM,omitempty"`
	BRTHDT          interface{} `json:"BRTH_DT,omitempty"`
	IDENTNO         interface{} `json:"IDENT_NO,omitempty"`
	EMPDT           string      `json:"EMP_DT,omitempty"`
	EMPEPDT         string      `json:"EMP_EP_DT,omitempty"`
	EMPTPCD         string      `json:"EMP_TP_CD,omitempty"`
	EMPTPDSC        string      `json:"EMP_TP_DSC,omitempty"`
	EMPCTRTPCD      string      `json:"EMP_CTR_TP_CD,omitempty"`
	EMPLCSTPCD      string      `json:"EMP_LCS_TP_CD,omitempty"`
	EMPSTCD         string      `json:"EMP_ST_CD,omitempty"`
	DEPTID          string      `json:"DEPT_ID,omitempty"`
	THDEPTNM        string      `json:"TH_DEPT_NM,omitempty"`
	ENDEPTNM        string      `json:"EN_DEPT_NM,omitempty"`
	BSNLINEDEPTID   string      `json:"BSN_LINE_DEPT_ID,omitempty"`
	THBSNLINEDEPTNM string      `json:"TH_BSN_LINE_DEPT_NM,omitempty"`
	ENBSNLINEDEPTNM string      `json:"EN_BSN_LINE_DEPT_NM,omitempty"`
	PRNDEPTID       string      `json:"PRN_DEPT_ID,omitempty"`
	THPRNDEPTNM     string      `json:"TH_PRN_DEPT_NM,omitempty"`
	THPRNDEPTABR    string      `json:"TH_PRN_DEPT_ABR,omitempty"`
	ENPRNDEPTNM     string      `json:"EN_PRN_DEPT_NM,omitempty"`
	ENPRNDEPTABR    string      `json:"EN_PRN_DEPT_ABR,omitempty"`
	RGONCD          int         `json:"RGON_CD,omitempty"`
	ZONCD           int         `json:"ZON_CD,omitempty"`
	JOBCD           string      `json:"JOB_CD,omitempty"`
	THJOBNM         string      `json:"TH_JOB_NM,omitempty"`
	ENJOBNM         string      `json:"EN_JOB_NM,omitempty"`
	CORPTTLCD       string      `json:"CORP_TTL_CD,omitempty"`
	THCORPTTLNM     interface{} `json:"TH_CORP_TTL_NM,omitempty"`
	ENCORPTTLNM     interface{} `json:"EN_CORP_TTL_NM,omitempty"`
	KBNKAREANO      string      `json:"KBNK_AREA_NO,omitempty"`
	NTWCNTRCD       string      `json:"NTW_CNTR_CD,omitempty"`
	DEPTCNTRCD      string      `json:"DEPT_CNTR_CD,omitempty"`
	UNITCNTRCD      string      `json:"UNIT_CNTR_CD,omitempty"`
	SUBUNITCNTRCD   string      `json:"SUB_UNIT_CNTR_CD,omitempty"`
	KBNKBRNO        string      `json:"KBNK_BR_NO,omitempty"`
	CNTRTPCD        string      `json:"CNTR_TP_CD,omitempty"`
	EMPOFCRTPCD     string      `json:"EMP_OFCR_TP_CD,omitempty"`
	EMAILADR        string      `json:"EMAIL_ADR,omitempty"`
	OFFCPH1         string      `json:"OFFC_PH1,omitempty"`
	OFFCPH1EXT      string      `json:"OFFC_PH1_EXT,omitempty"`
	OFFCPH2         string      `json:"OFFC_PH2,omitempty"`
	OFFCPH2EXT      string      `json:"OFFC_PH2_EXT,omitempty"`
	OFFCPH3         string      `json:"OFFC_PH3,omitempty"`
	OFFCPH3EXT      string      `json:"OFFC_PH3_EXT,omitempty"`
	FAX             string      `json:"FAX,omitempty"`
	FAXEXT          string      `json:"FAX_EXT,omitempty"`
	MBLPH1          string      `json:"MBL_PH1,omitempty"`
	MBLPH2          string      `json:"MBL_PH2,omitempty"`
	TMTDT           interface{} `json:"TMT_DT,omitempty"`
	GNDCD           string      `json:"GND_CD,omitempty"`
	ETHNCGRP        string      `json:"ETHNC_GRP,omitempty"`
	COTPCD          string      `json:"CO_TP_CD,omitempty"`
	COTPNM          string      `json:"CO_TP_NM,omitempty"`
	COTPCAT         string      `json:"CO_TP_CAT,omitempty"`
	COLOGOSPPATH    string      `json:"CO_LOGO_SP_PATH,omitempty"`
	SPVSRID         string      `json:"SPVSR_ID,omitempty"`
	MONKF           string      `json:"MONK_F,omitempty"`
	PROVCD          string      `json:"PROV_CD,omitempty"`
	DEPTCD          string      `json:"DEPT_CD,omitempty"`
	PRBTDT          string      `json:"PRBT_DT,omitempty"`
	PRBTF           string      `json:"PRBT_F,omitempty"`
	DSCLNF          string      `json:"DSCLN_F,omitempty"`
	CLMF            string      `json:"CLM_F,omitempty"`
	PNDGDOCF        string      `json:"PNDG_DOC_F,omitempty"`
	QUALFDSC        string      `json:"QUALF_DSC,omitempty"`
	PPNTMS          string      `json:"PPN_TMS,omitempty"`
	SRCSTMID        int         `json:"SRC_STM_ID,omitempty"`
	IsUser          bool        `json:"IsUser,omitempty"`
	EMPBuilding     string      `json:"EMP_Building,omitempty"`
	EMPSecretary    interface{} `json:"EMP_Secretary,omitempty"`
	EMPSecretaryTEL interface{} `json:"EMP_Secretary_TEL,omitempty"`
	EMPCoverImage   string      `json:"EMP_CoverImage,omitempty"`
	EMPAchievement  interface{} `json:"EMP_Achievement,omitempty"`
	EMPExperties    interface{} `json:"EMP_Experties,omitempty"`
	EMPLineID       string      `json:"EMP_LineID,omitempty"`
	EMPNicknameTH   string      `json:"EMP_Nickname_TH,omitempty"`
	EMPNicknameEN   string      `json:"EMP_Nickname_EN,omitempty"`
	EMPOFFCPH1      string      `json:"EMP_OFFC_PH1,omitempty"`
	EMPOFFCPH2      interface{} `json:"EMP_OFFC_PH2,omitempty"`
	EMPMBLPH1       string      `json:"EMP_MBL_PH1,omitempty"`
	EMPMBLPH2       interface{} `json:"EMP_MBL_PH2,omitempty"`
	EMPOFFCFAX      interface{} `json:"EMP_OFFC_FAX,omitempty"`
	EMPInterests    interface{} `json:"EMP_Interests,omitempty"`
	COMPNAME        interface{} `json:"COMPNAME,omitempty"`
	JOBLAYER        string      `json:"JOB_LAYER,omitempty"`
	ENCNTRNMABR     interface{} `json:"EN_CNTR_NM_ABR,omitempty"`
	COLOGO          string      `json:"CO_LOGO,omitempty"`
}

type DopaReq struct {
	Nid        string `json:"nid"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	BirthDate  string `json:"birth_date"`
	LaserID    string `json:"laser_id"`
}

type DopaRes struct {
	IsChild1          bool   `json:"isChild1,omitempty"`
	IsChild2          bool   `json:"isChild2,omitempty"`
	IsChild3          bool   `json:"isChild3,omitempty"`
	IsCheckMiddleName bool   `json:"isCheckMiddleName,omitempty"`
	FaChkTaxYear      bool   `json:"faChkTaxYear,omitempty"`
	MoChkTaxYear      bool   `json:"moChkTaxYear,omitempty"`
	IsFaName          bool   `json:"isFaName,omitempty"`
	IsMoName          bool   `json:"isMoName,omitempty"`
	IsConnectBO       bool   `json:"isConnectBO,omitempty"`
	IsConnectMoi      bool   `json:"isConnectMoi,omitempty"`
	IsError           bool   `json:"isError"`
	Desc              string `json:"desc,omitempty"`
	ErrorDesc         string `json:"errorDesc,omitempty"`
}
