package files

import (
	"encoding/csv"
	"io"
	"os"
)

func NewCSVWriter(dir string) io.Writer {
	file, _ := os.OpenFile(dir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer := csv.NewWriter(file)
	writer.Write([]string{"NIM TPB", "NIM Jurusan", "Username", "Nama", "Status", "Fakultas", "Jurusan", "Email ITB", "Email"})
	writer.Flush()
	return &csvWriter{

	}
}

type csvWriter struct {

}

func (c csvWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}
