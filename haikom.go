package gomodhaikom

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const ErrorFile string = "haikom.go"

var (
	ErrorUserClientGeneral     error = errors.New("UserClient: general error")
	ErrorUserClientNoContact   error = errors.New("UserClient: no contact with server")
	ErrorUserClientWrongStatus error = errors.New("UserClient: wrong status code")
	ErrorUserClientNotValid    error = errors.New("UserClient: user not valid")
)

type UserInterface interface {
	GetUser(token, client, requestid string) (User, error)
}

type UserService struct {
	User UserInterface
}

type HaikomUser struct {
	User     string
	Password string
	Project  string
	Url      string
}

func NewUserService(u UserInterface) *UserService {
	return &UserService{
		User: u,
	}
}

func (h HaikomUser) GetUser(token, client, requestid string) (User, error) {

	const ErrorFunc string = "CheckHaikomToken"

	ErrorHaikom := errors.New("haikom: can't get answer from haikom")

	msg := fmt.Sprintf("<request><token>%s</token></request>", token)

	cc, err := http.NewRequest(http.MethodPost, h.Url, strings.NewReader(msg))
	if checkError(err) {
		return User{}, ErrorHaikom
	}
	cc.Header.Add("project", h.Project)
	cc.Header.Add("username", h.User)
	cc.Header.Add("password", h.Password)
	cc.Header.Add("Content-Type", "text/xml")

	resp, err := http.DefaultClient.Do(cc)
	if checkError(err) {
		return User{}, ErrorHaikom
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return User{}, ErrorUserClientWrongStatus
	}

	resBody, err := io.ReadAll(resp.Body)
	if checkError(err) {
		return User{}, ErrorUserClientNotValid
	}

	cr := User{}
	userXml := CustomerResponse{}
	err = xml.Unmarshal([]byte(resBody), &userXml)
	if err != nil {
		return cr, ErrorUserClientGeneral
	}
	if !userXml.Valid {
		return cr, ErrorUserClientNotValid
	}
	cr.MapXml(&userXml.User)
	cr.ClientIdentifier = client
	return cr, nil
}

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}
