package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatehistoricalavailabilityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatehistoricalavailabilityDud struct { 
    


    

}

// Aggregatehistoricalavailability
type Aggregatehistoricalavailability struct { 
    // Weekly - All available week offsets from the historical start date.
    Weekly []int `json:"weekly"`


    // Yearly - All available historical year offsets from the forecast start date.
    Yearly []int `json:"yearly"`

}

// String returns a JSON representation of the model
func (o *Aggregatehistoricalavailability) String() string {
     o.Weekly = []int{0} 
     o.Yearly = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatehistoricalavailability) MarshalJSON() ([]byte, error) {
    type Alias Aggregatehistoricalavailability

    if AggregatehistoricalavailabilityMarshalled {
        return []byte("{}"), nil
    }
    AggregatehistoricalavailabilityMarshalled = true

    return json.Marshal(&struct {
        
        Weekly []int `json:"weekly"`
        
        Yearly []int `json:"yearly"`
        *Alias
    }{

        
        Weekly: []int{0},
        


        
        Yearly: []int{0},
        

        Alias: (*Alias)(u),
    })
}

