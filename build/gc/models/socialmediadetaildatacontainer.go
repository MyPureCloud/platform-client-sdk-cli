package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediadetaildatacontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediadetaildatacontainerDud struct { 
    


    

}

// Socialmediadetaildatacontainer
type Socialmediadetaildatacontainer struct { 
    // Interval
    Interval string `json:"interval"`


    // Messages
    Messages []Socialmediadetailmessagecontainer `json:"messages"`

}

// String returns a JSON representation of the model
func (o *Socialmediadetaildatacontainer) String() string {
    
     o.Messages = []Socialmediadetailmessagecontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediadetaildatacontainer) MarshalJSON() ([]byte, error) {
    type Alias Socialmediadetaildatacontainer

    if SocialmediadetaildatacontainerMarshalled {
        return []byte("{}"), nil
    }
    SocialmediadetaildatacontainerMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Messages []Socialmediadetailmessagecontainer `json:"messages"`
        *Alias
    }{

        


        
        Messages: []Socialmediadetailmessagecontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

