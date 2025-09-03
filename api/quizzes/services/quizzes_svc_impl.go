package services

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"ragamaya-api/api/quizzes/dto"
	"ragamaya-api/api/quizzes/repositories"
	"ragamaya-api/models"
	"ragamaya-api/pkg/config"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/helpers"
	"ragamaya-api/pkg/logger"
	"ragamaya-api/pkg/mapper"
	"time"

	storageDTO "ragamaya-api/api/storages/dto"
	storageService "ragamaya-api/api/storages/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate

	storageService storageService.CompServices
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate, storageService storageService.CompServices) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,

		storageService: storageService,
	}
}

func (s *CompServicesImpl) FindAllCategories(ctx *gin.Context) ([]dto.CategoryRes, *exceptions.Exception) {
	data, err := s.repo.FindAllCategories(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var result []dto.CategoryRes

	for _, v := range data {
		result = append(result, mapper.MapQuizCategoryMTO(v))
	}

	return result, nil
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.QuizReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCategory, err := s.repo.FindCategoryByName(ctx, tx, data.Category)
	if err != nil {
		if err.Status == http.StatusNotFound && existCategory == nil {
			existCategory = &models.QuizCategory{
				UUID: uuid.NewString(),
				Name: data.Category,
			}
			err = s.repo.CreateCategory(ctx, tx, *existCategory)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	input := mapper.MapQuizITM(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.SlugifyUnique(data.Title)
	input.Questions = helpers.FormatToJSON(data.Questions)
	input.CategoryUUID = existCategory.UUID

	err = s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) Search(ctx *gin.Context, data dto.SearchReq) ([]dto.QuizRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	result, err := s.repo.Search(ctx, s.DB, data)
	if err != nil {
		return nil, err
	}

	var output []dto.QuizRes

	for _, v := range result {
		output = append(output, mapper.MapQuizMTO(v))
	}

	return output, nil
}

func (s *CompServicesImpl) FindBySlug(ctx *gin.Context, slug string) (*dto.QuizPublicDetailRes, *exceptions.Exception) {
	result, err := s.repo.FindBySlug(ctx, s.DB, slug)
	if err != nil {
		return nil, err
	}

	output := mapper.MapQuizMTPDO(*result)
	return &output, nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.QuizDetailRes, *exceptions.Exception) {
	result, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}

	output := mapper.MapQuizMTDO(*result)
	return &output, nil
}

func (s *CompServicesImpl) Analyze(ctx *gin.Context, quizUUID string, data dto.AnalyzeReq) (*dto.AnalyzeRes, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		return nil, err
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCertificate, err := s.repo.FindCertificateByQuizUUIDandUserUUID(ctx, tx, quizUUID, userData.UUID)
	if err != nil {
		if err.Status != http.StatusNotFound {
			return nil, err
		}
	}

	if existCertificate != nil {
		return nil, exceptions.NewValidationException(fmt.Errorf("quiz already completed"))
	}

	quizData, err := s.repo.FindByUUID(ctx, tx, quizUUID)
	if err != nil {
		return nil, err
	}

	quiz := mapper.MapQuizMTDO(*quizData)

	if len(data.Answers) != quiz.TotalQuestions {
		return nil, exceptions.NewValidationException(fmt.Errorf("invalid answers length"))
	}

	correct := 0
	for i, q := range quiz.Questions {
		if data.Answers[i] == q.AnswerIndex {
			correct++
		}
	}

	var result dto.AnalyzeRes
	result.Score = (float32(correct) / float32(quiz.TotalQuestions)) * 100
	result.Status = dto.Failed
	if int(result.Score) >= quiz.MinimumScore {
		result.Status = dto.Success
	}

	err = s.repo.CreateAttempt(ctx, tx, models.QuizAttempt{
		QuizUUID: quizUUID,
		UserUUID: userData.UUID,
		Score:    result.Score,
		Status:   string(result.Status),
	})
	if err != nil {
		return nil, err
	}

	if result.Status == dto.Success {
		certificateData := &models.QuizCertificate{
			UUID:     uuid.NewString(),
			QuizUUID: quizUUID,
			UserUUID: userData.UUID,
			Score:    result.Score,
		}

		certificate, err := s.GenerateCertificate(dto.CertificateReq{
			UUID:     certificateData.UUID,
			UserName: userData.Name,
			QuizName: quiz.Title,
			Date:     time.Now().Format("2006-01-02"),
			Score:    result.Score,
			QRData:   config.GetFrontendBaseURL(),
		})
		if err != nil {
			return nil, err
		}

		uploadData, err := s.storageService.Create(ctx, storageDTO.FilesInput{
			OriginalFileName: `quiz-certificate-` + certificateData.UUID + `.pdf`,
			FileBuffer:       *certificate,
			Size:             helpers.FormatFileSize(int64(len(*certificate))),
			Extension:        `pdf`,
			MimeType:         `application/pdf`,
			MimeSubType:      `pdf`,
			Meta:             `{}`,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		certificateData.CertificateURL = uploadData.PublicURL
		err = s.repo.CreateCertificate(ctx, tx, *certificateData)
		if err != nil {
			return nil, err
		}

		certificateData, err = s.repo.FindCertificateByUUID(ctx, tx, certificateData.UUID)
		if err != nil {
			return nil, err
		}

		resultCertificate := mapper.MapCertificateMTO(*certificateData)
		result.Certificate = &resultCertificate
	}

	return &result, nil
}

func (s *CompServicesImpl) GenerateCertificate(data dto.CertificateReq) (*[]byte, *exceptions.Exception) {
	tmpl, err := template.ParseFiles(`static/templates/quiz_certificate.html`)
	if err != nil {
		return nil, exceptions.NewException(500, fmt.Sprintf("template parse error: %s", err))
	}

	var bufHTML bytes.Buffer
	if err := tmpl.Execute(&bufHTML, data); err != nil {
		logger.Error("template execute error: %v", err)
		return nil, exceptions.NewException(500, exceptions.ErrInternalServer)
	}

	tempFile, err := os.CreateTemp("", "cert_*.html")
	if err != nil {
		logger.Error("generate certificate error: %v", err)
		return nil, exceptions.NewException(500, exceptions.ErrInternalServer)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(bufHTML.Bytes()); err != nil {
		logger.Error("generate certificate error: %v", err)
		return nil, exceptions.NewException(500, exceptions.ErrInternalServer)
	}
	tempFile.Close()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Headless,
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-features", "TranslateUI"),
		chromedp.Flag("disable-ipc-flooding-protection", true),
	)

	chromePath := os.Getenv("CHROME_PATH")
	if chromePath == "" {
		chromePath = os.Getenv("CHROME_BIN")
	}
	if chromePath == "" {
		possiblePaths := []string{
			"/usr/bin/chromium-browser",
			"/usr/bin/chromium",
			"/usr/bin/google-chrome-stable",
			"/usr/bin/google-chrome",
		}
		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				chromePath = path
				break
			}
		}
	}

	if chromePath != "" {
		opts = append(opts, chromedp.ExecPath(chromePath))
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var pdfBuf []byte
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("file://" + tempFile.Name()),
		chromedp.WaitReady("body"),
		chromedp.Sleep(500 * time.Millisecond),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(11.7).
				WithPaperHeight(8.3).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				WithScale(1.0).
				Do(ctx)
			return err
		}),
	})

	if err != nil {
		logger.Error("generate certificate error: %v", err)
		return nil, exceptions.NewException(500, exceptions.ErrInternalServer)
	}

	if len(pdfBuf) == 0 {
		logger.Error("generate certificate error: empty PDF buffer")
		return nil, exceptions.NewException(500, exceptions.ErrInternalServer)
	}

	return &pdfBuf, nil
}

func (s *CompServicesImpl) Update(ctx *gin.Context, data dto.QuizUpdateReq) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	existCategory, err := s.repo.FindCategoryByName(ctx, tx, data.Category)
	if err != nil {
		if err.Status == http.StatusNotFound && existCategory == nil {
			existCategory = &models.QuizCategory{
				UUID: uuid.NewString(),
				Name: data.Category,
			}
			err = s.repo.CreateCategory(ctx, tx, *existCategory)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	input := mapper.MapQuizUTM(data)
	input.Slug = helpers.SlugifyUnique(data.Title)
	input.Questions = helpers.FormatToJSON(data.Questions)
	input.CategoryUUID = existCategory.UUID

	err = s.repo.Update(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) Delete(ctx *gin.Context, uuid string) *exceptions.Exception {
	return s.repo.Delete(ctx, s.DB, uuid)
}
