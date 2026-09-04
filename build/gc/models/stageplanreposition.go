package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplanrepositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplanrepositionDud struct { 
    

}

// Stageplanreposition
type Stageplanreposition struct { 
    // After - The ID of the Stageplan to place this Stageplan after. Omit or null to move to the front.
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Stageplanreposition) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplanreposition) MarshalJSON() ([]byte, error) {
    type Alias Stageplanreposition

    if StageplanrepositionMarshalled {
        return []byte("{}"), nil
    }
    StageplanrepositionMarshalled = true

    return json.Marshal(&struct {
        
        After string `json:"after"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

