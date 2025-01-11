package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestcompletedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestcompletedataDud struct { 
    


    


    


    


    

}

// Contestcompletedata
type Contestcompletedata struct { 
    // DateEnd - End date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateEnd time.Time `json:"dateEnd"`


    // Anonymization - Anonymization of the contest
    Anonymization string `json:"anonymization"`


    // Metrics - Metrics of the contest
    Metrics []Contestdatametrics `json:"metrics"`


    // Prizes - Prizes of the contest
    Prizes []Contestdataprizes `json:"prizes"`


    // Winners - Winners of the contest
    Winners []Contestdatawinners `json:"winners"`

}

// String returns a JSON representation of the model
func (o *Contestcompletedata) String() string {
    
    
     o.Metrics = []Contestdatametrics{{}} 
     o.Prizes = []Contestdataprizes{{}} 
     o.Winners = []Contestdatawinners{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestcompletedata) MarshalJSON() ([]byte, error) {
    type Alias Contestcompletedata

    if ContestcompletedataMarshalled {
        return []byte("{}"), nil
    }
    ContestcompletedataMarshalled = true

    return json.Marshal(&struct {
        
        DateEnd time.Time `json:"dateEnd"`
        
        Anonymization string `json:"anonymization"`
        
        Metrics []Contestdatametrics `json:"metrics"`
        
        Prizes []Contestdataprizes `json:"prizes"`
        
        Winners []Contestdatawinners `json:"winners"`
        *Alias
    }{

        


        


        
        Metrics: []Contestdatametrics{{}},
        


        
        Prizes: []Contestdataprizes{{}},
        


        
        Winners: []Contestdatawinners{{}},
        

        Alias: (*Alias)(u),
    })
}

