package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmessageresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmessageresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Chatmessageresponse
type Chatmessageresponse struct { 
    // Id - The id of the message
    Id string `json:"id"`


    // DateCreated - Message's created time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Message's last updated time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ToJid - Jid of message's recipient (roomJid or userJid)
    ToJid string `json:"toJid"`


    // Jid - Jid of message's sender (userJid)
    Jid string `json:"jid"`


    // Body - Message's body
    Body string `json:"body"`


    // Mentions - Message's mentions
    Mentions map[string]string `json:"mentions"`


    // Edited - If message was edited
    Edited bool `json:"edited"`


    // AttachmentDeleted - If message's attachment was deleted
    AttachmentDeleted bool `json:"attachmentDeleted"`


    // FileUri - URI of file attachment
    FileUri string `json:"fileUri"`


    // Thread - The id for a thread this message corresponds to
    Thread Entity `json:"thread"`


    // User - The user who sent the message
    User Addressableentityref `json:"user"`


    // ToUser - The receiving user of the message
    ToUser Addressableentityref `json:"toUser"`


    // Reactions - The emoji reactions to this message
    Reactions []Chatreaction `json:"reactions"`

}

// String returns a JSON representation of the model
func (o *Chatmessageresponse) String() string {
    
    
    
    
    
    
     o.Mentions = map[string]string{"": ""} 
    
    
    
    
    
    
     o.Reactions = []Chatreaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmessageresponse) MarshalJSON() ([]byte, error) {
    type Alias Chatmessageresponse

    if ChatmessageresponseMarshalled {
        return []byte("{}"), nil
    }
    ChatmessageresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ToJid string `json:"toJid"`
        
        Jid string `json:"jid"`
        
        Body string `json:"body"`
        
        Mentions map[string]string `json:"mentions"`
        
        Edited bool `json:"edited"`
        
        AttachmentDeleted bool `json:"attachmentDeleted"`
        
        FileUri string `json:"fileUri"`
        
        Thread Entity `json:"thread"`
        
        User Addressableentityref `json:"user"`
        
        ToUser Addressableentityref `json:"toUser"`
        
        Reactions []Chatreaction `json:"reactions"`
        *Alias
    }{

        


        


        


        


        


        


        
        Mentions: map[string]string{"": ""},
        


        


        


        


        


        


        


        
        Reactions: []Chatreaction{{}},
        

        Alias: (*Alias)(u),
    })
}

