package src

// Meta ...
type Meta struct {
	Key   string
	Value string `gorm:"type:text"`
}

// GetMeta ...
func (db *DB) GetMeta(key string) (*Meta, error) {
	var m Meta
	res := db.db.First(&m, "key = ?", key)

	if res.RecordNotFound() {
		return nil, nil
	}

	return &m, res.Error
}

// SaveMeta ...
func (db *DB) SaveMeta(meta Meta) error {
	var m Meta
	if err := db.db.Where(Meta{Key: meta.Key}).FirstOrCreate(&m).Error; err != nil {
		return err
	}
	err := db.db.Model(&m).Update("value", meta.Value).Error
	return err
}
