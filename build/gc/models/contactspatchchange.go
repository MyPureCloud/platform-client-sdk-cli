package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactspatchchangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactspatchchangeDud struct { 
    


    


    

}

// Contactspatchchange
type Contactspatchchange struct { 
    // Field - A JSONPath string, whose syntax is a strict subset of the JSONPath RFC 9535.  The root of the field string must be \"$.\" indicating a path from the root of the entity. You may only use dot-notation to access named fields. Examples: To select the `firstName` field of a Contact, use: \"$.firstName\".To access object fields, use the top level object field name: \"$.address\". To access nested field names, use the nested field name: \"$.address.city\". Note: trying to patch both nested fields and their parent field is not allowed and will result in a 409 error response.
    Field string `json:"field"`


    // Value - The value which is applied to the selected field for the patch. Acceptable types are String, Integer, Boolean, Array, Map.
    Value interface{} `json:"value"`


    // Action - The action of the operation.UpdateIfEmpty: Update if and only if the current value is emptyUpdate: Update the field unconditionally.UpdateIfExists: Update the field if and only if the existing field is not empty.AppendToCollection: Add new items to a collection, preserving existing data.Remove: Remove the current value unconditionally.RemoveFromCollection: Remove specified value from a collection.
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Contactspatchchange) String() string {
    
     o.Value = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactspatchchange) MarshalJSON() ([]byte, error) {
    type Alias Contactspatchchange

    if ContactspatchchangeMarshalled {
        return []byte("{}"), nil
    }
    ContactspatchchangeMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Value interface{} `json:"value"`
        
        Action string `json:"action"`
        *Alias
    }{

        


        
        Value: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

