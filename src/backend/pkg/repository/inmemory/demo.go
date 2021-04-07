package inmemory

import (
	"encoding/json"

	"github.com/backendlersGuideToFrontend-presentation/src/backend/pkg/repository"
	"github.com/google/uuid"
)

var demoDataSet string = `
{
	"4cbfe717-51a1-4414-9384-83e648dbd764": {
		"id": "4cbfe717-51a1-4414-9384-83e648dbd764",
		"name":"Waffles", 
		"species":"dog" , 
		"properties": {
			"furColor":"black"
		}
	}
}
`

func demoData(loadDemoData bool) map[uuid.UUID]repository.Animal {
	data := make(map[uuid.UUID]repository.Animal)

	if !loadDemoData {
		return data
	}

	if err := json.Unmarshal([]byte(demoDataSet), &data); err != nil {
		panic(err)
	}

	return data
}
