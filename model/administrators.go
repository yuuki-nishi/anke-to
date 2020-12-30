package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Administrators administratorsテーブルの構造体
type Administrators struct {
	QuestionnaireID int    `sql:"type:int(11);not null;primary_key;"`
	UserTraqid      string `sql:"type:char(32);not null;primary_key;"`
}

// InsertAdministrators アンケートの管理者を追加
func InsertAdministrators(questionnaireID int, administrators []string) error {
	var administrator Administrators
	var err error
	for _, v := range administrators {
		administrator = Administrators{
			QuestionnaireID: questionnaireID,
			UserTraqid:      v,
		}
		err = db.Create(&administrator).Error
		if err != nil {
			return fmt.Errorf("failed to insert administrators: %w", err)
		}
	}
	return nil
}

// DeleteAdministrators アンケートの管理者の削除
func DeleteAdministrators(questionnaireID int) error {
	err := db.
		Where("questionnaire_id = ?", questionnaireID).
		Delete(Administrators{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete administrators: %w", err)
	}
	return nil
}

// GetAdminQuestionnaireIDs 自分が管理者のアンケートの取得
func GetAdminQuestionnaireIDs(user string) ([]int, error) {
	questionnaireIDs := []int{}
	err := db.
		Model(&Administrators{}).
		Where("user_traqid = ?", user).
		Or("user_traqid = ?", "traP").
		Pluck("DISTINCT questionnaire_id", &questionnaireIDs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get questionnaire_id: %w", err)
	}
	return questionnaireIDs, nil
}

// CheckQuestionnaireAdmin 自分がアンケートの管理者か判定
func CheckQuestionnaireAdmin(userID string, questionnaireID int) (bool, error) {
	err := db.
		Where("user_traqid = ? AND questionnaire_id = ?", userID, questionnaireID).
		Find(&Administrators{}).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to get a administrator: %w", err)
	}

	return true, nil
}
