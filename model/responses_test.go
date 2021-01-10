package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v3"
)

func TestInsertResponses(t *testing.T) {
	t.Parallel()

	assertion := assert.New(t)
	_, questionID, responseID := insertTestResponses(t)

	type args struct {
		responseID    int
		responseMetas []*ResponseMeta
	}
	type expect struct {
		isErr bool
		err   error
	}

	type test struct {
		description string
		args
		expect
	}

	testCases := []test{
		{
			description: "valid",
			args: args{
				responseID: responseID,
				responseMetas: []*ResponseMeta{
					{QuestionID: questionID, Data: "リマインダーBOTを作った話"},
				},
			},
		},
		{
			description: "responseID not exist",
			args: args{
				responseID: -1,
				responseMetas: []*ResponseMeta{
					{QuestionID: questionID, Data: "リマインダーBOTを作った話"},
				},
			},
			expect: expect{
				isErr: true,
			},
		},
	}
	for _, testCase := range testCases {
		err := InsertResponses(testCase.args.responseID, testCase.args.responseMetas)

		if !testCase.expect.isErr {
			assertion.NoError(err, testCase.description, "no error")
		} else if testCase.expect.err != nil {
			assertion.EqualError(err, testCase.expect.err.Error(), testCase.description, "error")
		}
		if err != nil {
			continue
		}

		response := Response{}
		err = db.Where("response_id = ?", responseID).First(&response).Error
		if err != nil {
			t.Errorf("failed to get questionnaire(%s): %w", testCase.description, err)
		}

		assertion.Equal(testCase.args.responseID, response.ResponseID, testCase.description, "responseID")
		assertion.Equal(questionID, response.QuestionID, testCase.description, "questionID")
		assertion.Equal(testCase.args.responseMetas[0].Data, response.Body.ValueOrZero(), testCase.description, "Body")
		assertion.WithinDuration(time.Now(), response.ModifiedAt, 2*time.Second, testCase.description, "ModifiedAt")
		assertion.Equal(time.Time{}, response.DeletedAt.ValueOrZero(), 2*time.Second, testCase.description, "DeletedAt")
	}
}

func insertTestResponses(t *testing.T) (int, int, int) {
	questionnaireID, err := InsertQuestionnaire("第1回集会らん☆ぷろ募集アンケート", "第1回メンバー集会でのらん☆ぷろで発表したい人を募集します らん☆ぷろで発表したい人あつまれー！", null.NewTime(time.Now(), false), "public")
	require.NoError(t, err)

	err = InsertAdministrators(questionnaireID, []string{userOne})
	require.NoError(t, err)

	questionID, err := InsertQuestion(questionnaireID, 1, 1, "Text", "質問文", true)
	require.NoError(t, err)

	responseID, err := InsertRespondent(userTwo, questionnaireID, null.NewTime(time.Now(), true))
	require.NoError(t, err)
	return questionnaireID, questionID, responseID
}
