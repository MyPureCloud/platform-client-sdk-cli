package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnintentDud struct { 
    


    


    

}

// Reportingturnintent
type Reportingturnintent struct { 
    // Name - The name of the intent detected during this reporting turn.
    Name string `json:"name"`


    // Confidence - The confidence score of the intent detected during this reporting turn.
    Confidence float64 `json:"confidence"`


    // Slots - The slots detected during this reporting turn.
    Slots []Reportingturnintentslot `json:"slots"`

}

// String returns a JSON representation of the model
func (o *Reportingturnintent) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Slots = []Reportingturnintentslot{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnintent) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnintent

    if ReportingturnintentMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnintentMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Confidence float64 `json:"confidence"`
        
        Slots []Reportingturnintentslot `json:"slots"`
        
        *Alias
    }{
        

        

        

        

        

        
        Slots: []Reportingturnintentslot{{}},
        

        
        Alias: (*Alias)(u),
    })
}

