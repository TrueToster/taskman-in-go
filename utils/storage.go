package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Task struct {
	Id       int
	Title    string
	Date     string
	Priority int
	Details  string
	Done     bool
}

type Storage struct {
	Dir string
}

func NewStorage(dir string) *Storage {
	return &Storage{Dir: dir}
}

func (s *Storage) Init() error {
	return os.MkdirAll(s.Dir, 0755)
}

func (s *Storage) LoadAll() ([]Task, error) {
	entries, err := os.ReadDir(s.Dir)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, e := range entries {
		if filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.Dir, e.Name()))
		if err != nil {
			continue
		}
		var t Task
		if json.Unmarshal(data, &t) == nil {
			tasks = append(tasks, t)
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	return tasks, nil
}

func (s *Storage) NextID() int {
	tasks, _ := s.LoadAll()
	max := 0
	for _, t := range tasks {
		if t.Id > max {
			max = t.Id
		}
	}
	return max + 1
}

func (s *Storage) Save(t Task) error {
	data, _ := json.MarshalIndent(t, "", "  ")
	path := filepath.Join(s.Dir, fmt.Sprintf("%d.json", t.Id))
	return os.WriteFile(path, data, 0644)
}

func (s *Storage) Delete(id int) error {
	return os.Remove(filepath.Join(s.Dir, fmt.Sprintf("%d.json", id)))
}

func (s *Storage) Get(id int) (*Task, error) {
	data, err := os.ReadFile(filepath.Join(s.Dir, fmt.Sprintf("%d.json", id)))
	if err != nil {
		return nil, err
	}
	var t Task
	json.Unmarshal(data, &t)
	return &t, nil
}
