package main

import (
	"GoShare"
	"fmt"

	"github.com/kardianos/osext"
)

func main() {
	fmt.Println("main called")
	err := Shares.InitConfig()
	fmt.Println("main called")
	if err != nil {
		Shares.CreateFolder(Shares.LogFilePath)
		fmt.Println("config error ", err, Shares.LogFilePath)

	}
	// logconf := &golog.LoggerConfig{
	// 	Level:          golog.ToLogLevel("DEBUG"),
	// 	Verbosity:      golog.LDefault | golog.LHeaderFooter | golog.LFile,
	// 	FileRotateSize: (2 << 23), /*16MB*/
	// 	FileDepth:      5,
	// }
	// golog.NewLogger("log", "C:\\logs\\test.log", logconf)
	// logprovider := golog.GetLogger("log1")
	// logprovider := getLogProvider("log")
	// if logprovider == nil {
	// 	logprovider = getLogProvider("log")
	// }
	logprovider := Shares.GetLogProvider("log")
	fmt.Println("return")
	fmt.Println(logprovider)
	folderPath, _ := osext.ExecutableFolder()
	logprovider.Debug("Service Started : root folder " + folderPath)
}
