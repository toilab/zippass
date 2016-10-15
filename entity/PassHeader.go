package entity

import (
	"time"
)

// passのヘッダ情報
type PassHeader struct {
	PassId				int		`json:"pass_id"`							// パスのid
	PassName			string	`json:"pass_name"`							// パス名
	PassDescription		string	`json:"pass_description"`					// パス説明
	PassUrl				string	`json:"pass_url"`							// パスの実態ファイル格納場所（DriveのURL）
	OwnerUserId			int		`json:"owner_user_id"`						// パスの所有者ID
	CreatePassCnt		int		`json:"create_pass_cnt"`					// パスの作成回数
	CreateDt			time.Time	`json:"create_dt"`						// 作成日時
	CreateUserId		int		`json:"create_user_id"`						// 作成ユーザID
	UpdateDt			time.Time	`json:"update_dt"`						// 更新日時
	UpdateUserId		int		`json:"update_user_id"`						// 更新ユーザID
	DelFlg				bool	`json:"del_flg"`							// 論理削除フラグ
}

// コンストラクタ
func NewPassHeader(passId int, passName string, passDescription string, passUrl string,
				   ownerUserId int, createPassCnt int, createDt time.Time, createUserId int,
				   updateDt time.Time, updateUserId int, delFlg bool ) *PassHeader {
	ret := &PassHeader{
		PassId: passId,
		PassName: passName,
		PassDescription: passDescription,
		PassUrl: passUrl,
		OwnerUserId: ownerUserId,
		CreatePassCnt: createPassCnt,
		CreateDt: createDt,
		CreateUserId: createUserId,
		UpdateDt: updateDt,
		UpdateUserId: updateUserId,
		DelFlg: delFlg,
	}
	return ret
}



