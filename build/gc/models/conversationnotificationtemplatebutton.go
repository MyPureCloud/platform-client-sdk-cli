package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplatebuttonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplatebuttonDud struct { 
    


    


    


    


    


    

}

// Conversationnotificationtemplatebutton - Template button object
type Conversationnotificationtemplatebutton struct { 
    // VarType - Specifies the type of the button.
    VarType string `json:"type"`


    // Text - Button text message.
    Text string `json:"text"`


    // Index - index of the button in the list.
    Index int `json:"index"`


    // PhoneNumber - Button phone number.
    PhoneNumber string `json:"phoneNumber"`


    // Url - Button URL link.
    Url string `json:"url"`


    // Parameters - Template parameters for placeholders in the button.
    Parameters []Conversationnotificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplatebutton) String() string {
    
    
    
    
    
     o.Parameters = []Conversationnotificationtemplateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplatebutton) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplatebutton

    if ConversationnotificationtemplatebuttonMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplatebuttonMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Index int `json:"index"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        Url string `json:"url"`
        
        Parameters []Conversationnotificationtemplateparameter `json:"parameters"`
        *Alias
    }{

        


        


        


        


        


        
        Parameters: []Conversationnotificationtemplateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

