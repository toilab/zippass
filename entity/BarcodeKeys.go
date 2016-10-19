package entity

import (
)

// Barcode情報
type BarcodeKeys struct {
	PassId				int			`json:"pass_id"`						// パスのid
	BarcodeType         string	 	`json:"barcode_type"`					// バーコードタイプ(barcode/barcodes)
	AltText             string	 	`json:"alt_text"`						//
	Format       		string 		`json:"format"`							//
	Message        		string 		`json:"message"`						//
	MessageEncoding     string     	`json:"message_encoding"`				//
}

// コンストラクタ
func NewBarcodeKeys(passId int, barcodeType string, altText string,
					format string, message string, messageEncoding string) *BarcodeKeys {
	ret := &BarcodeKeys{
		PassId: passId,
		BarcodeType: barcodeType,
		AltText: altText,
		Format: format,
		Message: message,
		MessageEncoding: messageEncoding,
	}
	return ret
}
