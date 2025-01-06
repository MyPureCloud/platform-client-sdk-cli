package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatpresenceDud struct { 
    


    


    


    


    


    

}

// Chatpresence
type Chatpresence struct { 
    // Source
    Source string `json:"source"`


    // OrganizationPresence
    OrganizationPresence Organizationpresence `json:"organizationPresence"`


    // SystemPresence
    SystemPresence string `json:"systemPresence"`


    // Message
    Message string `json:"message"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // OutOfOffice
    OutOfOffice Outofoffice `json:"outOfOffice"`

}

// String returns a JSON representation of the model
func (o *Chatpresence) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatpresence) MarshalJSON() ([]byte, error) {
    type Alias Chatpresence

    if ChatpresenceMarshalled {
        return []byte("{}"), nil
    }
    ChatpresenceMarshalled = true

    return json.Marshal(&struct {
        
        Source string `json:"source"`
        
        OrganizationPresence Organizationpresence `json:"organizationPresence"`
        
        SystemPresence string `json:"systemPresence"`
        
        Message string `json:"message"`
        
        DateModified time.Time `json:"dateModified"`
        
        OutOfOffice Outofoffice `json:"outOfOffice"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

