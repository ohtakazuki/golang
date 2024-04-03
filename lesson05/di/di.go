package main

import "fmt"

// Loggerインターフェースを定義
type Logger interface {
	Log(message string)
}

// ConsoleLoggerはLoggerインターフェースを実装
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("ConsoleLogger:", message)
}

// FileLoggerはLoggerインターフェースを実装
type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Println("FileLogger:", message)
}

// Serviceはロガーに依存
type Service struct {
	logger Logger
}

func (s Service) DoSomething() {
	s.logger.Log("Service is doing something")
}

func main() {
	// ConsoleLoggerのインスタンスを作成
	consoleLogger := ConsoleLogger{}

	// FileLoggerのインスタンスを作成
	fileLogger := FileLogger{}

	// LoggerインターフェースにConsoleLoggerを注入してServiceを作成
	service1 := Service{logger: consoleLogger}
	service1.DoSomething()

	// LoggerインターフェースにFileLoggerを注入してServiceを作成
	service2 := Service{logger: fileLogger}
	service2.DoSomething()
}
