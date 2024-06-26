package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditfilterDud struct { 
    


    


    


    

}

// Auditfilter
type Auditfilter struct { 
    // Name - The name of the field by which to filter.
    Name string `json:"name"`


    // VarType - The type of the filter, DATE or STRING.
    VarType string `json:"type"`


    // Operator - The operation that the filter performs.
    Operator string `json:"operator"`


    // Values - The values to make the filter comparison against.
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Auditfilter) String() string {
    
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditfilter) MarshalJSON() ([]byte, error) {
    type Alias Auditfilter

    if AuditfilterMarshalled {
        return []byte("{}"), nil
    }
    AuditfilterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Operator string `json:"operator"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        


        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

