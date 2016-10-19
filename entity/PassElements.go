package entity

import (
	"time"
)

// pass情報
type PassElements struct {
	PassId				int		`json:"pass_id"`
	// Standard Keys
	Description			string	`json:"description", datastore:",noindex"`	// このpassの簡潔な説明
	FormatVersion		int		`json:"format_ver"`							// ファイルフォーマットVer 現在1固定
	OrganizationName	string	`json:"org_name"`							// 組織名
	PassTypeId			string	`json:"passtype_id"`						// パス型識別子
	SerialNumber		string	`json:"serial_num"`							// シリアルNo
	TeamId				string	`json:"team_id"`							// チーム識別子

	// Associated App Keys
	AppLaunchURL		string	`json:"app_launch_url"`						// 連動Appを起動する際にアクセスするURL
	AssociatedStoreId	[]int	`json:"associated_store_id"`				// 連動Appに採番されたiTunes Store item ID

	// Companion App Keys
	UserInfo			string	`json:"user_info"`							// 連動Appのためのカスタム情報

	// Expiration Keys
	ExpirationDate		time.Time	`json:"expirationDate"`					// 有効期限日時 実際にJSONに格納する際のフォーマットはW3C date, as a string
	Voided				bool	`json:"voided"`								// パスが有効かどうか. true:無効 false:有効

	// Relevance Keys
	MaxDistance			int			`json:"max_distance"`					// ビーコンや位置情報にパスが反応する発信地からの最大距離(m).この番号はパスの初期値距離と比較され、より小さい値が使用される
	RelevantDate		time.Time	`json:"relevant_date"`					// パスが反応する日時. 実際にJSONに格納する際のフォーマットはW3C date, as a string

	// Pass Style(Style Key + Pass Structure)
	StyleType			string		`json:"style_type"`						// このPassのStyleKey -> boardingPass, coupon, eventTicket, generic, storeCardからひとつ
	TransitType			string		`json:"transit_type"`					// 搭乗券タイプのときのみ使用

	// Field
	BackgroundColor		string		`json:"background_color"`
	ForegroundColor		string		`json:"foreground_color"`
	GroupingId			string		`json:"grouping_id"`
	LabelColor			string		`json:"label_color"`
	LogoText			string		`json:"logo_text"`
	SuppressStripShine	bool		`json:"suppress_strip_shine"`

	// Web
	AuthenticationToken	string		`json:"authentication_token"`
	WebServiceURL		string		`json:"web_service_url"`

	// NFC
	Message				string		`json:"message"`
	EncryptionPublicKey	string		`json:"encryption_public_key"`
}

