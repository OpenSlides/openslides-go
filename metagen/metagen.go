// Package metagen contains go definitions automatically created from the meta
// repository.
package metagen

import _ "embed" // needed for embeding

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"
//go:generate sh -c "cp ../meta/init_autoupdate_request.json ."

// InitAutoupdateRequest contains the embedded JSON payload for the initial
// autoupdate request.
//
//go:embed init_autoupdate_request.json
var InitAutoupdateRequest []byte
