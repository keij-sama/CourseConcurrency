package logger

type Logger interface {
	Ingo(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}
