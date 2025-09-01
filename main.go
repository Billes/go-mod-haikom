package gomodhaikom

import (
	"encoding/xml"
	"errors"
	"fmt"

	papertrail "github.com/Billes/go-mod-papertrail"
	"github.com/Billes/proddata/env"
	"github.com/Billes/proddata/models"
	"github.com/gofiber/fiber/v2"
)

const ErrorFile string = "clients/haikom.go"

var (
	ErrorUserClientGeneral   error = errors.New("UserClient: general error")
	ErrorUserClientNoContact error = errors.New("UserClient: no contact with server")
	ErrorUserClientNotValid  error = errors.New("UserClient: user not valid")
)

var UserClient *UserService

type UserInterface interface {
	GetUser(token, client, requestid string) (models.User, error)
}

type UserService struct {
	User UserInterface
}

type HaikomUser struct {
	user     string
	password string
}

func NewHaikomUserClient(u UserInterface) *UserService {
	return &UserService{
		User: u,
	}
}

func (h HaikomUser) GetUser(token, client, requestid string) (models.User, error) {

	const ErrorFunc string = "CheckHaikomToken"

	ErrorHaikom := errors.New("haikom: can't get answer from haikom")
	cr := models.User{}

	msg := fmt.Sprintf("<request><token>%s</token></request>", token)
	a := fiber.AcquireAgent()
	req := a.Request()
	req.SetRequestURI(env.HAIKOM_URL)
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.SetContentType(fiber.MIMETextXML)
	req.Header.Add("project", "billes")
	req.Header.Add("username", getEnv("HAIKOM_USER"))
	req.Header.Add("password", getEnv("HAIKOM_PASSWORD"))
	a = a.Body([]byte(msg))
	if err := a.Parse(); err != nil {
		papertrail.Error([]string{ErrorFile, ErrorFunc, "Parse()", requestid}, "can't connect to haikom", err.Error())
		return cr, ErrorHaikom
	}
	statusCode, resultBody, _ := a.String()
	if statusCode != fiber.StatusOK {
		papertrail.Error([]string{ErrorFile, ErrorFunc}, fmt.Sprintf("got statuscode %d", statusCode), "")
		return cr, ErrorUserClientNoContact
	}
	userXml := models.CustomerResponse{}
	err := xml.Unmarshal([]byte(resultBody), &userXml)
	if err != nil {
		papertrail.Error([]string{ErrorFile, ErrorFunc, "xml.Unmarshal", requestid}, "can't parse xml data", err.Error())
		return cr, ErrorUserClientGeneral
	}
	if !userXml.Valid {
		return cr, ErrorUserClientNotValid
	}
	cr.MapXml(&userXml.User)
	return cr, nil
}

func getEnv(name string) string {
	return env.GetEnv(name, "")
}
