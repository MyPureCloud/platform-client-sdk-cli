package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DurationDud struct { 
    


    


    


    


    

}

// Duration
type Duration struct { 
    // Seconds
    Seconds int `json:"seconds"`


    // Zero
    Zero bool `json:"zero"`


    // Nano
    Nano int `json:"nano"`


    // Negative
    Negative bool `json:"negative"`


    // Units
    Units []Temporalunit `json:"units"`

}

// String returns a JSON representation of the model
func (o *Duration) String() string {
    
    
    
    
     o.Units = []Temporalunit{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Duration) MarshalJSON() ([]byte, error) {
    type Alias Duration

    if DurationMarshalled {
        return []byte("{}"), nil
    }
    DurationMarshalled = true

    return json.Marshal(&struct {
        
        Seconds int `json:"seconds"`
        
        Zero bool `json:"zero"`
        
        Nano int `json:"nano"`
        
        Negative bool `json:"negative"`
        
        Units []Temporalunit `json:"units"`
        *Alias
    }{

        


        


        


        


        
        Units: []Temporalunit{{}},
        

        Alias: (*Alias)(u),
    })
}

