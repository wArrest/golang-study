package maps

type Dictionary map[string]string

type DictionaryErr string

const (
  ErrNotFound         = DictionaryErr("could not find the word you were looking for")
  ErrWordExists       = DictionaryErr("cannot add word because it already exists")
  ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
  return string(e)
}

func Search(dictionary Dictionary, key string) (string, error) {
  if val, ok := dictionary[key]; ok {
    return val, nil
  }
  return "", ErrNotFound
}

func (d Dictionary) Create(key string, value string) error {
  if _, ok := d[key]; ok {
    return ErrWordExists
  }
  d[key] = value
  return nil
}

func (d Dictionary) Update(key string, newValue string) error {
  if _, ok := d[key]; !ok {
    return ErrWordDoesNotExist
  }
  d[key] = newValue
  return nil
}
