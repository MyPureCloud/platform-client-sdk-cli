package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WarningMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WarningDud struct { 
    


    


    

}

// Warning
type Warning struct { 
    // Value - The value for the warning configuration.
    Value int `json:"value"`


    // RangeType - The range type for the warning configuration.
    RangeType string `json:"rangeType"`


    // Color - The color for the warning configuration in RGB hexadecimal format (for example \"#FF0000\" represents red).
    Color string `json:"color"`

}

// String returns a JSON representation of the model
func (o *Warning) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Warning) MarshalJSON() ([]byte, error) {
    type Alias Warning

    if WarningMarshalled {
        return []byte("{}"), nil
    }
    WarningMarshalled = true

    return json.Marshal(&struct {
        
        Value int `json:"value"`
        
        RangeType string `json:"rangeType"`
        
        Color string `json:"color"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

