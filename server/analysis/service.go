package analysis

import (
	"fmt"
	"strings"
	"time"

	"github.com/shahar05/cash-flow-viewer/core"
)

func GetPieByDate(d *core.DateRange) ([]CategorySum, error) {

	dateQuery := ""
	dateArgs := []interface{}{}
	if d != nil {
		dateQuery = `WHERE t.date BETWEEN $1 AND $2`
		dateArgs = append(dateArgs, d.StartDate)
		dateArgs = append(dateArgs, d.EndDate)
	}

	query := fmt.Sprintf(` SELECT c.name, SUM(t.amount)
						   FROM transactions t
						   JOIN categories c ON t.category_id = c.id
						   %s
						   GROUP BY c.name`, dateQuery)

	rows, err := DB.Query(query, dateArgs...) // , params.StartDate, params.EndDate
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var response []CategorySum

	// Iterate over the result set
	for rows.Next() {
		var categorySum CategorySum
		if err := rows.Scan(&categorySum.Name, &categorySum.Sum); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		response = append(response, categorySum)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return response, nil
}

func GetMetric(date *core.DateRange, category string, timePeriodStr string) (*core.Metric, error) {

	// Validate timePeriodStr using the helper function
	timePeriod, err := core.ValidateTimePeriod(timePeriodStr)
	if err != nil {
		timePeriod = core.Ptr(core.Month) // Set default TimePeriod
	}

	// Initialize query arguments
	queryArgs := []interface{}{*timePeriod}

	// Conditionally add date range and category conditions
	var conditions []string
	if date != nil {
		conditions = append(conditions, "t.date BETWEEN $2 AND $3")
		queryArgs = append(queryArgs, date.StartDate, date.EndDate)
	}

	// Correct index for category condition
	if strings.TrimSpace(category) != "" {
		// Adjust the argument position dynamically based on existing query args
		categoryIndex := len(queryArgs) + 1 // +1 because we start with timePeriod at $1
		conditions = append(conditions, fmt.Sprintf("c.name = $%d", categoryIndex))
		queryArgs = append(queryArgs, category)
	}

	// Build the WHERE clause based on conditions
	whereStr := ""
	if len(conditions) > 0 {
		whereStr = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Build the query using fmt.Sprintf for clarity
	query := fmt.Sprintf(`
		SELECT 
			DATE_TRUNC($1, t.date) AS time_period, 
			SUM(t.amount) AS total 
		FROM transactions t
		JOIN categories c ON t.category_id = c.id
		%s
		GROUP BY time_period
		ORDER BY time_period;
	`, whereStr)

	// Execute the query
	rows, err := DB.Query(query, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the query results
	var result []core.AggregatedValue
	for rows.Next() {
		var period time.Time
		var total float64
		if err := rows.Scan(&period, &total); err != nil {
			return nil, err
		}
		result = append(result, core.AggregatedValue{
			Timestamp: period,
			Value:     total,
		})
	}

	// Check if any error occurred during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the result
	return &core.Metric{
		Period:         *timePeriod,
		AggregatedData: result,
	}, nil
}

// GetMonthlyTransactions function retrieves monthly transaction totals
// func GetMonthlyTransactions() ([]MonthlyTransaction, error) {
// 	query := `
// 		SELECT DATE_TRUNC('month', date) AS month,
// 		       SUM(amount) AS total
// 		FROM transactions
// 		GROUP BY month
// 		ORDER BY month;`

// 	rows, err := DB.Query(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute query: %w", err)
// 	}
// 	defer rows.Close()

// 	var monthlyTransactions []MonthlyTransaction

// 	for rows.Next() {
// 		var mt MonthlyTransaction
// 		if err := rows.Scan(&mt.Month, &mt.Total); err != nil {
// 			return nil, fmt.Errorf("failed to scan row: %w", err)
// 		}
// 		monthlyTransactions = append(monthlyTransactions, mt)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("error during row iteration: %w", err)
// 	}

// 	return monthlyTransactions, nil
// }
