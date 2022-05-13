package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodereferenceDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Activitycodereference
type Activitycodereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SecondaryPresences - The secondary presences of this activity code.
    SecondaryPresences []Secondarypresence `json:"secondaryPresences"`


    

}

// String returns a JSON representation of the model
func (o *Activitycodereference) String() string {
    
    
     o.SecondaryPresences = []Secondarypresence{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodereference) MarshalJSON() ([]byte, error) {
    type Alias Activitycodereference

    if ActivitycodereferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        SecondaryPresences []Secondarypresence `json:"secondaryPresences"`
        *Alias
    }{

        


        


        
        SecondaryPresences: []Secondarypresence{{}},
        


        

        Alias: (*Alias)(u),
    })
}

