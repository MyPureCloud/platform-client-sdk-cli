package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QualityauditqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QualityauditqueryfilterDud struct { 
    


    

}

// Qualityauditqueryfilter
type Qualityauditqueryfilter struct { 
    // Property - Name of the property to filter.
    Property string `json:"property"`


    // Value - Value of the property to filter.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Qualityauditqueryfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qualityauditqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Qualityauditqueryfilter

    if QualityauditqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    QualityauditqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        Property string `json:"property"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

