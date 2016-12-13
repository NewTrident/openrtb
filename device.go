package openrtb

// 3.2.11 Object: Device
//
// This object provides information pertaining to the device through which the user is interacting. Device
// information includes its hardware, platform, location, and carrier data. The device can refer to a mobile
// handset, a desktop computer, set top box, or other digital device.
type Device struct {

	// Attribute:
	//   ua
	// Type:
	//   string; recommended
	// Description:
	//   Browser user agent string.
	UA string `json:"ua,omitempty"`

	// Attribute:
	//   geo
	// Type:
	//   object; recommended
	// Description:
	//   Location of the device assumed to be the user’s current
	//   location defined by a Geo object (Section 3.2.12).
	Geo *Geo `json:"geo,omitempty"`

	// Attribute:
	//   dnt
	// Type:
	//   integer; recommended
	// Description:
	//   Standard “Do Not Track” flag as set in the header by the
	//   browser, where 0 = tracking is unrestricted, 1 = do not track.
	DNT int8 `json:"dnt,omitempty"`

	// Attribute:
	//   lmt
	// Type:
	//   integer; recommended
	// Description:
	//   “Limit Ad Tracking” signal commercially endorsed (e.g., iOS,
	//   Android), where 0 = tracking is unrestricted, 1 = tracking must
	//   be limited per commercial guidelines.
	Lmt int8 `json:"lmt,omitempty"`

	// Attribute:
	//   ip
	// Type:
	//   string; recommended
	// Description:
	//   IPv4 address closest to device.
	IP string `json:"ip,omitempty"`

	// Attribute:
	//   ipv6
	// Type:
	//   string
	// Description:
	//   IP address closest to device as IPv6.
	IPv6 string `json:"ipv6,omitempty"`

	// Attribute:
	//   devicetype
	// Type:
	//   integer
	// Description:
	//   The general type of device. Refer to List 5.17.
	DeviceType int8 `json:"devicetype,omitempty"`

	// Attribute:
	//   make
	// Type:
	//   string
	// Description:
	//   Device make (e.g., “Apple”).
	Make string `json:"make,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model (e.g., “iPhone”).
	Model string `json:"model,omitempty"`

	// Attribute:
	//   os
	// Type:
	//   string
	// Description:
	//   Device operating system (e.g., “iOS”).
	OS string `json:"os,omitempty"`

	// Attribute:
	//   osv
	// Type:
	//   string
	// Description:
	//   Device operating system version (e.g., “3.1.2”).
	OSV string `json:"osv,omitempty"`

	// Attribute:
	//   hwv
	// Type:
	//   string
	// Description:
	//   Hardware version of the device (e.g., “5S” for iPhone 5S).
	HWV string `json:"hwv,omitempty"`

	// Attribute:
	//   h
	// Type:
	//   integer
	// Description:
	//   Physical height of the screen in pixels.
	H uint64 `json:"h,omitempty"`

	// Attribute:
	//   w
	// Type:
	//   integer
	// Description:
	//   Physical width of the screen in pixels.
	W uint64 `json:"w,omitempty"`

	// Attribute:
	//   ppi
	// Type:
	//   integer
	// Description:
	//   Screen size as pixels per linear inch.
	PPI uint64 `json:"ppi,omitempty"`

	// Attribute:
	//   pxratio
	// Type:
	//   float
	// Description:
	//   The ratio of physical pixels to device independent pixels.
	PxRatio float64 `json:"pxratio,omitempty"`

	// Attribute:
	//   js
	// Type:
	//   integer
	// Description:
	//   Support for JavaScript, where 0 = no, 1 = yes.
	JS       int8 `json:"js,omitempty"`
	Geofetch int8 `json:"geofetch,omitempty"`
	// Attribute:
	//   flashver
	// Type:
	//   string
	// Description:
	//   Version of Flash supported by the browser.
	FlashVer string `json:"flashver,omitempty"`

	// Attribute:
	//   language
	// Type:
	//   string
	// Description:
	//   Browser language using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`

	// Attribute:
	//   carrier
	// Type:
	//   string
	// Description:
	//   Carrier or ISP (e.g., “VERIZON”). “WIFI” is often used in mobile
	//   to indicate high bandwidth (e.g., video friendly vs. cellular).
	Carrier string `json:"carrier,omitempty"`

	// Attribute:
	//   connectiontype
	// Type:
	//   integer
	// Description:
	//   Network connection type. Refer to List 5.18.
	Connectiontype int8 `json:"connectiontype,omitempty"`

	// Attribute:
	//   ifa
	// Type:
	//   string
	// Description:
	//   ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`

	// Attribute:
	//   didsha1
	// Type:
	//   string
	// Description:
	//   Hardware device ID (e.g., IMEI); hashed via SHA1.
	DIDSHA1 string `json:"didsha1,omitempty"`

	// Attribute:
	//   didmd5
	// Type:
	//   string
	// Description:
	//  Hardware device ID (e.g., IMEI); hashed via MD5.
	DIDMD5 string `json:"didmd5,omitempty"`

	// Attribute:
	//   dpidsha1
	// Type:
	//   string
	// Description:
	//   Platform device ID (e.g., Android ID); hashed via SHA1.
	DPIDSHA1 string `json:"dpidsha1,omitempty"`

	// Attribute:
	//   dpidmd5
	// Type:
	//   string
	// Description:
	//   Platform device ID (e.g., Android ID); hashed via MD5.
	DPIDMD5 string `json:"dpidmd5,omitempty"`

	// Attribute:
	//   macsha1
	// Type:
	//   string
	// Description:
	//   MAC address of the device; hashed via SHA1.
	MACSHA1 string `json:"macsha1,omitempty"`

	// Attribute:
	//   macmd5
	// Type:
	//   string
	// Description:
	//   MAC address of the device; hashed via MD5.
	MACMD5 string `json:"macmd5,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext,omitempty"`
}
