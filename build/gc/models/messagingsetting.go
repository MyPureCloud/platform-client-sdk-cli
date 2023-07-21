package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingsettingDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Version string `json:"version"`


    CreatedBy Domainentityref `json:"createdBy"`


    UpdatedBy Domainentityref `json:"updatedBy"`


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingsetting - Messaging setting for messaging platform integrations
type Messagingsetting struct { 
    


    // Name - The messaging Setting profile name
    Name string `json:"name"`


    


    


    


    


    


    // Content - Configuration relating to message contents
    Content Contentsetting `json:"content"`


    // Event - Configuration relating to events which may occur
    Event Eventsetting `json:"event"`


    

}

// String returns a JSON representation of the model
func (o *Messagingsetting) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsetting) MarshalJSON() ([]byte, error) {
    type Alias Messagingsetting

    if MessagingsettingMarshalled {
        return []byte("{}"), nil
    }
    MessagingsettingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Content Contentsetting `json:"content"`
        
        Event Eventsetting `json:"event"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

