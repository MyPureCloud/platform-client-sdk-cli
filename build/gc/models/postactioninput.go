package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PostactioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PostactioninputDud struct { 
    


    


    


    


    


    

}

// Postactioninput - Definition of an Action to be created or updated.
type Postactioninput struct { 
    // Category - Category of action, Can be up to 256 characters long
    Category string `json:"category"`


    // Name - Name of action, Can be up to 256 characters long
    Name string `json:"name"`


    // IntegrationId - The ID of the integration this action is associated to
    IntegrationId string `json:"integrationId"`


    // Config - Configuration to support request and response processing
    Config Actionconfig `json:"config"`


    // Contract - Action contract
    Contract Actioncontractinput `json:"contract"`


    // Secure - Indication of whether or not the action is designed to accept sensitive data
    Secure bool `json:"secure"`

}

// String returns a JSON representation of the model
func (o *Postactioninput) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Postactioninput) MarshalJSON() ([]byte, error) {
    type Alias Postactioninput

    if PostactioninputMarshalled {
        return []byte("{}"), nil
    }
    PostactioninputMarshalled = true

    return json.Marshal(&struct {
        
        Category string `json:"category"`
        
        Name string `json:"name"`
        
        IntegrationId string `json:"integrationId"`
        
        Config Actionconfig `json:"config"`
        
        Contract Actioncontractinput `json:"contract"`
        
        Secure bool `json:"secure"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

