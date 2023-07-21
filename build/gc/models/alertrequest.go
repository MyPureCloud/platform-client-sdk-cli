package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertrequestDud struct { 
    


    


    


    


    

}

// Alertrequest
type Alertrequest struct { 
    // VarType - The action being taken on the alert.
    VarType string `json:"type"`


    // DateStart - The start date of the mute/snooze period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - The end date of the mute/snooze period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    // Unread - The fields need for an unread update requests
    Unread Unreadfields `json:"unread"`


    // ValidRequest
    ValidRequest bool `json:"validRequest"`

}

// String returns a JSON representation of the model
func (o *Alertrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertrequest) MarshalJSON() ([]byte, error) {
    type Alias Alertrequest

    if AlertrequestMarshalled {
        return []byte("{}"), nil
    }
    AlertrequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        Unread Unreadfields `json:"unread"`
        
        ValidRequest bool `json:"validRequest"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

