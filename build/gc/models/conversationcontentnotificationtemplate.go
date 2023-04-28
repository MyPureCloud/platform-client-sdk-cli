package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentnotificationtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentnotificationtemplateDud struct { 
    


    


    


    


    

}

// Conversationcontentnotificationtemplate - Template notification object.
type Conversationcontentnotificationtemplate struct { 
    // Id - The identifier of the message template in 'your-namespace@your-template-id/name' format. For External vendor (e.g WhatsApp), 'your-namespace@your-template-name'. For GenesysCloud canned response message template use 'cannedresponse' as your-namespace and use response ID as your-template-id (e.g. response ID=1234 then 'cannedresponse@1234')
    Id string `json:"id"`


    // Language - Template language.
    Language string `json:"language"`


    // Header - The template header.
    Header Conversationnotificationtemplateheader `json:"header"`


    // Body - The template body.
    Body Conversationnotificationtemplatebody `json:"body"`


    // Footer - The template footer.
    Footer Conversationnotificationtemplatefooter `json:"footer"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentnotificationtemplate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentnotificationtemplate) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentnotificationtemplate

    if ConversationcontentnotificationtemplateMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentnotificationtemplateMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Language string `json:"language"`
        
        Header Conversationnotificationtemplateheader `json:"header"`
        
        Body Conversationnotificationtemplatebody `json:"body"`
        
        Footer Conversationnotificationtemplatefooter `json:"footer"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

