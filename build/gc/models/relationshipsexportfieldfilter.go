package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportfieldfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportfieldfilterDud struct { 
    


    

}

// Relationshipsexportfieldfilter
type Relationshipsexportfieldfilter struct { 
    // Field - Field name to apply the filter
    Field string `json:"field"`


    // Value - Value to check field's value against
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Relationshipsexportfieldfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexportfieldfilter) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexportfieldfilter

    if RelationshipsexportfieldfilterMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportfieldfilterMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

