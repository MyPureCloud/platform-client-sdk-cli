package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagetypingeventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagetypingeventrequestDud struct { 
    


    

}

// Messagetypingeventrequest
type Messagetypingeventrequest struct { 
    // Typing - Typing event
    Typing Conversationeventtyping `json:"typing"`


    // DateSent - The time when the message typing event was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateSent time.Time `json:"dateSent"`

}

// String returns a JSON representation of the model
func (o *Messagetypingeventrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagetypingeventrequest) MarshalJSON() ([]byte, error) {
    type Alias Messagetypingeventrequest

    if MessagetypingeventrequestMarshalled {
        return []byte("{}"), nil
    }
    MessagetypingeventrequestMarshalled = true

    return json.Marshal(&struct {
        
        Typing Conversationeventtyping `json:"typing"`
        
        DateSent time.Time `json:"dateSent"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

