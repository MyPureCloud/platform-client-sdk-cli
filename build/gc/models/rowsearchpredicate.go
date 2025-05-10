package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RowsearchpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RowsearchpredicateDud struct { 
    


    


    


    

}

// Rowsearchpredicate
type Rowsearchpredicate struct { 
    // ColumnId - The decision table column identifier this predicate applies to
    ColumnId string `json:"columnId"`


    // Operator - The search operator for this predicate
    Operator string `json:"operator"`


    // Value - The value that will be searched for in rows. Exactly one of 'value' or 'special' can be used.
    Value string `json:"value"`


    // Special - The special value that will be searched for in rows. Exactly one of 'value' or 'special' can be used.
    Special string `json:"special"`

}

// String returns a JSON representation of the model
func (o *Rowsearchpredicate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Rowsearchpredicate) MarshalJSON() ([]byte, error) {
    type Alias Rowsearchpredicate

    if RowsearchpredicateMarshalled {
        return []byte("{}"), nil
    }
    RowsearchpredicateMarshalled = true

    return json.Marshal(&struct {
        
        ColumnId string `json:"columnId"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        Special string `json:"special"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

