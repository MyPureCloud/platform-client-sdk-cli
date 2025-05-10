package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContractualMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContractualDud struct { 
    


    

}

// Contractual
type Contractual struct { 
    // SchemaPropertyKey - The contract schema property key that describes the input value of this column.
    SchemaPropertyKey string `json:"schemaPropertyKey"`


    // Contractual - The nested contractual definition that is defined by a contract schema, if any.
    Contractual *Contractual `json:"contractual"`

}

// String returns a JSON representation of the model
func (o *Contractual) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contractual) MarshalJSON() ([]byte, error) {
    type Alias Contractual

    if ContractualMarshalled {
        return []byte("{}"), nil
    }
    ContractualMarshalled = true

    return json.Marshal(&struct {
        
        SchemaPropertyKey string `json:"schemaPropertyKey"`
        
        Contractual *Contractual `json:"contractual"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

