package manager

import (
    "fmt"

    "../models"
)

type RiderManager struct {
    riders map[string]models.Rider
}


func (manager RiderManager) AddRider(rider models.Rider) (err error)  {
    if _, exists := manager.riders[rider.Id]; exists {
        return fmt.Errorf("rider already exists")
    }
    manager.riders[rider.Id] = rider
    return
}

func (manager RiderManager) GetRider(riderId string) (rider models.Rider, err error) {
    if _, exists := manager.riders[riderId]; !exists {
        return rider, fmt.Errorf("rider doesn't exists")
    }
    return manager.riders[riderId], nil
}