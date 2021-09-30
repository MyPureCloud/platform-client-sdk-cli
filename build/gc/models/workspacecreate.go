package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkspacecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkspacecreateDud struct { 
    


    


    

}

// Workspacecreate
type Workspacecreate struct { 
    // Name - The workspace name
    Name string `json:"name"`


    // Bucket
    Bucket string `json:"bucket"`


    // Description
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Workspacecreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workspacecreate) MarshalJSON() ([]byte, error) {
    type Alias Workspacecreate

    if WorkspacecreateMarshalled {
        return []byte("{}"), nil
    }
    WorkspacecreateMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Bucket string `json:"bucket"`
        
        Description string `json:"description"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

