package port

type LoggerPort interface {
	PrintInfo(num string, group string, message string)
	PrintError(num string, group string, message string)
	PrintDebug(group string, num string, message string)
}
