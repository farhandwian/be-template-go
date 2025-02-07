package seeder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Features[PROPERTIES any] struct {
	Geometry   Geometry   `json:"geometry"`
	Properties PROPERTIES `json:"properties"`
}

type GeoJSON[PROPERTIES any] struct {
	Features []Features[PROPERTIES] `json:"features"`
}

func loadJSON[T any](filename string) (*T, error) {

	var objs T

	file, err := os.Open(filepath.Join("seeder", filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&objs); err != nil {
		return nil, err
	}

	return &objs, nil

}

func getTags[T any](excludes ...string) []string {
	var t T
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()
	if typ.Kind() != reflect.Struct {
		return nil
	}

	// Create a map for quick lookup of excluded fields
	excludeMap := make(map[string]bool)
	for _, exclude := range excludes {
		excludeMap[exclude] = true
	}

	var tags []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			// Split the tag on comma and take the first part
			// This handles cases like `json:"name,omitempty"`
			tagParts := strings.Split(tag, ",")
			tagName := tagParts[0]

			// Check if the tag should be excluded
			if !excludeMap[tagName] {
				tags = append(tags, tagName)
			}
		}
	}

	return tags
}

func mapGeoJSONToItems[T any](filename string) ([]T, error) {
	// Load the GeoJSON data using the provided LoadJSON function
	geoJSON, err := loadJSON[GeoJSON[map[string]any]](filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load GeoJSON: %w", err)
	}

	items := make([]T, len(geoJSON.Features))

	for i, feature := range geoJSON.Features {
		var item T
		itemValue := reflect.ValueOf(&item).Elem()

		// Process geometry
		if err := processGeometry(itemValue, feature.Geometry); err != nil {
			return nil, err
		}

		// Process properties
		if err := processProperties(itemValue, feature.Properties); err != nil {
			return nil, err
		}

		items[i] = item
	}

	return items, nil
}

func processGeometry(itemValue reflect.Value, geometry Geometry) error {
	geometryField := itemValue.FieldByName("Geometry")
	if geometryField.IsValid() && geometryField.Type() == reflect.TypeOf(datatypes.JSON{}) {
		geometryJSON, err := json.Marshal(geometry)
		if err != nil {
			return fmt.Errorf("failed to marshal geometry: %w", err)
		}
		geometryField.Set(reflect.ValueOf(datatypes.JSON(geometryJSON)))
	}
	return nil
}

func processProperties(itemValue reflect.Value, props map[string]any) error {
	itemType := itemValue.Type()
	for j := 0; j < itemType.NumField(); j++ {
		field := itemType.Field(j)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		fieldValue := itemValue.Field(j)
		if !fieldValue.CanSet() {
			continue
		}

		propValue, ok := props[jsonTag]
		if !ok {
			continue
		}

		if err := setField(fieldValue, propValue); err != nil {
			return fmt.Errorf("failed to set field %s: %w", field.Name, err)
		}
	}
	return nil
}

func setField(fieldValue reflect.Value, propValue any) error {
	switch fieldValue.Kind() {
	case reflect.String:
		if propValue != nil {
			fieldValue.SetString(fmt.Sprintf("%v", propValue))
		} else {
			fieldValue.Set(reflect.Zero(fieldValue.Type())) // Set to zero value (empty string)
		}
	case reflect.Bool:
		fieldValue.SetBool(getBool(propValue))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldValue.SetInt(int64(getFloat(propValue)))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldValue.SetUint(uint64(getFloat(propValue)))
	case reflect.Float32, reflect.Float64:
		fieldValue.SetFloat(getFloat(propValue))
	case reflect.Ptr:
		if propValue == nil {
			// Set the pointer field to nil when propValue is nil
			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		} else {
			ptrValue := reflect.New(fieldValue.Type().Elem())
			if err := setField(ptrValue.Elem(), propValue); err != nil {
				return err
			}
			fieldValue.Set(ptrValue)
		}
	case reflect.Slice:
		if slice, ok := propValue.([]any); ok {
			newSlice := reflect.MakeSlice(fieldValue.Type(), len(slice), len(slice))
			for i, v := range slice {
				if err := setField(newSlice.Index(i), v); err != nil {
					return err
				}
			}
			fieldValue.Set(newSlice)
		}
	default:
		return fmt.Errorf("unsupported field type: %v", fieldValue.Kind())
	}
	return nil
}

