package zapgoogle

// New create new ZapLogger
// func New(logName string) *logging.Logger {
// 	// log, err := zap.NewDevelopment(zap.Hooks())
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	client, err := logging.NewClient(context.Background(), "api-project-661531736098")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Initialize a logger
// 	return client.Logger(logName)

// 	// return log.WithOptions(zap.Hooks((&ZapGoogle{
// 	// 	Logger: client.Logger(logName),
// 	// }).Log))
// }

//Log logs a zap entry
// func (l *ZapGoogle) Log(e zapcore.Entry) error {
// 	fmt.Println("zap-google log")
// 	googleEntry := logging.Entry{
// 		//	LogName:   e.LoggerName,
// 		Severity: severity[e.Level],
// 		//	Timestamp: e.Time,
// 		Payload: parsePayload(e.Message),
// 	}
// 	fmt.Println("e.Stack : ", e.Stack)

// 	// if googleEntry.Severity == logging.Critical {
// 	// 	return l.Logger.LogSync(context.Background(), googleEntry)
// 	// }

// 	l.Logger.Log(googleEntry)
// 	//l.Logger.Flush()
// 	return nil
// }

// func parsePayload(msg string) map[string]interface{} {
// 	fmt.Println("parsePayload : ", msg)
// 	payload := make(map[string]interface{})
// 	p := make(map[string]interface{})
// 	if err := json.NewDecoder(strings.NewReader(msg)).Decode(&p); err != nil {
// 		payload["Payload"] = msg
// 	}

// 	for k, v := range p {
// 		payload[k] = v
// 	}
// 	return payload
// }
