package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryresponsemetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryresponsemetricDud struct { 
    


    

}

// Developmentactivityaggregatequeryresponsemetric
type Developmentactivityaggregatequeryresponsemetric struct { 
    // Metric - The metric this applies to
    Metric string `json:"metric"`


    // Stats - The aggregated values for this metric
    Stats Developmentactivityaggregatequeryresponsestatistics `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsemetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryresponsemetric) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryresponsemetric

    if DevelopmentactivityaggregatequeryresponsemetricMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryresponsemetricMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Stats Developmentactivityaggregatequeryresponsestatistics `json:"stats"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

