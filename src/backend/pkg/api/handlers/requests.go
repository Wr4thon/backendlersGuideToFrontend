package handlers

type (
	UpdateAnimalRequest struct {
		Properties map[string]interface{}
	}

	AddAnimalRequest struct {
		Name       string
		Species    string
		Properties map[string]interface{}
	}
)
