package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButtonresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButtonresponseDud struct { 
    


    


    


    

}

// Buttonresponse
type Buttonresponse struct { 
    // VarType - Button response type that captures Button and QuickReply type responses
    VarType string `json:"type"`


    // Text - Text to show inside the Button reply. This is also used as the response text after clicking on the Button.
    Text string `json:"text"`


    // Payload - Content of the textback payload after clicking a button
    Payload string `json:"payload"`


    // MessageType - Button response message type that captures QuickReply , Cards and Carousel .This is used  as label for Card selection
    MessageType string `json:"messageType"`

}

// String returns a JSON representation of the model
func (o *Buttonresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buttonresponse) MarshalJSON() ([]byte, error) {
    type Alias Buttonresponse

    if ButtonresponseMarshalled {
        return []byte("{}"), nil
    }
    ButtonresponseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Payload string `json:"payload"`
        
        MessageType string `json:"messageType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

