package entity

import (
	"time"
	"encoding/asn1"
)


// フィールド設定情報
type FieldKeys struct {
	PassId					int			`json:"pass_id"`
	FieldType				string		`json:"field_type"`
	AttributeValue			string			`json:"attribute_value"`								// BLEビーコンのmajor-ID 実際の型はuint16だが、datastoreが対応していないためDTOはintにする
	ChangeMessage			string			`json:"change_message"`								// BLEビーコンのminor-ID 実際の型はuint16だが、datastoreが対応していないためDTOはintにする
	DataDetectorTypes		[]string	`json:"data_detector_types"`						// BLEビーコンのユニークID
	Key						string		`json:"key"`						// BLEビーコンにパスが反応したときにロック画面に表示されるテキスト
	Label					[]string	`json:"label"`
	TextAlignment			string		`json:"text_alignment"`
	Value					string		`json:"value"`
	IgnoresTimeZone			bool		`json:"ignores_time_zone"`
	IsRelative				bool		`json:"is_relative"`
	TimeStyle				string		`json:"time_style"`
	NumberStyle				string		`json:"number_style"`
}

// コンストラクタ
func NewFieldKeys(passId int, fieldType string, attributeValue string, changeMessage string,
				  dataDetectorTypes []string, key string, label []string, textAlignment string,
				  value string, ignoresTimeZone bool, isRelative bool, timeStyle string, numberStyle string ) *FieldKeys {
	ret := &FieldKeys{
		PassId: passId,
		FieldType: fieldType,
		AttributeValue: attributeValue,
		ChangeMessage: changeMessage,
		DataDetectorTypes: dataDetectorTypes,
		Key: key,
		Label: label,
		TextAlignment: textAlignment,
		Value: value,
		IgnoresTimeZone: ignoresTimeZone,
		IsRelative: isRelative,
		TimeStyle: timeStyle,
		NumberStyle: numberStyle,
	}
	return ret
}

