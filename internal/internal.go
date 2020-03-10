package internal

func NewVocabulary(lang string) *Vocabulary {
	return &Vocabulary{
		data: make(map[string]string),
		lang: lang,
	}
}

type Vocabulary struct {
	data map[string]string
	lang string
}

func (v *Vocabulary) AsMap() map[string]string {
	return v.data
}

func (v *Vocabulary) Set(key string, data string) {
	v.data[key] = data
}

func (v *Vocabulary) Lookup(key string) (data string, ok bool) {
	if value, ok := v.data[key]; ok {
		return "\x02" + value, true
	}
	return "", false
}
