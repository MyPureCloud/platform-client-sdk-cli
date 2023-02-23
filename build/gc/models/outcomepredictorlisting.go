package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomepredictorlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomepredictorlistingDud struct { 
    

}

// Outcomepredictorlisting
type Outcomepredictorlisting struct { 
    // Entities
    Entities []Outcomepredictor `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Outcomepredictorlisting) String() string {
     o.Entities = []Outcomepredictor{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomepredictorlisting) MarshalJSON() ([]byte, error) {
    type Alias Outcomepredictorlisting

    if OutcomepredictorlistingMarshalled {
        return []byte("{}"), nil
    }
    OutcomepredictorlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Outcomepredictor `json:"entities"`
        *Alias
    }{

        
        Entities: []Outcomepredictor{{}},
        

        Alias: (*Alias)(u),
    })
}

