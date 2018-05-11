package model

type Home struct {
	HomeId        uint32   `json:"home_id"`
	Title         string   `json:"title"`
	AppVersionMin string   `json:"app_version_min"`
	AppVersionMax string   `json:"app_version_max"`
	DeviceIds     []string `json:"device_ids"`
	Timestamp     int64    `json:"timestamp"`
}
