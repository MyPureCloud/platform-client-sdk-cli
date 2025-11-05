package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrackingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrackingsettingsDud struct { 
    


    


    


    

}

// Trackingsettings - Settings that control how tracking data is collected and filtered.
type Trackingsettings struct { 
    // ExcludedQueryParameters - List of parameters to be excluded from the query string.
    ExcludedQueryParameters []string `json:"excludedQueryParameters"`


    // ShouldKeepUrlFragment - Whether or not to keep the URL fragment.
    ShouldKeepUrlFragment bool `json:"shouldKeepUrlFragment"`


    // SearchQueryParameters - List of query parameters used for search (e.g. 'query').
    SearchQueryParameters []string `json:"searchQueryParameters"`


    // IpFilters - IP address filtering configuration for tracking restrictions.
    IpFilters []Ipfilter `json:"ipFilters"`

}

// String returns a JSON representation of the model
func (o *Trackingsettings) String() string {
     o.ExcludedQueryParameters = []string{""} 
    
     o.SearchQueryParameters = []string{""} 
     o.IpFilters = []Ipfilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trackingsettings) MarshalJSON() ([]byte, error) {
    type Alias Trackingsettings

    if TrackingsettingsMarshalled {
        return []byte("{}"), nil
    }
    TrackingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ExcludedQueryParameters []string `json:"excludedQueryParameters"`
        
        ShouldKeepUrlFragment bool `json:"shouldKeepUrlFragment"`
        
        SearchQueryParameters []string `json:"searchQueryParameters"`
        
        IpFilters []Ipfilter `json:"ipFilters"`
        *Alias
    }{

        
        ExcludedQueryParameters: []string{""},
        


        


        
        SearchQueryParameters: []string{""},
        


        
        IpFilters: []Ipfilter{{}},
        

        Alias: (*Alias)(u),
    })
}

