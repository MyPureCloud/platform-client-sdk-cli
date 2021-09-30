package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletopicentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletopicentitylistingDud struct { 
    

}

// Availabletopicentitylisting
type Availabletopicentitylisting struct { 
    // Entities
    Entities []Availabletopic `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Availabletopicentitylisting) String() string {
    
    
     o.Entities = []Availabletopic{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletopicentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Availabletopicentitylisting

    if AvailabletopicentitylistingMarshalled {
        return []byte("{}"), nil
    }
    AvailabletopicentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Availabletopic `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Availabletopic{{}},
        

        
        Alias: (*Alias)(u),
    })
}

