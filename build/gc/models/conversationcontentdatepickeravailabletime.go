package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentdatepickeravailabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentdatepickeravailabletimeDud struct { 
    


    

}

// Conversationcontentdatepickeravailabletime - A datetimepicker available time that a user can take.
type Conversationcontentdatepickeravailabletime struct { 
    // Duration - The duration of the scheduling event in seconds.
    Duration int `json:"duration"`


    // DateTime - The date and time of the event being scheduled. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateTime time.Time `json:"dateTime"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentdatepickeravailabletime) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentdatepickeravailabletime) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentdatepickeravailabletime

    if ConversationcontentdatepickeravailabletimeMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentdatepickeravailabletimeMarshalled = true

    return json.Marshal(&struct {
        
        Duration int `json:"duration"`
        
        DateTime time.Time `json:"dateTime"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}
