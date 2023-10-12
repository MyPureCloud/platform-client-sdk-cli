package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemschemaDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Workitemschema
type Workitemschema struct { 
    


    // Name
    Name string `json:"name"`


    // Version - The version of the Workitem custom attribute schema.
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Workitemschema) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemschema) MarshalJSON() ([]byte, error) {
    type Alias Workitemschema

    if WorkitemschemaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemschemaMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

