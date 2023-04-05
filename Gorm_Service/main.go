package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) Create(user *User) error {
	return s.DB.Create(user).Error
}

func (s *UserService) GetAll() ([]User, error) {
	var users []User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(id uint) (*User, error) {
	var user User
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Update(user *User) error {
	return s.DB.Save(user).Error
}

func (s *UserService) Delete(user *User) error {
	return s.DB.Delete(user).Error
}

func main() {
	// connect to the database
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// migrate the schema
	db.AutoMigrate(&User{})

	// create a new UserService with the database connection
	userService := &UserService{DB: db}
	productService := &ProductService{db}

	api := &API{
		User:    userService,
		Product: productService,
	}
}

// // create a new HTTP server and handle requests
// http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
//     switch r.Method {
//     case "GET":
//         users, err := userService.GetAll()
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }
//         json.NewEncoder(w).Encode(users)
//     case "POST":
//         var user User
//         if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//             http.Error(w, err.Error(), http.StatusBadRequest)
//             return
//         }
//         if err := userService.Create(&user); err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//             return
//         }
//         w.WriteHeader(http.StatusCreated)
//         json.NewEncoder(w).Encode(user)
//     default:
//         http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
//     }
// })

// http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
//     id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }
//     user, err := userService.GetByID(uint(id))
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusNotFound)
//         return
//     }
//     switch r.Method {
