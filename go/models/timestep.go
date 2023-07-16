package models

// Timstep is the param for Tomorrow.io
// 
// Docs: https://docs.tomorrow.io/reference/weather-data-layers#timestep-availability
type Timestep string

const (
	Best Timestep = "best"
	// [-6h,+6h] 
	Minutely Timestep = "1m"
	// [-168h,+360h] | Free
	Hourly Timestep = "1h"
	// [-7d,+15d] | Free
	Daily Timestep = "1d"
)
