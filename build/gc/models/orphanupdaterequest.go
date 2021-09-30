package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrphanupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrphanupdaterequestDud struct { 
    


    


    

}

// Orphanupdaterequest
type Orphanupdaterequest struct { 
    // ArchiveDate - The orphan recording's archive date. Must be greater than 1 day from now if set. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ArchiveDate time.Time `json:"archiveDate"`


    // DeleteDate - The orphan recording's delete date. Must be greater than archiveDate if set, otherwise one day from now. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DeleteDate time.Time `json:"deleteDate"`


    // ConversationId - A conversation Id that this orphan's recording is to be attached to. If not present, the conversationId will be deduced from the recording media.
    ConversationId string `json:"conversationId"`

}

// String returns a JSON representation of the model
func (o *Orphanupdaterequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Orphanupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Orphanupdaterequest

    if OrphanupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    OrphanupdaterequestMarshalled = true

    return json.Marshal(&struct { 
        ArchiveDate time.Time `json:"archiveDate"`
        
        DeleteDate time.Time `json:"deleteDate"`
        
        ConversationId string `json:"conversationId"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

