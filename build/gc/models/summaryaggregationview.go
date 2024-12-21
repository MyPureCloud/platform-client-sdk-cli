package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryaggregationviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryaggregationviewDud struct { 
    


    


    


    

}

// Summaryaggregationview
type Summaryaggregationview struct { 
    // Target - Target metric name
    Target string `json:"target"`


    // Name - A unique name for this view. Must be distinct from other views and built-in metric names.
    Name string `json:"name"`


    // Function - Type of view you wish to create
    Function string `json:"function"`


    // VarRange - Range of numbers for slicing up data
    VarRange Aggregationrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Summaryaggregationview) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryaggregationview) MarshalJSON() ([]byte, error) {
    type Alias Summaryaggregationview

    if SummaryaggregationviewMarshalled {
        return []byte("{}"), nil
    }
    SummaryaggregationviewMarshalled = true

    return json.Marshal(&struct {
        
        Target string `json:"target"`
        
        Name string `json:"name"`
        
        Function string `json:"function"`
        
        VarRange Aggregationrange `json:"range"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

