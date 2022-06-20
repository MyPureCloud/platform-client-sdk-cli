package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmessagemetadatacontentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmessagemetadatacontentDud struct { 
    


    

}

// Conversationmessagemetadatacontent - Metadata information about a message content.
type Conversationmessagemetadatacontent struct { 
    // ContentType - Type of this content element.
    ContentType string `json:"contentType"`


    // SubType - Content subtype
    SubType string `json:"subType"`

}

// String returns a JSON representation of the model
func (o *Conversationmessagemetadatacontent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmessagemetadatacontent) MarshalJSON() ([]byte, error) {
    type Alias Conversationmessagemetadatacontent

    if ConversationmessagemetadatacontentMarshalled {
        return []byte("{}"), nil
    }
    ConversationmessagemetadatacontentMarshalled = true

    return json.Marshal(&struct {
        
        ContentType string `json:"contentType"`
        
        SubType string `json:"subType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

