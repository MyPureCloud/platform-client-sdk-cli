package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DivisionentityrefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DivisionentityrefDud struct { 
    


    


    


    

}

// Divisionentityref
type Divisionentityref struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // DateDivisionUpdated - The time the entity division was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDivisionUpdated time.Time `json:"dateDivisionUpdated"`

}

// String returns a JSON representation of the model
func (o *Divisionentityref) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Divisionentityref) MarshalJSON() ([]byte, error) {
    type Alias Divisionentityref

    if DivisionentityrefMarshalled {
        return []byte("{}"), nil
    }
    DivisionentityrefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        DateDivisionUpdated time.Time `json:"dateDivisionUpdated"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

