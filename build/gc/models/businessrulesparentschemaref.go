package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessrulesparentschemarefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessrulesparentschemarefDud struct { 
    


    


    


    

}

// Businessrulesparentschemaref
type Businessrulesparentschemaref struct { 
    // Id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // Version - DSS schema version number used by this decision table version
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Businessrulesparentschemaref) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessrulesparentschemaref) MarshalJSON() ([]byte, error) {
    type Alias Businessrulesparentschemaref

    if BusinessrulesparentschemarefMarshalled {
        return []byte("{}"), nil
    }
    BusinessrulesparentschemarefMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SelfUri string `json:"selfUri"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

