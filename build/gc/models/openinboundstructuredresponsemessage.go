package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundstructuredresponsemessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundstructuredresponsemessageDud struct { 
    


    


    

}

// Openinboundstructuredresponsemessage
type Openinboundstructuredresponsemessage struct { 
    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openinboundmessagemessagingchannel `json:"channel"`


    // ButtonResponse - Button response element.
    ButtonResponse Contentbuttonresponse `json:"buttonResponse"`


    // OriginatingMessageId - Id of original structured message that this messages responds to.
    OriginatingMessageId string `json:"originatingMessageId"`

}

// String returns a JSON representation of the model
func (o *Openinboundstructuredresponsemessage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundstructuredresponsemessage) MarshalJSON() ([]byte, error) {
    type Alias Openinboundstructuredresponsemessage

    if OpeninboundstructuredresponsemessageMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundstructuredresponsemessageMarshalled = true

    return json.Marshal(&struct {
        
        Channel Openinboundmessagemessagingchannel `json:"channel"`
        
        ButtonResponse Contentbuttonresponse `json:"buttonResponse"`
        
        OriginatingMessageId string `json:"originatingMessageId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

