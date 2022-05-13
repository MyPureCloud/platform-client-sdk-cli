package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AmazonlexrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AmazonlexrequestDud struct { 
    


    

}

// Amazonlexrequest
type Amazonlexrequest struct { 
    // RequestAttributes - AttributeName/AttributeValue pairs of User Defined Request Attributes to be sent to the amazon bot See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs
    RequestAttributes map[string]string `json:"requestAttributes"`


    // SessionAttributes - AttributeName/AttributeValue pairs of Session Attributes to be sent to the amazon bot. See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs
    SessionAttributes map[string]string `json:"sessionAttributes"`

}

// String returns a JSON representation of the model
func (o *Amazonlexrequest) String() string {
     o.RequestAttributes = map[string]string{"": ""} 
     o.SessionAttributes = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Amazonlexrequest) MarshalJSON() ([]byte, error) {
    type Alias Amazonlexrequest

    if AmazonlexrequestMarshalled {
        return []byte("{}"), nil
    }
    AmazonlexrequestMarshalled = true

    return json.Marshal(&struct {
        
        RequestAttributes map[string]string `json:"requestAttributes"`
        
        SessionAttributes map[string]string `json:"sessionAttributes"`
        *Alias
    }{

        
        RequestAttributes: map[string]string{"": ""},
        


        
        SessionAttributes: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

