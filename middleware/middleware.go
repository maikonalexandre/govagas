package middleware

import "github.com/maikonalexandre/govagas/config"

var logger *config.Logger

func InitializeHandler() {
	logger = config.GetLogger("middleware: ")
}
