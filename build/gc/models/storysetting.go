package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StorysettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StorysettingDud struct { 
    


    

}

// Storysetting
type Storysetting struct { 
    // Mention - Setting relating to Story Mentions
    Mention Inboundonlysetting `json:"mention"`


    // Reply - Setting relating to Story Replies
    Reply Inboundonlysetting `json:"reply"`

}

// String returns a JSON representation of the model
func (o *Storysetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Storysetting) MarshalJSON() ([]byte, error) {
    type Alias Storysetting

    if StorysettingMarshalled {
        return []byte("{}"), nil
    }
    StorysettingMarshalled = true

    return json.Marshal(&struct {
        
        Mention Inboundonlysetting `json:"mention"`
        
        Reply Inboundonlysetting `json:"reply"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

