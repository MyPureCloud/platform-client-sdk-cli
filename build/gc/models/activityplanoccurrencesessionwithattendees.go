package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencesessionwithattendeesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencesessionwithattendeesDud struct { 
    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrencesessionwithattendees
type Activityplanoccurrencesessionwithattendees struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Edited - Whether this session is edited
    Edited bool `json:"edited"`


    // Violations - The violations for this session
    Violations []Activityplansessionviolation `json:"violations"`


    // Attendees - The attendees associated with this session
    Attendees []Userreference `json:"attendees"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencesessionwithattendees) String() string {
    
    
     o.Violations = []Activityplansessionviolation{{}} 
     o.Attendees = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencesessionwithattendees) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencesessionwithattendees

    if ActivityplanoccurrencesessionwithattendeesMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencesessionwithattendeesMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Edited bool `json:"edited"`
        
        Violations []Activityplansessionviolation `json:"violations"`
        
        Attendees []Userreference `json:"attendees"`
        *Alias
    }{

        


        


        
        Violations: []Activityplansessionviolation{{}},
        


        
        Attendees: []Userreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

