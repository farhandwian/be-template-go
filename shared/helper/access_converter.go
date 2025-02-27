package helper

// GetUniqueAccessObjects iterates over the APIData entries in the ApiPrinter and returns a unique list of objects (from x-access-keto).
func GetUniqueAccessObjects(apiDataList []APIData) []string {
	objectSet := make(map[string]struct{})
	for _, apiData := range apiDataList {
		// We assume that the AccessKeto field is always set
		obj := apiData.AccessKeto.Object
		if obj != "" {
			objectSet[obj] = struct{}{}
		}
	}

	var objects []string
	for obj := range objectSet {
		objects = append(objects, obj)
	}
	return objects
}
