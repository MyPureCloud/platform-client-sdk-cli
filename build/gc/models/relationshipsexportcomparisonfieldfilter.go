package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelationshipsexportcomparisonfieldfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelationshipsexportcomparisonfieldfilterDud struct { 
    


    

}

// Relationshipsexportcomparisonfieldfilter
type Relationshipsexportcomparisonfieldfilter struct { 
    // Field - Field name to apply the filter
    Field string `json:"field"`


    // Value - Value to compare field's value against
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Relationshipsexportcomparisonfieldfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relationshipsexportcomparisonfieldfilter) MarshalJSON() ([]byte, error) {
    type Alias Relationshipsexportcomparisonfieldfilter

    if RelationshipsexportcomparisonfieldfilterMarshalled {
        return []byte("{}"), nil
    }
    RelationshipsexportcomparisonfieldfilterMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

