package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserstationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserstationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    WebRtcCallAppearances int `json:"webRtcCallAppearances"`

}

// Userstation
type Userstation struct { 
    


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // AssociatedUser
    AssociatedUser User `json:"associatedUser"`


    // AssociatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    AssociatedDate time.Time `json:"associatedDate"`


    // DefaultUser
    DefaultUser User `json:"defaultUser"`


    // ProviderInfo - Provider-specific info for this station, e.g. { \"edgeGroupId\": \"ffe7b15c-a9cc-4f4c-88f5-781327819a49\" }
    ProviderInfo map[string]string `json:"providerInfo"`


    

}

// String returns a JSON representation of the model
func (o *Userstation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ProviderInfo = map[string]string{"": ""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userstation) MarshalJSON() ([]byte, error) {
    type Alias Userstation

    if UserstationMarshalled {
        return []byte("{}"), nil
    }
    UserstationMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        AssociatedUser User `json:"associatedUser"`
        
        AssociatedDate time.Time `json:"associatedDate"`
        
        DefaultUser User `json:"defaultUser"`
        
        ProviderInfo map[string]string `json:"providerInfo"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        ProviderInfo: map[string]string{"": ""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

