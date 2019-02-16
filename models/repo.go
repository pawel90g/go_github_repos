package models

type Repo struct {
	Id                int32
	NodeId            string
	Name              string
	Full_Name         string
	Private           bool
	Owner             User
	Html_Url          string
	Description       string
	Fork              bool
	Url               string
	Created_At        string
	Updated_At        string
	Pushed_At         string
	Git_Url           string
	Size              int32
	Stargazers_Count  int32
	Watchers_Count    int32
	Language          string
	Has_Issues        bool
	Has_Projects      bool
	Has_Downloads     bool
	Has_Wiki          bool
	Has_Pages         bool
	Forks_Count       int32
	Open_Issues_Count int32
	Forks             int32
	Open_Issues       int32
	Watchers          int32
	Default_Branch    string
}
