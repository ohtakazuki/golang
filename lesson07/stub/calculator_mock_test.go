package stub

import (
	"testing"

	"github.com/golang/mock/gomock"
)

// MockLoggerインターフェースを生成
//go:generate mockgen -source=logger.go -destination=mock_logger_test.go -package=stub

func TestMockCalculator(t *testing.T) {
	var calc *Calculator
	var mockLogger *MockLogger

	// テストデータの準備
	setup := func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockLogger = NewMockLogger(ctrl)
		calc = NewCalculator(mockLogger)
	}

	// 加算のテスト
	t.Run("Addition", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(2): result = 2")
		mockLogger.EXPECT().Log("Add(3): result = 5")
		calc.Add(2)
		calc.Add(3)
		assertEqual(t, calc.Result(), 5)
	})

	// 減算のテスト
	t.Run("Subtraction", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(10): result = 10")
		mockLogger.EXPECT().Log("Subtract(3): result = 7")
		calc.Add(10)
		calc.Subtract(3)
		assertEqual(t, calc.Result(), 7)
	})

	// 乗算のテスト
	t.Run("Multiplication", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(2): result = 2")
		mockLogger.EXPECT().Log("Multiply(3): result = 6")
		calc.Add(2)
		calc.Multiply(3)
		assertEqual(t, calc.Result(), 6)
	})

	// 除算のテスト
	t.Run("Division", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(10): result = 10")
		mockLogger.EXPECT().Log("Divide(2): result = 5")
		calc.Add(10)
		err := calc.Divide(2)
		assertNoError(t, err)
		assertEqual(t, calc.Result(), 5)
	})

	// ゼロ除算のテスト
	t.Run("DivisionByZero", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(10): result = 10")
		mockLogger.EXPECT().Log("Divide(0): error = division by zero")
		calc.Add(10)
		err := calc.Divide(0)
		assertError(t, err)
	})

	// 複雑な計算のテスト
	t.Run("ComplexCalculation", func(t *testing.T) {
		setup(t)
		mockLogger.EXPECT().Log("Add(2): result = 2")
		mockLogger.EXPECT().Log("Multiply(3): result = 6")
		mockLogger.EXPECT().Log("Subtract(1): result = 5")
		mockLogger.EXPECT().Log("Divide(2): result = 2")
		calc.Add(2)
		calc.Multiply(3)
		calc.Subtract(1)
		err := calc.Divide(2)
		assertNoError(t, err)
		assertEqual(t, calc.Result(), 2)
	})
}
