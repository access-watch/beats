package httpout

type config struct {
  Endpoint string             `config:"endpoint"`
  Headers  map[string]string  `config:"headers"`
}

var (
  defaultConfig = config{
    Endpoint: "https://access.watch/api/1.0/log",
    Headers:  nil,
  }
)

func (c *config) Validate() error {

  return nil
}
