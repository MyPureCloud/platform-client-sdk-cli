package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklisttransferconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklisttransferconfigDud struct { 
    

}

// Checklisttransferconfig
type Checklisttransferconfig struct { 
    // Enabled - Whether checklist transfer data is visible to the receiving agent. Defaults to false when not set; newly created Copilots have this enabled by default.
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Checklisttransferconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklisttransferconfig) MarshalJSON() ([]byte, error) {
    type Alias Checklisttransferconfig

    if ChecklisttransferconfigMarshalled {
        return []byte("{}"), nil
    }
    ChecklisttransferconfigMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

