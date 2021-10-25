package models


// Define data to use in all packages 
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}