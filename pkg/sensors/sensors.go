//go:build !darwin

package sensors

import (
	"github.com/nezhahq/agent/model"
	"github.com/shirou/gopsutil/v4/sensors"
)

func GetTemperatures() ([]*model.SensorTemperature, error) {
	temperatures, err := sensors.SensorsTemperatures()
	if err != nil {
		return nil, err
	}

	tempStat := []*model.SensorTemperature{}
	for _, t := range temperatures {
		if t.Temperature > 0 {
			tempStat = append(tempStat, &model.SensorTemperature{
				Name:        t.SensorKey,
				Temperature: t.Temperature,
			})
		}
	}

	return tempStat, nil
}
