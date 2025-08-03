package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TemporalunitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TemporalunitDud struct { 
    


    


    


    

}

// Temporalunit
type Temporalunit struct { 
    // DurationEstimated
    DurationEstimated bool `json:"durationEstimated"`


    // Duration
    Duration Duration `json:"duration"`


    // TimeBased
    TimeBased bool `json:"timeBased"`


    // DateBased
    DateBased bool `json:"dateBased"`

}

// String returns a JSON representation of the model
func (o *Temporalunit) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Temporalunit) MarshalJSON() ([]byte, error) {
    type Alias Temporalunit

    if TemporalunitMarshalled {
        return []byte("{}"), nil
    }
    TemporalunitMarshalled = true

    return json.Marshal(&struct {
        
        DurationEstimated bool `json:"durationEstimated"`
        
        Duration Duration `json:"duration"`
        
        TimeBased bool `json:"timeBased"`
        
        DateBased bool `json:"dateBased"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

