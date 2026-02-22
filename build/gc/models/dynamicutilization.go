package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicutilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicutilizationDud struct { }

// Dynamicutilization
type Dynamicutilization struct { }

// String returns a JSON representation of the model
func (o *Dynamicutilization) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicutilization) MarshalJSON() ([]byte, error) {
    type Alias Dynamicutilization

    if DynamicutilizationMarshalled {
        return []byte("{}"), nil
    }
    DynamicutilizationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

