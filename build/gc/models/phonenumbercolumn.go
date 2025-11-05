package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PhonenumbercolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PhonenumbercolumnDud struct { 
    


    


    

}

// Phonenumbercolumn
type Phonenumbercolumn struct { 
    // ColumnName
    ColumnName string `json:"columnName"`


    // VarType
    VarType string `json:"type"`


    // CallableTimeColumnName
    CallableTimeColumnName string `json:"callableTimeColumnName"`

}

// String returns a JSON representation of the model
func (o *Phonenumbercolumn) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Phonenumbercolumn) MarshalJSON() ([]byte, error) {
    type Alias Phonenumbercolumn

    if PhonenumbercolumnMarshalled {
        return []byte("{}"), nil
    }
    PhonenumbercolumnMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        VarType string `json:"type"`
        
        CallableTimeColumnName string `json:"callableTimeColumnName"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

