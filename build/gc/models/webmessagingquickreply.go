package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingquickreplyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingquickreplyDud struct { 
    


    


    


    

}

// Webmessagingquickreply - Quick reply object
type Webmessagingquickreply struct { 
    // Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
    Text string `json:"text"`


    // Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
    Payload string `json:"payload"`


    // Image - URL of an image associated with the quick reply.
    Image string `json:"image"`


    // Action - Specifies the type of action that is triggered upon clicking the quick reply.
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Webmessagingquickreply) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingquickreply) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingquickreply

    if WebmessagingquickreplyMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingquickreplyMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Image string `json:"image"`
        
        Action string `json:"action"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

