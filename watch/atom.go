package watch

import (
	"encoding/xml"
	"errors"
)

type Atom struct {
	EntryList []*Entry `xml:"entry"`
}

type Entry struct {
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Title   string `xml:"title"`
}

func atomGetLatest(data []byte) (*Entry, error) {
	atom := &Atom{}
	err := xml.Unmarshal(data, atom)
	if err != nil {
		return nil, err
	}
	if len(atom.EntryList) < 1 {
		return nil, errors.New("No Entries")
	}
	return atom.EntryList[0], nil
}
