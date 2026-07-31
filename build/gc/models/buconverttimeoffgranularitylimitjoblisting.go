package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuconverttimeoffgranularitylimitjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuconverttimeoffgranularitylimitjoblistingDud struct { 
    

}

// Buconverttimeoffgranularitylimitjoblisting
type Buconverttimeoffgranularitylimitjoblisting struct { 
    // Entities
    Entities []Buconverttimeofflimitgranularityjobresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Buconverttimeoffgranularitylimitjoblisting) String() string {
     o.Entities = []Buconverttimeofflimitgranularityjobresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buconverttimeoffgranularitylimitjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Buconverttimeoffgranularitylimitjoblisting

    if BuconverttimeoffgranularitylimitjoblistingMarshalled {
        return []byte("{}"), nil
    }
    BuconverttimeoffgranularitylimitjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buconverttimeofflimitgranularityjobresponse `json:"entities"`
        *Alias
    }{

        
        Entities: []Buconverttimeofflimitgranularityjobresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

