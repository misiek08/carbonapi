package cairo

import (
	"github.com/go-graphite/carbonapi/expr/functions/cairo/png"
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/types"
)

type cairo struct {
	interfaces.FunctionBase
}

func GetOrder() interfaces.Order {
	return interfaces.Any
}

func New(configFile string) []interfaces.FunctionMetadata {
	res := make([]interfaces.FunctionMetadata, 0)
	f := &cairo{}
	functions := []string{"color", "stacked", "areaBetween", "alpha", "dashed", "drawAsInfinite", "secondYAxis", "lineWidth", "threshold"}
	for _, n := range functions {
		res = append(res, interfaces.FunctionMetadata{Name: n, F: f})
	}
	return res
}

func (f *cairo) Do(ctx interfaces.FunctionCallContext) ([]*types.MetricData, error) {
	return png.EvalExprGraph(ctx)
}

func (f *cairo) Description() map[string]types.FunctionDescription {
	return png.Description()
}
