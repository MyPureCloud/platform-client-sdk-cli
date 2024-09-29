package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertablepresencesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertablepresencesDud struct { 
    

}

// Alertablepresences
type Alertablepresences struct { 
    // AlertablePresences - The list of alertable system presences.
    AlertablePresences []string `json:"alertablePresences"`

}

// String returns a JSON representation of the model
func (o *Alertablepresences) String() string {
     o.AlertablePresences = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertablepresences) MarshalJSON() ([]byte, error) {
    type Alias Alertablepresences

    if AlertablepresencesMarshalled {
        return []byte("{}"), nil
    }
    AlertablepresencesMarshalled = true

    return json.Marshal(&struct {
        
        AlertablePresences []string `json:"alertablePresences"`
        *Alias
    }{

        
        AlertablePresences: []string{""},
        

        Alias: (*Alias)(u),
    })
}

