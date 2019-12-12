package smartSummarize

import (
	"testing"

	"github.com/go-graphite/carbonapi/expr/helper"
	"github.com/go-graphite/carbonapi/expr/metadata"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
	th "github.com/go-graphite/carbonapi/tests"
)

func init() {
	md := New("")
	evaluator := th.EvaluatorFromFunc(md[0].F)
	metadata.SetEvaluator(evaluator)
	helper.SetEvaluator(evaluator)
	for _, m := range md {
		metadata.RegisterFunction(m.Name, m.F)
	}
}

func TestEvalSummarize(t *testing.T) {
	tests := []th.SummarizeEvalTestItem{
		//{
		//	"smartSummarize(metric1,'1hour','sum')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", generateValues(1800, 1800 + 3.5 * 3600, 1), 1, 1800)},
		//	},
		//	[]float64{6478200, 19438200, 32398200, 45358200},
		//	"smartSummarize(metric1,'1hour','sum')",
		//	3600,
		//	1800,
		//	5400,
		//},
		{
			"smartSummarize(metric1,'1hour','sum','1y')",
			map[parser.MetricRequest][]*types.MetricData{
				{"metric1", 0, 1}: {types.MakeMetricData("metric1", generateValues(0, 1800 + 3.5 * 3600, 1), 1, 0)},
			},
			[]float64{6478200, 19438200, 32398200, 45358200},
			"smartSummarize(metric1,'1hour','sum','1y')",
			3600,
			1800,
			5400,
		},
		//{
		//	"smartSummarize(metric1,'5s')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{
		//			1, 2, 3, 4, 5,
		//		}, 10, now32)},
		//	},
		//	[]float64{1, 2, 3, 4, 5},
		//	"smartSummarize(metric1,'5s')",
		//	10,
		//	now32,
		//	now32 + 50,
		//},
		//{
		//	"smartSummarize(metric1,'5s','avg')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 1, 2, 3, math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()}, 1, now32)},
		//	},
		//	[]float64{1, 2, 3, 4, 5, 2, math.NaN()},
		//	"smartSummarize(metric1,'5s','avg')",
		//	5,
		//	now32,
		//	now32 + 35,
		//},
		//{
		//	"smartSummarize(metric1,'5s','max')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{1, 2, 3, 4.5, 5},
		//	"smartSummarize(metric1,'5s','max')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','min')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{0, 1, 1.5, 2, 5},
		//	"smartSummarize(metric1,'5s','min')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','last')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{1, 2, 3, 4.5, 5},
		//	"smartSummarize(metric1,'5s','last')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','p50')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{0.5, 1.5, 2, 3, 5},
		//	"smartSummarize(metric1,'5s','p50')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','p25')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{0, 1, 2, 3, 5},
		//	"smartSummarize(metric1,'5s','p25')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','p99.9')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{1, 2, 3, 4.498, 5},
		//	"smartSummarize(metric1,'5s','p99.9')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'5s','p100.1')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
		//	"smartSummarize(metric1,'5s','p100.1')",
		//	5,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'1s','p50')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5}, 1, now32)},
		//	},
		//	[]float64{1, 0, 0, 0.5, 1, 2, 1, 1, 1.5, 2, 3, 2, 2, 1.5, 3, 4, 3, 2, 3, 4.5, 5, 5, 5, 5, 5},
		//	"smartSummarize(metric1,'1s','p50')",
		//	1,
		//	now32,
		//	now32 + 25*1,
		//},
		//{
		//	"smartSummarize(metric1,'10min')",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{
		//			1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
		//			3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
		//			5, 5, 5, 5, 5}, 60, tenThirtyTwo)},
		//	},
		//	[]float64{11, 31, 33},
		//	"smartSummarize(metric1,'10min')",
		//	600,
		//	tenThirty,
		//	tenThirty + 30*60,
		//},
		//{
		//	"smartSummarize(metric1,'10min','sum',true)",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{
		//			1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
		//			3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
		//			5, 5, 5, 5, 5}, 60, tenThirtyTwo)},
		//	},
		//	[]float64{15, 35, 25},
		//	"smartSummarize(metric1,'10min','sum',true)",
		//	600,
		//	tenThirtyTwo,
		//	tenThirtyTwo + 25*60,
		//},
		//{
		//	"smartSummarize(metric1,'10min','sum',true)",
		//	map[parser.MetricRequest][]*types.MetricData{
		//		{"metric1", 0, 1}: {types.MakeMetricData("metric1", []float64{
		//			1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
		//			3, 3, 3, 3, 3, 4, 4, 4, 4, 4,
		//			5, 5, 5, 5, 5}, 60, tenThirtyTwo)},
		//	},
		//	[]float64{15, 35, 25},
		//	"smartSummarize(metric1,'10min','sum',true)",
		//	600,
		//	tenThirtyTwo,
		//	tenThirtyTwo + 25*60,
		//},
	}

	for _, tt := range tests {
		th.TestSummarizeEvalExpr(t, &tt)
	}
}

func generateValues(start, stop, step int64) (values []float64) {
	for i := start; i < stop; i += step {
		values = append(values, float64(i))
	}
	return
}
