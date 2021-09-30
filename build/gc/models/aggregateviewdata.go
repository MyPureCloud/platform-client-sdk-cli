package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregateviewdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregateviewdataDud struct { 
    


    

}

// Aggregateviewdata
type Aggregateviewdata struct { 
    // Name
    Name string `json:"name"`


    // Stats
    Stats Statisticalsummary `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Aggregateviewdata) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregateviewdata) MarshalJSON() ([]byte, error) {
    type Alias Aggregateviewdata

    if AggregateviewdataMarshalled {
        return []byte("{}"), nil
    }
    AggregateviewdataMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Stats Statisticalsummary `json:"stats"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

