package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactcolumnconditionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactcolumnconditionsettingsDud struct { 
    


    


    


    

}

// Contactcolumnconditionsettings
type Contactcolumnconditionsettings struct { 
    // ColumnName - The name of the contact list column to evaluate.
    ColumnName string `json:"columnName"`


    // Operator - The operator to use when comparing values.
    Operator string `json:"operator"`


    // Value - The value to compare against the contact's data.
    Value string `json:"value"`


    // ValueType - The data type the value should be treated as.
    ValueType string `json:"valueType"`

}

// String returns a JSON representation of the model
func (o *Contactcolumnconditionsettings) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactcolumnconditionsettings) MarshalJSON() ([]byte, error) {
    type Alias Contactcolumnconditionsettings

    if ContactcolumnconditionsettingsMarshalled {
        return []byte("{}"), nil
    }
    ContactcolumnconditionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        
        ValueType string `json:"valueType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

