package strategy

import (
    "../constants"
    "../models"
)


type IFare interface {
    GetPrice(rider models.Rider, origin models.Coordinate, destination models.Coordinate) float64
}

type BasicFare struct {
}

func (b BasicFare) GetPrice(rider models.Rider, origin models.Coordinate, destination models.Coordinate) float64 {
    return destination.GetCoordinates()[1]-origin.GetCoordinates()[1] * constants.PRICE_PER_KM
}

type SpecialVoucherFare struct {

}