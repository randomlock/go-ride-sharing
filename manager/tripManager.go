package manager

import (
    "fmt"

    "../models"
    "../strategy"
)

type TripManager struct {
    trips map[string]*models.Trip
    tripsPerRider map[string][]models.Rider
    driverManager DriverManager
    FareStrategy strategy.IFare
    DriverAllocationStrategy strategy.IDriverAllocation
}

func NewTripManager(driverManager DriverManager, fareStrategy strategy.IFare, driverAllocationStrategy strategy.IDriverAllocation) *TripManager {
    return &TripManager{driverManager: driverManager, FareStrategy: fareStrategy, DriverAllocationStrategy: driverAllocationStrategy}
}

func (t *TripManager) NewTrip(rider models.Rider, origin models.Coordinate, destination models.Coordinate) (string, error) {
    if !t.driverManager.HasAvailableDriver() {
        return "", fmt.Errorf("no drivers are available")
    }
    if t.IsRiding(rider.Id) {
        return "", fmt.Errorf("rider is already riding. Please finish the current trip")
    }

    if t.IsDestinationValid(origin, destination) {
        return "", fmt.Errorf("destination should be greater than origin")
    }

    driverID := t.DriverAllocationStrategy.GetDriver(t.driverManager.GetAvailableDrivers(), rider)
    driver, err := t.driverManager.GetDriver(driverID)
    if err != nil {
        return "", fmt.Errorf("no driver found")
    }
    price := t.FareStrategy.GetPrice(rider, origin, destination)

    trip := models.NewTrip(driver, rider, origin, destination, price)
    t.trips[trip.Id] = trip
    driver.UpdateStatus(models.STATUS_IN_TRANSIT)
    return trip.Id, nil
}

func (t *TripManager) End(driverId string) error  {
    if _, exists := t.driverManager.drivers[driverId]; !exists {
        return fmt.Errorf("driver doesn't exist")
    }
    driver := t.driverManager.drivers[driverId]
    if !driver.CanEnd() {
        return fmt.Errorf("driver cannot end as it is not in valid state")
    }

    if _, exists := t.trips[driver.CurrentTripId]; !exists {
        return fmt.Errorf("trip doesn't exist for the driver")
    }
    trip := t.trips[driver.CurrentTripId]
    trip.End()
    driver.UpdateStatus(models.STATUS_AVAILABLE)
    return nil
}


func (t TripManager) Withdraw(tripId string) error  {
    if _, exists := t.trips[tripId]; !exists {
        return fmt.Errorf("trip doesn't exist")
    }
    trip := t.trips[tripId]
    if !trip.CanWithdraw() {
        return fmt.Errorf("trip cannot end as it is not in valid state - %s", trip.Status)
    }
    trip.Withdraw()
    return nil
}

func (t *TripManager) IsRiding(riderId string) bool {
    if trips, exists :=  t.tripsPerRider[riderId]; exists {
        for _, trip := range trips {
            if _, exists := t.trips[trip.Id]; exists && t.trips[trip.Id].IsTripFinished() {
                return true
            }
        }
    }
    return false
}

func (t *TripManager) IsDestinationValid(origin models.Coordinate, destination models.Coordinate) bool {
    return destination.FindDistanceFrom(origin) > float64(0)
}


