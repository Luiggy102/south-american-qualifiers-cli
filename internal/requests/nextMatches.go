package requests

import (
	"encoding/json"
	"io"
	"net/http"

	e "github.com/Luigy102/south-american-qualifiers-cli/internal/extras"
	t "github.com/Luigy102/south-american-qualifiers-cli/types"
)

func nextMatches() t.NextMatches {
	url := "https://conmebol-api.vercel.app/api/matches"
	res, err := http.Get(url)
	e.Check(err)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("conmebol-api not aviliable")
	}

	body, err := io.ReadAll(res.Body)
	e.Check(err)

	var nextMatches t.NextMatches
	err = json.Unmarshal(body, &nextMatches)
	e.Check(err)

	return nextMatches
}
