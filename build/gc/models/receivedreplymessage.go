package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReceivedreplymessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReceivedreplymessageDud struct { 
    


    

}

// Receivedreplymessage
type Receivedreplymessage struct { 
    // Title - Text to show in the title.
    Title string `json:"title"`


    // Subtitle - Text to show in the subtitle.
    Subtitle string `json:"subtitle"`

}

// String returns a JSON representation of the model
func (o *Receivedreplymessage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Receivedreplymessage) MarshalJSON() ([]byte, error) {
    type Alias Receivedreplymessage

    if ReceivedreplymessageMarshalled {
        return []byte("{}"), nil
    }
    ReceivedreplymessageMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

