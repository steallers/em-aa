package constant

import "os"

const (
	EmployeeNotFound = "employee with id (%d) not found"
	UnDescribeErrors = "something with server please call your admin"
)

func GetAppName() string {
	return os.Getenv("APP_NAME")
}

func AppVersion() string {
	return os.Getenv("APP_VERSION")
}

func PluginFunction() string {
	return os.Getenv("PLUGIN_VERSION")
}
