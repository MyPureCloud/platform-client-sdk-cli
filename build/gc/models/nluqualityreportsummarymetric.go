package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NluqualityreportsummarymetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NluqualityreportsummarymetricDud struct { 
    


    

}

// Nluqualityreportsummarymetric
type Nluqualityreportsummarymetric struct { 
    // Name - The name of the metric. e.g. recall, f1_score
    Name string `json:"name"`


    // Value - The value of the metric
    Value float32 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummarymetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nluqualityreportsummarymetric) MarshalJSON() ([]byte, error) {
    type Alias Nluqualityreportsummarymetric

    if NluqualityreportsummarymetricMarshalled {
        return []byte("{}"), nil
    }
    NluqualityreportsummarymetricMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value float32 `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

