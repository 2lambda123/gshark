package vars

import (
	//"github.com/madneal/gshark/logger"
	"gopkg.in/ini.v1"
)

const (
	DefaultMaxConcurrentIndexers = 2
	PageStep                     = 5
	SearchNum                    = 25
	GITLAB                       = "gitlab"
)

var (
	REPO_PATH             string
	MAX_INDEXERS          int
	HTTP_HOST             string
	HTTP_PORT             int
	MAX_Concurrency_REPOS int
	DEBUG_MODE            bool
	PAGE_SIZE             = 10
	SCKEY                 string
	GOBUSTER              string
	SUBDOMAIN_WORDLIST    string
	ENABLE_SUBDOMAIN      bool
)

var (
	Cfg *ini.File
)