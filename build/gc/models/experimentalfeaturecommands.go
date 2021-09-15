package models

type ExperimentalFeature string

const (
	DummyCommand       ExperimentalFeature = "dummy_command"
	AlternativeFormats ExperimentalFeature = "alternative_formats"
	TransformData      ExperimentalFeature = "transform_data"
)

func (f ExperimentalFeature) String() string {
	return string(f)
}

func (f ExperimentalFeature) Description() string {
	switch f {
	case DummyCommand:
		return "Dummy command description..."
	case AlternativeFormats:
		return "Send and receive data as alternative format to JSON"
	case TransformData:
		return "Transform output data with a Go template"
	default:
		return ""
	}
}
