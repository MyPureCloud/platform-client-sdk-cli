package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnotificationtemplateheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnotificationtemplateheaderDud struct { 
    


    


    


    

}

// Conversationnotificationtemplateheader - Template header object.
type Conversationnotificationtemplateheader struct { 
    // VarType - Template header type.
    VarType string `json:"type"`


    // Text - Header text. For WhatsApp, ignored.
    Text string `json:"text"`


    // Media - Media template header image.
    Media Conversationcontentattachment `json:"media"`


    // Parameters - Template parameters for placeholders in template.
    Parameters []Conversationnotificationtemplateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Conversationnotificationtemplateheader) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Parameters = []Conversationnotificationtemplateparameter{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnotificationtemplateheader) MarshalJSON() ([]byte, error) {
    type Alias Conversationnotificationtemplateheader

    if ConversationnotificationtemplateheaderMarshalled {
        return []byte("{}"), nil
    }
    ConversationnotificationtemplateheaderMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Media Conversationcontentattachment `json:"media"`
        
        Parameters []Conversationnotificationtemplateparameter `json:"parameters"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Parameters: []Conversationnotificationtemplateparameter{{}},
        

        
        Alias: (*Alias)(u),
    })
}

