package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingcampaigndivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingcampaigndivisionviewDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Messagingcampaigndivisionview
type Messagingcampaigndivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Messagingcampaigndivisionview) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingcampaigndivisionview) MarshalJSON() ([]byte, error) {
    type Alias Messagingcampaigndivisionview

    if MessagingcampaigndivisionviewMarshalled {
        return []byte("{}"), nil
    }
    MessagingcampaigndivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

