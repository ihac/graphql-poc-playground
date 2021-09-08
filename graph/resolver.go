package graph

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/hotstar/graphql-poc-playground/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	companyStorage map[string]*model.Company
	userStorage    map[string]*model.User

	maxUserID    int
	maxCompanyID int

	userMutex    sync.Mutex
	companyMutex sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		companyStorage: make(map[string]*model.Company),
		userStorage:    make(map[string]*model.User),
	}
}

func (r *Resolver) ListAllCompanies() ([]*model.Company, error) {
	fmt.Println("COMMAND: select * from companies")

	var companies []*model.Company
	for _, company := range r.companyStorage {
		companies = append(companies, company)
	}

	return companies, nil
}

func (r *Resolver) ListAllUsers() ([]*model.User, error) {
	fmt.Println("COMMAND: select * from users")

	var users []*model.User
	for _, user := range r.userStorage {
		users = append(users, user)
	}
	return users, nil
}

func (r *Resolver) ListAllUsersByCompany(companyID string) ([]*model.User, error) {
	fmt.Printf("COMMAND: select * from users where company = %s\n", companyID)

	var users []*model.User
	for _, user := range r.userStorage {
		if user.CompanyID == companyID {
			users = append(users, user)
		}
	}
	return users, nil
}

func (r *Resolver) ListAllUsersByCompanies(companyIDs []string) (map[string][]*model.User, error) {
	fmt.Printf("COMMAND: select * from users where company in (%s)\n", strings.Join(companyIDs, ", "))

	companyIDMap := make(map[string]struct{})
	for _, companyID := range companyIDs {
		companyIDMap[companyID] = struct{}{}
	}

	users := make(map[string][]*model.User)
	for _, user := range r.userStorage {
		if _, ok := companyIDMap[user.CompanyID]; !ok {
			continue
		}
		users[user.CompanyID] = append(users[user.CompanyID], user)
	}
	return users, nil
}

func (r *Resolver) GetCompanyByID(companyID string) (*model.Company, error) {
	fmt.Printf("COMMAND: select * from companies where ID = %s\n", companyID)

	company, ok := r.companyStorage[companyID]
	if !ok {
		return nil, fmt.Errorf("company %s does not exist", companyID)
	}
	return company, nil
}

func (r *Resolver) GetUserByID(userID string) (*model.User, error) {
	fmt.Printf("COMMAND: select * from users where ID = %s\n", userID)

	user, ok := r.userStorage[userID]
	if !ok {
		return nil, fmt.Errorf("user %s does not exist", userID)
	}
	return user, nil
}

func (r *Resolver) InsertUser(new *model.NewUser) (*model.User, error) {
	if _, ok := r.companyStorage[new.CompanyID]; !ok {
		return nil, fmt.Errorf("company %s does not exist", new.CompanyID)
	}
	r.userMutex.Lock()
	r.maxUserID++
	r.userMutex.Unlock()

	fmt.Printf("COMMAND: insert into users (name, age, companyId) values (%s, %d, %s)\n", new.Name, new.Age, new.CompanyID)

	id := strconv.Itoa(r.maxUserID)
	user := model.User{
		ID:        id,
		Name:      new.Name,
		Age:       new.Age,
		CompanyID: new.CompanyID,
	}
	r.userStorage[id] = &user
	r.companyStorage[user.CompanyID].Employees = append(r.companyStorage[user.CompanyID].Employees, &user)
	return &user, nil
}

func (r *Resolver) InsertCompany(new *model.NewCompany) (*model.Company, error) {
	r.companyMutex.Lock()
	r.maxCompanyID++
	r.companyMutex.Unlock()

	fmt.Printf("COMMAND: insert into companies (name, address) values (%s, %s)\n", new.Name, new.Address)

	id := strconv.Itoa(r.maxCompanyID)
	company := model.Company{
		ID:        id,
		Name:      new.Name,
		Address:   new.Address,
		Employees: []*model.User{},
	}
	r.companyStorage[id] = &company
	return &company, nil
}
