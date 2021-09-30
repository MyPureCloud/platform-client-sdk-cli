package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuickreplyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuickreplyDud struct { 
    


    


    


    


    

}

// Quickreply
type Quickreply struct { 
    // Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
    Text string `json:"text"`


    // Payload - Content of the textback payload after clicking a quick reply
    Payload string `json:"payload"`


    // Url - The location of the image file associated with quick reply
    Url string `json:"url"`


    // Action - Specifies the type of action that is triggered upon clicking the quick reply. Currently, the only supported action is \"Message\" which sends a message using the quick reply text.
    Action string `json:"action"`


    // IsSelected - Indicates if the quick reply option is selected by end customer
    IsSelected bool `json:"isSelected"`

}

// String returns a JSON representation of the model
func (o *Quickreply) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Quickreply) MarshalJSON() ([]byte, error) {
    type Alias Quickreply

    if QuickreplyMarshalled {
        return []byte("{}"), nil
    }
    QuickreplyMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        Url string `json:"url"`
        
        Action string `json:"action"`
        
        IsSelected bool `json:"isSelected"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

