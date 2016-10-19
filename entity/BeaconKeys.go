package entity


// ビーコン機能情報
type Beacon struct {
	PassId				int		`json:"pass_id"`
	Major				int		`json:"major"`								// BLEビーコンのmajor-ID 実際の型はuint16だが、datastoreが対応していないためDTOはintにする
	Minor				int		`json:"minor"`								// BLEビーコンのminor-ID 実際の型はuint16だが、datastoreが対応していないためDTOはintにする
	ProximityUUID		string	`json:"proximity_uuid"`						// BLEビーコンのユニークID
	RelevantText		string	`json:"relevant_text"`						// BLEビーコンにパスが反応したときにロック画面に表示されるテキスト
}

// Beaconコンストラクタ
func NewBeacon(passId int, major int, minor int, proximityUUID string, relevantText string) *Beacon {
	ret := &Beacon{
		PassId: passId,
		Major: major,
		Minor: minor,
		ProximityUUID: proximityUUID,
		RelevantText: relevantText,
	}
	return ret
}

