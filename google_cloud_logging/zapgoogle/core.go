package zapgoogle

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"

	"cloud.google.com/go/logging"
	"go.uber.org/zap/zapcore"
)

// Core implements the interface utilized by zap
type Core struct {
	zapcore.LevelEnabler
	encoder zapcore.Encoder
	writer  *Writer
	fields  []zapcore.Field
}

// Writer implements customer options
type Writer struct {
	ProjectName string // project of the google cloud log
	LogName     string // name of the google cloud log
	Logger      *logging.Logger
}

// New initializes the logger
func New(enab zapcore.LevelEnabler, encoder zapcore.Encoder, writer *Writer) *Core {
	if writer == nil {
		panic("writer cannot be nil")
	}

	if writer.LogName == "" {
		panic("log cannot be missing")
	}

	if writer.ProjectName == "" {
		panic("log cannot be missing")
	}

	return &Core{
		LevelEnabler: enab,
		encoder:      encoder,
		writer: &Writer{
			ProjectName: writer.ProjectName,
			LogName:     writer.LogName,
			Logger:      initLogger(writer.ProjectName, writer.LogName),
		},
	}
}

func initLogger(projectName, logName string) *logging.Logger {
	client, err := logging.NewClient(context.Background(), projectName)
	if err != nil {
		panic(err)
	}

	return client.Logger(logName)
}

// With handles operations
func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	// Clone core.
	clone := *c

	// Clone encoder.
	clone.encoder = c.encoder.Clone()

	// Clone and append fields.
	clone.fields = make([]zapcore.Field, len(c.fields)+len(fields))
	copy(clone.fields, c.fields)
	copy(clone.fields[len(c.fields):], fields)

	// Done.
	return &clone
}

// Check handles operations
func (c *Core) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checked.AddCore(entry, c)
	}
	return checked
}

// Write handles the writer operations to google cloud
func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// Generate the message.
	buffer, err := c.encoder.EncodeEntry(entry, fields)
	if err != nil {
		return errors.Wrap(err, "failed to encode log entry")
	}

	p := make(map[string]interface{})
	if err := json.NewDecoder(strings.NewReader(buffer.String())).Decode(&p); err != nil {
		return err
	}

	g := logging.Entry{
		Severity:  severity[entry.Level],
		Timestamp: entry.Time,
		//Payload:   buffer.String(),
		Payload: p,
	}

	if g.Severity == logging.Critical || g.Severity == logging.Error {
		go func() {
			if err := c.writer.Logger.LogSync(context.Background(), g); err != nil {
				return
			}
		}()
		return nil
	}

	c.writer.Logger.Log(g)
	return nil
}

// Sync not supported
func (c *Core) Sync() error {
	return nil
}
