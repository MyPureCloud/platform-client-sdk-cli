package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalpagemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalpagemetadataDud struct { 
    


    

}

// Externalpagemetadata
type Externalpagemetadata struct { 
    // Name - The name of the external page
    Name string `json:"name"`


    // ProfilePictureUrl - The profile picture URL of the external page
    ProfilePictureUrl string `json:"profilePictureUrl"`

}

// String returns a JSON representation of the model
func (o *Externalpagemetadata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalpagemetadata) MarshalJSON() ([]byte, error) {
    type Alias Externalpagemetadata

    if ExternalpagemetadataMarshalled {
        return []byte("{}"), nil
    }
    ExternalpagemetadataMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ProfilePictureUrl string `json:"profilePictureUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

