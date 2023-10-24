package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cache struct {
	entities map[int]Entity
}

func NewCache() *Cache {
	return &Cache{
		entities: make(map[int]Entity),
	}
}

func (c *Cache) AddEntity(entity Entity) error {
	c.entities[entity.GetID()] = entity
	return nil
}

func (c *Cache) GetByID(id int) (Entity, error) {
	entity, exists := c.entities[id]
	if !exists {
		return nil, fmt.Errorf("entity not found")
	}
	return entity, nil
}

func (c *Cache) GetAll() []Entity {
	entities := []Entity{}
	
	for _, val := range c.entities {
		entities = append(entities, val)
	}

	return entities
}

func (c *Cache) Init(path string, s Entity) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
			return err
	}

	var intermediateMap []map[string]interface{}
  err = json.Unmarshal(bytes, &intermediateMap)
	if err != nil {
		return err
	}
	
	for _, value := range intermediateMap {
		entry, _ := s.Convert(value)
		
		c.entities[entry.GetID()] = entry
		fmt.Println(entry)
	}

	return nil
}
