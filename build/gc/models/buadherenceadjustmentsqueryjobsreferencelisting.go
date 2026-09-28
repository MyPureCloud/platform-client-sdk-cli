package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuadherenceadjustmentsqueryjobsreferencelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuadherenceadjustmentsqueryjobsreferencelistingDud struct { 
    

}

// Buadherenceadjustmentsqueryjobsreferencelisting
type Buadherenceadjustmentsqueryjobsreferencelisting struct { 
    // Entities
    Entities []Buadherenceadjustmentsqueryjobsreference `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Buadherenceadjustmentsqueryjobsreferencelisting) String() string {
     o.Entities = []Buadherenceadjustmentsqueryjobsreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buadherenceadjustmentsqueryjobsreferencelisting) MarshalJSON() ([]byte, error) {
    type Alias Buadherenceadjustmentsqueryjobsreferencelisting

    if BuadherenceadjustmentsqueryjobsreferencelistingMarshalled {
        return []byte("{}"), nil
    }
    BuadherenceadjustmentsqueryjobsreferencelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buadherenceadjustmentsqueryjobsreference `json:"entities"`
        *Alias
    }{

        
        Entities: []Buadherenceadjustmentsqueryjobsreference{{}},
        

        Alias: (*Alias)(u),
    })
}

