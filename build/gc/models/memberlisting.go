package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MemberlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MemberlistingDud struct { 
    

}

// Memberlisting
type Memberlisting struct { 
    // Entities
    Entities []Member `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Memberlisting) String() string {
    
    
     o.Entities = []Member{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Memberlisting) MarshalJSON() ([]byte, error) {
    type Alias Memberlisting

    if MemberlistingMarshalled {
        return []byte("{}"), nil
    }
    MemberlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Member `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Member{{}},
        

        
        Alias: (*Alias)(u),
    })
}

