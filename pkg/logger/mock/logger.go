package mock_logger

import "github.com/dragonfish/go/v2/pkg/logger"

type MockLogger struct{}

func (l *MockLogger) Debug(args ...interface{}) {}

func (l *MockLogger) Debugf(format string, args ...interface{}) {}

func (l *MockLogger) Info(args ...interface{}) {}

func (l *MockLogger) Infof(format string, args ...interface{}) {}

func (l *MockLogger) Warn(args ...interface{}) {}

func (l *MockLogger) Warnf(format string, args ...interface{}) {}

func (l *MockLogger) Error(args ...interface{}) {}

func (l *MockLogger) Errorf(format string, args ...interface{}) {}

func (l *MockLogger) Fatal(args ...interface{}) {}

func (l *MockLogger) Fatalf(format string, args ...interface{}) {}

func (l *MockLogger) Panic(args ...interface{}) {}

func (l *MockLogger) Panicf(format string, args ...interface{}) {}

func (l *MockLogger) WithField(key string, value interface{}) logger.Logger {
	return l
}

func (l *MockLogger) WithFields(fields map[string]string) logger.Logger {
	return l
}

func (l *MockLogger) WithError(err error) logger.Logger {
	return l
}
