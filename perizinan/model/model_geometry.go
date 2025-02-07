package model

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Geometry struct {
	Type        GeometryType `json:"type"`
	Coordinates []float64    `json:"coordinates"`
}

type GeometryType string

const GeometryTypePoint GeometryType = "Point"

func NewGeometryPoint(lng, lat float64) Geometry {
	return Geometry{
		Type:        GeometryTypePoint,
		Coordinates: []float64{lng, lat},
	}
}

func (g Geometry) GormDataType() string {
	return "geometry"
}

func (g Geometry) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_GeomFromText(?)",
		Vars: []any{fmt.Sprintf("POINT(%f %f)", g.Coordinates[0], g.Coordinates[1])},
	}
}

func (g *Geometry) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	// Check if we have enough bytes
	if len(b) < 25 {
		return fmt.Errorf("invalid POINT data: %v", b)
	}

	// Skip first byte (0 or 1 indicating little/big endian)
	// Next 4 bytes are for the SRID (usually 0)
	// Then comes the actual point data

	// Check if it's a POINT (type should be 1)
	if b[5] != 0x01 {
		return fmt.Errorf("not a POINT geometry")
	}

	// Read X and Y coordinates (doubles)
	x := math.Float64frombits(binary.LittleEndian.Uint64(b[9:17]))
	y := math.Float64frombits(binary.LittleEndian.Uint64(b[17:25]))

	g.Type = "Point"
	g.Coordinates = []float64{x, y} // [longitude, latitude]

	return nil
}
