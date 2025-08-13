package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v3"
)

type PackageEntry struct {
	NugetName      string `yaml:"ng"`
	GithubUserRepo string `yaml:"gh"`
}

type YamlData map[string][]PackageEntry

type NuGetRegistration struct {
	Mode  string `json:"mode"`
	Items []struct {
		Items []struct {
			CatalogEntry struct {
				Published string `json:"published"`
				Downloads int    `json:"downloads"`
			} `json:"catalogEntry"`
		} `json:"items"`
	} `json:"items"`
}

type GitHubRepoDocument struct {
	StargazersCount int `json:"stargazers_count"`
	WatchersCount   int `json:"watchers_count"`
}

type NuGetInfoContainer struct {
	TotalDownloads int64
}

type NugetSearchResult struct {
	Data []struct {
		Id             string `json:"@id"`
		TotalDownloads int64  `json:"totalDownloads"`
		Versions       []struct {
			Version   string `json:"version"`
			Downloads int64  `json:"downloads"`
			Id        string `json:"@id"`
		} `json:"versions"`
	} `json:"data"`
}

func main_nugets(data []byte, outFile *os.File) {
	var yamlData YamlData
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		log.Fatal(err)
	}

	out := bufio.NewWriter(outFile)
	fmt.Fprintln(out, "## Packages")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "[cols=5*]")
	fmt.Fprintln(out, "|===")
	fmt.Fprintln(out, "|category 2+|NuGet         2+|GitHub")
	fmt.Fprintln(out, "|           |name|tot dl     |name|stars")

	for _, category := range slices.Sorted(maps.Keys(yamlData)) {
		packages := yamlData[category]
		fmt.Fprintf(out, ".%d+| %s\n", len(packages), category)
		for _, pkg := range packages {
			fmt.Printf("processing %s - %s\n", pkg.NugetName, pkg.GithubUserRepo)

			data, err := getNuGetInfo(pkg.NugetName)
			_ = data
			if err != nil {
				fmt.Printf("  Package: %s - error: %v\n", pkg.NugetName, err)
				continue
			}

			fmt.Fprintf(out, "|https://www.nuget.org/packages/%s[%s]\n", pkg.NugetName, pkg.NugetName)
			// optional add ("image:https://img.shields.io/nuget/dt/%s?label=[]", pkg.NugetName)
			fmt.Fprintf(out, ">|%s\n", humanizeSep(data.TotalDownloads))

			fmt.Fprintf(out, "|https://github.com/%s[%s]\n", pkg.GithubUserRepo, pkg.GithubUserRepo)
			if len(pkg.GithubUserRepo) > 0 {
				ghData := getGithubData(pkg.GithubUserRepo)
				fmt.Fprintf(out, ">|%s", humanizeSep(ghData.StargazersCount))
			} else {
				fmt.Fprint(out, "|")
			}
			days := -1.
			totalDownloads := -1.
			avg := float64(totalDownloads) / days
			_ = avg

			// fmt.Fprintf("|%s  %s: total downloads=%d, published=%s, per day avg=%.2f\n",
			// 	pkg.NugetName, totalDownloads, publishedDate.Format("2006-01-02"), avg)
			fmt.Fprintln(out)
			fmt.Println()
		}
		fmt.Fprintln(out)
	}

	fmt.Fprintln(out, "|===")

	out.Flush()
}

// adds thousand separator
func humanizeSep[T constraints.Integer](num T) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", num)
}

// add SI suffix k, M, B
func humanizeSiSuf[T constraints.Integer](num T) string {
	numf := float64(num)

	bil := 1e9
	if numf > bil {
		return fmt.Sprintf("%.0fB", numf/bil)
	}

	mil := 1e6
	if numf > mil {
		return fmt.Sprintf("%.0fM", numf/mil)
	}
	kilo := float64(1e3)
	if numf > kilo {
		return fmt.Sprintf("%.0fk", numf/kilo)
	}

	return fmt.Sprint(num)
}

func getGithubData(userRepo string) GitHubRepoDocument {
	url := fmt.Sprint("https://api.github.com/repos/", userRepo)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return GitHubRepoDocument{StargazersCount: -1, WatchersCount: -1}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("GET %s returned %d\n", url, resp.StatusCode)
		return GitHubRepoDocument{StargazersCount: -2, WatchersCount: -2}
	}

	var doc GitHubRepoDocument
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		fmt.Printf("failed to parse json: %s\n", err)
		return GitHubRepoDocument{StargazersCount: -3, WatchersCount: -3}
	}

	return doc
}

func getNuGetInfo(name string) (NuGetInfoContainer, error) {
	var res NuGetInfoContainer

	data, err := getGenerelNugetInfo(name)
	_ = data
	if err != nil {
		return res, err
	}

	searchUrl := fmt.Sprintf("https://azuresearch-usnc.nuget.org/query?q=%s&prerelease=false&semVerLevel=2.0.0", name)

	resp, err := http.Get(searchUrl)
	if err != nil {
		fmt.Println("GET ", searchUrl, " failed with code: ", resp.StatusCode)
		return res, err
	}

	var searchRes NugetSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&searchRes); err != nil {
		return res, err
	}

	desiredId := fmt.Sprintf("https://api.nuget.org/v3/registration5-gz-semver2/%s/index.json", strings.ToLower(name))
	for _, item := range searchRes.Data {
		if item.Id == desiredId {
			res.TotalDownloads = item.TotalDownloads
			return res, nil
		}
	}

	return res, nil
}

func getGenerelNugetInfo(pkg string) (NuGetRegistration, error) {
	var res NuGetRegistration

	url := fmt.Sprintf("https://api.nuget.org/v3/registration5-semver1/%s/index.json", strings.ToLower(pkg))
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	fmt.Println("GET ", url, " returned status: ", resp.StatusCode)
	if resp.StatusCode != 200 {
		return res, fmt.Errorf("nuget api returned %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
