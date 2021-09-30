package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdaymetriclistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdaymetriclistingDud struct { 
    

}

// Workdaymetriclisting
type Workdaymetriclisting struct { 
    // Entities
    Entities []Workdaymetric `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Workdaymetriclisting) String() string {
    
    
     o.Entities = []Workdaymetric{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdaymetriclisting) MarshalJSON() ([]byte, error) {
    type Alias Workdaymetriclisting

    if WorkdaymetriclistingMarshalled {
        return []byte("{}"), nil
    }
    WorkdaymetriclistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Workdaymetric `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Workdaymetric{{}},
        

        
        Alias: (*Alias)(u),
    })
}

