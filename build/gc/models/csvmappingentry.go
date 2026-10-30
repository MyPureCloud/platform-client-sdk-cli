package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CsvmappingentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CsvmappingentryDud struct { 
    


    


    

}

// Csvmappingentry
type Csvmappingentry struct { 
    // SourceField - CSV field to map data from
    SourceField string `json:"sourceField"`


    // TargetField - Json path to map data to
    TargetField string `json:"targetField"`


    // ValueFormat - How the CSV cell value is interpreted at the target path. When omitted, the cell is treated as a string.
    ValueFormat string `json:"valueFormat"`

}

// String returns a JSON representation of the model
func (o *Csvmappingentry) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Csvmappingentry) MarshalJSON() ([]byte, error) {
    type Alias Csvmappingentry

    if CsvmappingentryMarshalled {
        return []byte("{}"), nil
    }
    CsvmappingentryMarshalled = true

    return json.Marshal(&struct {
        
        SourceField string `json:"sourceField"`
        
        TargetField string `json:"targetField"`
        
        ValueFormat string `json:"valueFormat"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

