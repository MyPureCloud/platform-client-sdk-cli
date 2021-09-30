package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticssessionmetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticssessionmetricDud struct { 
    


    


    

}

// Analyticssessionmetric
type Analyticssessionmetric struct { 
    // EmitDate - Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EmitDate time.Time `json:"emitDate"`


    // Name - Unique name of this metric
    Name string `json:"name"`


    // Value - The metric value
    Value int `json:"value"`

}

// String returns a JSON representation of the model
func (o *Analyticssessionmetric) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticssessionmetric) MarshalJSON() ([]byte, error) {
    type Alias Analyticssessionmetric

    if AnalyticssessionmetricMarshalled {
        return []byte("{}"), nil
    }
    AnalyticssessionmetricMarshalled = true

    return json.Marshal(&struct { 
        EmitDate time.Time `json:"emitDate"`
        
        Name string `json:"name"`
        
        Value int `json:"value"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

