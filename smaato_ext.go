package openrtb

type SmaatoExt struct {
	Udi              *UniqueDevice `json:"udi,omitempty"`
	Carriername      string        `json:"carriername,omitempty"`
	XUidh            string        `json:"x_uidh,omitempty"`
	Operaminibrowser int8          `json:"operaminibrowser,omitempty"`
}
