package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentquickreplyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentquickreplyDud struct { 
    


    


    


    


    

}

// Contentquickreply - Quick reply object.
type Contentquickreply struct { 
    // Id - A unique ID assigned to the quick reply (Deprecated).
    Id string `json:"id"`


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
func (o *Contentquickreply) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentquickreply) MarshalJSON() ([]byte, error) {
    type Alias Contentquickreply

    if ContentquickreplyMarshalled {
        return []byte("{}"), nil
    }
    ContentquickreplyMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Image string `json:"image"`
        
        Action string `json:"action"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

