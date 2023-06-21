package customlogger

import (
    "log"
    "os"
    "sync"
)

type loggers struct {
    filename string
    InfoLogger    *log.Logger
    WarningLogger *log.Logger
    ErrorLogger   *log.Logger
}

var logger *loggers
var once sync.Once

func GetInstance() *loggers {
    once.Do(func() {
        logger = createLogger("/var/logs/<%= baseName %>.log")
    })
    return logger
}

func createLogger(fname string) *loggers {
    file,err :=os.OpenFile(fname,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
    if(err!=nil){
		log.Fatal(err)
	}
    return &loggers{
        filename: fname,
        InfoLogger: log.New(file,"INFO:",log.LstdFlags|log.Lshortfile),
        WarningLogger: log.New(file,"WARNING:",log.LstdFlags|log.Lshortfile),
        ErrorLogger: log.New(file,"ERROR:",log.LstdFlags|log.Lshortfile)}
}

// uncomment if you want logs with in file
func Printfun(level string,msg string) {
    // logger := GetInstance()
    // if(level=="info"){
    // logger.InfoLogger.Println(msg)
    // }

    // if(level=="warn"){
    //  logger.WarningLogger.Println(msg)
    // }

    // if(level=="error"){
    // logger.ErrorLogger.Println(msg)
    // }
    log.Println(msg)
}