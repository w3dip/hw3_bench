package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	//"regexp"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fileContents, err := ioutil.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}

	//r := regexp.MustCompile("@")
	//seenBrowsers := []string{}
	seenBrowsers := make(map[string]interface{})
	uniqueBrowsers := 0
	//foundUsers := ""

	//lines := strings.Split(string(fileContents), "\n")

	//users := make([]map[string]interface{}, len(lines))
	i := 0
	scanner := bufio.NewScanner(file)
	buf := &bytes.Buffer{}
	fmt.Fprintln(out, "found users:")
	user := make(map[string]interface{})
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		//for _, line := range lines {
		line := scanner.Text()
		// fmt.Printf("%v %v\n", err, line)
		//err := json.Unmarshal([]byte(line), &users[indx])
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		//users = append(users, user)
		i++

		isAndroid := false
		isMSIE := false

		browsers, ok := user["browsers"].([]interface{})
		if !ok {
			// log.Println("cant cast browsers")
			continue
		}

		//buf := &bytes.Buffer{};

		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			//notSeenBefore := false;
			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				//notSeenBefore = true
				if _, ok := seenBrowsers[browser]; !ok && (isAndroid || isMSIE) {
					seenBrowsers[browser] = struct{}{}
					uniqueBrowsers++
				}
			}

			if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				//notSeenBefore = true
				if _, ok := seenBrowsers[browser]; !ok && (isAndroid || isMSIE) {
					seenBrowsers[browser] = struct{}{}
					uniqueBrowsers++
				}
			}
			//for _, item := range seenBrowsers {
			//	if item == browser {
			//		notSeenBefore = false
			//	}
			//}
			//if notSeenBefore {
			//	// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
			//	seenBrowsers = append(seenBrowsers, browser)
			//	uniqueBrowsers++
			//}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.ReplaceAll(user["email"].(string), "@", " [at] ")
		fmt.Fprintf(buf, "[%d] %s <%s>\n", i-1, user["name"], email)
		//buf.WriteString()
	}

	out.Write(buf.Bytes())
	fmt.Fprintln(out, "\nTotal unique browsers", uniqueBrowsers)
}

func main() {
	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()
	fmt.Println(fastResult)
}
