package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Casequeryjobresponse
type Casequeryjobresponse struct { 
    


    // State - The state of the query job.
    State string `json:"state"`


    // DateStarted - The date the job was started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // DateFinished - The date the job finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateFinished time.Time `json:"dateFinished"`


    // VarError - The error associated with the query job when the state is Failed.
    VarError Casequeryjoberror `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Casequeryjobresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobresponse

    if CasequeryjobresponseMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        DateFinished time.Time `json:"dateFinished"`
        
        VarError Casequeryjoberror `json:"error"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

