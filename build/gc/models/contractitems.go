package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContractitemsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContractitemsDud struct { 
    


    

}

// Contractitems
type Contractitems struct { 
    // VarType
    VarType []string `json:"type"`


    // Pattern
    Pattern string `json:"pattern"`

}

// String returns a JSON representation of the model
func (o *Contractitems) String() string {
     o.VarType = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contractitems) MarshalJSON() ([]byte, error) {
    type Alias Contractitems

    if ContractitemsMarshalled {
        return []byte("{}"), nil
    }
    ContractitemsMarshalled = true

    return json.Marshal(&struct {
        
        VarType []string `json:"type"`
        
        Pattern string `json:"pattern"`
        *Alias
    }{

        
        VarType: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

