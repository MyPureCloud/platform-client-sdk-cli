package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightstrendtotalitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightstrendtotalitemDud struct { 
    

}

// Insightstrendtotalitem
type Insightstrendtotalitem struct { 
    // Trends - Trends for the metric
    Trends Insightstrends `json:"trends"`

}

// String returns a JSON representation of the model
func (o *Insightstrendtotalitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightstrendtotalitem) MarshalJSON() ([]byte, error) {
    type Alias Insightstrendtotalitem

    if InsightstrendtotalitemMarshalled {
        return []byte("{}"), nil
    }
    InsightstrendtotalitemMarshalled = true

    return json.Marshal(&struct {
        
        Trends Insightstrends `json:"trends"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

