package models

type User struct {
	Login        string
	Id           int32
	Node_ID      string
	Avatar_Url   string
	Gravatar_ID  string
	Url          string
	Type         string
	Site_Admin   bool
	Name         string
	Company      string
	Blog         string
	Location     string
	Email        string
	Hireable     string
	Bio          string
	Public_Repos int32
	Public_Gists int32
	Followers    int32
	Following    int32
	Created_At   string
	Updated_At   string
}
