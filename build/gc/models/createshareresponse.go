package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateshareresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateshareresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Createshareresponse
type Createshareresponse struct { 
    


    // Name
    Name string `json:"name"`


    // SharedEntityType
    SharedEntityType string `json:"sharedEntityType"`


    // SharedEntity
    SharedEntity Domainentityref `json:"sharedEntity"`


    // MemberType
    MemberType string `json:"memberType"`


    // Member
    Member Domainentityref `json:"member"`


    // SharedBy
    SharedBy Domainentityref `json:"sharedBy"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // Succeeded
    Succeeded []Share `json:"succeeded"`


    // Failed
    Failed []Share `json:"failed"`


    

}

// String returns a JSON representation of the model
func (o *Createshareresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Succeeded = []Share{{}} 
    
    
    
     o.Failed = []Share{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createshareresponse) MarshalJSON() ([]byte, error) {
    type Alias Createshareresponse

    if CreateshareresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateshareresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        SharedEntityType string `json:"sharedEntityType"`
        
        SharedEntity Domainentityref `json:"sharedEntity"`
        
        MemberType string `json:"memberType"`
        
        Member Domainentityref `json:"member"`
        
        SharedBy Domainentityref `json:"sharedBy"`
        
        Workspace Domainentityref `json:"workspace"`
        
        Succeeded []Share `json:"succeeded"`
        
        Failed []Share `json:"failed"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Succeeded: []Share{{}},
        

        

        
        Failed: []Share{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

