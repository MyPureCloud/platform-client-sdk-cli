package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShrinkageoverridesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShrinkageoverridesDud struct { 
    


    

}

// Shrinkageoverrides
type Shrinkageoverrides struct { 
    // Clear - Set true to clear the shrinkage interval overrides
    Clear bool `json:"clear"`


    // Values - List of interval shrinkage overrides
    Values []Shrinkageoverride `json:"values"`

}

// String returns a JSON representation of the model
func (o *Shrinkageoverrides) String() string {
    
     o.Values = []Shrinkageoverride{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shrinkageoverrides) MarshalJSON() ([]byte, error) {
    type Alias Shrinkageoverrides

    if ShrinkageoverridesMarshalled {
        return []byte("{}"), nil
    }
    ShrinkageoverridesMarshalled = true

    return json.Marshal(&struct {
        
        Clear bool `json:"clear"`
        
        Values []Shrinkageoverride `json:"values"`
        *Alias
    }{

        


        
        Values: []Shrinkageoverride{{}},
        

        Alias: (*Alias)(u),
    })
}

