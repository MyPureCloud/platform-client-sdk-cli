package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateoutboundmessagingconversationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateoutboundmessagingconversationresponseDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Createoutboundmessagingconversationresponse
type Createoutboundmessagingconversationresponse struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Createoutboundmessagingconversationresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createoutboundmessagingconversationresponse) MarshalJSON() ([]byte, error) {
    type Alias Createoutboundmessagingconversationresponse

    if CreateoutboundmessagingconversationresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateoutboundmessagingconversationresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

