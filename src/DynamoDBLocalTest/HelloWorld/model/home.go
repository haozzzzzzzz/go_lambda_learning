package model

import "time"

type HomeOnlineStateType uint8

const (
	HomeOnlineStateOffline HomeOnlineStateType = 0
	HomeOnlineStateOline   HomeOnlineStateType = 1
	HomeOnlineStateDefault HomeOnlineStateType = HomeOnlineStateOffline
)

// 计数器
type Counter struct {
	Name  string `json:"name"`
	Count uint32 `json:"count"`
}

// 首页模型
type Home struct {
	PartitionKey        uint32              `json:"partition_key"`         // 分区键
	HomeId              uint32              `json:"home_id"`               // 首页ID
	Title               string              `json:"title"`                 // 标题
	Style               uint32              `json:"style"`                 // 首页风格样式
	AppVersionMin       string              `json:"app_version_min"`       // app最小版本
	AppVersionMax       string              `json:"app_version_max"`       // app最大版本
	EffectiveTime       time.Time           `json:"effective_time"`        // 生效时间
	OnlineState         HomeOnlineStateType `json:"online_state"`          // 在线状态
	LastUpdateUid       string              `json:"last_update_uid"`       // 上次更新uid
	LastUpdateUsername  string              `json:"last_update_username"`  // 上次更新用户名
	LastUpdateTime      time.Time           `json:"last_update_time"`      // 上次更新时间
	LastUpdateTimestamp int64               `json:"last_update_timestamp"` // 上次更新时间戳
	DeviceIds           []string            `json:"device_ids"`            // 设备ID数组
	CreateTime          time.Time           `json:"create_time"`           // 当前时间
}

// 主题模型
type Theme struct {
	PartitionKey        uint32    `json:"partition_key"`         // 分区键
	ThemeId             uint32    `json:"theme_id"`              // 主题ID
	Title               string    `json:"title"`                 // 主题名称
	CoverPic            string    `json:"cover_pic"`             // 封面图
	Description         string    `json:"description"`           // 描述
	Type                uint8     `json:"type"`                  // 主题类型
	LastUpdateUid       string    `json:"last_update_uid"`       // 上次更新的用户ID
	LastUpdateUsername  string    `json:"last_update_username"`  // 上次更新的用户名
	LastUpdateTime      time.Time `json:"last_update_time"`      // 上次更新的时间
	LastUpdateTimestamp int64     `json:"last_update_timestamp"` // 上次更新时间戳
	CreateTime          time.Time `json:"create_time"`           // 记录创建时间
}

// 首页主题
type HomeTheme struct {
	HomeId    uint32 `json:"home_id"`    // 首页ID
	ThemeId   uint32 `json:"theme_id"`   // 主题ID
	SortOrder uint32 `json:"sort_order"` // 排序序号
}

// 主题资源类型的类型
type ThemeResourceTypeType uint8

const (
	ThemeResourceTypeMovie      ThemeResourceTypeType = 1 // 电影
	ThemeResourceTypeShortVideo ThemeResourceTypeType = 2 // 短视频
	ThemeResourceTypeMV         ThemeResourceTypeType = 3 // MV
)

// 主题资源
type ThemeResource struct {
	ThemeId    uint32    `json:"theme_id"`    // 主题ID
	ResourceId string    `json:"resource_id"` // 资源ID
	SortOrder  uint32    `json:"sort_order"`  // 排序序号
	Type       uint32    `json:"type"`        // 资源类型
	CreateTime time.Time `json:"create_time"` // 记录创建时间
}
