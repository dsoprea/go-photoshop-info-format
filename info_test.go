package photoshopinfo

import (
	"os"
	"testing"

	"github.com/dsoprea/go-logging"
)

func TestReadPhotoshop30InfoRecord(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	pir, err := ReadPhotoshop30InfoRecord(f)
	log.PanicIf(err)

	if pir.ImageResourceId != 0x0404 {
		t.Fatalf("Image resource ID not correct: (0x%04x)", pir.ImageResourceId)
	}
}

func TestReadPhotoshop30Info(t *testing.T) {
	filepath := GetTestDataFilepath()
	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	pirIndex, err := ReadPhotoshop30Info(f)
	log.PanicIf(err)

	if len(pirIndex) != 1 {
		t.Fatalf("Expected exactly one record.")
	}

	pir := pirIndex[0x0404]

	if pir.ImageResourceId != 0x0404 {
		t.Fatalf("Image resource ID not correct: (0x%04x)", pir.ImageResourceId)
	}
}

func TestPhotoshop30InfoRecord_String(t *testing.T) {
	pir := Photoshop30InfoRecord{
		RecordType:      "abc",
		ImageResourceId: 123,
		Name:            "def",
		Data:            []byte("ghijkl"),
	}

	s := pir.String()

	if s != "RECORD-TYPE=[abc] IMAGE-RESOURCE-ID=[0x007b] NAME=[def] DATA-SIZE=(6)" {
		t.Fatalf("String representation not correct: [%s]", s)
	}
}
