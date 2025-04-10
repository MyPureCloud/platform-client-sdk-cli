package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EnrichfieldruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EnrichfieldruleDud struct { 
    


    


    

}

// Enrichfieldrule
type Enrichfieldrule struct { 
    // Field - A restricted JSONPath naming the specific field this combining behavior should apply to. You may use dot-notation for named fields, and array indexing for lists, but nothing more sophisticated (e.g. wildcards, sublists, filter expressions, etc). For example, to target the `firstName` field of a Contact, this should be \"$.firstName\".
    Field string `json:"field"`


    // Action - The behavior for how to combine data from the request body and the database.
    Action string `json:"action"`


    // ArrayAction - The behavior for how to combine items in array field from the request body and the database.
    ArrayAction string `json:"arrayAction"`

}

// String returns a JSON representation of the model
func (o *Enrichfieldrule) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Enrichfieldrule) MarshalJSON() ([]byte, error) {
    type Alias Enrichfieldrule

    if EnrichfieldruleMarshalled {
        return []byte("{}"), nil
    }
    EnrichfieldruleMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Action string `json:"action"`
        
        ArrayAction string `json:"arrayAction"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

