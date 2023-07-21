package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingsettingpatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingsettingpatchrequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingsettingpatchrequest - Messaging setting for messaging platform integrations
type Messagingsettingpatchrequest struct { 
    


    // Name - The messaging Setting profile name
    Name string `json:"name"`


    // Content - Settings relating to message contents
    Content Contentsetting `json:"content"`


    // Event - Settings relating to events which may occur
    Event Eventsetting `json:"event"`


    

}

// String returns a JSON representation of the model
func (o *Messagingsettingpatchrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingsettingpatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Messagingsettingpatchrequest

    if MessagingsettingpatchrequestMarshalled {
        return []byte("{}"), nil
    }
    MessagingsettingpatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Content Contentsetting `json:"content"`
        
        Event Eventsetting `json:"event"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

