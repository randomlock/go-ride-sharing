package models

type Status string

const STATUS_AVAILABLE Status = "AVAILABLE"
const STATUS_IN_TRANSIT Status = "IN-TRANSIT"

type Driver struct {
    User
    licenseNumber int
    CurrentTripId string
    Status Status
}

func NewDriver(user User, licenseNumber int) *Driver {
    return &Driver{User: user, licenseNumber: licenseNumber, Status: STATUS_AVAILABLE}
}

func (d *Driver) UpdateStatus(status Status)  {
    d.Status = status
}

func (d *Driver) IsAvailable() bool  {
    return d.Status == STATUS_AVAILABLE
}

func (d *Driver) CanEnd() bool {
    return d.Status == STATUS_IN_TRANSIT
}
