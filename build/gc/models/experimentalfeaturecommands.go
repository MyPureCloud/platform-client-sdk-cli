package models

type ExperimentalFeature string

const(
	DummyCommand ExperimentalFeature = "dummy_command"
)

func (f ExperimentalFeature) String() string {
	return string(f)
}

func (f ExperimentalFeature) Description() string {
	switch f {
	case DummyCommand:
		return "Dummy command description..."
	default:
		return ""
	}
}