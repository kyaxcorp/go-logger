package paths

import (
	"github.com/kyaxcorp/go-config"
	"github.com/kyaxcorp/go-helper/file"
	"github.com/kyaxcorp/go-helper/filesystem"
	fsPath "github.com/kyaxcorp/go-helper/filesystem/path"
	"github.com/kyaxcorp/go-helper/folder"
)

// cum sa fac ca valoarea interfetilor date sa ajunga in alta parte?!!...
// De asemenea aceste interfete pot sa le mut in alta parte!!! dar cum insasi metodele vor fi apelate?!...

type Ddqdqw interface {
	GetLogsPath() string
	GetApplicationErrorLogsPath() string
	GetApplicationLogsPath() string
	GetLogsPathForChannels(optFolder string) string
	GetLogsPathForClients(optFolder string) string
	GetDatabasePath(optFolder string) string
	GetLogsPathForServers(optFolder string) string
}

func GetLogsPath() string {
	LogsPath := config.GetConfig().Logging.LogsPath
	//log.Println("LOGS PATH",LogsPath);
	logsPath, _err := fsPath.GenRealPath(LogsPath, true)
	if _err != nil {
		return ""
	}

	// Create the backup folder
	if !folder.Exists(logsPath) {
		folder.MkDir(logsPath)
	}
	return logsPath
}

// GetApplicationErrorPath -> gets the path where error logs will be stored from the entire app
func GetApplicationErrorLogsPath() string {
	return file.FilterPath(GetLogsPath() + "errors" + filesystem.DirSeparator())
}

func GetApplicationLogsPath() string {
	return file.FilterPath(GetLogsPath() + "application" + filesystem.DirSeparator())
}

// GetLogsPathForChannels -> channels are additional logging based on the configuration provided in the config file
func GetLogsPathForChannels(optFolder string) string {
	_path := file.FilterPath(GetLogsPath() + "channels" + filesystem.DirSeparator())
	if optFolder != "" {
		_path += file.FilterPath(optFolder) + filesystem.DirSeparator()
	}
	return _path
}

// GetLogsPathForClients -> gets the path for clients, param: optional folder
func GetLogsPathForClients(optFolder string) string {
	_path := file.FilterPath(GetLogsPath() + "clients" + filesystem.DirSeparator())
	if optFolder != "" {
		_path += file.FilterPath(optFolder) + filesystem.DirSeparator()
	}
	return _path
}

func GetDatabasePath(optFolder string) string {
	_path := file.FilterPath(GetLogsPath() + "db" + filesystem.DirSeparator())
	if optFolder != "" {
		_path += file.FilterPath(optFolder) + filesystem.DirSeparator()
	}
	return _path
}

// GetLogsPathForServers -> gets the path for servers, param: optional folder
func GetLogsPathForServers(optFolder string) string {
	_path := file.FilterPath(GetLogsPath() + "servers" + filesystem.DirSeparator())
	if optFolder != "" {
		_path += file.FilterPath(optFolder) + filesystem.DirSeparator()
	}
	return _path
}
