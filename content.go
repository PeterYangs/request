package request

import (
	"encoding/json"
	"time"
)

type content struct {
	content []byte
	request *request
}

func (c content) ToString() string {

	return string(c.content)
}

// ToJsonMap 转map
func (c content) ToJsonMap() (map[string]interface{}, error) {

	var jsons map[string]interface{}

	err := json.Unmarshal(c.content, &jsons)

	if err != nil {

		return map[string]interface{}{}, err
	}

	return jsons, nil

}

// ToJsonStruct 转结构体
func (c content) ToJsonStruct(st interface{}) error {

	err := json.Unmarshal(c.content, st)

	if err != nil {

		return err
	}

	return nil
}

// Time 获取响应时间
func (c content) Time() time.Duration {

	return c.request.responseTime
}
