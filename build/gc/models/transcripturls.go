package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscripturlsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscripturlsDud struct { 
    


    


    

}

// Transcripturls
type Transcripturls struct { 
    // Conversation - The Conversation Reference
    Conversation Addressableentityref `json:"conversation"`


    // CommunicationId - The Communication ID
    CommunicationId string `json:"communicationId"`


    // Urls - List of Transcript URLs
    Urls []Segmenturl `json:"urls"`

}

// String returns a JSON representation of the model
func (o *Transcripturls) String() string {
    
    
     o.Urls = []Segmenturl{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripturls) MarshalJSON() ([]byte, error) {
    type Alias Transcripturls

    if TranscripturlsMarshalled {
        return []byte("{}"), nil
    }
    TranscripturlsMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Addressableentityref `json:"conversation"`
        
        CommunicationId string `json:"communicationId"`
        
        Urls []Segmenturl `json:"urls"`
        *Alias
    }{

        


        


        
        Urls: []Segmenturl{{}},
        

        Alias: (*Alias)(u),
    })
}

