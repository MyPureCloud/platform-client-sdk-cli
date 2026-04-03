package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternaleventsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternaleventsresponseDud struct { 
    

}

// Externaleventsresponse
type Externaleventsresponse struct { 
    // UnprocessedEntities - List of events that failed processing.
    UnprocessedEntities []Unprocessedexternalevent `json:"unprocessedEntities"`

}

// String returns a JSON representation of the model
func (o *Externaleventsresponse) String() string {
     o.UnprocessedEntities = []Unprocessedexternalevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externaleventsresponse) MarshalJSON() ([]byte, error) {
    type Alias Externaleventsresponse

    if ExternaleventsresponseMarshalled {
        return []byte("{}"), nil
    }
    ExternaleventsresponseMarshalled = true

    return json.Marshal(&struct {
        
        UnprocessedEntities []Unprocessedexternalevent `json:"unprocessedEntities"`
        *Alias
    }{

        
        UnprocessedEntities: []Unprocessedexternalevent{{}},
        

        Alias: (*Alias)(u),
    })
}

