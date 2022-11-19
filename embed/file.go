package embed

import _ "embed"

//go:generate go run ../cmd/cazilla

//go:embed mozilla_included.pem
var MozillaIncludedCAPEM []byte
