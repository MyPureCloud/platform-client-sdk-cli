package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnintentslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnintentslotDud struct { 
    


    


    


    

}

// Reportingturnintentslot
type Reportingturnintentslot struct { 
    // Name - The name of the slot.
    Name string `json:"name"`


    // Value - The value of the slot.
    Value string `json:"value"`


    // VarType - The NLU entity type of the slot (either builtin or user defined)
    VarType string `json:"type"`


    // Confidence - The confidence score this slot received during detection.
    Confidence float64 `json:"confidence"`

}

// String returns a JSON representation of the model
func (o *Reportingturnintentslot) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnintentslot) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnintentslot

    if ReportingturnintentslotMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnintentslotMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Value string `json:"value"`
        
        VarType string `json:"type"`
        
        Confidence float64 `json:"confidence"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

