package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagefooterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagefooterDud struct { 
    


    

}

// Messagefooter
type Messagefooter struct { 
    // VarType - Defines the content type of the footer in message
    VarType string `json:"type"`


    // Content - Content associated with the footer in the message
    Content string `json:"content"`

}

// String returns a JSON representation of the model
func (o *Messagefooter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagefooter) MarshalJSON() ([]byte, error) {
    type Alias Messagefooter

    if MessagefooterMarshalled {
        return []byte("{}"), nil
    }
    MessagefooterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Content string `json:"content"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

