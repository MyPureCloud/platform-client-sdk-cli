package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MembergroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MembergroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Membergroup
type Membergroup struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // VarType - The type of group, e.g. TEAM, etc.
    VarType string `json:"type"`


    // MemberCount - The number of members in this group
    MemberCount int `json:"memberCount"`


    

}

// String returns a JSON representation of the model
func (o *Membergroup) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Membergroup) MarshalJSON() ([]byte, error) {
    type Alias Membergroup

    if MembergroupMarshalled {
        return []byte("{}"), nil
    }
    MembergroupMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        VarType string `json:"type"`
        
        MemberCount int `json:"memberCount"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

