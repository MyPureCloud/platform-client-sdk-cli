package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoomupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoomupdaterequestDud struct { 
    


    


    


    

}

// Roomupdaterequest
type Roomupdaterequest struct { 
    // Description - Room's description
    Description string `json:"description"`


    // Subject - Room's subject
    Subject string `json:"subject"`


    // PinnedMessageIds - Room's pinned messages
    PinnedMessageIds []string `json:"pinnedMessageIds"`


    // OwnerIds - Room's owners
    OwnerIds []string `json:"ownerIds"`

}

// String returns a JSON representation of the model
func (o *Roomupdaterequest) String() string {
    
    
     o.PinnedMessageIds = []string{""} 
     o.OwnerIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Roomupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Roomupdaterequest

    if RoomupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    RoomupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Subject string `json:"subject"`
        
        PinnedMessageIds []string `json:"pinnedMessageIds"`
        
        OwnerIds []string `json:"ownerIds"`
        *Alias
    }{

        


        


        
        PinnedMessageIds: []string{""},
        


        
        OwnerIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

