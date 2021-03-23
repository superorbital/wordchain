/*
Copyright © 2021 SuperOrbital, LLC <info@superorbital.io>

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
package words

import (
	"fmt"
	"os"

	"github.com/markbates/pkger"
)

type WordCollection struct {
	Lists []WordGroup
}

type WordGroup struct {
	Type   string
	Length int
	Words  []string
}

type DataFile interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

///func findLength(wc WordCollection, value int) [10]int {
///	lists := [10]int{999, 999, 999, 999, 999, 999, 999, 999, 999, 999}
///	count := 0
///	for k, v := range wc.Lists {
///		if v.Length == value {
///			lists[count] = k
///			count += 1
///		}
///	}
///	return lists
///}

func GetList(file string) (DataFile, error) {
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Could not open file: %s", file)
			os.Exit(1)
		}
		return f, nil
	} else {
		f, err := pkger.Open("/data/words.json")
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}
