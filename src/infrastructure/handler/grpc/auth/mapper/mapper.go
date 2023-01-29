package mapper

import (
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/lumialvarez/go-grpc-auth-service/src/internal/user"
)

type Mapper struct {
}

func (m Mapper) ToDomainRegisterRequest(registerReq *pb.RegisterRequest) *user.User {
	return user.NewUser(0, registerReq.Name, registerReq.UserName, registerReq.Email, registerReq.Password, "", m.toDomainRole(registerReq.Role), true, []user.Notification{})
}

func (m Mapper) ToDTORegisterResponse(domainUser *user.User) *pb.RegisterResponse {
	return &pb.RegisterResponse{UserId: domainUser.Id()}
}

func (m Mapper) ToDomainLoginRequest(loginReq *pb.LoginRequest) *user.User {
	return user.NewUser(0, "", loginReq.UserName, "", loginReq.Password, "", "", true, []user.Notification{})
}

func (m Mapper) ToDTOLoginResponse(domainUser *user.User) *pb.LoginResponse {
	return &pb.LoginResponse{
		Token:    domainUser.Token(),
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     m.toDTORole(domainUser.Role()),
	}
}

func (m Mapper) ToDomainValidateRequest(validateReq *pb.ValidateRequest) *user.User {
	return user.NewUser(0, "", "", "", "", validateReq.Token, "", true, []user.Notification{})
}

func (m Mapper) ToDTOValidateResponse(domainUser *user.User) *pb.ValidateResponse {
	return &pb.ValidateResponse{
		UserId:   domainUser.Id(),
		UserName: domainUser.UserName(),
		Role:     m.toDTORole(domainUser.Role()),
	}
}

func (m Mapper) ToDTOListResponse(domainUsers *[]user.User) *pb.ListResponse {
	var dto pb.ListResponse
	for _, domainUser := range *domainUsers {
		tmp := pb.ListResponse_UserList{
			UserId:   domainUser.Id(),
			Name:     domainUser.Name(),
			UserName: domainUser.UserName(),
			Email:    domainUser.Email(),
			Role:     m.toDTORole(domainUser.Role()),
			Status:   domainUser.Status(),
		}
		dto.Users = append(dto.Users, &tmp)
	}
	return &dto
}

func (m Mapper) ToDomainUpdateRequest(updateReq *pb.UpdateRequest) *user.User {
	dto := updateReq.User
	return user.NewUser(dto.UserId, dto.Name, dto.UserName, dto.Email, dto.Password, "", m.toDomainRole(dto.Role), dto.Status, []user.Notification{})
}

func (m Mapper) ToDTOCurrentResponse(domainUser *user.User) *pb.CurrentResponse {
	var dto pb.CurrentResponse
	for _, domainNotif := range domainUser.Notifications() {
		tmp := pb.CurrentResponse_UserNotification{
			Id:     domainNotif.Id(),
			Title:  domainNotif.Title(),
			Detail: domainNotif.Detail(),
			Date:   domainNotif.Date().String(),
			Read:   domainNotif.Read(),
		}
		dto.Notifications = append(dto.Notifications, &tmp)
	}

	dto.UserId = domainUser.Id()
	dto.Name = domainUser.Name()
	dto.UserName = domainUser.UserName()
	dto.Email = domainUser.Email()
	dto.Role = m.toDTORole(domainUser.Role())

	return &dto
}

func (m Mapper) toDomainRole(dtoRole string) user.Role {
	return user.Role(dtoRole)
}

func (m Mapper) toDTORole(domainRole user.Role) string {
	return string(domainRole)
}
