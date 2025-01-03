package storage

import (
	"github.com/fatihesergg/go_blog/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Seed User
	db.AutoMigrate(&model.User{})
	var user model.User
	if err := db.First(&user).Error; err != nil && err == gorm.ErrRecordNotFound {
		password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		db.Create(&model.User{Name: "Admin", LastName: "Admin", UserName: "admin", Email: "admin@admin", Password: string(password)})
	}
	db.AutoMigrate(&model.Post{})

	PostgresStore = NewPostgreSqlStorage(*db)
}

var PostgresStore Repository

type Repository interface {
	GetAllPosts() ([]model.Post, error)
	AddPost(model.Post) error
	GetPostById(int) (model.Post, error)
	UpdatePost(model.Post) error
	DeletePost(int) error
	GetUserByEmail(string) (model.User, error)
	CreateUser(model.User) error
	GetUserById(int) (model.User, error)
	SearchPost(string) ([]model.Post, error)
}

type PostgreSqlStorage struct {
	DB gorm.DB
}

func NewPostgreSqlStorage(db gorm.DB) *PostgreSqlStorage {
	return &PostgreSqlStorage{DB: db}
}

func (s *PostgreSqlStorage) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	err := s.DB.Find(&posts).Error
	return posts, err
}

func (s *PostgreSqlStorage) AddPost(post model.Post) error {
	return s.DB.Create(&post).Error
}

func (s *PostgreSqlStorage) GetPostById(id int) (model.Post, error) {
	var post model.Post
	err := s.DB.First(&post, id).Error
	return post, err
}

func (s *PostgreSqlStorage) UpdatePost(post model.Post) error {
	return s.DB.Save(&post).Error
}

func (e *PostgreSqlStorage) DeletePost(id int) error {
	return e.DB.Delete(&model.Post{}, id).Error
}

func (s *PostgreSqlStorage) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (s *PostgreSqlStorage) CreateUser(user model.User) error {
	return s.DB.Create(&user).Error
}

func (s *PostgreSqlStorage) GetUserById(id int) (model.User, error) {
	var user model.User
	err := s.DB.First(&user, id).Error
	return user, err
}

func (s *PostgreSqlStorage) SearchPost(query string) ([]model.Post, error) {
	var posts []model.Post
	err := s.DB.Model(&model.Post{}).Where("UPPER(title) LIKE UPPER(?) OR UPPER(content) LIKE UPPER(?)", "%"+query+"%", "%"+query+"%").Find(&posts).Error
	return posts, err
}
