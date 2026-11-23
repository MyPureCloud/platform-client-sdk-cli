package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdatatypefacetslimitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdatatypefacetslimitDud struct { 
    


    


    


    

}

// Capabilitiesdatatypefacetslimit
type Capabilitiesdatatypefacetslimit struct { 
    // StringValue
    StringValue string `json:"stringValue"`


    // BooleanValue
    BooleanValue bool `json:"booleanValue"`


    // IntegerValue
    IntegerValue int `json:"integerValue"`


    // DecimalValue
    DecimalValue float64 `json:"decimalValue"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdatatypefacetslimit) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdatatypefacetslimit) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdatatypefacetslimit

    if CapabilitiesdatatypefacetslimitMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdatatypefacetslimitMarshalled = true

    return json.Marshal(&struct {
        
        StringValue string `json:"stringValue"`
        
        BooleanValue bool `json:"booleanValue"`
        
        IntegerValue int `json:"integerValue"`
        
        DecimalValue float64 `json:"decimalValue"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

