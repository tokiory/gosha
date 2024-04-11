package task

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Task struct {
  Name string `json:"name"`
  Done bool `json:"done"`
  Id string `json:"id"`
}

func (t Task) Serialize() string {
  var encodedBytes, _ = json.Marshal(t)
  return string(encodedBytes)
}

func New(name string) Task {
  id := uuid.New()
  return Task {
    Name: name,
    Done: false,
    Id: id.String(),
  }
}
