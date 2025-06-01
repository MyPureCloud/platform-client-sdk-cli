package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationformresponsecomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationformresponsecomponentDud struct { 
    


    

}

// Conversationformresponsecomponent - A response component from a form
type Conversationformresponsecomponent struct { 
    // Id - The id of the component in the original message.
    Id string `json:"id"`


    // Component - The content object capturing component response from the original message.
    Component Conversationformresponsecontent `json:"component"`

}

// String returns a JSON representation of the model
func (o *Conversationformresponsecomponent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationformresponsecomponent) MarshalJSON() ([]byte, error) {
    type Alias Conversationformresponsecomponent

    if ConversationformresponsecomponentMarshalled {
        return []byte("{}"), nil
    }
    ConversationformresponsecomponentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Component Conversationformresponsecontent `json:"component"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

