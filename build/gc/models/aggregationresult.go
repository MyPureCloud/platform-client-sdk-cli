package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregationresultDud struct { 
    


    


    


    


    

}

// Aggregationresult
type Aggregationresult struct { 
    // VarType
    VarType string `json:"type"`


    // Dimension - For termFrequency aggregations
    Dimension string `json:"dimension"`


    // Metric - For numericRange aggregations
    Metric string `json:"metric"`


    // Count
    Count int `json:"count"`


    // Results
    Results []Aggregationresultentry `json:"results"`

}

// String returns a JSON representation of the model
func (o *Aggregationresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Results = []Aggregationresultentry{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregationresult) MarshalJSON() ([]byte, error) {
    type Alias Aggregationresult

    if AggregationresultMarshalled {
        return []byte("{}"), nil
    }
    AggregationresultMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Metric string `json:"metric"`
        
        Count int `json:"count"`
        
        Results []Aggregationresultentry `json:"results"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Results: []Aggregationresultentry{{}},
        

        
        Alias: (*Alias)(u),
    })
}

