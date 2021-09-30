package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedraftinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedraftinputDud struct { 
    


    


    


    


    


    

}

// Updatedraftinput - Definition of an Action Draft to be created or updated.
type Updatedraftinput struct { 
    // Category - Category of action, Can be up to 256 characters long
    Category string `json:"category"`


    // Name - Name of action, Can be up to 256 characters long
    Name string `json:"name"`


    // Config - Configuration to support request and response processing
    Config Actionconfig `json:"config"`


    // Contract - Action contract
    Contract Actioncontractinput `json:"contract"`


    // Secure - Indication of whether or not the action is designed to accept sensitive data
    Secure bool `json:"secure"`


    // Version - Version of current Draft
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Updatedraftinput) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedraftinput) MarshalJSON() ([]byte, error) {
    type Alias Updatedraftinput

    if UpdatedraftinputMarshalled {
        return []byte("{}"), nil
    }
    UpdatedraftinputMarshalled = true

    return json.Marshal(&struct { 
        Category string `json:"category"`
        
        Name string `json:"name"`
        
        Config Actionconfig `json:"config"`
        
        Contract Actioncontractinput `json:"contract"`
        
        Secure bool `json:"secure"`
        
        Version int `json:"version"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

