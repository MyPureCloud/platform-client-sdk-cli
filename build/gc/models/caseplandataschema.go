package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseplandataschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseplandataschemaDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Caseplandataschema
type Caseplandataschema struct { 
    // Id - The id of the schema.
    Id string `json:"id"`


    // Version - The version of the schema.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Caseplandataschema) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseplandataschema) MarshalJSON() ([]byte, error) {
    type Alias Caseplandataschema

    if CaseplandataschemaMarshalled {
        return []byte("{}"), nil
    }
    CaseplandataschemaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

