/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package windows

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// downloadFile saves a remote URL to a local temp file, and returns its path.
// It's the caller's responsibility to clean up the temp file when done.
func downloadFile(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to download from %q: %w", url, err)
	}
	defer response.Body.Close()

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("unable to create temp file: %w", err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, response.Body)
	return tempFile.Name(), err
}
