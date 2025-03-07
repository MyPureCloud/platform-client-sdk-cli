package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlattenedjsonschemadocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlattenedjsonschemadocumentDud struct { 
    


    

}

// Flattenedjsonschemadocument - JSON schema that defines the transformed result that will be sent back to the caller. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned.
type Flattenedjsonschemadocument struct { 
    // Schema
    Schema Jsonschemadocument `json:"schema"`


    // ArrayProperties - Properties in the original document that were arrays
    ArrayProperties []string `json:"arrayProperties"`

}

// String returns a JSON representation of the model
func (o *Flattenedjsonschemadocument) String() string {
    
     o.ArrayProperties = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flattenedjsonschemadocument) MarshalJSON() ([]byte, error) {
    type Alias Flattenedjsonschemadocument

    if FlattenedjsonschemadocumentMarshalled {
        return []byte("{}"), nil
    }
    FlattenedjsonschemadocumentMarshalled = true

    return json.Marshal(&struct {
        
        Schema Jsonschemadocument `json:"schema"`
        
        ArrayProperties []string `json:"arrayProperties"`
        *Alias
    }{

        


        
        ArrayProperties: []string{""},
        

        Alias: (*Alias)(u),
    })
}

