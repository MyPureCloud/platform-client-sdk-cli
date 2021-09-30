package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientappMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientappDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    IntegrationType Integrationtype `json:"integrationType"`


    Notes string `json:"notes"`


    


    Config Clientappconfigurationinfo `json:"config"`


    ReportedState Integrationstatusinfo `json:"reportedState"`


    Attributes map[string]string `json:"attributes"`


    SelfUri string `json:"selfUri"`

}

// Clientapp - Details for a ClientApp
type Clientapp struct { 
    


    


    


    


    // IntendedState - Configured state of the integration.
    IntendedState string `json:"intendedState"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Clientapp) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientapp) MarshalJSON() ([]byte, error) {
    type Alias Clientapp

    if ClientappMarshalled {
        return []byte("{}"), nil
    }
    ClientappMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        IntendedState string `json:"intendedState"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

