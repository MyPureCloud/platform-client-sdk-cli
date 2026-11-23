package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RfcmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RfcmetadataDud struct { 
    


    

}

// Rfcmetadata
type Rfcmetadata struct { 
    // RfcName - The reason for contact name.
    RfcName string `json:"rfcName"`


    // ConversationCount - The number of conversations.
    ConversationCount int `json:"conversationCount"`

}

// String returns a JSON representation of the model
func (o *Rfcmetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rfcmetadata) MarshalJSON() ([]byte, error) {
    type Alias Rfcmetadata

    if RfcmetadataMarshalled {
        return []byte("{}"), nil
    }
    RfcmetadataMarshalled = true

    return json.Marshal(&struct {
        
        RfcName string `json:"rfcName"`
        
        ConversationCount int `json:"conversationCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

