package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SecurityprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SecurityprofileDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Securityprofile
type Securityprofile struct { 
    


    // Name
    Name string `json:"name"`


    // Permissions
    Permissions []string `json:"permissions"`


    

}

// String returns a JSON representation of the model
func (o *Securityprofile) String() string {
    
    
    
    
    
    
    
    
     o.Permissions = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Securityprofile) MarshalJSON() ([]byte, error) {
    type Alias Securityprofile

    if SecurityprofileMarshalled {
        return []byte("{}"), nil
    }
    SecurityprofileMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Permissions []string `json:"permissions"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Permissions: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

