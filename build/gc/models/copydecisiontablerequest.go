package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopydecisiontablerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopydecisiontablerequestDud struct { 
    


    

}

// Copydecisiontablerequest
type Copydecisiontablerequest struct { 
    // Name - The name of the newly created decision table. Must be unique
    Name string `json:"name"`


    // Description - The description of newly created decision table.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Copydecisiontablerequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copydecisiontablerequest) MarshalJSON() ([]byte, error) {
    type Alias Copydecisiontablerequest

    if CopydecisiontablerequestMarshalled {
        return []byte("{}"), nil
    }
    CopydecisiontablerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

