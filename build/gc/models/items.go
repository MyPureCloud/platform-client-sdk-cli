package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ItemsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ItemsDud struct { 
    


    

}

// Items
type Items struct { 
    // VarType
    VarType string `json:"type"`


    // Pattern
    Pattern string `json:"pattern"`

}

// String returns a JSON representation of the model
func (o *Items) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Items) MarshalJSON() ([]byte, error) {
    type Alias Items

    if ItemsMarshalled {
        return []byte("{}"), nil
    }
    ItemsMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Pattern string `json:"pattern"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

