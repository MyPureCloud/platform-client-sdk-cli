package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListingDud struct { 
    


    


    


    

}

// Listing
type Listing struct { 
    // Entities
    Entities []Csvsettings `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Listing) String() string {
     o.Entities = []Csvsettings{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listing) MarshalJSON() ([]byte, error) {
    type Alias Listing

    if ListingMarshalled {
        return []byte("{}"), nil
    }
    ListingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Csvsettings `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Csvsettings{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

