package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChanneltopicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChanneltopicDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Channeltopic
type Channeltopic struct { 
    // Id
    Id string `json:"id"`


    // State
    State string `json:"state"`


    // RejectionReason
    RejectionReason string `json:"rejectionReason"`


    

}

// String returns a JSON representation of the model
func (o *Channeltopic) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Channeltopic) MarshalJSON() ([]byte, error) {
    type Alias Channeltopic

    if ChanneltopicMarshalled {
        return []byte("{}"), nil
    }
    ChanneltopicMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        State string `json:"state"`
        
        RejectionReason string `json:"rejectionReason"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

