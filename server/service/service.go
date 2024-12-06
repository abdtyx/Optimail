package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	pb "github.com/abdtyx/Optimail/micro-user/dto"
	"google.golang.org/grpc"

	"github.com/abdtyx/Optimail/server/config"
	"github.com/abdtyx/Optimail/server/gpt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	cfg *config.Config
	db  *gorm.DB
	gpt *gpt.GPTCore

	userclient pb.UserClient
	conn       *grpc.ClientConn
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
	s.conn, err = grpc.NewClient(s.cfg.MicroUser.GrpcAddr)
	if err != nil {
		panic("failed to connect micro-user: " + err.Error())
	}
	s.userclient = pb.NewUserClient(s.conn)

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
		err = db.Close()
		if err != nil {
			return err
		}
	}
	if s.conn != nil {
		s.conn.Close()
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

type summarizeReq struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

func (s *Service) Summarize(c *gin.Context) {
	// Get content from mail server
	var req summarizeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "email and content required")
		return
	}
	log.Println(req)

	// auth (skip now)

	// query db
	var idReq = &pb.UsersGetIdByEmailRequest{
		Email: req.Email,
	}
	idResp, err := s.userclient.UsersGetIdByEmail(context.Background(), idReq)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "id not found")
		return
	}
	var settingsReq = &pb.UsersGetSettingsRequest{
		Id: idResp.Id,
	}
	settingsResp, err := s.userclient.UsersGetSettings(context.Background(), settingsReq)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "settings not found")
		return
	}

	// call chatgpt
	content := settingsResp.Settings.SummaryPrompt + "\nEmail:\n" + req.Content
	gptresp, err := s.gpt.Chat(content)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "gpt failed")
	}

	// handle gpt's resp
	var createSummaryRequest = &pb.UsersCreateSummaryRequest{
		Id:      idResp.Id,
		Summary: gptresp,
	}
	_, err = s.userclient.UsersCreateSummary(context.Background(), createSummaryRequest)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "update db failed")
	}

	// ack
	c.JSON(http.StatusOK, gin.H{
		"message": "Summary done",
	})
}

type emphasizeReq struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

func (s *Service) Emphasize(c *gin.Context) {
	// Get content from mail server
	var req emphasizeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: /emphasize", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "email and content required")
		return
	}
	log.Println(req)

	// auth (skip now)

	// query db
	var idReq = &pb.UsersGetIdByEmailRequest{
		Email: req.Email,
	}
	idResp, err := s.userclient.UsersGetIdByEmail(context.Background(), idReq)
	if err != nil {
		log.Println("**ERROR**: /emphasize", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "id not found")
		return
	}
	var settingsReq = &pb.UsersGetSettingsRequest{
		Id: idResp.Id,
	}
	settingsResp, err := s.userclient.UsersGetSettings(context.Background(), settingsReq)
	if err != nil {
		log.Println("**ERROR**: /emphasize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "settings not found")
		return
	}

	// call chatgpt
	content := settingsResp.Settings.EmphasisPrompt + "\nEmail:\n" + req.Content
	gptresp, err := s.gpt.Chat(content)
	if err != nil {
		log.Println("**ERROR**: /emphasize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "gpt failed")
	}

	// handle gpt's resp
	var createEmphasisRequest = &pb.UsersCreateEmphasisRequest{
		Id:      idResp.Id,
		Summary: gptresp,
	}
	_, err = s.userclient.UsersCreateEmphasis(context.Background(), createEmphasisRequest)
	if err != nil {
		log.Println("**ERROR**: /summarize", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "update db failed")
	}

	// ack
	c.JSON(http.StatusOK, gin.H{
		"message": "Emphasis done",
	})
}

/*
	gRPC APIs -> micro-user
*/

func (s *Service) CreateUser(c *gin.Context) {
	var req pb.UsersCreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: CreateUser:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "Username and password required")
		return
	}
	resp, err := s.userclient.UsersCreateUser(context.Background(), &req)
	if err != nil {
		log.Println("**ERROR**: UsersCreateUser:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      resp.Id,
		"pk":      resp.Pk,
		"message": resp.Message,
	})
}

