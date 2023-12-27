package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	names "github.com/rotblauer/cattracks-names"
)

// TestAliasName tests basic usage of AliasOrName
// and proves that we don't have an import cycle issue.
func TestAliasName(t *testing.T) {
	ryeSanitized := names.SanitizeName(names.AliasOrName("Rye13"))
	if ryeSanitized != "rye" {
		t.Errorf("expected rye, got %s", ryeSanitized)
	}
}

func ExampleCLI_Modify() {
	mock := `
{"type":"Feature","id":1,"geometry":{"type":"Point","coordinates":[-93.25535583496094,44.98938751220703]},"properties":{"Accuracy":16.58573341369629,"Activity":"Stationary","Elevation":256.4858703613281,"Heading":347.74163818359375,"Name":"Rye13","Pressure":95.55844116210938,"Speed":0,"Time":"2023-12-08T10:04:10.017Z","UUID":"05C63745-BFA3-4DE3-AF2F-CDE2173C0E11","UnixTime":1702029850,"Version":"V.customizableCatTrackHat"}}
{"type":"Feature","id":1,"geometry":{"type":"Point","coordinates":[-93.25535583496094,44.98938751220703]},"properties":{"Accuracy":16.58573341369629,"Activity":"Stationary","Elevation":256.4858703613281,"Heading":347.74163818359375,"Name":"Rye13","Pressure":95.55844116210938,"Speed":0,"Time":"2023-12-08T10:04:10.017Z","UUID":"05C63745-BFA3-4DE3-AF2F-CDE2173C0E11","UnixTime":1702029850,"Version":"V.customizableCatTrackHat"}}
`
	reader := bytes.NewBuffer([]byte(mock))
	writer := os.Stdout

	err := modify(reader, writer, "properties.Name", true, false)
	if err != nil {
		log.Fatalln(err)
	}

	// Output:
	// {"type":"Feature","id":1,"geometry":{"type":"Point","coordinates":[-93.25535583496094,44.98938751220703]},"properties":{"Accuracy":16.58573341369629,"Activity":"Stationary","Elevation":256.4858703613281,"Heading":347.74163818359375,"Name":"rye","Pressure":95.55844116210938,"Speed":0,"Time":"2023-12-08T10:04:10.017Z","UUID":"05C63745-BFA3-4DE3-AF2F-CDE2173C0E11","UnixTime":1702029850,"Version":"V.customizableCatTrackHat"}}
	// {"type":"Feature","id":1,"geometry":{"type":"Point","coordinates":[-93.25535583496094,44.98938751220703]},"properties":{"Accuracy":16.58573341369629,"Activity":"Stationary","Elevation":256.4858703613281,"Heading":347.74163818359375,"Name":"rye","Pressure":95.55844116210938,"Speed":0,"Time":"2023-12-08T10:04:10.017Z","UUID":"05C63745-BFA3-4DE3-AF2F-CDE2173C0E11","UnixTime":1702029850,"Version":"V.customizableCatTrackHat"}}
}
