package user_test

import (
	"errors"
	"fmt"
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
	fmt.Println("---------")
	testUser.Use()
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
	err := testUser.Use()

	if err != dummyError {
		t.Fail()
	}
}
