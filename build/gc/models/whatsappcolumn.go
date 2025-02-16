package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappcolumnDud struct { 
    


    

}

// Whatsappcolumn
type Whatsappcolumn struct { 
    // ColumnName - The name of the whatsApp column.
    ColumnName string `json:"columnName"`


    // VarType - Indicates the type of the whatsApp column. For example, 'work' or 'personal'.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Whatsappcolumn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappcolumn) MarshalJSON() ([]byte, error) {
    type Alias Whatsappcolumn

    if WhatsappcolumnMarshalled {
        return []byte("{}"), nil
    }
    WhatsappcolumnMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

