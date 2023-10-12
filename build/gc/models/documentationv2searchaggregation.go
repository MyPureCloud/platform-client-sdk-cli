package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Documentationv2searchaggregationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Documentationv2searchaggregationDud struct { 
    


    


    


    


    


    

}

// Documentationv2searchaggregation
type Documentationv2searchaggregation struct { 
    // Field - The field used for aggregation
    Field string `json:"field"`


    // Name - The name of the aggregation. The response aggregation uses this name.
    Name string `json:"name"`


    // VarType - The type of aggregation to perform
    VarType string `json:"type"`


    // Value - A value to use for aggregation
    Value string `json:"value"`


    // Size - The number aggregations results to return out of the entire result set
    Size int `json:"size"`


    // Order - The order in which aggregation results are sorted
    Order []string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Documentationv2searchaggregation) String() string {
    
    
    
    
    
     o.Order = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentationv2searchaggregation) MarshalJSON() ([]byte, error) {
    type Alias Documentationv2searchaggregation

    if Documentationv2searchaggregationMarshalled {
        return []byte("{}"), nil
    }
    Documentationv2searchaggregationMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        Size int `json:"size"`
        
        Order []string `json:"order"`
        *Alias
    }{

        


        


        


        


        


        
        Order: []string{""},
        

        Alias: (*Alias)(u),
    })
}

