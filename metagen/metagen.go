// Package metagen contains go definitions automatically created from the meta
// repository.
package metagen

import _ "embed"

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"
//go:generate sh -c "cp ../meta/init_autoupdate_request.json ."

//go:embed init_autoupdate_request.json
var InitAutoupdateRequest []byte
