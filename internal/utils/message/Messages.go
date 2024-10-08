package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	STS000 = "STS000"
	STS001 = "STS001"
	STS002 = "STS002"

	ERR000 = "ERR000"
	ERR001 = "ERR001"
	ERR002 = "ERR002"
	ERR003 = "ERR003"
	ERR004 = "ERR004"
	ERR005 = "ERR005"
	ERR006 = "ERR006"
	ERR007 = "ERR007"
	ERR008 = "ERR008"
	ERR009 = "ERR009"
	ERR010 = "ERR010"
	ERR011 = "ERR011"
	ERR012 = "ERR012"
	ERR013 = "ERR013"
	ERR014 = "ERR014"
	ERR015 = "ERR015"
	ERR016 = "ERR016"
	ERR017 = "ERR017"
	ERR018 = "ERR018"
	ERR019 = "ERR019"
	ERR020 = "ERR020"
	ERR021 = "ERR021"
	ERR022 = "ERR022"
	ERR023 = "ERR023"
	ERR024 = "ERR024"
	ERR026 = "ERR026"
	ERR027 = "ERR027"
	ERR028 = "ERR028"
	ERR029 = "ERR029"
	ERR030 = "ERR030"
	ERR031 = "ERR031"
	ERR032 = "ERR032"
	ERR033 = "ERR033"
	ERR034 = "ERR034"
	ERR035 = "ERR035"
	ERR036 = "ERR036"
	ERR037 = "ERR037"
	ERR038 = "ERR038"
	ERR039 = "ERR039"
	ERR040 = "ERR040"
	ERR041 = "ERR041"
	ERR042 = "ERR042"
	ERR043 = "ERR043"
	ERR044 = "ERR044"
	ERR045 = "ERR045"
	ERR046 = "ERR046"
	ERR047 = "ERR047"
	ERR048 = "ERR048"
	ERR049 = "ERR049"
	ERR050 = "ERR050"
	ERR051 = "ERR051"
	ERR052 = "ERR052"
	ERR053 = "ERR053"
	ERR054 = "ERR054"
	ERR055 = "ERR055"
	ERR056 = "ERR056"
	ERR057 = "ERR057"
	ERR999 = "ERR999"

	USER_ID  = "ユーザID"
	PASSWORD = "パスワード"
)

var MsgFlags = map[string]string{
	// 正常系
	STS000: "正常終了",
	STS001: "データが取得出来ませんでした",
	STS002: "データが存在しません",

	// エラー系
	ERR000: "ログインエラー。",
	ERR001: "必須項目です",
	ERR002: "%sは%d文字以下で入力してください",
	ERR003: "%sは8文字以上256文字以内で入力してください",
	ERR004: "アクセストークンがありません",
	ERR005: "アクセストークンが違います",
	ERR006: "APIが見つかりません",
	ERR007: "数字のみです",
	ERR008: "ユーザーに権限がありません",
	ERR009: "パスワードがちがいます",
	ERR010: "年を入れてください",
	ERR011: "月は、1〜12内で入力してください",
	ERR012: "日は、1〜31内で入力してください",
	ERR013: "名前がおかしいです",
	ERR014: "1(公休日),2(指定休日)を入れてください",
	ERR015: "アクセストークンが有効期限が過ぎています。",
	ERR016: "削除できません。",
	//ERR016: "ログインできません。",

	//ERR001: "通信タイムアウト",
	//ERR002: "パラメータエラー",
	//ERR003: "すでに登録済みの言語コード組み合わせパターンです。",
	//ERR004: "%sが存在しません",
	//ERR005: "パラメータのフォーマットが不正です",
	//ERR006: "%sは%d文字以下で入力してください",
	//ERR007: "ファイル形式が不正の為インポート出来ません",
	//ERR008: "新パスワードと確認用が一致しません",
	//ERR009: "パスワード変更をしているユーザーはログインしているユーザーではありません",
	//ERR010: "こちらのメールアドレスは登録されておりません",
	//ERR011: "ユーザーIDもしくはパスワードが一致しません",
	//ERR013: "値が範囲外です",
	//ERR014: "無効な言語です",
	//ERR015: "指定の%sは登録済みです",
	//ERR016: "開始日From≦開始日Toとなるように入力してください",
	//ERR017: "",
	//ERR018: "取り込むファイルを選択してください",
	//ERR019: "ヘッダーの項目名が不正の為インポート出来ません",
	//ERR020: "設定項目が不足しているもしくは、超過している為インポート出来ません",
	//ERR021: "ヘッダーの項目がない為インポート出来ません",
	//ERR022: "取り込むファイルの内容が不正の為インポート出来ません",
	//ERR023: "%s行目%s列目にブランクの%sがある為インポート出来ません",
	//ERR024: "%s行目%s列目の%s形式が不正の為インポート出来ません",
	//ERR026: "{1}の説明を編集の前に{0}の説明を保存してください",
	//ERR027: "送信日を選択してください",
	//ERR028: "過去日が指定されています。送信日は未来日を指定してください",
	//ERR029: "<h><p><br/>style (style=)のhtmlタグを指定してください",
	//ERR030: "システム一括設定されています",
	//ERR031: "WEB一括設定されています",
	//ERR032: "管理画面一括設定されています",
	//ERR033: "指定言語一括設定されています",
	//ERR034: "アカウントを無効にしてから削除してください。",
	//ERR035: "メールアドレスの形式が不正です",
	//ERR036: "すでにご登録頂いているメールアドレスです",
	//ERR037: "権限が選択されていません",
	//ERR038: "設定されているアカウントがあるため削除できません",
	//ERR039: "%sは8文字以上256文字以内で入力してください",
	//ERR040: "メール送信失敗",
	//ERR041: "%sには特殊文字を含めないでください",
	//ERR042: "リンクの有効期限切れています。もう一度再設定を行ってください",
	//ERR043: "パスワードが正しくありません",
	//ERR044: "指定言語は契約している言語でなければなりません",
	//ERR045: "%sは%sより大きい値を入力できません。",
	//ERR046: "%sに日本語が入力されません",
	//ERR047: "指定のURLはキャッシュ対象外設定済です。",
	//ERR048: "%sには半角英数字と記号(+-_.!$#/)のみが使用可能です。",
	//ERR049: "指定のURLは翻訳除外対象設定済です。",
	//ERR050: "同一権限名称の登録は出来ません。",
	//ERR051: "辞書登録が完了しませんでした。",
	//ERR052: "IP制限エラー",
	//ERR053: "リクエストヘッダに%sがありません。",
	//ERR054: "ファイルのインポートに失敗しました。",
	//ERR999: "その他システムエラー",
}

func GetMsg(code string, params ...interface{}) string {
	msg, ok := MsgFlags[code]
	if ok {
		if len(params) > 0 {
			msg = fmt.Sprintf(msg, params...)
		}
		return msg
	}

	return ""
}

func GetApiNameAndTime(c *gin.Context, msg string) string {
	retMsg := ""
	apiName := c.Request.URL.String()

	t := time.Now()
	retMsg = fmt.Sprintf("%s [%s %s]", msg, apiName, t.Format("2006/01/02 15:04:05"))

	return retMsg
}
