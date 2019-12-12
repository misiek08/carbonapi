package smartSummarize

import (
	"fmt"
	"github.com/go-graphite/carbonapi/expr/consolidations"
	"github.com/go-graphite/carbonapi/expr/helper"
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
	pb "github.com/go-graphite/protocol/carbonapi_v3_pb"
	"math"
)

type smartSummarize struct {
	interfaces.FunctionBase
}

func GetOrder() interfaces.Order {
	return interfaces.Any
}

func New(configFile string) []interfaces.FunctionMetadata {
	res := make([]interfaces.FunctionMetadata, 0)
	f := &smartSummarize{}
	functions := []string{"smartSummarize"}
	for _, n := range functions {
		res = append(res, interfaces.FunctionMetadata{Name: n, F: f})
	}
	return res
}

// smartSummarize(seriesList, intervalString, alignToInterval=False)
func (f *smartSummarize) Do(e parser.Expr, from, until int64, values map[parser.MetricRequest][]*types.MetricData) ([]*types.MetricData, error) {
	// TODO(dgryski): make sure the arrays are all the same 'size'
	args, err := helper.GetSeriesArg(e.Args()[0], from, until, values)
	if err != nil {
		return nil, err
	}

	bucketSizeInt32, err := e.GetIntervalArg(1, 1)
	if err != nil {
		return nil, err
	}
	bucketSize := int64(bucketSizeInt32)

	summarizeFunction, err := e.GetStringNamedOrPosArgDefault("func", 2, "sum")
	if err != nil {
		return nil, err
	}

	alignToInterval, err := e.GetStringNamedOrPosArgDefault("alignTo", 3, "")
	if err != nil {
		return nil, err
	}

	start := args[0].StartTime
	stop := args[0].StopTime
	if alignToInterval != "" {
		interval, err := parser.IntervalString(alignToInterval, 1)
		if err != nil {
			return nil, err
		}
		start = helper.AlignStartToInterval(start, stop, int64(interval))
	}

	buckets := helper.GetBuckets(start, stop, bucketSize)
	results := make([]*types.MetricData, 0, len(args))
	for _, arg := range args {

		name := fmt.Sprintf("smartSummarize(%s,'%s'", arg.Name, e.Args()[1].StringValue())
		name += ")"

		r := types.MetricData{FetchResponse: pb.FetchResponse{
			Name:              name,
			Values:            make([]float64, buckets, buckets+1),
			StepTime:          bucketSize,
			StartTime:         start,
			StopTime:          stop,
			ConsolidationFunc: "sum",
		}}

		t := arg.StartTime // unadjusted
		bucketEnd := start + bucketSize
		values := make([]float64, 0, bucketSize/arg.StepTime)
		ridx := 0
		bucketItems := 0
		for _, v := range arg.Values {
			bucketItems++
			if !math.IsNaN(v) {
				values = append(values, v)
			}

			t += arg.StepTime

			if t >= stop {
				break
			}

			if t >= bucketEnd {
				rv := consolidations.SummarizeValues(summarizeFunction, values)

				r.Values[ridx] = rv
				ridx++
				bucketEnd += bucketSize
				bucketItems = 0
				values = values[:0]
			}
		}

		// last partial bucket
		if bucketItems > 0 {
			rv := consolidations.SummarizeValues(summarizeFunction, values)
			r.Values[ridx] = rv
		}

		results = append(results, &r)
	}
	return results, nil
}

// Description is auto-generated description, based on output of https://github.com/graphite-project/graphite-web
func (f *smartSummarize) Description() map[string]types.FunctionDescription {
	return map[string]types.FunctionDescription{
		"smartSummarize": {
			Description: "Estimate hit counts from a list of time series.\n\nThis function assumes the values in each time series represent\nhits per second.  It calculates hits per some larger interval\nsuch as per day or per hour.  This function is like summarize(),\nexcept that it compensates automatically for different time scales\n(so that a similar graph results from using either fine-grained\nor coarse-grained records) and handles rarely-occurring events\ngracefully.",
			Function:    "smartSummarize(seriesList, intervalString, alignToInterval=False)",
			Group:       "Transform",
			Module:      "graphite.render.functions",
			Name:        "smartSummarize",
			Params: []types.FunctionParam{
				{
					Name:     "seriesList",
					Required: true,
					Type:     types.SeriesList,
				},
				{
					Name:     "intervalString",
					Required: true,
					Suggestions: types.NewSuggestions(
						"10min",
						"1h",
						"1d",
					),
					Type: types.Interval,
				},
				{
					Default: types.NewSuggestion(false),
					Name:    "alignToInterval",
					Type:    types.Boolean,
				},
			},
		},
	}
}
