package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
)

// runs before the main application
func init() {
	// load Configuration file
	//	Shares.InitConfig()
	// ensure that the log location exist and create if it does not
	//	Shares.CreateFolder(Shares.LogFilePath)
	log.Print("blah blah lah")
	// initializae logging
	// logconf := &golog.LoggerConfig{
	// 	Level:          golog.ToLogLevel("DEBUG"),
	// 	Verbosity:      golog.LDefault | golog.LHeaderFooter | golog.LFile,
	// 	FileRotateSize: (2 << 23), /*16MB*/
	// 	FileDepth:      5,
	// }
	// golog.NewLogger("log", Shares.LogFilePath+"\\"+Shares.ServiceName, logconf)
}

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	const svcName = "goexampleservice"

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if !isIntSess {
		runService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(svcName, true)
		return
	case "install":
		err = installService(svcName, "my service")
	case "remove":
		err = removeService(svcName)
	case "start":
		err = startService(svcName)
	case "stop":
		err = controlService(svcName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return
}
