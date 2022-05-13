package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkdaypointstrendMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkdaypointstrendDud struct { 
    DateStartWorkday time.Time `json:"dateStartWorkday"`


    DateEndWorkday time.Time `json:"dateEndWorkday"`


    User Userreference `json:"user"`


    DayOfWeek string `json:"dayOfWeek"`


    AveragePoints float64 `json:"averagePoints"`


    Trend []Workdaypointstrenditem `json:"trend"`

}

// Workdaypointstrend
type Workdaypointstrend struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Workdaypointstrend) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workdaypointstrend) MarshalJSON() ([]byte, error) {
    type Alias Workdaypointstrend

    if WorkdaypointstrendMarshalled {
        return []byte("{}"), nil
    }
    WorkdaypointstrendMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

