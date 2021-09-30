package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrusteeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrusteeDud struct { 
    Id string `json:"id"`


    


    


    DateCreated time.Time `json:"dateCreated"`


    


    CreatedBy Orguser `json:"createdBy"`


    Organization Organization `json:"organization"`


    SelfUri string `json:"selfUri"`

}

// Trustee
type Trustee struct { 
    


    // Enabled - If disabled no trustee user will have access, even if they were previously added.
    Enabled bool `json:"enabled"`


    // UsesDefaultRole - Denotes if trustee uses admin role by default.
    UsesDefaultRole bool `json:"usesDefaultRole"`


    


    // DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Trustee) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustee) MarshalJSON() ([]byte, error) {
    type Alias Trustee

    if TrusteeMarshalled {
        return []byte("{}"), nil
    }
    TrusteeMarshalled = true

    return json.Marshal(&struct { 
        
        
        Enabled bool `json:"enabled"`
        
        UsesDefaultRole bool `json:"usesDefaultRole"`
        
        
        
        DateExpired time.Time `json:"dateExpired"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

