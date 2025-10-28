package gomodhaikom

import (
	"encoding/xml"
)

type User struct {
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Email            string `json:"email"`
	Locale           string `json:"locale"`
	Group            string `json:"group"`
	Role             string `json:"role"`
	ClientId         string `json:"clientId"`         //i.e 1
	ClientIdentifier string `json:"clientIdentifier"` //i.e billes
	ClientName       string `json:"clientName"`       //i.e Billes
	Access           string `json:"access"`
}

func (u *User) MapXml(dbm *UserXml) {
	u.Email = dbm.Email
	u.Firstname = dbm.Firstname
	u.Lastname = dbm.Lastname
	u.ClientId = dbm.CustomerId
	u.ClientName = dbm.Customername
	for _, prop := range dbm.Properties {
		u.Locale = getProperty("locale", prop.Property)
		u.Group = getProperty("bmsGroups", prop.Property)
		u.Role = getProperty("role", prop.Property)
	}
}

func getProperty(key string, list []Property) string {
	var returnValue string
	for _, value := range list {
		if value.Name == key {
			returnValue = value.Value
			break
		}
	}
	return returnValue
}

// XML part of haikom
type CustomerResponse struct {
	Response xml.Name `xml:"response"`
	Valid    bool     `xml:"valid"`
	User     UserXml  `xml:"user"`
}

type CustomersResponse struct {
	Response xml.Name `xml:"response"`
	Users    UsersXml `xml:"users"`
}

type UsersXml struct {
	Users xml.Name  `xml:"users"`
	User  []UserXml `xml:"user"`
}

type UserXml struct {
	XMLName      xml.Name     `xml:"user"`
	CustomerId   string       `xml:"customerid"`
	Username     string       `xml:"username"`
	Email        string       `xml:"email"`
	Firstname    string       `xml:"firstname"`
	Lastname     string       `xml:"lastname"`
	Customername string       `xml:"customername"`
	Office       string       `xml:"office"`
	OfficeId     string       `xml:"officeid"`
	Phone        string       `xml:"phone"`
	Mobilephone  string       `xml:"mobilephone"`
	Properties   []Properties `xml:"properties"`
}

type Properties struct {
	XMLName  xml.Name   `xml:"properties"`
	Property []Property `xml:"property"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}
