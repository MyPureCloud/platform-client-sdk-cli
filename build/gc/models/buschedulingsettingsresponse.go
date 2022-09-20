package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulingsettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulingsettingsresponseDud struct { 
    

}

// Buschedulingsettingsresponse
type Buschedulingsettingsresponse struct { 
    // MessageSeverities - Schedule generation message severity configuration
    MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`

}

// String returns a JSON representation of the model
func (o *Buschedulingsettingsresponse) String() string {
     o.MessageSeverities = []Schedulermessagetypeseverity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulingsettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Buschedulingsettingsresponse

    if BuschedulingsettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    BuschedulingsettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`
        *Alias
    }{

        
        MessageSeverities: []Schedulermessagetypeseverity{{}},
        

        Alias: (*Alias)(u),
    })
}

