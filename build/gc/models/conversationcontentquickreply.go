package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentquickreplyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentquickreplyDud struct { 
    


    


    


    


    

}

// Conversationcontentquickreply - Quick reply object.
type Conversationcontentquickreply struct { 
    // Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
    Text string `json:"text"`


    // Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
    Payload string `json:"payload"`


    // Image - URL of an image associated with the quick reply.
    Image string `json:"image"`


    // Action - Specifies the type of action that is triggered upon clicking the quick reply.
    Action string `json:"action"`


    // SummaryText - Summary of what the quick reply relates to.
    SummaryText string `json:"summaryText"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentquickreply) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentquickreply) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentquickreply

    if ConversationcontentquickreplyMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentquickreplyMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Image string `json:"image"`
        
        Action string `json:"action"`
        
        SummaryText string `json:"summaryText"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

