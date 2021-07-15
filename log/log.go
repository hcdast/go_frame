package log

import (
	"github.com/op/go-logging"
	"os"
)

var (
	Logger = logging.MustGetLogger("tsl_frame")
	format = logging.MustStringFormatter(
		"%{color} %{level:.4s} %{id:03x} %{time:2006-01-02 15:04:05.000} %{shortfile}\t%{shortfunc}\t> %{message}%{color:reset}",
	)
)

func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backend2Formatter)
}
