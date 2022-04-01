package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateresponseassetrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateresponseassetrequestDud struct { 
    


    


    

}

// Createresponseassetrequest
type Createresponseassetrequest struct { 
    // Name - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
    Name string `json:"name"`


    // DivisionId - Division to associate to this asset. Can only be used with this division.
    DivisionId string `json:"divisionId"`


    // ContentMd5 - Content MD-5 of the file to upload
    ContentMd5 string `json:"contentMd5"`

}

// String returns a JSON representation of the model
func (o *Createresponseassetrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createresponseassetrequest) MarshalJSON() ([]byte, error) {
    type Alias Createresponseassetrequest

    if CreateresponseassetrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateresponseassetrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        DivisionId string `json:"divisionId"`
        
        ContentMd5 string `json:"contentMd5"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

