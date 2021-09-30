package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditquerysortMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditquerysortDud struct { 
    


    

}

// Auditquerysort
type Auditquerysort struct { 
    // Name - Name of the property to sort.
    Name string `json:"name"`


    // SortOrder - Sort Order
    SortOrder string `json:"sortOrder"`

}

// String returns a JSON representation of the model
func (o *Auditquerysort) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditquerysort) MarshalJSON() ([]byte, error) {
    type Alias Auditquerysort

    if AuditquerysortMarshalled {
        return []byte("{}"), nil
    }
    AuditquerysortMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        SortOrder string `json:"sortOrder"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

