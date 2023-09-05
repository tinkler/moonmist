package mlog

type Level uint8

var (
	ConsoleLevel           Level = L_SILENT
	LoggerFileLevel        Level = L_NORMAL
	ConsoleMessageSeparate       = "\n<<<<<<\n"
)

const (
	_        Level = iota
	L_SILENT       // print ML_ERR ML_TITLE message
	L_NORMAL       // print ML_ERR ML_TITLE ML_WARN message
	L_LOG          // print ML_ERR ML_TITLE ML_WARN ML_INFO message
	L_DEBUG        // print ML_ERR ML_TITLE ML_WARN ML_INFO ML_DEBUG message
)

type messageLevel uint8

const (
	_ messageLevel = iota
	ML_ERR
	ML_WARN
	ML_INFO
	ML_DEBUG
	ML_TITLE
)
