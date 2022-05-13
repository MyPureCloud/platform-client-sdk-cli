package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictiveroutingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictiveroutingDud struct { 
    

}

// Predictiverouting
type Predictiverouting struct { 
    // RespectSkills - A switch used to determine if agent skills will be considered.
    RespectSkills bool `json:"respectSkills"`

}

// String returns a JSON representation of the model
func (o *Predictiverouting) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictiverouting) MarshalJSON() ([]byte, error) {
    type Alias Predictiverouting

    if PredictiveroutingMarshalled {
        return []byte("{}"), nil
    }
    PredictiveroutingMarshalled = true

    return json.Marshal(&struct {
        
        RespectSkills bool `json:"respectSkills"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

