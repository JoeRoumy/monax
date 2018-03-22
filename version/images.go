// +build !arm

package version

import (
	"fmt"
)

const (
	DATA_VERSION      = "0.17"
	DB_VERSION        = "0.17.0"
	KEYS_VERSION      = "0.17.0"
	COMPILERS_VERSION = "0.18.1"
)

var (
	DefaultRegistry = "quay.io"
	BackupRegistry  = ""

	ImageData      = fmt.Sprintf("monax/data:%s", DATA_VERSION)
	ImageDB        = fmt.Sprintf("monax/db:%s", DB_VERSION)
	ImageKeys      = fmt.Sprintf("monax/keys:%s", KEYS_VERSION)
	ImageCompilers = fmt.Sprintf("monax/compilers:%s", COMPILERS_VERSION)
)
