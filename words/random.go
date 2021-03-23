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
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	types "github.com/superorbital/wordchain/types"
)

func Random(prefs types.Preferences) error {
	f, err := GetList(prefs.WordFile)
	if err != nil {
		return err
	}
	defer f.Close()
	wc := new(WordCollection)
	err = json.NewDecoder(f).Decode(wc)
	if err != nil {
		return err
	}
	//index := findLength(*wc, int(prefs.Length))
	chain := make([]string, len(prefs.Type))
	if prefs.Seed == "" {
		rand.Seed(time.Now().Unix())
	} else {
		h := md5.New()
		_, err := io.WriteString(h, prefs.Seed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Could not generate random seed")
			os.Exit(1)
		}
		var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
		rand.Seed(int64(seed))
	}
	for i := range chain {
		for idx := range wc.Lists {
			if (wc.Lists[idx].Type == prefs.Type[i]) && (wc.Lists[idx].Length == int(prefs.Length)) {
				randomIndex := rand.Intn(len(wc.Lists[idx].Words))
				chain[i] = wc.Lists[idx].Words[randomIndex]
			}
		}
		if chain[i] == "" {
			fmt.Fprintf(os.Stderr, "[ERROR] Could not print chain. Ensure the word list has %v lists available each with a word length of %d", prefs.Type, prefs.Length)
			os.Exit(1)
		}
	}

	pre := ""
	post := ""
	if prefs.Prepend != "" {
		pre = prefs.Prepend + prefs.Divider
	}
	if prefs.Postpend != "" {
		post = prefs.Divider + prefs.Postpend
	}
	fmt.Println(pre + strings.Join(chain[:], prefs.Divider) + post)

	return nil
}
