package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicutilizationruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicutilizationruleDud struct { }

// Dynamicutilizationrule
type Dynamicutilizationrule struct { }

// String returns a JSON representation of the model
func (o *Dynamicutilizationrule) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicutilizationrule) MarshalJSON() ([]byte, error) {
    type Alias Dynamicutilizationrule

    if DynamicutilizationruleMarshalled {
        return []byte("{}"), nil
    }
    DynamicutilizationruleMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{
        Alias: (*Alias)(u),
    })
}

