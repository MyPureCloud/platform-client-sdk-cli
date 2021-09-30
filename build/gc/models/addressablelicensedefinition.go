package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddressablelicensedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddressablelicensedefinitionDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Addressablelicensedefinition
type Addressablelicensedefinition struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Addressablelicensedefinition) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addressablelicensedefinition) MarshalJSON() ([]byte, error) {
    type Alias Addressablelicensedefinition

    if AddressablelicensedefinitionMarshalled {
        return []byte("{}"), nil
    }
    AddressablelicensedefinitionMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

