package semv

import (
	"strings"
        "fmt"
        "net/http"
	"github.com/blang/semver"
        "encoding/json"
        "io/ioutil"
)

// TagCmd for tag list
var defaultVersion = "0.0.0"

// List struct
type List struct {
	data semver.Versions
}

// GetList returns List
func GetList(url string) (*List, error) {
	list, err := getVersions(url)
	if err != nil {
		return nil, err
	}
	return &List{data: list}, nil
}

// FindSimilar finds similar one
func (l *List) FindSimilar(v semver.Version) *Semv {
	newSemv := &Semv{}
	for _, vv := range l.data {
		if vv.Major == v.Major && vv.Minor == v.Minor && vv.Patch == v.Patch {
			newSemv = MustNew(vv.String())
		}
	}
	return newSemv
}

// String to string
func (l *List) String() string {
	var ll []string
	for _, v := range l.data {
		ll = append(ll, defaultTagPrefix+v.String())
	}
	return strings.Join(ll, "\n")
}

// Latest gets latest version from List
func (l *List) Latest() *Semv {
	if l.data.Len() > 0 {
		return &Semv{data: l.data[len(l.data)-1]}
	}
	return &Semv{data: semver.MustParse(defaultVersion)}
}

// WithoutPreRelease excludes pre-release
func (l *List) WithoutPreRelease() *List {
	var list semver.Versions
	for _, v := range l.data {
		if len(v.Pre) > 0 {
			continue
		}
		list = append(list, v)
	}
	return &List{data: list}
}

// OnlyPreRelease includes only pre-release
func (l *List) OnlyPreRelease() *List {
	var list semver.Versions
	for _, v := range l.data {
		if len(v.Pre) == 0 {
			continue
		}
		list = append(list, v)
	}
	return &List{data: list}
}


type Tags []struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID string `json:"node_id"`
}

func getContent(url string) ([]byte, error) {

    client := http.Client{}
    req , err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("Cannot initiate http client: %v", err)
    }
    req.Header.Set("Accept", "application/vnd.github.v3+json")

    resp , err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("Read body: %v", err)
    }

    return data, nil
}

// getVersions executes git tag as command
func getVersions(url string) (semver.Versions, error) {

        body, err := getContent(url + "/tags")
        if err != nil {
            return nil, fmt.Errorf("Request error: %v", err)
        }

        var tags Tags
        err = json.Unmarshal(body, &tags)
        if err != nil {
            return nil, fmt.Errorf("Cannot Unmarshal tags: %v", err)
        }

	var list semver.Versions
	for i := 0; i < len(tags); i++  {
		if defaultTagPrefix != "" {
			trimmed := strings.TrimPrefix(tags[i].Name, defaultTagPrefix)
			if trimmed == tags[i].Name {
				continue
			}
		}
		sv, err := semver.ParseTolerant(tags[i].Name)
		if err != nil {
			continue
		}
		list = append(list, sv)
	}
	semver.Sort(list)

	return list, nil
}
