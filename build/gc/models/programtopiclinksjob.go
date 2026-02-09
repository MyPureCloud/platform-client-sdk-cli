package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramtopiclinksjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramtopiclinksjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Programtopiclinksjob
type Programtopiclinksjob struct { 
    


    // State
    State string `json:"state"`


    // ProgramId
    ProgramId string `json:"programId"`


    // CreatedBy
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Programtopiclinksjob) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programtopiclinksjob) MarshalJSON() ([]byte, error) {
    type Alias Programtopiclinksjob

    if ProgramtopiclinksjobMarshalled {
        return []byte("{}"), nil
    }
    ProgramtopiclinksjobMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        ProgramId string `json:"programId"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

