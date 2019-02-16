package models

type Email struct {
	Email      string
	Verified   bool
	Primary    bool
	Visibility string
}
