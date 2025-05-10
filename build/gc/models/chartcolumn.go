package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChartcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChartcolumnDud struct { 
    


    

}

// Chartcolumn
type Chartcolumn struct { 
    // Id - Column Id
    Id string `json:"id"`


    // ColumnType - Type of column
    ColumnType string `json:"columnType"`

}

// String returns a JSON representation of the model
func (o *Chartcolumn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chartcolumn) MarshalJSON() ([]byte, error) {
    type Alias Chartcolumn

    if ChartcolumnMarshalled {
        return []byte("{}"), nil
    }
    ChartcolumnMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ColumnType string `json:"columnType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

