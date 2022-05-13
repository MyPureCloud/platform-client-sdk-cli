package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChannelDud struct { 
    


    


    

}

// Channel
type Channel struct { 
    // ConnectUri
    ConnectUri string `json:"connectUri"`


    // Id
    Id string `json:"id"`


    // Expires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Expires time.Time `json:"expires"`

}

// String returns a JSON representation of the model
func (o *Channel) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Channel) MarshalJSON() ([]byte, error) {
    type Alias Channel

    if ChannelMarshalled {
        return []byte("{}"), nil
    }
    ChannelMarshalled = true

    return json.Marshal(&struct {
        
        ConnectUri string `json:"connectUri"`
        
        Id string `json:"id"`
        
        Expires time.Time `json:"expires"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

