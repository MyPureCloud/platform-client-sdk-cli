package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GuestcategoryresponselistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GuestcategoryresponselistingDud struct { 
    


    


    


    


    

}

// Guestcategoryresponselisting
type Guestcategoryresponselisting struct { 
    // Entities
    Entities []Guestresponsecategory `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SessionId
    SessionId string `json:"sessionId"`

}

// String returns a JSON representation of the model
func (o *Guestcategoryresponselisting) String() string {
     o.Entities = []Guestresponsecategory{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Guestcategoryresponselisting) MarshalJSON() ([]byte, error) {
    type Alias Guestcategoryresponselisting

    if GuestcategoryresponselistingMarshalled {
        return []byte("{}"), nil
    }
    GuestcategoryresponselistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Guestresponsecategory `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SessionId string `json:"sessionId"`
        *Alias
    }{

        
        Entities: []Guestresponsecategory{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

