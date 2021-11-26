package manager

import (
    "fmt"

    "../models"
)

type DriverManager struct {
    drivers map[string]models.Driver
}


func (manager DriverManager) AddDriver(driver models.Driver) (err error)  {
    if _, exists := manager.drivers[driver.Id]; exists {
        return fmt.Errorf("driver already exists")
    }
    manager.drivers[driver.Id] = driver
    return
}

func (manager DriverManager) HasAvailableDriver() bool {
    for _, driver := range manager.drivers {
        if driver.IsAvailable() {
            return true
        }
    }
    return false
}

func (manager DriverManager) GetAvailableDrivers() (drivers []models.Driver)  {
    for _, driver := range manager.drivers {
        if driver.IsAvailable() {
            drivers = append(drivers, driver)
        }
    }
    return
}

func (manager DriverManager) GetDriver(driverId string) (driver models.Driver, err error) {
    if _, exists := manager.drivers[driverId]; !exists {
        return driver, fmt.Errorf("driver doesn't exists")
    }
    return manager.drivers[driverId], nil
}
