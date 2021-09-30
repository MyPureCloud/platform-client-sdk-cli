package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactphonenumbercolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactphonenumbercolumnDud struct { 
    


    


    

}

// Contactphonenumbercolumn
type Contactphonenumbercolumn struct { 
    // ColumnName - The name of the phone column.
    ColumnName string `json:"columnName"`


    // VarType - Indicates the type of the phone column. For example, 'cell' or 'home'.
    VarType string `json:"type"`


    // CallableTimeColumn - A column that indicates the timezone to use for a given contact when checking callable times. Not allowed if 'automaticTimeZoneMapping' is set to true.
    CallableTimeColumn string `json:"callableTimeColumn"`

}

// String returns a JSON representation of the model
func (o *Contactphonenumbercolumn) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactphonenumbercolumn) MarshalJSON() ([]byte, error) {
    type Alias Contactphonenumbercolumn

    if ContactphonenumbercolumnMarshalled {
        return []byte("{}"), nil
    }
    ContactphonenumbercolumnMarshalled = true

    return json.Marshal(&struct { 
        ColumnName string `json:"columnName"`
        
        VarType string `json:"type"`
        
        CallableTimeColumn string `json:"callableTimeColumn"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

