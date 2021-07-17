package genre

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gusti-andika/movie.go/config"
	"github.com/gusti-andika/movie.go/movie"
)

type Genre struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (g Genre) NodeId() string {
	return fmt.Sprint(g.Id)
}

func (g Genre) NodeName() string {
	return g.Name
}

func (g Genre) MoviesURL() string {
	return fmt.Sprintf("%s/discover/movie?api_key=%s&with_genres=%d", config.BASE_URL, config.API_KEY, g.Id)
}

func (g *Genre) Movies() ([]movie.Movie, error) {
	url := g.MoviesURL()
	res, err := http.Get(url)
	if err != nil {
		return nil, err

	}

	defer res.Body.Close()

	data := struct {
		Page int           `json:"page"`
		Data []movie.Movie `json:"results"`
	}{}

	//ss, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(ss))

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data.Data, nil
}

func List() ([]Genre, error) {
	url := fmt.Sprintf("%s/genre/movie/list?api_key=%s", config.BASE_URL, config.API_KEY)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data := struct {
		Genres []Genre `json:genres`
	}{}

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&data); err != nil {
		return nil, err
	}

	return data.Genres, nil
}
