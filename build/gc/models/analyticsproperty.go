package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticspropertyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticspropertyDud struct { 
    


    


    

}

// Analyticsproperty
type Analyticsproperty struct { 
    // Property - User-defined rather than intrinsic system-observed values. These are tagged onto segments by other components within PureCloud or by API users directly.  This is the name of the user-defined property.
    Property string `json:"property"`


    // PropertyType - Indicates what the data type is (e.g. integer vs string) and therefore how to evaluate what would constitute a match
    PropertyType string `json:"propertyType"`


    // Value - What property value to match against
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Analyticsproperty) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsproperty) MarshalJSON() ([]byte, error) {
    type Alias Analyticsproperty

    if AnalyticspropertyMarshalled {
        return []byte("{}"), nil
    }
    AnalyticspropertyMarshalled = true

    return json.Marshal(&struct {
        
        Property string `json:"property"`
        
        PropertyType string `json:"propertyType"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

