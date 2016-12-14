package explorepackages

func LoggerConfigToShowLogs() {
	mLog := CreateLogger("main")
	mLog.Infof("Start")
	mLog.Debug("Its working")
	mLog.Errorf("Critical error")
}
