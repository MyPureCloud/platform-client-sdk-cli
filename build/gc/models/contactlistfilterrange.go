package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistfilterrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistfilterrangeDud struct { 
    


    


    


    


    

}

// Contactlistfilterrange
type Contactlistfilterrange struct { 
    // Min - The minimum value of the range. Required for the operator BETWEEN.
    Min string `json:"min"`


    // Max - The maximum value of the range. Required for the operator BETWEEN.
    Max string `json:"max"`


    // MinInclusive - Whether or not to include the minimum in the range.
    MinInclusive bool `json:"minInclusive"`


    // MaxInclusive - Whether or not to include the maximum in the range.
    MaxInclusive bool `json:"maxInclusive"`


    // InSet - A set of values that the contact data should be in. Required for the IN operator.
    InSet []string `json:"inSet"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilterrange) String() string {
    
    
    
    
     o.InSet = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistfilterrange) MarshalJSON() ([]byte, error) {
    type Alias Contactlistfilterrange

    if ContactlistfilterrangeMarshalled {
        return []byte("{}"), nil
    }
    ContactlistfilterrangeMarshalled = true

    return json.Marshal(&struct {
        
        Min string `json:"min"`
        
        Max string `json:"max"`
        
        MinInclusive bool `json:"minInclusive"`
        
        MaxInclusive bool `json:"maxInclusive"`
        
        InSet []string `json:"inSet"`
        *Alias
    }{

        


        


        


        


        
        InSet: []string{""},
        

        Alias: (*Alias)(u),
    })
}

