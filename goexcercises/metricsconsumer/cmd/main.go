package main

import (
	"metricsconsumer"
	"time"
)

func main() {
	snapshot := metricsconsumer.NamespaceSnapshot{
		Namespace: "lbs",
		CPUCores:  12,
		MemoryMB:  330,
		PodCount:  10,
		Timestamp: time.Time{},
	}
	writer := metricsconsumer.TerminalWriter{}
	renderer := metricsconsumer.NewTerminalWriterRenderer[metricsconsumer.NamespaceSnapshot](writer, metricsconsumer.NamespaceSnapshotConverter)
	_ = renderer.RenderAll([]metricsconsumer.NamespaceSnapshot{snapshot})
}
