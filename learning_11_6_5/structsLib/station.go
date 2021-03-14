// Station struct Cтанция метро
package structsLib

type Station struct {
	Id          string
	Name        string
	Description string
	MetroLine   *Line
	Mute        bool
}
