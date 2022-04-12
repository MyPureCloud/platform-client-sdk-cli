package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingsettingreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingsettingreferenceDud struct { 
    


    Name string `json:"name"`


    SelfUri string `json:"selfUri"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    CreatedBy Domainentityref `json:"createdBy"`


    UpdatedBy Domainentityref `json:"updatedBy"`


    


    

}

// Messagingsettingreference - Messaging Setting for messaging platform integrations
type Messagingsettingreference struct { 
    // Id - The messaging Setting unique identifier associated with this integration
    Id string `json:"id"`


    


    


    


    


    // Version - Version number
    Version string `json:"version"`


    


    


    // Content - Settings relating to message contents
    Content Contentsetting `json:"content"`


    // Event - Settings relating to events which may occur
    Event Eventsetting `json:"event"`

}

// String returns a JSON representation of the model
func (o *Messagingsettingreference) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsettingreference) MarshalJSON() ([]byte, error) {
    type Alias Messagingsettingreference

    if MessagingsettingreferenceMarshalled {
        return []byte("{}"), nil
    }
    MessagingsettingreferenceMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        
        
        
        
        
        
        
        
        Version string `json:"version"`
        
        
        
        
        
        Content Contentsetting `json:"content"`
        
        Event Eventsetting `json:"event"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

