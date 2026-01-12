package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentunavailabletimesvalidationjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentunavailabletimesvalidationjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Queryagentunavailabletimesvalidationjobresponse
type Queryagentunavailabletimesvalidationjobresponse struct { 
    


    // Agent - The agent who started the job
    Agent Userreference `json:"agent"`


    // Status - Status of the job
    Status string `json:"status"`


    // Result - Validation results if status == 'Complete'
    Result Unavailabletimesvalidationresult `json:"result"`


    // VarError - Error details if status == 'Error'. The error occurs if the validation job has failed
    VarError Errorbody `json:"error"`


    

}

// String returns a JSON representation of the model
func (o *Queryagentunavailabletimesvalidationjobresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentunavailabletimesvalidationjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Queryagentunavailabletimesvalidationjobresponse

    if QueryagentunavailabletimesvalidationjobresponseMarshalled {
        return []byte("{}"), nil
    }
    QueryagentunavailabletimesvalidationjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        Status string `json:"status"`
        
        Result Unavailabletimesvalidationresult `json:"result"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

