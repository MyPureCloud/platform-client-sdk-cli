package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkflowtargetsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkflowtargetsettingsDud struct { 
    

}

// Workflowtargetsettings
type Workflowtargetsettings struct { 
    // DataFormat - The data format to use when invoking target.
    DataFormat string `json:"dataFormat"`

}

// String returns a JSON representation of the model
func (o *Workflowtargetsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workflowtargetsettings) MarshalJSON() ([]byte, error) {
    type Alias Workflowtargetsettings

    if WorkflowtargetsettingsMarshalled {
        return []byte("{}"), nil
    }
    WorkflowtargetsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DataFormat string `json:"dataFormat"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

