package antivirus_blocker

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func Test_GetIniBase64(t *testing.T) {
	file := "D:/git/celluloid/bin/anvir.ini"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("read file %s error %v", file, err)
	}
	var buf bytes.Buffer
	err = gzipWrite(&buf, []byte(b))
	if err != nil {
		t.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	t.Log(str)
}
