# logrus_dingding
dingding robot notification

[doc](https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.lmydCN&treeId=257&articleId=105735&docType=1)

code:
```golang
package main

import (
	"github.com/sirupsen/logrus"
	ding "github.com/kazaff/logrus_dingding"
)

func main(){
	log := logrus.New()
	log.Hooks.Add(ding.NewDingRobot("app name", "robot hook link"))

	log.WithField("a","b").Fatal("hola!")
}
```
