package main

type EtsItems struct {
	Items []EtsItem `xml:"Row"`
}

type EtsItem struct {
	NtinId       int    `xml:"NtinId"`
	Serial       string `xml:"Serial"`
	Status       int    `xml:"Status"`
	ParentNtinId int    `xml:"ParentNtinId"`
	ParentSerial string `xml:"ParentSerial"`
	Type         int    `xml:"Type"`
	Notification int    `xml:"Notification"`
	Flags        string `xml:"Flags"`
	HelperCode   string `xml:"HelperCode"`
}
