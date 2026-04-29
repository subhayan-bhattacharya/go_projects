package main

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"os"
	datamodel "plugin-registry/datamodel"
	errors "plugin-registry/errors"
	plugins "plugin-registry/plugin"
	"time"
)

func readFile() ([]datamodel.SensorData, error) {
	var sensorData []datamodel.SensorData
	data, err := os.ReadFile("data.json")
	if err != nil {
		return sensorData, &errors.DataReadError{
			err.Error(),
		}
	}
	err = json.Unmarshal(data, &sensorData)
	if err != nil {
		return sensorData, &errors.DataReadError{
			err.Error(),
		}
	}
	return sensorData, nil
}

func main() {
	sensorData, _ := readFile()
	context := context2.Background()
	plugin1 := plugins.Sanitizer[*datamodel.SensorData]{}
	plugin2 := plugins.NewMemoryStore[*datamodel.SensorData]()
	timeoutOption := plugins.WithTimeOut[*datamodel.SensorData](30 * time.Second)
	pluginPipeline := plugins.NewPluginPipeline[*datamodel.SensorData](timeoutOption)
	pluginPipeline.AddPlugin(plugin1)
	pluginPipeline.AddPlugin(plugin2)
	for i := range sensorData {
		err := pluginPipeline.Execute(context, &sensorData[i])
		if err != nil {
			fmt.Printf("the pipeline execution resulted in an error %v\n", err)
		}
	}
	fmt.Println("what has been saved in the map")
	fmt.Println(plugin2.Db)
}
