package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tama-jp/rss/internal/frameworks/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogBase struct {
	FileName   string `toml:"file_name"`
	MaxSize    int    `toml:"max_size"`
	MaxBackups int    `toml:"max_backups"`
	MaxAge     int    `toml:"max_age"`
}

func NewLogger(config *config.Config) (*LogBase, error) {

	fmt.Println("LogFileName", config.General.LogFileName)

	var logBase = &LogBase{
		FileName:   config.Logger.FileName,
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	}

	return logBase, nil

}

func (loggerData *LogBase) PrintInfo(num string, group string, message string) {
	//ファイル取得
	//ファイルは無ければ生成(os.O_CREATE)、書き込み(os.O_WRONLY)、追記モード(os.O_APPEND)、権限は0666

	// ログの出力先とローテーションの設定を行う
	log.SetOutput(&lumberjack.Logger{
		Filename:   loggerData.FileName,   // ログファイルの名前
		MaxSize:    loggerData.MaxSize,    // ログファイルの最大サイズ（MB）
		MaxBackups: loggerData.MaxBackups, // バックアップファイルの最大数
		MaxAge:     loggerData.MaxAge,     // バックアップファイルの最大保存日数
	})

	// ログのフォーマットを設定（時間を含める）
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{
		"group": group,
		"num":   num,
	}).Info(message)
}

func (loggerData *LogBase) PrintError(num string, group string, message string) {
	//ファイル取得
	//ファイルは無ければ生成(os.O_CREATE)、書き込み(os.O_WRONLY)、追記モード(os.O_APPEND)、権限は0666

	// ログの出力先とローテーションの設定を行う
	log.SetOutput(&lumberjack.Logger{
		Filename:   loggerData.FileName,   // ログファイルの名前
		MaxSize:    loggerData.MaxSize,    // ログファイルの最大サイズ（MB）
		MaxBackups: loggerData.MaxBackups, // バックアップファイルの最大数
		MaxAge:     loggerData.MaxAge,     // バックアップファイルの最大保存日数
	})

	// ログのフォーマットを設定（時間を含める）
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{
		"group": group,
		"num":   num,
	}).Error(message)
}

func (loggerData *LogBase) PrintDebug(num string, group string, message string) {
	//ファイル取得
	//ファイルは無ければ生成(os.O_CREATE)、書き込み(os.O_WRONLY)、追記モード(os.O_APPEND)、権限は0666

	// ログの出力先とローテーションの設定を行う
	log.SetOutput(&lumberjack.Logger{
		Filename:   loggerData.FileName,   // ログファイルの名前
		MaxSize:    loggerData.MaxSize,    // ログファイルの最大サイズ（MB）
		MaxBackups: loggerData.MaxBackups, // バックアップファイルの最大数
		MaxAge:     loggerData.MaxAge,     // バックアップファイルの最大保存日数
	})

	// ログのフォーマットを設定（時間を含める）
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{
		"group": group,
		"num":   num,
	}).Debug(message)
}
