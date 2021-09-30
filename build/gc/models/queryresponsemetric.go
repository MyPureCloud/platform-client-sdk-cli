package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresponsemetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresponsemetricDud struct { 
    


    

}

// Queryresponsemetric
type Queryresponsemetric struct { 
    // Metric - The metric this applies to
    Metric string `json:"metric"`


    // Stats - The aggregated values for this metric
    Stats Queryresponsestats `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Queryresponsemetric) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresponsemetric) MarshalJSON() ([]byte, error) {
    type Alias Queryresponsemetric

    if QueryresponsemetricMarshalled {
        return []byte("{}"), nil
    }
    QueryresponsemetricMarshalled = true

    return json.Marshal(&struct { 
        Metric string `json:"metric"`
        
        Stats Queryresponsestats `json:"stats"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

