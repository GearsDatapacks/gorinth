package gorinth

import (
	"fmt"
)

func format(log string, logType string, values ...any) string {
	return "[Gorinth] " + logType + " - " + fmt.Sprintf(log, values...)
}

type gorinthError string

func (err gorinthError) Error() string {
	return format(string(err), "Error")
}

func makeError(message string, values ...any) gorinthError {
	return gorinthError(fmt.Sprintf(message, values...))
}

func logWarning(warn string, values ...any) {
	fmt.Println(format(warn, "Warning", values...))
}

// func logInfo(info string, values ...any) {
// 	fmt.Println(format(info, "Info", values))
// 	os.Exit(1)
// }
