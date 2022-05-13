package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EducationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EducationDud struct { 
    


    


    


    


    

}

// Education
type Education struct { 
    // School
    School string `json:"school"`


    // FieldOfStudy
    FieldOfStudy string `json:"fieldOfStudy"`


    // Notes - Notes about education has a 2000 character limit
    Notes string `json:"notes"`


    // DateStart - Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // DateEnd - Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`

}

// String returns a JSON representation of the model
func (o *Education) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Education) MarshalJSON() ([]byte, error) {
    type Alias Education

    if EducationMarshalled {
        return []byte("{}"), nil
    }
    EducationMarshalled = true

    return json.Marshal(&struct {
        
        School string `json:"school"`
        
        FieldOfStudy string `json:"fieldOfStudy"`
        
        Notes string `json:"notes"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

