package entity

import (

)

// pass情報
type PassInfo struct {
	PassTypeID			string	`json:"passtype_id"`						// パス型識別子
	SerialNumber		string	`json:"serial_num"`							// シリアルNo
	TeamID				string	`json:"team_id"`							// チーム識別子
	FormatVersion		int		`json:"format_ver"`							// ファイルフォーマットVer 現在1固定
	Description			string	`json:"description", datastore:",noindex"`	// このpassの簡潔な説明
	OrganizationName	string	`json:"org_name"`							// 組織名
}