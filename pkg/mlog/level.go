package mlog

type level uint8

var (
	ConsoleLevel           level = L_SILENT
	LoggerFileLevel        level = L_NORMAL
	ConsoleMessageSeparate       = "\n<<<<<<\n"
)

const (
	_        level = iota
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
