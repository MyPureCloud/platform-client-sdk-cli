package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunityenrollmentsresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunityenrollmentsresultDud struct { 
    


    


    

}

// Queryopportunityenrollmentsresult
type Queryopportunityenrollmentsresult struct { 
    // NextStartDate - The start date to use for the next query to retrieve additional results in ISO-8601 format. Null if there are no more results
    NextStartDate time.Time `json:"nextStartDate"`


    // Enrollments - The enrollments for the query operation
    Enrollments []Queryopportunityenrollmentresult `json:"enrollments"`


    // Opportunities - The referenced opportunities when expand=opportunities is specified
    Opportunities []Queryenrollmentopportunityresult `json:"opportunities"`

}

// String returns a JSON representation of the model
func (o *Queryopportunityenrollmentsresult) String() string {
    
     o.Enrollments = []Queryopportunityenrollmentresult{{}} 
     o.Opportunities = []Queryenrollmentopportunityresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunityenrollmentsresult) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunityenrollmentsresult

    if QueryopportunityenrollmentsresultMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunityenrollmentsresultMarshalled = true

    return json.Marshal(&struct {
        
        NextStartDate time.Time `json:"nextStartDate"`
        
        Enrollments []Queryopportunityenrollmentresult `json:"enrollments"`
        
        Opportunities []Queryenrollmentopportunityresult `json:"opportunities"`
        *Alias
    }{

        


        
        Enrollments: []Queryopportunityenrollmentresult{{}},
        


        
        Opportunities: []Queryenrollmentopportunityresult{{}},
        

        Alias: (*Alias)(u),
    })
}

