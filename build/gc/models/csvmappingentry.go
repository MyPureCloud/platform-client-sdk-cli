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
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

