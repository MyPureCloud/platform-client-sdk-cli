package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatebenefitassessmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatebenefitassessmentrequestDud struct { 
    

}

// Createbenefitassessmentrequest
type Createbenefitassessmentrequest struct { 
    // QueueIds - The list of queue ids that are to be assessed for Predictive Routing benefit.
    QueueIds []string `json:"queueIds"`

}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentrequest) String() string {
     o.QueueIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createbenefitassessmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Createbenefitassessmentrequest

    if CreatebenefitassessmentrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatebenefitassessmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        QueueIds []string `json:"queueIds"`
        *Alias
    }{

        
        QueueIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

