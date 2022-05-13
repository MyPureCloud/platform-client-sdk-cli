package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ArchitectjobmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ArchitectjobmessageDud struct { 
    


    


    

}

// Architectjobmessage
type Architectjobmessage struct { 
    // DateTime - The DateTime when the message was generated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateTime time.Time `json:"dateTime"`


    // VarType - The message type.
    VarType string `json:"type"`


    // Text - The text of the message.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Architectjobmessage) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Architectjobmessage) MarshalJSON() ([]byte, error) {
    type Alias Architectjobmessage

    if ArchitectjobmessageMarshalled {
        return []byte("{}"), nil
    }
    ArchitectjobmessageMarshalled = true

    return json.Marshal(&struct {
        
        DateTime time.Time `json:"dateTime"`
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

