package chain

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ChainDef struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name           string `yaml:"name"`
	PromptTemplate string `yaml:"prompt_template"`
	Model          string `yaml:"model"`
	ExtractKey     string `yaml:"extract_key"`   // pull this key from JSON response
	UseOutputOf    string `yaml:"use_output_of"` // inject prior step's output as {{.PriorOutput}}
	Parallel       bool   `yaml:"parallel"`      // run concurrently with adjacent parallel steps
}

type Loader struct {
	dir    string
	chains map[string]ChainDef
}

func NewLoader(dir string) (*Loader, error) { panic("fake") }

func (l *Loader) loadAll() error { panic("fake") }

// chains dir optional

func (l *Loader) load(name string) error { panic("fake") }

func (l *Loader) Get(name string) (ChainDef, bool) { panic("fake") }

func (l *Loader) Names() []string { panic("fake") }
