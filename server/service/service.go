package service

import (
	"log"
	"net/http"

	pb "github.com/abdtyx/Optimail/micro-user/dto"

	"github.com/abdtyx/Optimail/server/config"
	"github.com/abdtyx/Optimail/server/gpt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	cfg        *config.Config
	db         *gorm.DB
	userclient *pb.UserClient
	gpt        *gpt.GPTCore
}

func New(cfg *config.Config) *Service {
	svc := &Service{cfg: cfg}
	err := svc.init()
	if err != nil {
		return nil
	}
	return svc
}

func (s *Service) init() error {
	// gorm
	db, err := gorm.Open(mysql.Open(s.cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to open database connection, error: " + err.Error())
	}
	s.db = db.Debug()

	// micro-user client

	// gpt
	s.gpt = gpt.New(s.cfg.ChatGPT)

	return nil
}

func (s *Service) Close() error {
	if s.db != nil {
		db, err := s.db.DB()
		if err != nil {
			return err
		}
		return db.Close()
	}
	return nil
}

/*
	Homepage
*/

func (s *Service) RootHandler(c *gin.Context) {
	// return webpage
	c.String(http.StatusOK, "Welcome abdtyx's Optimail")
}

/*
	Main Functionality
*/

func (s *Service) Summarize(c *gin.Context) {
	// Get content from mail server
	// var req dto.SummaryRequest
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	log.Println("**ERROR**: /summarize", err)
	// 	c.JSON(http.StatusInternalServerError, "Server error")
	// 	return
	// }
	// log.Println(req)

	// auth

	// query db
	// setting, err := userclient.GetSettingsBySecret(req.Secret)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	// 		"error": gin.H{
	// 			"code":    "BadServiceGetSettingsBySecretRequest",
	// 			"message": "Bad credentials",
	// 		},
	// 	})
	// }

	// call chatgpt
	content := setting.Prompt + req.Content
	gptresp, err := s.gpt.Chat(content)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "gpt failed",
		})
	}

	// handle gpt's resp

	// ack
	c.JSON(http.StatusOK, gin.H{
		"message": "Summary done",
	})
}

func (s *Service) Emphasize(c *gin.Context) {
	var req dto.EmphasizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: /emphasize", err)
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	log.Println(req)

	// auth

	// query db

	// call chatgpt

	// handle gpt's resp

	// construct resp

	resp := dto.EmphasizeResponse{
		Emphasis: "TODO:",
	}
	c.JSON(http.StatusOK, resp)
}

/*
	gRPC APIs -> micro-user
*/

func (s *Service) CreateUser(c *gin.Context) {

}

func (s *Service) Login(c *gin.Context) {

}

func (s *Service) Logout(c *gin.Context) {

}

func (s *Service) ChangePwd(c *gin.Context) {

}

func (s *Service) GetSettings(c *gin.Context) {

}

func (s *Service) UpdateSettings(c *gin.Context) {

}

func (s *Service) ResetSecret(c *gin.Context) {

}
