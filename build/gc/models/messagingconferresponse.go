package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingconferresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingconferresponseDud struct { 
    


    


    


    

}

// Messagingconferresponse
type Messagingconferresponse struct { 
    // Conversation - conversation ID.
    Conversation Addressableentityref `json:"conversation"`


    // CommunicationId - The internal communication ID.
    CommunicationId string `json:"communicationId"`


    // PeerCommunicationId - The peer internal communication ID.
    PeerCommunicationId string `json:"peerCommunicationId"`


    // CommandId - Command ID for a given request.
    CommandId string `json:"commandId"`

}

// String returns a JSON representation of the model
func (o *Messagingconferresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingconferresponse) MarshalJSON() ([]byte, error) {
    type Alias Messagingconferresponse

    if MessagingconferresponseMarshalled {
        return []byte("{}"), nil
    }
    MessagingconferresponseMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Addressableentityref `json:"conversation"`
        
        CommunicationId string `json:"communicationId"`
        
        PeerCommunicationId string `json:"peerCommunicationId"`
        
        CommandId string `json:"commandId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

