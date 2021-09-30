package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuintradayforecastdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuintradayforecastdataDud struct { 
    


    

}

// Buintradayforecastdata
type Buintradayforecastdata struct { 
    // Offered - The number of interactions routed into the queues in the selected planning groups for the given media type for an agent to answer
    Offered float64 `json:"offered"`


    // AverageHandleTimeSeconds - The average handle time in seconds an agent spent handling interactions
    AverageHandleTimeSeconds float64 `json:"averageHandleTimeSeconds"`

}

// String returns a JSON representation of the model
func (o *Buintradayforecastdata) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buintradayforecastdata) MarshalJSON() ([]byte, error) {
    type Alias Buintradayforecastdata

    if BuintradayforecastdataMarshalled {
        return []byte("{}"), nil
    }
    BuintradayforecastdataMarshalled = true

    return json.Marshal(&struct { 
        Offered float64 `json:"offered"`
        
        AverageHandleTimeSeconds float64 `json:"averageHandleTimeSeconds"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

