package log

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.Logger
)

func init() {
	defaultLogger = New(WithLevel(zapcore.InfoLevel))
}

func Default() *zap.Logger {
	return defaultLogger
}

func GetLogger(c *gin.Context) *zap.SugaredLogger {
	return FromContext(c.Request.Context()).Sugar()
}

func Init(opts ...Option) *zap.Logger {
	defaultLogger = New(opts...)
	return defaultLogger
}

func New(opts ...Option) *zap.Logger {
	opt := option{}
	opts = append(defaultOptions, opts...)
	for _, o := range opts {
		o(&opt)
	}

	jsonEncoder := encoder()

	core := zapcore.NewTee()

	if !opt.disableConsole {
		core = zapcore.NewTee(
			// to stdout
			zapcore.NewCore(
				jsonEncoder,
				zapcore.NewMultiWriteSyncer(zapcore.Lock(os.Stdout)),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= opt.level && lvl < zapcore.ErrorLevel
				}),
			),
			// to stderr
			zapcore.NewCore(
				jsonEncoder,
				zapcore.NewMultiWriteSyncer(zapcore.Lock(os.Stderr)),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= opt.level && lvl >= zapcore.ErrorLevel
				}),
			),
		)
	}

	if opt.filename != "" {
		writer := buildLumberjackWriter(opt.filename)
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(
				jsonEncoder,
				zapcore.AddSync(writer),
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= opt.level
				}),
			),
		)
	}

	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(opt.callerSkip),
	}
	if opt.errorStacktrace {
		zapOptions = append(zapOptions, zap.AddStacktrace(zapcore.ErrorLevel))
	}
	logger := zap.New(core, zapOptions...)
	return logger
}

type option struct {
	level           zapcore.Level
	callerSkip      int
	errorStacktrace bool
	disableConsole  bool
	filename        string
}

// ------------------- Options -------------------
var defaultOptions = []Option{
	WithErrorStacktrace(false),
	WithFilename("./logs/runtime.log"),
}

type Option func(*option)

// AddCallerSkip adds the given number of callers skipped to the options.
func AddCallerSkip(skip int) Option {
	return func(opt *option) {
		opt.callerSkip += skip
	}
}

// WithErrorStacktrace sets whether to record stacktraces on errors.
func WithErrorStacktrace(on bool) Option {
	return func(opt *option) {
		opt.errorStacktrace = on
	}
}

// WithLevel sets the log level.
func WithLevel(level zapcore.Level) Option {
	return func(opt *option) {
		opt.level = level
	}
}

// DisableConsole disables logging to console.
func DisableConsole() Option {
	return func(opt *option) {
		opt.disableConsole = true
	}
}

// WithFilename sets the log filename.
func WithFilename(filename string) Option {
	return func(opt *option) {
		opt.filename = filename
	}
}

func encoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	recordTimeFormat := time.RFC3339Nano
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func buildLumberjackWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename: filename,
	}
}
