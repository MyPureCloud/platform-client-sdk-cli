package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ColumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ColumnDud struct { 
    


    


    


    

}

// Column
type Column struct { 
    // ColumnName - Column name. Mandatory for Fixed position/length file format.
    ColumnName string `json:"columnName"`


    // ColumnNumber - 0 based column number in delimited file format
    ColumnNumber int `json:"columnNumber"`


    // StartPosition - Zero-based position of the first column's character. Mandatory for Fixed position/length file format.
    StartPosition int `json:"startPosition"`


    // Length - Column width. Mandatory for Fixed position/length file format.
    Length int `json:"length"`

}

// String returns a JSON representation of the model
func (o *Column) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Column) MarshalJSON() ([]byte, error) {
    type Alias Column

    if ColumnMarshalled {
        return []byte("{}"), nil
    }
    ColumnMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        ColumnNumber int `json:"columnNumber"`
        
        StartPosition int `json:"startPosition"`
        
        Length int `json:"length"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

