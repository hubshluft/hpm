package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetPackageInformation(pkg string) {
	var packageData map[string]interface{}
	pkgURL := db + strings.TrimSpace(string(pkg)) + ".json"

	fmt.Printf("[1/1] connecting to a given URL\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", pkgURL, nil)
	if err != nil {
		fmt.Printf("[0/1] error while connecting to a given URL: %s\n", err)
		return
	} else {
		fmt.Printf("[1/1] creating a URL request\n")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[0/1] error while creating a URL request: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("[0/1] bad status: %s\n", resp.Status)
			return
		} else {
			bodyText, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n%s\n\n", string(bodyText))

			err = json.Unmarshal(bodyText, &packageData)
			if err != nil {
				fmt.Printf("[0/1] error while unmarshaling a JSON file: %s\n", err)
				return
			} else {
				version, ok := packageData["version"].(string)
				if !ok {
					fmt.Printf("[0/1] error while parsing JSON file (version)\n")
					return
				}

				maintainer, ok := packageData["maintainer"].(string)
				if !ok {
					fmt.Printf("[0/1] error while parsing JSON file (maintainer)\n")
					return
				}

				dependencies, ok := packageData["dependencies"].([]interface{})
				if !ok {
					fmt.Printf("[0/1] Error while parsing JSON file (dependencies)\n")
					return
				}

				source, ok := packageData["source"].(string)
				if !ok {
					fmt.Printf("[0/1] error while parsing JSON file (source)\n")
					return
				}

				path, ok := packageData["path"].(string)
				if !ok {
					fmt.Printf("[0/1] error while parsing JSON file (source)\n")
					return
				}

				Sync(pkg, version, maintainer, dependencies, source, path)
			}
		}
	}
}
