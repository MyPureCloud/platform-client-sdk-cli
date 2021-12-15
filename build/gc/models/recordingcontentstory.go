package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingcontentstoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingcontentstoryDud struct { 
    


    


    

}

// Recordingcontentstory - Story object.
type Recordingcontentstory struct { 
    // VarType - Type of ephemeral story attachment.
    VarType string `json:"type"`


    // Url - URL to the ephemeral story.
    Url string `json:"url"`


    // ReplyToId - ID of the ephemeral story being replied to.
    ReplyToId string `json:"replyToId"`

}

// String returns a JSON representation of the model
func (o *Recordingcontentstory) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingcontentstory) MarshalJSON() ([]byte, error) {
    type Alias Recordingcontentstory

    if RecordingcontentstoryMarshalled {
        return []byte("{}"), nil
    }
    RecordingcontentstoryMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Url string `json:"url"`
        
        ReplyToId string `json:"replyToId"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

