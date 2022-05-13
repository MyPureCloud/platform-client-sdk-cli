package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcknowledgescreenrecordingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcknowledgescreenrecordingrequestDud struct { 
    


    


    

}

// Acknowledgescreenrecordingrequest
type Acknowledgescreenrecordingrequest struct { 
    // ParticipantJid
    ParticipantJid string `json:"participantJid"`


    // RoomId
    RoomId string `json:"roomId"`


    // ConversationId
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Acknowledgescreenrecordingrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acknowledgescreenrecordingrequest) MarshalJSON() ([]byte, error) {
    type Alias Acknowledgescreenrecordingrequest

    if AcknowledgescreenrecordingrequestMarshalled {
        return []byte("{}"), nil
    }
    AcknowledgescreenrecordingrequestMarshalled = true

    return json.Marshal(&struct {
        
        ParticipantJid string `json:"participantJid"`
        
        RoomId string `json:"roomId"`
        
        ConversationId string `json:"conversationId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

