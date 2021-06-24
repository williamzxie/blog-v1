package impl

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/enum"
	"blog-go-gin/models/model"
	"fmt"
	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"net/smtp"
	"sync"
	"time"
)

type UserAuthServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserAuthServiceImpl) Register(user *pb.User) error {
	//TODO 检测账号是否存在

	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		//新增用户信息
		userId, err := model.AddUserInfo(tx, &model.UserInfo{
			Email:      user.GetUsername(),
			Nickname:   fmt.Sprintf("用户%s", common.GetRandomString()),
			Avatar:     common.DefaultAvatar,
			CreateTime: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		//新增用户账号
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			return err
		}
		err = model.AddUserAuth(tx, &model.UserAuth{
			UserInfoID: userId,
			Username:   user.GetUsername(),
			Password:   string(hashPwd),
			CreateTime: time.Now().Unix(),
			LoginType:  enum.EMAIL.GetLoginType(),
		})
		if err != nil {
			return err
		}
		//绑定用户角色
		err = model.AddUserRole(tx, &model.UserRole{
			UserID: userId,
			RoleID: enum.Admin.GetRoleId(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserAuthServiceImpl) Login(user *pb.User) error {
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPwd)) //验证（对比）
	//if err != nil {
	//	fmt.Println("pwd wrong")
	//} else {
	//	fmt.Println("pwd ok")
	//}
	panic("implement me")
}

func (u *UserAuthServiceImpl) GetLoginCode(username string) error {
	//TODO 检验账号是否合法
	//生成六位随机验证码发送
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	logging.Logger.Debug(code)
	//将验证码通过邮件发送给用户
	mail := email.NewEmail()
	//设置发送方的邮箱
	mail.From = common.Sender
	// 设置接收方的邮箱
	mail.To = []string{username}
	//设置主题
	mail.Subject = common.EmailSubject
	//设置文件发送的内容
	mail.Text = []byte("您的验证码为 " + code + " 有效期15分钟，请不要告诉他人哦！")
	//设置服务器相关的配置
	err := mail.Send(common.EmailServerAddr, smtp.PlainAuth("", common.Sender, common.EmailAuthorizationCode, common.EmailHost))
	if err != nil {
		logging.Logger.Error(err)
		return err
	}
	// 将验证码存入redis，设置过期时间为15分钟
	common.RedisUtil.SetEx(common.CodeKey+username, code, common.CodeExpireTime, time.Second)
	return nil
}
