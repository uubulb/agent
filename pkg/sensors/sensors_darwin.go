//go:build darwin

package sensors

import (
	"sort"

	"github.com/nezhahq/agent/model"
	"github.com/uubulb/darwinsmc"
)

func GetTemperatures() ([]*model.SensorTemperature, error) {
	sensorMap, err := darwinsmc.GetThermals()
	if err != nil {
		return nil, err
	}

	tempStat := []*model.SensorTemperature{}
	for key, value := range sensorMap {
		if value > 0 {
			tempStat = append(tempStat, &model.SensorTemperature{
				Name:        key,
				Temperature: value,
			})
		}
	}

	sort.Slice(tempStat, func(i, j int) bool {
		return tempStat[i].Name < tempStat[j].Name
	})

	return tempStat, nil
}
