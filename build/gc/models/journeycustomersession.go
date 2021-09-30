package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneycustomersessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneycustomersessionDud struct { 
    


    

}

// Journeycustomersession
type Journeycustomersession struct { 
    // Id - An ID of a Customer/User's session within the Journey System at a point-in-time
    Id string `json:"id"`


    // VarType - The type of the Customer/User's session within the Journey System (e.g. web, app)
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Journeycustomersession) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeycustomersession) MarshalJSON() ([]byte, error) {
    type Alias Journeycustomersession

    if JourneycustomersessionMarshalled {
        return []byte("{}"), nil
    }
    JourneycustomersessionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

