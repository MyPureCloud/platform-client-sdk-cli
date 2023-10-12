package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivityquerymetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivityquerymetricDud struct { 
    


    

}

// Flowactivityquerymetric
type Flowactivityquerymetric struct { 
    // Metric - The requested metric
    Metric string `json:"metric"`


    // Details - Flag for including observation details for this metric in the response
    Details bool `json:"details"`

}

// String returns a JSON representation of the model
func (o *Flowactivityquerymetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivityquerymetric) MarshalJSON() ([]byte, error) {
    type Alias Flowactivityquerymetric

    if FlowactivityquerymetricMarshalled {
        return []byte("{}"), nil
    }
    FlowactivityquerymetricMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Details bool `json:"details"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

