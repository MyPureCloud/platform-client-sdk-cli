package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenrichmediasettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenrichmediasettingDud struct { 
    


    

}

// Openrichmediasetting
type Openrichmediasetting struct { 
    // QuickReply - Setting relating to quick replies
    QuickReply Outboundonlysetting `json:"quickReply"`


    // Cards - Setting relating to cards
    Cards Outboundonlysetting `json:"cards"`

}

// String returns a JSON representation of the model
func (o *Openrichmediasetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openrichmediasetting) MarshalJSON() ([]byte, error) {
    type Alias Openrichmediasetting

    if OpenrichmediasettingMarshalled {
        return []byte("{}"), nil
    }
    OpenrichmediasettingMarshalled = true

    return json.Marshal(&struct {
        
        QuickReply Outboundonlysetting `json:"quickReply"`
        
        Cards Outboundonlysetting `json:"cards"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

