package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

var (
	getSalaryQuery = sq.Select("assignment_id, NVL(salary,0) AS salary, date_from, fte,pay_basis_id").
		From(`APPS.XXPER_MTS_HR_5482_ASSIGNM_PV A
OUTER APPLY (
	SELECT MAX(change_date) AS last_change_salary_date
	FROM (
		SELECT
			change_date,
			proposed_salary_n - LAG(proposed_salary_n, 1, 0) OVER (ORDER BY change_date) AS sal_diff,
			LAG(proposed_salary_n, 1, 0) OVER (ORDER BY change_date) AS sal_prev
		FROM apps.xxper_mts_hr_5482_salaries_v sal
		WHERE sal.assignment_id = A.assignment_id
	) t
	WHERE t.sal_diff != 0 AND t.sal_prev != 0
) s`).PlaceholderFormat(sq.Colon)
)

func main() {
	assignments := []string{"403819", "641395"}
	orSt := make(sq.Or, 0, len(assignments))

	for k := range assignments {
		orSt = append(orSt, sq.Eq{"assignment_id": assignments[k]})
	}

	query, args, err := getSalaryQuery.Where(orSt).
		Where("effective_start_date <= SYSDATE").Where("effective_end_date >= SYSDATE").ToSql()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(query)
	fmt.Println(args)
}
