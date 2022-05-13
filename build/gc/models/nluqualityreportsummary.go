package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluqualityreportsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluqualityreportsummaryDud struct { 
    

}

// Nluqualityreportsummary
type Nluqualityreportsummary struct { 
    // Metrics - The list of metrics in the summary
    Metrics []Nluqualityreportsummarymetric `json:"metrics"`

}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummary) String() string {
     o.Metrics = []Nluqualityreportsummarymetric{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluqualityreportsummary) MarshalJSON() ([]byte, error) {
    type Alias Nluqualityreportsummary

    if NluqualityreportsummaryMarshalled {
        return []byte("{}"), nil
    }
    NluqualityreportsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Nluqualityreportsummarymetric `json:"metrics"`
        *Alias
    }{

        
        Metrics: []Nluqualityreportsummarymetric{{}},
        

        Alias: (*Alias)(u),
    })
}

