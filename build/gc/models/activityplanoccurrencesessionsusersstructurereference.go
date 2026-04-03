package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencesessionsusersstructurereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencesessionsusersstructurereferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrencesessionsusersstructurereference
type Activityplanoccurrencesessionsusersstructurereference struct { 
    


    // Sessions - The sessions to delete from this activity plan occurrence
    Sessions []Activityplansessionstructurereference `json:"sessions"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencesessionsusersstructurereference) String() string {
     o.Sessions = []Activityplansessionstructurereference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencesessionsusersstructurereference) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencesessionsusersstructurereference

    if ActivityplanoccurrencesessionsusersstructurereferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencesessionsusersstructurereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Sessions []Activityplansessionstructurereference `json:"sessions"`
        *Alias
    }{

        


        
        Sessions: []Activityplansessionstructurereference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

