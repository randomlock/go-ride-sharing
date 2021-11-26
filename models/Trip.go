package models

type TripStatus string

const (
	TRIP_STATUS_COMPLETED TripStatus = "Completed"
	TRIP_STATUS_INITIATED TripStatus = "Initiated"
	TRIP_STATUS_WITHDRAW TripStatus = "WithDraw Trip"
)


type Trip struct {
    Id string
    Driver Driver
    Rider Rider
    Origin Coordinate
    Destination Coordinate
    Status TripStatus
    Price float64
}

func NewTrip(driver Driver, rider Rider, origin Coordinate, destination Coordinate, price float64) *Trip {
    return &Trip{Driver: driver, Rider: rider, Origin: origin, Destination: destination, Status: TRIP_STATUS_INITIATED}
}

func (t *Trip) End()  {
    t.Status = TRIP_STATUS_COMPLETED
}

func (t *Trip) IsTripFinished() bool {
    return t.Status == TRIP_STATUS_COMPLETED || t.Status == TRIP_STATUS_WITHDRAW
}

func (t *Trip) Withdraw()  {
    t.Status = TRIP_STATUS_WITHDRAW
}

func (t *Trip) CanEnd() bool {
    return t.Status == TRIP_STATUS_INITIATED
}

func (t *Trip) CanWithdraw() bool {
    return t.Status == TRIP_STATUS_INITIATED
}


