package seeders

import "gorm.io/gorm"

type Registry struct {
	db *gorm.DB
}

type SeederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) SeederRegistry {
	return &Registry{db: db}
}

func (r *Registry) Run() {
	RunRoleSeeder(r.db)
	RunUserSeeder(r.db)
}
