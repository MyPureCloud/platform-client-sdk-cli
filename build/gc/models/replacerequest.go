package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReplacerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReplacerequestDud struct { 
    


    


    

}

// Replacerequest
type Replacerequest struct { 
    // ChangeNumber
    ChangeNumber int `json:"changeNumber"`


    // Name
    Name string `json:"name"`


    // AuthToken
    AuthToken string `json:"authToken"`

}

// String returns a JSON representation of the model
func (o *Replacerequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Replacerequest) MarshalJSON() ([]byte, error) {
    type Alias Replacerequest

    if ReplacerequestMarshalled {
        return []byte("{}"), nil
    }
    ReplacerequestMarshalled = true

    return json.Marshal(&struct { 
        ChangeNumber int `json:"changeNumber"`
        
        Name string `json:"name"`
        
        AuthToken string `json:"authToken"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

