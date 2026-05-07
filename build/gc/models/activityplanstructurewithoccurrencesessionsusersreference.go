package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanstructurewithoccurrencesessionsusersreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanstructurewithoccurrencesessionsusersreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanstructurewithoccurrencesessionsusersreference
type Activityplanstructurewithoccurrencesessionsusersreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Occurrences - The occurrences to delete from this activity plan
    Occurrences []Activityplanoccurrencesessionsusersstructurereference `json:"occurrences"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanstructurewithoccurrencesessionsusersreference) String() string {
    
     o.Occurrences = []Activityplanoccurrencesessionsusersstructurereference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanstructurewithoccurrencesessionsusersreference) MarshalJSON() ([]byte, error) {
    type Alias Activityplanstructurewithoccurrencesessionsusersreference

    if ActivityplanstructurewithoccurrencesessionsusersreferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanstructurewithoccurrencesessionsusersreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Occurrences []Activityplanoccurrencesessionsusersstructurereference `json:"occurrences"`
        *Alias
    }{

        


        
        Occurrences: []Activityplanoccurrencesessionsusersstructurereference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

