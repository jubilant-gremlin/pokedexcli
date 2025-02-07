package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
)

func (c *Client) MonInfo (url string) (Pokemon, error) {
    // check cache first
    dat, ok := c.cache.Get(url)
    if !ok {
        // not in cache, request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return Pokemon{}, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return Pokemon{}, err
        }
        defer resp.Body.Close()

        dat, err := io.ReadAll(resp.Body)
        if err != nil {
            return Pokemon{}, err
        }

        // add to cache
        mon := Pokemon{}
        err = json.Unmarshal(dat, &mon)
        if err != nil {
            return Pokemon{}, err
        }
        return mon, nil
    }
    // in cache, continue
    mon := Pokemon{}
    err := json.Unmarshal(dat, &mon)
    if err != nil {
        return Pokemon{}, err
    }
    return mon, nil
}
