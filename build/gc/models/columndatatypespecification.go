package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ColumndatatypespecificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ColumndatatypespecificationDud struct { 
    


    


    


    


    

}

// Columndatatypespecification
type Columndatatypespecification struct { 
    // ColumnName - The column name of a column selected for dynamic queueing
    ColumnName string `json:"columnName"`


    // ColumnDataType - The data type of the column selected for dynamic queueing (TEXT, NUMERIC or TIMESTAMP)
    ColumnDataType string `json:"columnDataType"`


    // Min - The minimum length of the numeric column selected for dynamic queueing
    Min int `json:"min"`


    // Max - The maximum length of the numeric column selected for dynamic queueing
    Max int `json:"max"`


    // MaxLength - The maximum length of the text column selected for dynamic queueing
    MaxLength int `json:"maxLength"`

}

// String returns a JSON representation of the model
func (o *Columndatatypespecification) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Columndatatypespecification) MarshalJSON() ([]byte, error) {
    type Alias Columndatatypespecification

    if ColumndatatypespecificationMarshalled {
        return []byte("{}"), nil
    }
    ColumndatatypespecificationMarshalled = true

    return json.Marshal(&struct {
        
        ColumnName string `json:"columnName"`
        
        ColumnDataType string `json:"columnDataType"`
        
        Min int `json:"min"`
        
        Max int `json:"max"`
        
        MaxLength int `json:"maxLength"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

