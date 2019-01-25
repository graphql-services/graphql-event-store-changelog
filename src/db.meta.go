package src

// Meta ...
type Meta struct {
	Name  string
	Value string `gorm:"type:text"`
}

// GetMeta ...
func (db *DB) GetMeta(name string) (*Meta, error) {
	var m Meta
	res := db.db.First(&m, "name = ?", name)

	if res.RecordNotFound() {
		return nil, nil
	}

	return &m, res.Error
}

// SaveMeta ...
func (db *DB) SaveMeta(name, value string) error {
	var m Meta
	if err := db.db.Where(Meta{Name: name}).FirstOrCreate(&m).Error; err != nil {
		return err
	}
	err := db.db.Model(&m).Update("value", value).Error
	return err
}
