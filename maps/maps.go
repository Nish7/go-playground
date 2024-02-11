package maps

type Dictionary map[string]string

var ErrNotFound = DictionaryErr("could not find the word you are looking for")
var ErrWordExists = DictionaryErr("word already exists")
var ErrWordDoesNotExists = DictionaryErr("word does not exists")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, exists := d[key]

	if !exists {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	if err == nil {
		return ErrWordExists
	}

	d[key] = val
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExists
	}

	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExists
	}

	delete(d, key)
	return nil
}
