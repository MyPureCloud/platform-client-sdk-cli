package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntegrationeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntegrationeventDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    CorrelationId string `json:"correlationId"`


    Timestamp time.Time `json:"timestamp"`


    Level string `json:"level"`


    EventCode string `json:"eventCode"`


    Message Messageinfo `json:"message"`


    Entities []Evententity `json:"entities"`


    ContextAttributes map[string]string `json:"contextAttributes"`


    


    User User `json:"user"`

}

// Integrationevent - Describes an event that has happened related to an integration
type Integrationevent struct { 
    


    


    


    


    


    


    


    


    


    // DetailMessage - Message with additional details about the event. (e.g. an exception cause.)
    DetailMessage Messageinfo `json:"detailMessage"`


    

}

// String returns a JSON representation of the model
func (o *Integrationevent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Integrationevent) MarshalJSON() ([]byte, error) {
    type Alias Integrationevent

    if IntegrationeventMarshalled {
        return []byte("{}"), nil
    }
    IntegrationeventMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        DetailMessage Messageinfo `json:"detailMessage"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

