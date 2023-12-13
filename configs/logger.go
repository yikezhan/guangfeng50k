package configs

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

//只能输出结构化日志，但是性能要高于 SugaredLogger
var Log *zap.Logger

//可以输出 结构化日志、非结构化日志。性能茶语 zap.Logger，具体可见上面的的单元测试
var SLog *zap.SugaredLogger

// 初始化日志 logger
func InitLog(logLevel zapcore.Level) {
	config := getEncodeConfig()
	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})
	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})
	// 获取io.Writer的实现
	infoWriter := getInfoWriter()
	warnWriter := getErrorWriter()
	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel),                         //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel),                         //warn及以上写入errPath
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel), //同时将日志输出到控制台，NewJSONEncoder 是结构化输出
	)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SLog = Log.Sugar()
}

func getEncodeConfig() zapcore.EncoderConfig {
	// 使用zap提供的 NewProductionEncoderConfig
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 时间的key
	encoderConfig.TimeKey = "time"
	// 级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 显示调用者信息
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回json 格式的 日志编辑器
	return encoderConfig
}

func getInfoWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   viper.GetString("log.info.filename"),
		MaxSize:    viper.GetInt("log.info.max_size"),    //最大M数，超过则切割
		MaxBackups: viper.GetInt("log.info.max_backups"), //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     viper.GetInt("log.info.max_age"),     //保存30天
		Compress:   false,                                //是否压缩
	}
}

func getErrorWriter() io.Writer {
	return &lumberjack.Logger{
		Filename:   viper.GetString("log.error.filename"),
		MaxSize:    viper.GetInt("log.error.max_size"),    //最大M数，超过则切割
		MaxBackups: viper.GetInt("log.error.max_backups"), //最大文件保留数，超过就删除最老的日志文件
		MaxAge:     viper.GetInt("log.error.max_age"),     //保存30天
		Compress:   false,                                 //是否压缩
	}
}
