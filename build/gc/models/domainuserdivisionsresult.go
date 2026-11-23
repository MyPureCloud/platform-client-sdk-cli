package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainuserdivisionsresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainuserdivisionsresultDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainuserdivisionsresult
type Domainuserdivisionsresult struct { 
    


    // Name
    Name string `json:"name"`


    // Permission - The permission string being queried for.
    Permission string `json:"permission"`


    // Divisions - The list of division IDs the subject has the queried permission in.
    Divisions []string `json:"divisions"`


    

}

// String returns a JSON representation of the model
func (o *Domainuserdivisionsresult) String() string {
    
    
     o.Divisions = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainuserdivisionsresult) MarshalJSON() ([]byte, error) {
    type Alias Domainuserdivisionsresult

    if DomainuserdivisionsresultMarshalled {
        return []byte("{}"), nil
    }
    DomainuserdivisionsresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Permission string `json:"permission"`
        
        Divisions []string `json:"divisions"`
        *Alias
    }{

        


        


        


        
        Divisions: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

