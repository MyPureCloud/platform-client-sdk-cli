package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeattributionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeattributionlistingDud struct { 
    


    

}

// Outcomeattributionlisting
type Outcomeattributionlisting struct { 
    // Entities
    Entities []Outcomeattributionrequest `json:"entities"`


    // PercentFailedThreshold - Optional percent failed threshold for validation errors; if reached will halt the job. Default is 5 percent, allowed values 0 to 100.
    PercentFailedThreshold int `json:"percentFailedThreshold"`

}

// String returns a JSON representation of the model
func (o *Outcomeattributionlisting) String() string {
     o.Entities = []Outcomeattributionrequest{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeattributionlisting) MarshalJSON() ([]byte, error) {
    type Alias Outcomeattributionlisting

    if OutcomeattributionlistingMarshalled {
        return []byte("{}"), nil
    }
    OutcomeattributionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Outcomeattributionrequest `json:"entities"`
        
        PercentFailedThreshold int `json:"percentFailedThreshold"`
        *Alias
    }{

        
        Entities: []Outcomeattributionrequest{{}},
        


        

        Alias: (*Alias)(u),
    })
}

