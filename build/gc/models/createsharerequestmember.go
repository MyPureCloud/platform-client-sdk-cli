package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatesharerequestmemberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatesharerequestmemberDud struct { 
    


    

}

// Createsharerequestmember
type Createsharerequestmember struct { 
    // MemberType
    MemberType string `json:"memberType"`


    // Member
    Member Memberentity `json:"member"`

}

// String returns a JSON representation of the model
func (o *Createsharerequestmember) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createsharerequestmember) MarshalJSON() ([]byte, error) {
    type Alias Createsharerequestmember

    if CreatesharerequestmemberMarshalled {
        return []byte("{}"), nil
    }
    CreatesharerequestmemberMarshalled = true

    return json.Marshal(&struct {
        
        MemberType string `json:"memberType"`
        
        Member Memberentity `json:"member"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

