package data

type IPBlockList struct {
	Firstseen  string `json:"firstseen"`
	DstIP      string `json:"dstip"`
	DstPort    int    `json:"dstport"`
	LastOnline string `json:"lastonline"`
	Malware    string `json:"malware"`
}

type Users struct {
	Users []User `xml:"User"`
}

type User struct {
	Name   string  `json:"name" xml:"Name"`
	Age    int     `json:"age" xml:"Age"`
	Email  string  `json:"email" xml:"Email"`
	Gender bool    `json:"gender" xml:"Gender"`
	Test   float64 `json:"test" xml:"Test"`
}
