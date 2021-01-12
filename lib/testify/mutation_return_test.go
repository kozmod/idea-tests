package testify

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type User struct {
	Name  string
	Email string
	ID    uuid.UUID
}

type repo interface {
	FindByID(uuid.UUID) (User, error)
	Create(user User)
}

type Service struct {
	repo
}

type repoMock struct {
	mock.Mock
	cache map[uuid.UUID]User
}

func newRepoMock() *repoMock {
	cache := make(map[uuid.UUID]User)
	return &repoMock{cache: cache}
}

func (m *repoMock) Create(user User) {
	m.Called(user)
	m.cache[user.ID] = user
}

func (m *repoMock) FindByID(uuid uuid.UUID) (User, error) {
	args := m.Called(uuid)
	res := args[0]
	if res == nil {
		return m.cache[uuid], args.Error(1)
	}
	return res.(User), args.Error(1)
}

func (s *Service) Register(newUser User) (User, error) {
	newUser.ID = uuid.New()
	s.repo.Create(newUser)
	createdUser, err := s.FindByID(newUser.ID)
	if err != nil {
		return User{}, err
	}
	return createdUser, nil
}

func TestMutationByCustomMock(t *testing.T) {
	repo := newRepoMock()
	s := Service{repo: repo}
	user := User{Name: "TestName", Email: "TestEmail"}

	repo.On("Create", mock.Anything)
	// I need to make sure FindByID function return value is a user with created UUID
	repo.On("FindByID", mock.Anything).Return(nil, nil)

	createdUser, err := s.Register(user)

	assert.NoError(t, err)
	assert.IsType(t, createdUser.ID, uuid.UUID{})
}
