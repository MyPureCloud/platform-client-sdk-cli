package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DivisioneddomainentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DivisioneddomainentityrefDud struct { 
    


    


    

}

// Divisioneddomainentityref
type Divisioneddomainentityref struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Divisioneddomainentityref) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Divisioneddomainentityref) MarshalJSON() ([]byte, error) {
    type Alias Divisioneddomainentityref

    if DivisioneddomainentityrefMarshalled {
        return []byte("{}"), nil
    }
    DivisioneddomainentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

