package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type RequestDuty struct {
	Users []string `json:"duty_users"`
	//Parameter
	//0 - duty only in workdays
	//1 - duty both workdays and weekdays
	//2 - duty 3 days, if duty weekday is friday
	Parameter int      `json:"parameter"`
	Tasks     []string `json:"duty_tasks"`
}

func makeRequest(requsers, reqTasks []string, parameter int) (string, error) {
	req := &RequestDuty{
		Users:     requsers,
		Parameter: parameter,
		Tasks:     reqTasks,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	onDutyService := os.Getenv("ON_DUTY_SERVICE")
	url := onDutyService + "/duty"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("Status code is not 200")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	message := string(body)
	return message, nil
}

func (ba *BotAPI) addDuty(params string) (message string) {
	///comedian add_duty @name1 @name2 @name3 , <parameter>, tasks...
	parameters := strings.Split(params, ",")
	if len(parameters) < 3 {
		return message
	}
	users := parameters[0]
	parameter, err := strconv.Atoi(parameters[1])
	if err != nil {
		return message
	}
	tasks := parameters[2]

	requsers := strings.Split(users, "")
	reqTasks := strings.Split(tasks, "")
	message, err = makeRequest(requsers, reqTasks, parameter)
	if err != nil {
		logrus.Errorf("Error making request to on-duty service: %v", err)
		return message
	}
	return message
}
