//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"rumahdermawan/backedn-rdi/app"
	"rumahdermawan/backedn-rdi/controller/activity"
	period3 "rumahdermawan/backedn-rdi/controller/period"
	user3 "rumahdermawan/backedn-rdi/controller/user"
	"rumahdermawan/backedn-rdi/repository/activities"
	"rumahdermawan/backedn-rdi/repository/period"
	"rumahdermawan/backedn-rdi/repository/user"
	activities2 "rumahdermawan/backedn-rdi/service/activities"
	period2 "rumahdermawan/backedn-rdi/service/period"
	user2 "rumahdermawan/backedn-rdi/service/user"
)

var userSet = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
	user2.NewUserService,
	wire.Bind(new(user2.UserService), new(*user2.UserServiceImpl)),
	user3.NewServiceUserToken,
	wire.Bind(new(user3.UserToken), new(*user3.UserTokenImpl)),
	user3.NewUserController,
	wire.Bind(new(user3.UserController), new(*user3.UserControllerImpl)),
)
var periodSet = wire.NewSet(
	period.NewPeriodRepository,
	wire.Bind(new(period.PeriodRepository), new(*period.PeriodRepositoryImpl)),
	period2.NewPeriodService,
	wire.Bind(new(period2.PeriodService), new(*period2.PeriodServiceImpl)),
	period3.NewPeriodController,
	wire.Bind(new(period3.PeriodController), new(*period3.PeriodControllerImpl)),
)

var activitySet = wire.NewSet(
	activities.NewActivityRepository,
	wire.Bind(new(activities.ActivitiyRepository), new(*activities.ActivityRepositoryImpl)),
	activities2.NewActivityService,
	wire.Bind(new(activities2.ActivityService), new(*activities2.ActivityServiceImpl)),
	activity.NewActivityController,
	wire.Bind(new(activity.ActivityController), new(*activity.ActivityControllerImpl)),
)

func InitializedServer() *gin.Engine {
	wire.Build(
		app.NewDb,
		validator.New,
		userSet,
		periodSet,
		activitySet,
		app.NewRouter,
	)
	return nil
}
