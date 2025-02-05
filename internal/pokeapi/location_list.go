package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
)

// ListLocations - 
func (c *Client) ListLocations(pageURL *string) (resourceList, error) {
    url := BaseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }
// check cache first
    dat, ok := c.cache.Get(url)
    if !ok {
        // not found in cache, request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return resourceList{}, err
        }   

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return resourceList{}, err
        }
        defer resp.Body.Close()

        dat, err := io.ReadAll(resp.Body)
        if err != nil {
            return resourceList{}, err
        }

        // add to cache
        c.cache.Add(url, dat)
        locs := resourceList{}
        err = json.Unmarshal(dat, &locs)
        if err != nil {
            return resourceList{}, err
        }
        return locs, nil
    }

    locs := resourceList{}
    err := json.Unmarshal(dat, &locs)
    if err != nil {
        return resourceList{}, err
    }
    return locs, nil
}
