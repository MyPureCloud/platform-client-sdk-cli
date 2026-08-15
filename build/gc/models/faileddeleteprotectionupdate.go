package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaileddeleteprotectionupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaileddeleteprotectionupdateDud struct { 
    


    


    

}

// Faileddeleteprotectionupdate
type Faileddeleteprotectionupdate struct { 
    // Conversation
    Conversation Conversationreference `json:"conversation"`


    // ErrorMessage
    ErrorMessage string `json:"errorMessage"`


    // ErrorCode
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Faileddeleteprotectionupdate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faileddeleteprotectionupdate) MarshalJSON() ([]byte, error) {
    type Alias Faileddeleteprotectionupdate

    if FaileddeleteprotectionupdateMarshalled {
        return []byte("{}"), nil
    }
    FaileddeleteprotectionupdateMarshalled = true

    return json.Marshal(&struct {
        
        Conversation Conversationreference `json:"conversation"`
        
        ErrorMessage string `json:"errorMessage"`
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

