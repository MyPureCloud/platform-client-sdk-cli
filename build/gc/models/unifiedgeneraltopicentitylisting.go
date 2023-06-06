package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnifiedgeneraltopicentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnifiedgeneraltopicentitylistingDud struct { 
    

}

// Unifiedgeneraltopicentitylisting
type Unifiedgeneraltopicentitylisting struct { 
    // Entities
    Entities []Unifiedgeneraltopic `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Unifiedgeneraltopicentitylisting) String() string {
     o.Entities = []Unifiedgeneraltopic{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unifiedgeneraltopicentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Unifiedgeneraltopicentitylisting

    if UnifiedgeneraltopicentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UnifiedgeneraltopicentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Unifiedgeneraltopic `json:"entities"`
        *Alias
    }{

        
        Entities: []Unifiedgeneraltopic{{}},
        

        Alias: (*Alias)(u),
    })
}

