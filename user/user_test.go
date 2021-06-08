package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/scalent-gayatri/testing/mocks"
	"github.com/scalent-gayatri/testing/user"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
