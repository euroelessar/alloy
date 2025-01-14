package otelcolconvert

import (
	"fmt"

	"github.com/grafana/alloy/internal/component/otelcol"
	"github.com/grafana/alloy/internal/component/otelcol/processor/deltatocumulative"
	"github.com/grafana/alloy/internal/converter/diag"
	"github.com/grafana/alloy/internal/converter/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/pipeline"
)

func init() {
	converters = append(converters, deltatocumulativeProcessorConverter{})
}

type deltatocumulativeProcessorConverter struct{}

func (deltatocumulativeProcessorConverter) Factory() component.Factory {
	return deltatocumulativeprocessor.NewFactory()
}

func (deltatocumulativeProcessorConverter) InputComponentName() string {
	return "otelcol.processor.deltatocumulative"
}

func (deltatocumulativeProcessorConverter) ConvertAndAppend(state *State, id componentstatus.InstanceID, cfg component.Config) diag.Diagnostics {
	var diags diag.Diagnostics

	label := state.AlloyComponentLabel()

	args := toDeltatocumulativeProcessor(state, id, cfg.(*deltatocumulativeprocessor.Config))
	block := common.NewBlockWithOverride([]string{"otelcol", "processor", "deltatocumulative"}, label, args)

	diags.Add(
		diag.SeverityLevelInfo,
		fmt.Sprintf("Converted %s into %s", StringifyInstanceID(id), StringifyBlock(block)),
	)

	state.Body().AppendBlock(block)
	return diags
}

func toDeltatocumulativeProcessor(state *State, id componentstatus.InstanceID, cfg *deltatocumulativeprocessor.Config) *deltatocumulative.Arguments {
	var (
		nextMetrics = state.Next(id, pipeline.SignalMetrics)
		nextLogs    = state.Next(id, pipeline.SignalLogs)
		nextTraces  = state.Next(id, pipeline.SignalTraces)
	)

	return &deltatocumulative.Arguments{
		MaxStale:   cfg.MaxStale,
		MaxStreams: cfg.MaxStreams,
		Output: &otelcol.ConsumerArguments{
			Metrics: ToTokenizedConsumers(nextMetrics),
			Logs:    ToTokenizedConsumers(nextLogs),
			Traces:  ToTokenizedConsumers(nextTraces),
		},
		DebugMetrics: common.DefaultValue[deltatocumulative.Arguments]().DebugMetrics,
	}
}
