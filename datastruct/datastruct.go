package datastruct

import (
	"time"
)

// Data ...
type Data struct {
	Datetime  time.Now
	VehicleID string
	Speed     float32
}

t := time.Now()
	dt := t.Format("01.02.2006 15:04:05")
	fmt.Println(dt)