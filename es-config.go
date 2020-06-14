package main

type LoaderConfig struct {
	url         string
	template    string
	write_alias string
	read_alias  string
	shards      int
	object      string
	box_type    string
}

func NewLoader() *LoaderConfig {
	return &LoaderConfig{
		url:         "http://localhost:9200",
		read_alias:  "cyberprobe",
		write_alias: "active-cyberprobe",
		template:    "active-cyberprobe",
		shards:      1,
	}
}

func (lc LoaderConfig) Url(url string) *LoaderConfig {
	lc.url = url
	return &lc
}

func (lc LoaderConfig) ReadAlias(val string) *LoaderConfig {
	lc.read_alias = val
	return &lc
}

func (lc LoaderConfig) WriteAlias(val string) *LoaderConfig {
	lc.write_alias = val
	return &lc
}

func (lc LoaderConfig) Template(val string) *LoaderConfig {
	lc.template = val
	return &lc
}

func (lc LoaderConfig) Shards(val int) *LoaderConfig {
	lc.shards = val
	return &lc
}

func (lc LoaderConfig) BoxType(val string) *LoaderConfig {
	lc.box_type = val
	return &lc
}

func (lc LoaderConfig) Build() (*Loader, error) {
	l := &Loader{
		LoaderConfig: lc,
	}
	err := l.Init()
	if err != nil {
		return nil, err
	}
	return l, nil
}
