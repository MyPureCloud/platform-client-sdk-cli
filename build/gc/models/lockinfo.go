package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LockinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LockinfoDud struct { 
    


    


    


    

}

// Lockinfo
type Lockinfo struct { 
    // LockedBy
    LockedBy Domainentityref `json:"lockedBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateExpires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`


    // Action
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Lockinfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lockinfo) MarshalJSON() ([]byte, error) {
    type Alias Lockinfo

    if LockinfoMarshalled {
        return []byte("{}"), nil
    }
    LockinfoMarshalled = true

    return json.Marshal(&struct { 
        LockedBy Domainentityref `json:"lockedBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        Action string `json:"action"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

