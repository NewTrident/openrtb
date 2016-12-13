package openrtb

type UniqueDevice struct {
	IFA           string `json:"ifa,omitempty"`
	ANDROIDID     string `json:"androidid,omitempty"`
	ANDROIDIDMD5  string `json:"androididmd5,omitempty"`
	ANDROIDIDSHA1 string `json:"androididsha1,omitempty"`
	IMEI          string `json:"imei,omitempty"`
	IMEIMD5       string `json:"imeimd5,omitempty"`
	IMEISHA1      string `json:"imeisha1,omitempty"`
	UDIDMD5       string `json:"udidmd5,omitempty"`
	UDIDSHA1      string `json:"udidsha1,omitempty"`
	MACMD5        string `json:"macmd5,omitempty"`
	MACSHA1       string `json:"macsha1,omitempty"`
	ODIN          string `json:"odin,omitempty"`
	OPENUDID      string `json:"openudid,omitempty"`
	IDFA          string `json:"idfa,omitempty"`
	IDFAMD5       string `json:"idfamd5,omitempty"`
	IDFASHA1      string `json:"idfasha1,omitempty"`
	IDFATRACKING  int8   `json:"idfatracking,omitempty"`
	GOOGLEADID    string `json:"googleadid,omitempty"`
	GOOGLEDNT     int8   `json:"googlednt,omitempty"`
	BBID          string `json:"bbid,omitempty"`
	ATUID         string `json:"atuid,omitempty"`
}
