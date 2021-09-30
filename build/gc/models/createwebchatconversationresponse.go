package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatewebchatconversationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatewebchatconversationresponseDud struct { 
    


    


    


    

}

// Createwebchatconversationresponse
type Createwebchatconversationresponse struct { 
    // Id - Chat Conversation identifier
    Id string `json:"id"`


    // Jwt - The JWT that you can use to identify subsequent calls on this conversation
    Jwt string `json:"jwt"`


    // EventStreamUri - The URI which provides the conversation event stream.
    EventStreamUri string `json:"eventStreamUri"`


    // Member - Chat Member
    Member Webchatmemberinfo `json:"member"`

}

// String returns a JSON representation of the model
func (o *Createwebchatconversationresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createwebchatconversationresponse) MarshalJSON() ([]byte, error) {
    type Alias Createwebchatconversationresponse

    if CreatewebchatconversationresponseMarshalled {
        return []byte("{}"), nil
    }
    CreatewebchatconversationresponseMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Jwt string `json:"jwt"`
        
        EventStreamUri string `json:"eventStreamUri"`
        
        Member Webchatmemberinfo `json:"member"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

