package data

type IPBlockList struct {
	Firstseen	string	`json:"firstseen"`
	DstIP		string	`json:"dstip"`
	DstPort		int		`json:"dstport"`
	LastOnline	string	`json:"lastonline"`
	Malware		string	`json:"malware"`
}

type User struct {
	Name	string	`json:"name"`
	Age		int		`json:"age"`
	Email	string	`json:"email"`
	Gender	bool	`json:"gender"`
}