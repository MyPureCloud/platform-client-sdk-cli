package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediapublicmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediapublicmetadataDud struct { 
    


    


    


    

}

// Opensocialmediapublicmetadata - Information about a public message.
type Opensocialmediapublicmetadata struct { 
    // RootId - The id of the root public message.
    RootId string `json:"rootId"`


    // ReplyToId - The id of the message this public message is replying to.
    ReplyToId string `json:"replyToId"`


    // Source - The source of the public message. Useful when there could be more than location. Channel specific, e.g., for Facebook it's a source page.
    Source string `json:"source"`


    // Url - The URL of the social post on the native platform.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Opensocialmediapublicmetadata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediapublicmetadata) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediapublicmetadata

    if OpensocialmediapublicmetadataMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediapublicmetadataMarshalled = true

    return json.Marshal(&struct {
        
        RootId string `json:"rootId"`
        
        ReplyToId string `json:"replyToId"`
        
        Source string `json:"source"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

