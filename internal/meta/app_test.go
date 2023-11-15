// Package meta app file defines variables.
//
// provides variables to store application info and getters to give read access
package meta_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/meta"
	catalogModel "github.com/cthiel77/server-demo/internal/meta/model/catalog"
)

func TestGetAppNameInit(t *testing.T) {
	value := meta.GetAppName()
	expected := meta.App.Name
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAppDescInit(t *testing.T) {
	value := meta.GetAppDescription()
	expected := meta.App.Description
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestGetAppLicenseInit(t *testing.T) {
	value := meta.GetAppLicense()
	expected := meta.Extended.License
	if value != expected {
		t.Errorf("wrong value %s, expected %s", value, expected)
	}
}

func TestInitCatalog(t *testing.T) {
	cat := catalogModel.App{}
	e := meta.InitAppCatalog(`{
		"global": {
			"key": "quote_lib",
			"label": "Zitat Bibliothek",
			"description": "Die Bibliothek stellt alle notwendigen Funktionen und Models zur Verfügung, um Zitat Funktionalität in anderen Anwendungen zu implementieren",
			"prerequisites": [
				{
					"description": "\"quotes.json\" Datei im Verzeichnis des Binaries ablegen. (Beispiel siehe: /config/quotes.json)",
					"implementation": {
						"duration": 10,
						"developers_count": 1
					}
				}
			]
		},
		"features": [
			{
				"key": "quote_cli_output",
				"label": "Zitat Kommandozeilen Ausgabe",
				"description": "Gibt das Zitat des Tages auf der Kommandozeile aus",
				"prerequisites": []
			},
			{
				"key": "quote_api",
				"label": "Zitat Api",
				"description": "Stellt einen Endpunkt zur Verfügung der das Zitat des Tages im Json Format zurückgibt",
				"prerequisites": []
			},
			{
				"key": "quote_mailer",
				"label": "Zitat Mailer",
				"description": "Verschickt das Zitat des Tages an eine fest definierte Empfängerliste",
				"prerequisites": [
					{
						"description": "smtp Daten in der Konfigurations datei anlegen und Empfänger Adressen anlegen (Beispiel siehe: /config/go-srv-quote-of-the-day.json)",
						"implementation": {
							"duration": 10,
							"developers_count": 1
						}
					}
				]
			}
		],
		"metrics": {
			"creation": {
				"duration": 2400,
				"developers_count": 1
			},
			"implementation": {
				"duration": 30,
				"developers_count": 1
			}
		}
	}`, &cat)

	if e != nil {
		t.Errorf("%+v", e)
	}

	t.Logf("%+v", cat)
}