func getBool(v any) bool {
	if val, ok := v.(bool); ok {
		return val
	}
	return false
}

func getFloat(v any) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		x, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0
		}
		return x
	default:
		return 0
	}
}

func mapJSONToItems[T any](filename string) ([]T, error) {
	// Load the JSON data using the provided LoadJSON function
	data, err := loadJSON[[]map[string]any](filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load JSON: %w", err)
	}

	items := make([]T, len(*data))

	for i, item := range *data {
		var structItem T
		itemValue := reflect.ValueOf(&structItem).Elem()

		// Process coordinate field
		if err := processCoordinate(itemValue, item); err != nil {
			return nil, err
		}

		// process sensor type
		if err := processSensorType(itemValue, item); err != nil {
			return nil, err
		}

		// Process other fields
		if err := processFields(itemValue, item); err != nil {
			return nil, err
		}

		items[i] = structItem
	}

	return items, nil
}

func processSensorType(itemValue reflect.Value, item map[string]any) error {
	sensorType, ok := item["sensor_type"]
	if !ok {
		return nil
	}

	sensorTypeField := itemValue.FieldByName("SensorType")
	if !sensorTypeField.IsValid() {
		return nil
	}

	// Check if the field is datatypes.JSON
	if sensorTypeField.Type() == reflect.TypeOf(datatypes.JSON{}) {
		sensorTypeJSON, err := json.Marshal(sensorType)
		if err != nil {
			return fmt.Errorf("failed to marshal sensor_type: %w", err)
		}

		sensorTypeField.Set(reflect.ValueOf(datatypes.JSON(sensorTypeJSON)))
	}

	return nil
}
func processCoordinate(itemValue reflect.Value, item map[string]any) error {
	coordinateField, ok := item["coordinate"]
	if !ok {
		return nil // No coordinate field, skip
	}

	coordinateStr, ok := coordinateField.(string)
	if !ok {
		return fmt.Errorf("coordinate field is not a string")
	}

	coords := strings.Split(coordinateStr, " ")
	if len(coords) != 2 {
		return fmt.Errorf("invalid coordinate format")
	}

	longitude := getFloat(coords[0])

	latitude := getFloat(coords[1])

	geometry := Geometry{
		Type:        "Point",
		Coordinates: []float64{longitude, latitude},
	}

	geometryJSON, err := json.Marshal(geometry)
	if err != nil {
		return fmt.Errorf("failed to marshal geometry: %w", err)
	}

	geometryField := itemValue.FieldByName("Geometry")
	if geometryField.IsValid() && geometryField.Type() == reflect.TypeOf(datatypes.JSON{}) {
		geometryField.Set(reflect.ValueOf(datatypes.JSON(geometryJSON)))
	}

	return nil
}

func processFields(itemValue reflect.Value, item map[string]any) error {
	itemType := itemValue.Type()
	for j := 0; j < itemType.NumField(); j++ {
		field := itemType.Field(j)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "coordinate" || jsonTag == "sensor_type" {
			continue
		}

		fieldValue := itemValue.Field(j)
		if !fieldValue.CanSet() {
			continue
		}

		propValue, ok := item[jsonTag]
		if !ok {
			continue
		}

		if err := setField(fieldValue, propValue); err != nil {
			return fmt.Errorf("failed to set field %s: %w", field.Name, err)
		}
	}
	return nil
}

func Seeder[T any](db *gorm.DB, filename string, chunkSize int) error {
	items, err := mapGeoJSONToItems[T](filename)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return fmt.Errorf("no valid data to insert")
	}

	tags := getTags[T]("id")

	return insertInChunks(db, items, tags, chunkSize)
}

func SeederWithoutGeoJSON[T any](db *gorm.DB, filename string, chunkSize int) error {
	items, err := mapJSONToItems[T](filename)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return fmt.Errorf("no valid data to insert")
	}

	tags := getTags[T]("id", "coordinate")

	return insertInChunks(db, items, tags, chunkSize)
}

func insertInChunks[T any](db *gorm.DB, items []T, tags []string, chunkSize int) error {
	totalItems := len(items)
	for i := 0; i < totalItems; i += chunkSize {
		end := i + chunkSize
		if end > totalItems {
			end = totalItems
		}

		chunk := items[i:end]

		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(tags),
		}).Create(&chunk).Error

		if err != nil {
			return fmt.Errorf("error creating data chunk %d-%d: %v", i, end-1, err)
		}

	}

	return nil
}
