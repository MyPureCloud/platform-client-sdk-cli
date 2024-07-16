package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InfrastructureascodejobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InfrastructureascodejobDud struct { 
    Id string `json:"id"`


    


    AcceleratorId string `json:"acceleratorId"`


    DateSubmitted time.Time `json:"dateSubmitted"`


    SubmittedBy Userreference `json:"submittedBy"`


    Status string `json:"status"`


    ErrorInfo Errorinfo `json:"errorInfo"`


    Results string `json:"results"`


    RollbackResults string `json:"rollbackResults"`


    SelfUri string `json:"selfUri"`

}

// Infrastructureascodejob - Information about a CX infrastructure as code job
type Infrastructureascodejob struct { 
    


    // DryRun - Whether or not the job was a dry run
    DryRun bool `json:"dryRun"`


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Infrastructureascodejob) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Infrastructureascodejob) MarshalJSON() ([]byte, error) {
    type Alias Infrastructureascodejob

    if InfrastructureascodejobMarshalled {
        return []byte("{}"), nil
    }
    InfrastructureascodejobMarshalled = true

    return json.Marshal(&struct {
        
        DryRun bool `json:"dryRun"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

