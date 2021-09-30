package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatesharerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatesharerequestDud struct { 
    


    


    


    


    

}

// Createsharerequest
type Createsharerequest struct { 
    // SharedEntityType - The share entity type
    SharedEntityType string `json:"sharedEntityType"`


    // SharedEntity - The entity that will be shared
    SharedEntity Sharedentity `json:"sharedEntity"`


    // MemberType
    MemberType string `json:"memberType"`


    // Member - The member that will have access to this share. Only required if a list of members is not provided.
    Member Sharedentity `json:"member"`


    // Members
    Members []Createsharerequestmember `json:"members"`

}

// String returns a JSON representation of the model
func (o *Createsharerequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Members = []Createsharerequestmember{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createsharerequest) MarshalJSON() ([]byte, error) {
    type Alias Createsharerequest

    if CreatesharerequestMarshalled {
        return []byte("{}"), nil
    }
    CreatesharerequestMarshalled = true

    return json.Marshal(&struct { 
        SharedEntityType string `json:"sharedEntityType"`
        
        SharedEntity Sharedentity `json:"sharedEntity"`
        
        MemberType string `json:"memberType"`
        
        Member Sharedentity `json:"member"`
        
        Members []Createsharerequestmember `json:"members"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Members: []Createsharerequestmember{{}},
        

        
        Alias: (*Alias)(u),
    })
}

