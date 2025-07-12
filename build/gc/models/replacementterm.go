package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReplacementtermMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReplacementtermDud struct { 
    


    


    

}

// Replacementterm
type Replacementterm struct { 
    // VarType - Replacement term type
    VarType string `json:"type"`


    // ExistingValue
    ExistingValue string `json:"existingValue"`


    // UpdatedValue
    UpdatedValue string `json:"updatedValue"`

}

// String returns a JSON representation of the model
func (o *Replacementterm) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Replacementterm) MarshalJSON() ([]byte, error) {
    type Alias Replacementterm

    if ReplacementtermMarshalled {
        return []byte("{}"), nil
    }
    ReplacementtermMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ExistingValue string `json:"existingValue"`
        
        UpdatedValue string `json:"updatedValue"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

