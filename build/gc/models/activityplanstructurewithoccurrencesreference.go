package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanstructurewithoccurrencesreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanstructurewithoccurrencesreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanstructurewithoccurrencesreference
type Activityplanstructurewithoccurrencesreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Occurrences - The occurrences to delete from this activity plan
    Occurrences []Activityplanoccurrencereference `json:"occurrences"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanstructurewithoccurrencesreference) String() string {
    
     o.Occurrences = []Activityplanoccurrencereference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanstructurewithoccurrencesreference) MarshalJSON() ([]byte, error) {
    type Alias Activityplanstructurewithoccurrencesreference

    if ActivityplanstructurewithoccurrencesreferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanstructurewithoccurrencesreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Occurrences []Activityplanoccurrencereference `json:"occurrences"`
        *Alias
    }{

        


        
        Occurrences: []Activityplanoccurrencereference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

