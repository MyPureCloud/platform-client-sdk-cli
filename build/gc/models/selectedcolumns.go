package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelectedcolumnsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelectedcolumnsDud struct { 
    


    

}

// Selectedcolumns
type Selectedcolumns struct { 
    // ColumnOrder - Indicates the order/position of the selected column
    ColumnOrder int `json:"columnOrder"`


    // ColumnName - Indicates enum name of the column from the export view
    ColumnName string `json:"columnName"`

}

// String returns a JSON representation of the model
func (o *Selectedcolumns) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selectedcolumns) MarshalJSON() ([]byte, error) {
    type Alias Selectedcolumns

    if SelectedcolumnsMarshalled {
        return []byte("{}"), nil
    }
    SelectedcolumnsMarshalled = true

    return json.Marshal(&struct {
        
        ColumnOrder int `json:"columnOrder"`
        
        ColumnName string `json:"columnName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

