package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionresponselistingDud struct { 
    

}

// Outcomeattributionresponselisting
type Outcomeattributionresponselisting struct { 
    // Entities
    Entities []Outcomeattributionresultsresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Outcomeattributionresponselisting) String() string {
     o.Entities = []Outcomeattributionresultsresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionresponselisting

    if OutcomeattributionresponselistingMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Outcomeattributionresultsresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Outcomeattributionresultsresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

