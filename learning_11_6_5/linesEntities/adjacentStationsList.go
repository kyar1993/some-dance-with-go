package linesEntities

import (
	"learning_11_6_5/structsLib"
)

var AdjacentStationsList = map[string][]*structsLib.Station{
	"Dostoyevskaya":              []*structsLib.Station{Vladimirskaya},
	"Vladimirskaya":              []*structsLib.Station{Dostoyevskaya},
	"Spasskaya":                  []*structsLib.Station{SennayaPloschad, Sadovaya},
	"SennayaPloschad":            []*structsLib.Station{Sadovaya, Spasskaya},
	"Sadovaya":                   []*structsLib.Station{Spasskaya, SennayaPloschad},
	"PloschadAlexandraNevskogo1": []*structsLib.Station{PloschadAlexandraNevskogo2},
	"PloschadAlexandraNevskogo2": []*structsLib.Station{PloschadAlexandraNevskogo1},
}
