package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommonallalertupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommonallalertupdaterequestDud struct { 
    


    

}

// Commonallalertupdaterequest
type Commonallalertupdaterequest struct { 
    // VarType - The action to take
    VarType string `json:"type"`


    // Unread - The fields need for an unread update requests
    Unread Unreadfields `json:"unread"`

}

// String returns a JSON representation of the model
func (o *Commonallalertupdaterequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commonallalertupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Commonallalertupdaterequest

    if CommonallalertupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    CommonallalertupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Unread Unreadfields `json:"unread"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

