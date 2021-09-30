package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationemergencynumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationemergencynumberDud struct { 
    


    


    

}

// Locationemergencynumber
type Locationemergencynumber struct { 
    // E164
    E164 string `json:"e164"`


    // Number
    Number string `json:"number"`


    // VarType - The type of emergency number.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Locationemergencynumber) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationemergencynumber) MarshalJSON() ([]byte, error) {
    type Alias Locationemergencynumber

    if LocationemergencynumberMarshalled {
        return []byte("{}"), nil
    }
    LocationemergencynumberMarshalled = true

    return json.Marshal(&struct { 
        E164 string `json:"e164"`
        
        Number string `json:"number"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