func (s *Service) Login(c *gin.Context) {
	var req pb.UsersLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: Login:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "Username and password required")
		return
	}
	resp, err := s.userclient.UsersLogin(context.Background(), &req)
	if err != nil {
		log.Println("**ERROR**: UsersLogin:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	// generate jwt
	jwtstr, err := genJwt(resp.Id, time.Now().Add(2*time.Hour), []byte(s.cfg.JWTKey))
	if err != nil {
		log.Println("**ERROR**: Login:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.Header("Authentication", jwtstr)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func (s *Service) Logout(c *gin.Context) {
	jwtstr := c.GetHeader("Authentication")
	id, err := verifyJwt(jwtstr, []byte(s.cfg.JWTKey))
	if err != nil {
		log.Println("**ERROR**: Logout:", err)
		c.AbortWithStatusJSON(http.StatusForbidden, err)
		return
	}
	var req = &pb.UsersLogoutRequest{
		Id: id,
	}
	_, err = s.userclient.UsersLogout(context.Background(), req)
	if err != nil {
		log.Println("**ERROR**: UsersLogout:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func (s *Service) ChangePwd(c *gin.Context) {
	var req pb.UsersChangePwdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("**ERROR**: ChangePwd:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "username ,old_password, and new_password required")
		return
	}
	_, err := s.userclient.UsersChangePwd(context.Background(), &req)
	if err != nil {
		log.Println("**ERROR**: UsersChangePwd:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed",
	})
}

func (s *Service) GetSettings(c *gin.Context) {
	jwtstr := c.GetHeader("Authentication")
	id, err := verifyJwt(jwtstr, []byte(s.cfg.JWTKey))
	if err != nil {
		log.Println("**ERROR**: GetSettings:", err)
		c.AbortWithStatusJSON(http.StatusForbidden, err)
		return
	}
	var req = &pb.UsersGetSettingsRequest{
		Id: id,
	}
	resp, err := s.userclient.UsersGetSettings(context.Background(), req)
	if err != nil {
		log.Println("**ERROR**: UsersGetSettings:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"settings": resp,
		"message":  "OK",
	})
}

func (s *Service) UpdateSettings(c *gin.Context) {
	jwtstr := c.GetHeader("Authentication")
	id, err := verifyJwt(jwtstr, []byte(s.cfg.JWTKey))
	if err != nil {
		log.Println("**ERROR**: Logout:", err)
		c.AbortWithStatusJSON(http.StatusForbidden, err)
		return
	}

	var req = &pb.UsersUpdateSettingsRequest{
		Id: id,
	}
	if err := c.ShouldBindJSON(req.Settings); err != nil {
		log.Println("**ERROR**: UpdateSettings:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "settings required")
		return
	}

	_, err = s.userclient.UsersUpdateSettings(context.Background(), &req)
	if err != nil {
		log.Println("**ERROR**: UsersUpdateSettings:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Settings updated",
	})
}

/*
	JWT
*/

func hmacEval(key []byte, msg []byte) ([]byte, error) {
	mac := hmac.New(sha512.New, key)
	mac.Write(msg)
	res := mac.Sum(nil)

	return res, nil
}

type jwt struct {
	ID     string    // uuid
	Expire time.Time // valid through
}

/*
	Simplified JWT. Only provide Integrity, no Authenticity.
*/

func genJwt(id string, expire time.Time, key []byte) (jwtstr string, err error) {
	jwtstr = ""
	j := jwt{
		ID:     id,
		Expire: expire,
	}
	b, err := json.Marshal(j)
	if err != nil {
		return
	}
	// generate MAC tag
	mac, err := hmacEval(key, b)
	if err != nil {
		return
	}
	return hex.EncodeToString(append(b, mac...)), nil
}

func verifyJwt(jwtstr string, key []byte) (id string, err error) {
	jwtbyte := []byte(jwtstr)
	// check length
	l := len(jwtbyte)
	if l <= 64 {
		return "", errors.New("invalid jwt")
	}
	// retrieve jwt byte
	b := jwtbyte[:l-64]
	// generate MAC tag
	mac, err := hmacEval(key, b)
	if err != nil {
		return "", err
	}
	// compare MAC tags
	if !hmac.Equal(mac, jwtbyte[l-64:]) {
		return "", errors.New("invalid jwt")
	}
	// retrieve ID
	var j jwt
	if err := json.Unmarshal(b, &j); err != nil {
		return "", err
	}
	return j.ID, nil
}
