package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationbuttoncomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationbuttoncomponentDud struct { 
    


    

}

// Conversationbuttoncomponent - Structured template button object.
type Conversationbuttoncomponent struct { 
    // Title - Text to show inside the button.
    Title string `json:"title"`


    // Actions - The button actions.
    Actions Conversationcontentactions `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Conversationbuttoncomponent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationbuttoncomponent) MarshalJSON() ([]byte, error) {
    type Alias Conversationbuttoncomponent

    if ConversationbuttoncomponentMarshalled {
        return []byte("{}"), nil
    }
    ConversationbuttoncomponentMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Actions Conversationcontentactions `json:"actions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

