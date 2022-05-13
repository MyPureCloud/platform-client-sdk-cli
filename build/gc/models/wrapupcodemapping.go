package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupcodemappingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupcodemappingDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Wrapupcodemapping
type Wrapupcodemapping struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // DefaultSet - The default set of wrap-up flags. These will be used if there is no entry for a given wrap-up code in the mapping.
    DefaultSet []string `json:"defaultSet"`


    // Mapping - A map from wrap-up code identifiers to a set of wrap-up flags.
    Mapping map[string][]string `json:"mapping"`


    

}

// String returns a JSON representation of the model
func (o *Wrapupcodemapping) String() string {
    
    
     o.DefaultSet = []string{""} 
     o.Mapping = map[string][]string{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupcodemapping) MarshalJSON() ([]byte, error) {
    type Alias Wrapupcodemapping

    if WrapupcodemappingMarshalled {
        return []byte("{}"), nil
    }
    WrapupcodemappingMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        DefaultSet []string `json:"defaultSet"`
        
        Mapping map[string][]string `json:"mapping"`
        *Alias
    }{

        


        


        


        


        


        
        DefaultSet: []string{""},
        


        
        Mapping: map[string][]string{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

