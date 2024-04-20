package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartmetricresultvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartmetricresultvalueDud struct { 
    Value int `json:"value"`


    GroupByAttributes []Groupbyattribute `json:"groupByAttributes"`

}

// Journeyviewchartmetricresultvalue
type Journeyviewchartmetricresultvalue struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Journeyviewchartmetricresultvalue) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartmetricresultvalue) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartmetricresultvalue

    if JourneyviewchartmetricresultvalueMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartmetricresultvalueMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

