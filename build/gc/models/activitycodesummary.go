package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodesummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodesummaryDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Activitycodesummary
type Activitycodesummary struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // SecondaryPresences - The secondary presences of this activity code.
    SecondaryPresences []Secondarypresence `json:"secondaryPresences"`


    

}

// String returns a JSON representation of the model
func (o *Activitycodesummary) String() string {
    
    
     o.SecondaryPresences = []Secondarypresence{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodesummary) MarshalJSON() ([]byte, error) {
    type Alias Activitycodesummary

    if ActivitycodesummaryMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodesummaryMarshalled = true

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

