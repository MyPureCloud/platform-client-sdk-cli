package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailcolumnDud struct { 
    


    

}

// Emailcolumn
type Emailcolumn struct { 
    // ColumnName - The name of the email column.
    ColumnName string `json:"columnName"`


    // VarType - Indicates the type of the email column. For example, 'work' or 'personal'.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Emailcolumn) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailcolumn) MarshalJSON() ([]byte, error) {
    type Alias Emailcolumn

    if EmailcolumnMarshalled {
        return []byte("{}"), nil
    }
    EmailcolumnMarshalled = true

    return json.Marshal(&struct { 
        ColumnName string `json:"columnName"`
        
        VarType string `json:"type"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

