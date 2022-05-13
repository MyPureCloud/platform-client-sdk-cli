package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatemetricdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatemetricdataDud struct { 
    


    


    

}

// Aggregatemetricdata
type Aggregatemetricdata struct { 
    // Metric
    Metric string `json:"metric"`


    // Qualifier
    Qualifier string `json:"qualifier"`


    // Stats
    Stats Statisticalsummary `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Aggregatemetricdata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatemetricdata) MarshalJSON() ([]byte, error) {
    type Alias Aggregatemetricdata

    if AggregatemetricdataMarshalled {
        return []byte("{}"), nil
    }
    AggregatemetricdataMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Qualifier string `json:"qualifier"`
        
        Stats Statisticalsummary `json:"stats"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

