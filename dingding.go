package logrus_dingding

import (
	"github.com/sirupsen/logrus"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	"errors"
	"time"
)

type DingHook struct {
	AppName string
	Robot string
}

func NewDingRobot(appname string, host string)*DingHook{
	return &DingHook{
		AppName: appname,
		Robot: host,
	}
}

func (hook *DingHook) Fire(entry *logrus.Entry) error{

	fields, _:= json.Marshal(entry.Data)
	fields, _ = json.Marshal(string(fields[:]))	// 为了转义
	tmp := strings.TrimRight(string(fields[1:]), "\"")	// 截断前后的双引号

	message := `{"msgtype": "markdown","markdown": {"title": "`+entry.Level.String()+`","text": "#### ` + entry.Level.String() + `\n ![screenshot](http://lorempixel.com/400/200/cats?t=`+time.Now().Format(time.RFC850)+`)\n\n `+ entry.Message + `\n\n `+tmp+`"}}`
	jsonStr := []byte(message)
	req, err := http.NewRequest("POST", hook.Robot, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var ddResponce struct {
		Errcode int
		Errmsg string
	}
	err = json.Unmarshal(body, &ddResponce)
	if err != nil {
		return err
	}

	if ddResponce.Errcode != 0 {
		return errors.New(ddResponce.Errmsg)
	}

	return nil
}

func (hook *DingHook) Levels()[]logrus.Level{
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}
