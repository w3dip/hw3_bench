package main

import (
	"bufio"
	"bytes"
	//"hw3_bench/user"

	//"encoding/json"
	json "encoding/json"
	"fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"io"
	"os"
	//"regexp"
	"strings"
)

//easyjson:json
type User struct {
	Browsers []string `json:"browsers"`
	Company  string   `json:"company"`
	Country  string   `json:"country"`
	Email    string   `json:"email"`
	Job      string   `json:"job"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
}

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAe1574d6DecodeHw3BenchUser(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "company":
			out.Company = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAe1574d6EncodeHw3BenchUser(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"company\":"
		out.RawString(prefix)
		out.String(string(in.Company))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe1574d6EncodeHw3BenchUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe1574d6EncodeHw3BenchUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe1574d6DecodeHw3BenchUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe1574d6DecodeHw3BenchUser(l, v)
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]interface{})
	uniqueBrowsers := 0
	i := 0
	//scanner := bufio.NewScanner(file)
	reader := bufio.NewReader(file)
	buf := &bytes.Buffer{}
	fmt.Fprintln(out, "found users:")
	//user := make(map[string]interface{})
	user := User{}
	//for scanner.Scan() {
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		//line := scanner.Text()
		lexer := &jlexer.Lexer{Data: line}
		user.UnmarshalEasyJSON(lexer)
		//err := json.Unmarshal([]byte(line), &user)
		//if err != nil {
		//	panic(err)
		//}
		i++

		isAndroid := false
		isMSIE := false

		//browsers, ok := user["browsers"].([]interface{})
		//if !ok {
		//	continue
		//}

		//for _, browserRaw := range browsers {
		for _, browser := range user.Browsers {
			//browser, ok := browserRaw.(string)
			//if !ok {
			//	continue
			//}
			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				if _, ok := seenBrowsers[browser]; !ok && (isAndroid || isMSIE) {
					seenBrowsers[browser] = struct{}{}
					uniqueBrowsers++
				}
			}

			if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				if _, ok := seenBrowsers[browser]; !ok && (isAndroid || isMSIE) {
					seenBrowsers[browser] = struct{}{}
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		//email := strings.ReplaceAll(user["email"].(string), "@", " [at] ")
		email := strings.ReplaceAll(user.Email, "@", " [at] ")
		//fmt.Fprintf(buf, "[%d] %s <%s>\n", i-1, user["name"], email)
		fmt.Fprintf(buf, "[%d] %s <%s>\n", i-1, user.Name, email)
	}

	out.Write(buf.Bytes())
	fmt.Fprintln(out, "\nTotal unique browsers", uniqueBrowsers)
}

//func main() {
//	fastOut := new(bytes.Buffer)
//	FastSearch(fastOut)
//	fastResult := fastOut.String()
//	fmt.Println(fastResult)
//}
