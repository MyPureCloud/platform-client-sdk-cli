package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PosttextmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PosttextmessageDud struct { 
    


    


    

}

// Posttextmessage
type Posttextmessage struct { 
    // VarType - Message type
    VarType string `json:"type"`


    // Text - Message text. If type is structured, used as fallback for clients that do not support particular structured content
    Text string `json:"text"`


    // Content - A list of content elements in message
    Content []Messagecontent `json:"content"`

}

// String returns a JSON representation of the model
func (o *Posttextmessage) String() string {
    
    
     o.Content = []Messagecontent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Posttextmessage) MarshalJSON() ([]byte, error) {
    type Alias Posttextmessage

    if PosttextmessageMarshalled {
        return []byte("{}"), nil
    }
    PosttextmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Content []Messagecontent `json:"content"`
        *Alias
    }{

        


        


        
        Content: []Messagecontent{{}},
        

        Alias: (*Alias)(u),
    })
}

