package metricsconsumer

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TerminalWriter struct{}

func (t *TerminalWriter) WriteTable(headers []string, rows [][]string) error {
	headerString := strings.Join(headers, ",")
	fmt.Println(headerString)
	for _, row := range rows {
		rowString := strings.Join(row, ",")
		fmt.Println(rowString)
	}
	return nil
}

type PrometheusEmitter struct{}

func (p *PrometheusEmitter) EmitGauge(metricName string, labels map[string]string, value float64) {
	// In a real system this would register and expose a Prometheus gauge.
	// For this exercise we just print what would have been emitted.
	fmt.Printf("GAUGE %s{", metricName)
	for k, v := range labels {
		fmt.Printf(`%s="%s" `, k, v)
	}
	fmt.Printf("} %f\n", value)
}

type JSONExporter struct{}

func (j *JSONExporter) ExportJSON(v any) error {
	// In production this would write to a file or an HTTP sink.
	// Here we marshal and print.
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
