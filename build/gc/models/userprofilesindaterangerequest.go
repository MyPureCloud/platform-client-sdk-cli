package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserprofilesindaterangerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserprofilesindaterangerequestDud struct { 
    


    

}

// Userprofilesindaterangerequest
type Userprofilesindaterangerequest struct { 
    // StartWorkday - Start work day in ISO-8601 format used in the date range.
    StartWorkday string `json:"startWorkday"`


    // EndWorkday - End work day in ISO-8601 format used in the date range.
    EndWorkday string `json:"endWorkday"`

}

// String returns a JSON representation of the model
func (o *Userprofilesindaterangerequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userprofilesindaterangerequest) MarshalJSON() ([]byte, error) {
    type Alias Userprofilesindaterangerequest

    if UserprofilesindaterangerequestMarshalled {
        return []byte("{}"), nil
    }
    UserprofilesindaterangerequestMarshalled = true

    return json.Marshal(&struct {
        
        StartWorkday string `json:"startWorkday"`
        
        EndWorkday string `json:"endWorkday"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

