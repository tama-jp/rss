package message

const (
	API_LIFE_TIME_ACCESS_TOKEN = 30 * 24 * 60 * 60
	//API_LIFE_TIME_ACCESS_TOKEN        = 30
	REGEX_CHARACTER_NO_SPECIAL  = "^([一-龠ぁ-ゔァ-ヴーa-zA-Z0-9ａ-ｚＡ-Ｚ０-９々〆〤\\s\\/　]+)$"
	REGEX_IGNORE_JAPANESS       = "^([^一-龠ぁ-ゔァ-ヴ]+)$"
	REGEX_HAS_CHARACTER_SPECIAL = "^([一-龠ぁ-ゔァ-ヴーa-zA-Z0-9ａ-ｚＡ-Ｚ０-９々〆〤\\\\s\\\\/　.,-。、／ー・]+)$"
	REGEX_EMAIL                 = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT2                = "2006-01-02T15:04:05Z07:00"

	DATE_TIME_FORMAT                  = "2006-01-02 15:04:05"
	DATE_TIME_FORMAT_YYYYMMDD_HHMM    = "2006-01-02 15:04"
	DATE_TIME_FORMAT_MILLI_SEC_FORMAT = "2006-01-02 15:04:05.000"
	TIME_FORMAT_HHMM                  = "15:04"
	REGEX_VALID_USER_ID               = `^[A-Za-z0-9\+\-\._!$#]+$`
	NOTIFY_MEDIA_ALL                  = "all"
	NOTIFY_MEDIA_ADMIN                = "admin"
	NOTIFY_MEDIA_MAIL                 = "mail"
	CONTRACT_NOTICES_BOTH             = "both"
	CONTRACT_NOTICES_MAIL             = "mail"
	CONTRACT_NOTICES_INFO             = "info"
	X_USER_NAME                       = "x-user-name"
	X_PASSWORD                        = "x-password"

	DATABASE_DATA_NOT_FOUND = "database data not found"
)
