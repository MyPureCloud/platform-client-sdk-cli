package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonecolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonecolumnDud struct { 
    


    

}

// Phonecolumn
type Phonecolumn struct { 
    // ColumnName - The name of the phone column.
    ColumnName string `json:"columnName"`


    // VarType - The type of the phone column. For example, 'cell' or 'home'.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Phonecolumn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonecolumn) MarshalJSON() ([]byte, error) {
    type Alias Phonecolumn

    if PhonecolumnMarshalled {
        return []byte("{}"), nil
    }
    PhonecolumnMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

