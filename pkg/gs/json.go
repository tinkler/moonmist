package gs

import (
	"io"
	"net/http"

	"github.com/tinkler/moonmist/pkg/jsonz"
	"github.com/tinkler/moonmist/pkg/jsonz/cjson"
)

var json jsonz.Json = cjson.Json

func Bind(r *http.Request, v any) error {
	reqByt, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(reqByt, v)
}
