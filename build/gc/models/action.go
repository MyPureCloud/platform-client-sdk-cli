package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Action
type Action struct { 
    


    // Name
    Name string `json:"name"`


    // IntegrationId - The ID of the integration for which this action is associated
    IntegrationId string `json:"integrationId"`


    // Category - Category of Action
    Category string `json:"category"`


    // Contract - Action contract
    Contract Actioncontract `json:"contract"`


    // Version - Version of this action
    Version int `json:"version"`


    // Secure - Indication of whether or not the action is designed to accept sensitive data
    Secure bool `json:"secure"`


    // Config - Configuration to support request and response processing
    Config Actionconfig `json:"config"`


    

}

// String returns a JSON representation of the model
func (o *Action) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Action) MarshalJSON() ([]byte, error) {
    type Alias Action

    if ActionMarshalled {
        return []byte("{}"), nil
    }
    ActionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        IntegrationId string `json:"integrationId"`
        
        Category string `json:"category"`
        
        Contract Actioncontract `json:"contract"`
        
        Version int `json:"version"`
        
        Secure bool `json:"secure"`
        
        Config Actionconfig `json:"config"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

