// +build !cairo

package png

import (
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/types"
	"net/http"
)

const HaveGraphSupport = false

func EvalExprGraph(ctx interfaces.FunctionCallContext) ([]*types.MetricData, error) {
	return nil, nil
}

func MarshalPNG(params PictureParams, results []*types.MetricData) []byte {
	return nil
}

func MarshalSVG(params PictureParams, results []*types.MetricData) []byte {
	return nil
}

func MarshalPNGRequest(r *http.Request, results []*types.MetricData, templateName string) []byte {
	return nil
}

func MarshalSVGRequest(r *http.Request, results []*types.MetricData, templateName string) []byte {
	return nil
}

func Description() map[string]types.FunctionDescription {
	return nil
}
