package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryresponsestatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryresponsestatisticsDud struct { 
    


    


    


    

}

// Developmentactivityaggregatequeryresponsestatistics
type Developmentactivityaggregatequeryresponsestatistics struct { 
    // Count - The count for this metric
    Count int `json:"count"`


    // Min - The minimum value in this metric
    Min float32 `json:"min"`


    // Max - The maximum value in this metric
    Max float32 `json:"max"`


    // Sum - The total of the values for this metric
    Sum float32 `json:"sum"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsestatistics) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryresponsestatistics) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryresponsestatistics

    if DevelopmentactivityaggregatequeryresponsestatisticsMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryresponsestatisticsMarshalled = true

    return json.Marshal(&struct {
        
        Count int `json:"count"`
        
        Min float32 `json:"min"`
        
        Max float32 `json:"max"`
        
        Sum float32 `json:"sum"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

