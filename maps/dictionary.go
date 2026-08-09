package maps

import "errors"

type Dictionary map[string]string 

var ErrDictionaryUnknownKey error = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error){
	if _, exists := d[key]; !exists {
		return "", ErrDictionaryUnknownKey	
	}
	return d[key], nil
}
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
