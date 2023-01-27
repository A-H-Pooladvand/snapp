package models

import (
	"errors"
	"gorm.io/gorm"
	"order/pkg/mysql"
	"order/pkg/str"
)

type Model struct {
	gorm.Model
	query *gorm.DB
	items any
}

func DB() *Model {
	return NewModel()
}

func NewModel() *Model {
	return &Model{
		query: mysql.New(),
	}
}

func (m *Model) DB() *Model {
	return DB()
}

func (m *Model) Get() *Model {
	m.query.Find(m.items)

	return m
}

func (m *Model) WhereEq(field, value, op string) *Model {
	if "" == op {
		op = "="
	}
	op = str.Wrap(op, " ")
	field = str.Wrap(field, "`")

	query := field + op + "?"

	m.query = m.query.Where(query, value)

	return m
}

func (m *Model) Result(model any) *Model {
	m.items = model

	return m
}

func (m *Model) When(condition bool, callback func(model *Model) *Model) *Model {
	if condition {
		return callback(m)
	}

	return m
}

func (m *Model) Create(attrs any) *Model {
	m.query = m.query.Create(attrs)

	return m
}

func (m *Model) First() *Model {
	m.query = m.query.First(m.items)

	return m
}

func (m *Model) IsEmpty() bool {
	err := m.query.Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (m *Model) IsNotEmpty() bool {
	return !m.IsEmpty()
}

func (m *Model) Find(conds ...any) *Model {
	m.query = m.query.First(m.items, conds...)

	return m
}

func (m *Model) Select(query any, args ...any) *Model {
	m.query = m.query.Select(query, args...)

	return m
}

func (m *Model) Update(model any, attributes any) bool {
	m.query = m.query.Model(model).Updates(attributes)

	return m.query.RowsAffected > 0 && m.query.Error == nil
}

func (m *Model) Delete(value any, conds ...any) bool {

	m.query = m.query.Delete(value, conds...)

	return m.query.RowsAffected > 0 && m.query.Error == nil
}

func (m *Model) Preload(query string, args ...any) *Model {
	m.query = m.query.Preload(query, args...)

	return m
}
