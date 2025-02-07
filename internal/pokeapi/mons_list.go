package pokeapi

import (
    "encoding/json" 
    "io"
    "net/http"
)

func (c *Client) ListMons(url string) (locationArea, error) {
    // check cache first
    dat, ok := c.cache.Get(url)
    if !ok {
        // not in cache, request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return locationArea{}, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return locationArea{}, err
        }
        defer resp.Body.Close()

        dat, err := io.ReadAll(resp.Body)
        if err != nil {
            return locationArea{}, err
        }

        c.cache.Add(url, dat)
        mons := locationArea{}
        err = json.Unmarshal(dat, &mons)
        if err != nil {
            return locationArea{}, err
        }

        return mons, nil
        }
    // in cache, do not request
    mons := locationArea{}
    err := json.Unmarshal(dat, &mons)
    if err != nil {
        return locationArea{}, err
    }

    return mons, nil
}
