package search

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	def, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, val string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = val
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
