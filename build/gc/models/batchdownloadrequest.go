package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchdownloadrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchdownloadrequestDud struct { 
    


    

}

// Batchdownloadrequest
type Batchdownloadrequest struct { 
    // ConversationId - Conversation id requested
    ConversationId string `json:"conversationId"`


    // RecordingId - Recording id requested, optional.  Leave null for all recordings on the conversation
    RecordingId string `json:"recordingId"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchdownloadrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchdownloadrequest

    if BatchdownloadrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchdownloadrequestMarshalled = true

    return json.Marshal(&struct { 
        ConversationId string `json:"conversationId"`
        
        RecordingId string `json:"recordingId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

