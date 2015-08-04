package models

import (
	"time"
)

type Sessions struct {
	Id     string
	Expiry time.Time
}
