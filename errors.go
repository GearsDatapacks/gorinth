package gorinth

import (
	"fmt"
)

func format(log string, logType string, values ...any) string {
	return "[Gorinth] " + logType + " - " + fmt.Sprintf(log, values...)
}

type Error string

func (err Error) Error() string {
	return format(string(err), "Error")
}

func makeError(message string, values ...any) Error {
	return Error(fmt.Sprintf(message, values...))
}

// func logError(err string, values ...any) {
// 	fmt.Println(format(err, "Error", values))
// 	os.Exit(1)
// }

func logWarning(warn string, values ...any) {
	fmt.Println(format(warn, "Warning", values...))
}

// func logInfo(info string, values ...any) {
// 	fmt.Println(format(info, "Info", values))
// 	os.Exit(1)
// }
