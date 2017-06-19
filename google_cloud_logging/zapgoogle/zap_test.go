package zapgoogle_test

import (
	"testing"

	"github.com/chadit/GoSamples/google_cloud_logging/zapgoogle"
	"github.com/pkg/errors"

	"time"

	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestMap(t *testing.T) {
	// Initialize Zap.
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapgoogle.New(zapcore.DebugLevel, encoder, &zapgoogle.Writer{ProjectName: "api-project-661531736098", LogName: "my-log"})
	logger := zap.New(core, zap.Development(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Log.
	logger.Error(fmt.Sprintf("nuked %v", time.Now()), zap.String("subsystem", "example"))
	logger.Fatal(fmt.Sprintf("zzzz %v", time.Now()), zap.String("subsystem", "example"))

	//	logger.Fatal("boom boom", zap.String("subsystem", "example"))
	//logger.Fatal("boom boom", zap.Float64("keyCount", 100))
	fmt.Println("done")
	// Sync.
	fmt.Println(errors.Wrap(logger.Sync(), "failed to sync packets to Sentry"))

	time.Sleep(4 * time.Second)
}
