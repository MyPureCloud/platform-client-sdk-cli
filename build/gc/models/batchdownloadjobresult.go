package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchdownloadjobresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchdownloadjobresultDud struct { 
    


    


    


    


    

}

// Batchdownloadjobresult
type Batchdownloadjobresult struct { 
    // ConversationId - Conversation id of the result
    ConversationId string `json:"conversationId"`


    // RecordingId - Recording id of the result
    RecordingId string `json:"recordingId"`


    // ResultUrl - URL of results... HTTP GET from this location to download results for this item
    ResultUrl string `json:"resultUrl"`


    // ContentType - Content type of this result
    ContentType string `json:"contentType"`


    // ErrorMsg - An error message, in case of failed processing will indicate the cause of the failure
    ErrorMsg string `json:"errorMsg"`

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobresult) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchdownloadjobresult) MarshalJSON() ([]byte, error) {
    type Alias Batchdownloadjobresult

    if BatchdownloadjobresultMarshalled {
        return []byte("{}"), nil
    }
    BatchdownloadjobresultMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        RecordingId string `json:"recordingId"`
        
        ResultUrl string `json:"resultUrl"`
        
        ContentType string `json:"contentType"`
        
        ErrorMsg string `json:"errorMsg"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

