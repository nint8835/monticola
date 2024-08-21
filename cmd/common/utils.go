package common

import (
	"log/slog"
)

func CheckError(err error, msg string) {
	if err != nil {
		slog.Error(
			msg,
			slog.String("error", err.Error()),
		)
	}
}
