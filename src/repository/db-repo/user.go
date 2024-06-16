package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/models"
	"github.com/Orololuwa/collect_am-api/src/repository"
	"gorm.io/gorm"
)

type userOrm struct {
	DB *sql.DB
	dbGorm *gorm.DB
}
func NewUserDBRepo(db *driver.DB) repository.UserDBRepo {
	return &userOrm{
		DB: db.SQL,
		dbGorm: db.Gorm,
	}
}

type testUserDBRepo struct {
	DB *sql.DB
}
func NewUserTestingDBRepo() repository.UserDBRepo {
	return &testUserDBRepo{
	}
}


func (m *userOrm) CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var newId int
	var err error

	queryFields := ""
	queryPlaceholders := ""
	var args []interface{}

	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		value := userValue.Field(i)
		tagValue := field.Tag.Get("db")

		// Check if the field is from an embedded struct
		if field.Anonymous {
			for j := 0; j < value.NumField(); j++ {
				embeddedField := value.Type().Field(j)
				embeddedValue := value.Field(j)
				embeddedTagValue := embeddedField.Tag.Get("db")

				if embeddedValue.IsZero() || embeddedTagValue == "" {
					continue
				}

				if queryFields == "" {
					queryFields += embeddedTagValue
				} else {
					queryFields += ", " + embeddedTagValue
				}

				if queryPlaceholders == "" {
					queryPlaceholders += "$" + strconv.Itoa(len(args)+1)
				} else {
					queryPlaceholders += ", $" + strconv.Itoa(len(args)+1)
				}

				args = append(args, embeddedValue.Interface())
			}
			continue
		}

		if value.IsZero() || tagValue == "" {
			continue
		}

		if queryFields == "" {
			queryFields += tagValue
		} else {
			queryFields += ", " + tagValue
		}

		if queryPlaceholders == "" {
			queryPlaceholders += "$" + strconv.Itoa(len(args)+1)
		} else {
			queryPlaceholders += ", $" + strconv.Itoa(len(args)+1)
		}

		args = append(args, value.Interface())
	}

	query := fmt.Sprintf(`
		INSERT INTO users 
			(%s)
		VALUES 
			(%s)
		RETURNING id
	`, queryFields, queryPlaceholders)

	if tx != nil {
		err = tx.QueryRowContext(ctx, query, args...).Scan(&newId)
	} else {
		err = m.DB.QueryRowContext(ctx, query, args...).Scan(&newId)
	}

	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (m *userOrm) GetAUser(ctx context.Context, tx *sql.Tx, u models.User) (*models.User, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user models.User

    // Prepare the base query
    query := `
        SELECT id, first_name, last_name, email, phone, password, created_at, updated_at
        FROM users
        WHERE 1=1
    `

    var args []interface{}

    userType := reflect.TypeOf(u)
    userValue := reflect.ValueOf(u)

    for i := 0; i < userType.NumField(); i++ {
        field := userType.Field(i)
        value := userValue.Field(i)
		tagValue := field.Tag.Get("db")

        if value.IsZero() || tagValue == "" {
            continue
        }

        switch value.Interface().(type) {
        case int, int64:
            query += " AND " + tagValue + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        case string:
            query += " AND " + tagValue + " = $" + strconv.Itoa(len(args)+1)
            args = append(args, value.Interface())
        // Add more cases as needed for other types
        }
    }

    // Execute the query
    var err error
    if tx != nil {
        err = tx.QueryRowContext(ctx, query, args...).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Phone,
			&user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    } else {
        err = m.DB.QueryRowContext(ctx, query, args...).Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Phone,
			&user.Password,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
    }

    if err == sql.ErrNoRows {
        return nil, nil // No rows found, return nil
    }

    if err != nil {
        return &user, err
    }

    return &user, nil
}

func (m *userOrm) GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var users = make([]models.User, 0)

	query := `
		SELECT (id, first_name, last_name, email, phone, created_at, updated_at)
		from users
	`

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query)
	}else{
		rows, err = m.DB.QueryContext(ctx, query)
	}
	if err != nil {
		return users, err
	}

	for rows.Next(){
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (m *userOrm) UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error){
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		UPDATE 
			users set (first_name, last_name) = ($1, $2)
		WHERE
			id = $3
	`

	var err error
	if tx != nil{
		_, err = tx.ExecContext(ctx, query, firstName, lastName, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, firstName, lastName, id)
	}

	if err != nil{
		return  err
	}

	return nil
}

func (m *userOrm) DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := "DELETE FROM users WHERE id = $1"

	var err error 

	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	}else{
		_, err = m.DB.ExecContext(ctx, query, id)
	}

    if err != nil {
        return err
    }

    return nil
}


func (o *userOrm) GetOneByID(id uint) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByEmail(email string) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("email = ?", email).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByPhone(phone string) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("phone = ?", phone).First(&user)
	return user, result.Error
}

func (o *userOrm) InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error) {
	db := o.dbGorm
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Create(&user)
	return user.ID, result.Error
}

func (o *userOrm) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	db := o.dbGorm
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Model(&user).Updates(&user)
	return result.Error
}