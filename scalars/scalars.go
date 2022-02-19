package scalars

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
)

type Json map[string]interface{}

func (js *Json) UnmarshalGQL(v interface{}) error {
	data, ok := v.(map[string]interface{})
	if !ok {
		return fmt.Errorf("type recieved is not of type map[string]interface{}")
	}
	*js = data
	return nil
}

func (js Json) MarshalGQL(w io.Writer) {
	b, err := json.Marshal(js)
	fmt.Println(err)
	w.Write(b)
}

func (js Json) Value() (driver.Value, error) {
	return json.Marshal(js)
}

func (js *Json) Scan(v interface{}) error {
	data, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("type recieved is not of type map[string]interface{}")
	}
	json.Unmarshal(data, js)
	return nil
}
