package requests

import (
	"encoding/json"
	"io"
	"net/http"

	e "github.com/Luigy102/south-american-qualifiers-cli/internal/extras"
	t "github.com/Luigy102/south-american-qualifiers-cli/types"
)

func PreviousMatches() t.PreviousMatches {
	url := "https://conmebol-api.vercel.app/api/results"
	res, err := http.Get(url)
	e.Check(err)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("conmebol-api not aviliable")
	}

	body, err := io.ReadAll(res.Body)
	e.Check(err)

	var previousMatches t.PreviousMatches
	err = json.Unmarshal(body, &previousMatches)
	e.Check(err)

	return previousMatches

}
