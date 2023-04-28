package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightstrendmetricitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightstrendmetricitemDud struct { 
    


    

}

// Insightstrendmetricitem
type Insightstrendmetricitem struct { 
    // Metric - The gamification metric for the trend
    Metric Addressableentityref `json:"metric"`


    // Trends - Trends for the metric
    Trends Insightstrends `json:"trends"`

}

// String returns a JSON representation of the model
func (o *Insightstrendmetricitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightstrendmetricitem) MarshalJSON() ([]byte, error) {
    type Alias Insightstrendmetricitem

    if InsightstrendmetricitemMarshalled {
        return []byte("{}"), nil
    }
    InsightstrendmetricitemMarshalled = true

    return json.Marshal(&struct {
        
        Metric Addressableentityref `json:"metric"`
        
        Trends Insightstrends `json:"trends"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

