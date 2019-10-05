package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type ManagerInfo struct {
	Name     string
	Age      int
	HireDate time.Time
}

func (c ManagerInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ManagerInfo) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
