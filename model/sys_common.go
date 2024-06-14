package model

type CommonModel struct {
	RuntimePath string
}

type APPModel struct {
	DBPath      string
	LogPath     string
	LogSaveName string
	LogFileExt  string
}
