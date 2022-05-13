package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumericrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumericrangeDud struct { 
    


    


    


    

}

// Numericrange
type Numericrange struct { 
    // Gt - Greater than
    Gt float32 `json:"gt"`


    // Gte - Greater than or equal to
    Gte float32 `json:"gte"`


    // Lt - Less than
    Lt float32 `json:"lt"`


    // Lte - Less than or equal to
    Lte float32 `json:"lte"`

}

// String returns a JSON representation of the model
func (o *Numericrange) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numericrange) MarshalJSON() ([]byte, error) {
    type Alias Numericrange

    if NumericrangeMarshalled {
        return []byte("{}"), nil
    }
    NumericrangeMarshalled = true

    return json.Marshal(&struct {
        
        Gt float32 `json:"gt"`
        
        Gte float32 `json:"gte"`
        
        Lt float32 `json:"lt"`
        
        Lte float32 `json:"lte"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

