package httpout

import (
  "bytes"
  "encoding/json"
  "net/http"

  "github.com/elastic/beats/libbeat/common"
  "github.com/elastic/beats/libbeat/common/op"
  "github.com/elastic/beats/libbeat/logp"
  "github.com/elastic/beats/libbeat/outputs"
)

func init() {
  outputs.RegisterOutputPlugin("http", New)
}

type httpOutput struct {
  config config
}

func New(_ string, config *common.Config, _ int) (outputs.Outputer, error) {
  c := &httpOutput{config: defaultConfig}
  err := config.Unpack(&c.config)
  if err != nil {
    return nil, err
  }

  return c, nil
}

func (out *httpOutput) init(config config) error {
  return nil
}

func (out *httpOutput) Close() error {
  return nil
}

func (out *httpOutput) PublishEvent(
  sig op.Signaler,
  opts outputs.Options,
  data outputs.Data,
) error {
  var jsonEvent []byte
  var err error

  client := &http.Client{}

  jsonEvent, err = json.Marshal(data.Event)
  if err != nil {
    logp.Err("Fail to convert the event to JSON (%v): %#v", err, data.Event)
    op.SigCompleted(sig)
    return err
  }

  req, err := http.NewRequest("POST", out.config.Endpoint, bytes.NewBuffer(jsonEvent))
  if err != nil {
    op.SigCompleted(sig)
    return err
  }

  for k, v := range out.config.Headers {
    req.Header.Set(k, v)
  }

  resp, err := client.Do(req)
  if err != nil {
    op.SigCompleted(sig)
    return err
  }

  defer resp.Body.Close()

  op.SigCompleted(sig)
  return nil
}
