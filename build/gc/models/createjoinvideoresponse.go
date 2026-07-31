package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatejoinvideoresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatejoinvideoresponseDud struct { 
    


    


    

}

// Createjoinvideoresponse
type Createjoinvideoresponse struct { 
    // CommunicationId - The communication id for the video or modified by the command.
    CommunicationId string `json:"communicationId"`


    // ConversationId - The conversation id for the conversation created or modified by the command.
    ConversationId string `json:"conversationId"`


    // JoinCode - The join code for the video conference. Only returned by the voice-to-video upgrade endpoint (POST /conversations/videos/{conversationId}/agentconference/communications/{communicationId}); not populated by POST /conversations/videos. Valid until the voice-to-video offer expires (default 5 minutes) or until used by a guest. One-time use.
    JoinCode string `json:"joinCode"`

}

// String returns a JSON representation of the model
func (o *Createjoinvideoresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createjoinvideoresponse) MarshalJSON() ([]byte, error) {
    type Alias Createjoinvideoresponse

    if CreatejoinvideoresponseMarshalled {
        return []byte("{}"), nil
    }
    CreatejoinvideoresponseMarshalled = true

    return json.Marshal(&struct {
        
        CommunicationId string `json:"communicationId"`
        
        ConversationId string `json:"conversationId"`
        
        JoinCode string `json:"joinCode"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

