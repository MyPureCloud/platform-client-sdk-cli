package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulingsettingsDud struct { 
    

}

// Buschedulingsettings
type Buschedulingsettings struct { 
    // MessageSeverities - Schedule generation message severity configuration
    MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`

}

// String returns a JSON representation of the model
func (o *Buschedulingsettings) String() string {
     o.MessageSeverities = []Schedulermessagetypeseverity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulingsettings) MarshalJSON() ([]byte, error) {
    type Alias Buschedulingsettings

    if BuschedulingsettingsMarshalled {
        return []byte("{}"), nil
    }
    BuschedulingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`
        *Alias
    }{

        
        MessageSeverities: []Schedulermessagetypeseverity{{}},
        

        Alias: (*Alias)(u),
    })
}

