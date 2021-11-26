package strategy

import (
    "math"

    "../models"
)

type IDriverAllocation interface {
    GetDriver(drivers []models.Driver, rider models.Rider) string
}

type NearestDriverAllocation struct {
}

func (d *NearestDriverAllocation) GetDriver(drivers []models.Driver, rider models.Rider) string {
    minDistance := math.MaxFloat64
    var targetDriver models.Driver
    for _, driver := range drivers {
        if !driver.IsAvailable() {
            continue
        }
        distance := rider.Coordinate.FindDistanceFrom(driver.Coordinate)
        if distance < minDistance {
            minDistance = distance
            targetDriver = driver
        }
    }
    return targetDriver.Id
}
