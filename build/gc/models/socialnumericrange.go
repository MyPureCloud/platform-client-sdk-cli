package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialnumericrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialnumericrangeDud struct { 
    


    


    


    


    

}

// Socialnumericrange
type Socialnumericrange struct { 
    // Eq - Equals to
    Eq float32 `json:"eq"`


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
func (o *Socialnumericrange) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialnumericrange) MarshalJSON() ([]byte, error) {
    type Alias Socialnumericrange

    if SocialnumericrangeMarshalled {
        return []byte("{}"), nil
    }
    SocialnumericrangeMarshalled = true

    return json.Marshal(&struct {
        
        Eq float32 `json:"eq"`
        
        Gt float32 `json:"gt"`
        
        Gte float32 `json:"gte"`
        
        Lt float32 `json:"lt"`
        
        Lte float32 `json:"lte"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

