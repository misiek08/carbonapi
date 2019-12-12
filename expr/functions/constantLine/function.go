package constantLine

import (
	"fmt"
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/types"
	pb "github.com/go-graphite/protocol/carbonapi_v3_pb"
)

type constantLine struct {
	interfaces.FunctionBase
}

func GetOrder() interfaces.Order {
	return interfaces.Any
}

func New(configFile string) []interfaces.FunctionMetadata {
	res := make([]interfaces.FunctionMetadata, 0)
	f := &constantLine{}
	functions := []string{"constantLine"}
	for _, n := range functions {
		res = append(res, interfaces.FunctionMetadata{Name: n, F: f})
	}
	return res
}

func (f *constantLine) Do(ctx interfaces.FunctionCallContext) ([]*types.MetricData, error) {
	value, err := ctx.E.GetFloatArg(0)

	if err != nil {
		return nil, err
	}
	p := types.MetricData{
		FetchResponse: pb.FetchResponse{
			Name:              fmt.Sprintf("%g", value),
			StartTime:         ctx.From,
			StopTime:          ctx.Until,
			StepTime:          ctx.Until - ctx.From,
			Values:            []float64{value, value},
			ConsolidationFunc: "max",
		},
	}

	return []*types.MetricData{&p}, nil
}

// Description is auto-generated description, based on output of https://github.com/graphite-project/graphite-web
func (f *constantLine) Description() map[string]types.FunctionDescription {
	return map[string]types.FunctionDescription{
		"constantLine": {
			Description: "Takes a float F.\n\nDraws a horizontal line at value F across the graph.\n\nExample:\n\n.. code-block:: none\n\n  &target=constantLine(123.456)",
			Function:    "constantLine(value)",
			Group:       "Special",
			Module:      "graphite.render.functions",
			Name:        "constantLine",
			Params: []types.FunctionParam{
				{
					Name:     "value",
					Required: true,
					Type:     types.Float,
				},
			},
		},
	}
}
