package loginaccountmapper

import (
	"context"
	"falcon/model/db/loginaccout"
	"falcon/pkg/meta/dbu"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	S loginaccout.Schema
	m dbu.Meta
)

func init() {
	m.Init(&S)
}

type Mapper struct {
	*dbu.Mapper
	*loginaccout.Schema
}

func New(db *gorm.DB) *Mapper {
	return &Mapper{
		Mapper: dbu.New(db, &m),
		Schema: &S,
	}
}

func (m *Mapper) OneByID(ctx context.Context, id int64) (*loginaccout.Schema, error) {
	var row loginaccout.Schema
	err := m.DB.Model(&loginaccout.Schema{}).
		Where(m.Tag(&m.Id).Eq(), id).
		Take(&row).
		Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func (m *Mapper) CreateAccount(ctx context.Context, phone, password string) (bool, error) {
	var count int64
	err := m.DB.Model(&loginaccout.Schema{}).
		Where(m.Tag(&m.Phone).Eq(), phone).
		Count(&count).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if count > 0 {
		return false, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	row := &loginaccout.Schema{
		Email:    "",
		Phone:    phone,
		Password: string(hashedPassword),
	}

	err = m.DB.Create(row).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Mapper) PasswordLogin(ctx context.Context, phone, password string) (*loginaccout.Schema, error) {
	var row loginaccout.Schema
	err := m.DB.Model(&loginaccout.Schema{}).
		Where(m.Tag(&m.Phone).V(), phone).
		Take(&row).
		Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(password))
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, nil
	}
	return &row, nil
}
