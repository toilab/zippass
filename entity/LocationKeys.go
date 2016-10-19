package entity

import (
)

// 地理情報機能情報
type Location struct {
	PassId				int		`json:"pass_id"`
	Altitude			int64	`json:"altitude"`							// 高度
	Latitude			int64	`json:"latitude"`							// 緯度
	Longitude			int64	`json:"longitude"`							// 経度
	RelevantText		string	`json:"relevant_text"`						// 地理情報にパスが反応したときにロック画面に表示されるテキスト
}

// Locationコンストラクタ
func NewLocation(altitude int64, latitude int64, longitude int64, relevantText string) *Location {
	ret := &Location{
		Altitude: altitude,
		Latitude: latitude,
		Longitude: longitude,
		RelevantText: relevantText,
	}
	return ret
}
