package mask

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	DefaultMaskChar = "*"
)

type Masks struct {
	MaskChar string   `yaml:"maskChar"`
	Values   []string `yaml:"values"`
}

func (m *Masks) Add(r string) {
	if !m.contains(r) {
		m.Values = append(m.Values, r)
	}
}

func (m *Masks) contains(r string) bool {
	for _, val := range m.Values {
		if strings.EqualFold(val, r) {
			return true
		}
	}
	return false
}

func LoadMasks() *Masks {
	var m Masks
	if err := viper.Unmarshal(&m); err != nil {
		m = Masks{
			MaskChar: DefaultMaskChar,
			Values:   make([]string, 0),
		}
		m.Save()
	}
	return &m
}

func (m *Masks) Save() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Err: Could not locate home directory.")
		os.Exit(1)
	}
	content, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println("Err: Could not store configuration.")
		os.Exit(1)
	}
	f, err := os.OpenFile(path.Join(home, ".mask"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Print("Err: Could not access file `.mask` in home directory")
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		fmt.Print("Err: Could not store configuration in `.mask` file")
		os.Exit(1)
	}
}

func (r *Masks) Remove(v string) bool {
	for i, existing := range r.Values {
		if strings.EqualFold(existing, v) {
			r.Values = append(r.Values[:i], r.Values[i+1:]...)
			return true
		}
	}
	return false
}
