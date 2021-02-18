package schedule_generator

import (
	"database/sql"

	"github.com/jorensjongers/scheduler/backend/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type DbController struct {
	db *sql.DB
}

func newDBController() DbController {
	db := createDB()
	return DbController{db}
}

// GetModelParameters -- returns the model parameters a stored in the DB
func (c DbController) GetModelParameters() (model.ModelParameters, error) {

	balanceMinimum, err := c.getMinBalanceScore()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting min balance score from db")
	}

	shiftTypeParams, err := c.getShiftTypeParams()
	if err != nil {
		return model.ModelParameters{}, errors.Wrap(err, "failed getting shift type parameters from db")
	}

	result := model.ModelParameters{
		BalanceMinimum:  int32(balanceMinimum),
		ShiftTypeParams: shiftTypeParams,
	}

	return result, nil
}

// SetModelParameters -- updates the model parameters in teh DB
func (c DbController) SetModelParameters() error {

	return nil
}

func (c DbController) getMinBalanceScore() (int, error) {

	minBalanceScoreQuery := `
		SELECT score
		FROM min_balance_score
		WHERE id = 1
	`

	var score int
	if err := c.db.QueryRow(minBalanceScoreQuery).Scan(&score); err != nil {
		return 0, err
	}

	return score, nil
}

func (c DbController) getShiftTypeParams() ([]model.ShiftTypeModelParameters, error) {

	shiftTypeParamsQuery := `
		SELECT shift_type, fairness_weight, included_in_balance
		FROM shift_type_params
	`

	rows, err := c.db.Query(shiftTypeParamsQuery)
	if err != nil {
		return []model.ShiftTypeModelParameters{}, err
	}

	result := []model.ShiftTypeModelParameters{}
	for rows.Next() {
		var (
			shiftType         string
			fairnessWeight    float32
			includedInBalance bool
		)

		if err := rows.Scan(&shiftType, &fairnessWeight, &includedInBalance); err != nil {
			return []model.ShiftTypeModelParameters{}, err
		}

		stp := model.ShiftTypeModelParameters{
			ShiftType:         model.ShiftType(shiftType),
			FairnessWeight:    fairnessWeight,
			IncludedInBalance: includedInBalance,
		}

		result = append(result, stp)
	}

	if err := rows.Err(); err != nil {
		return []model.ShiftTypeModelParameters{}, err
	}

	return result, nil
}
