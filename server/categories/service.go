package categories

import (
	"fmt"
	"strings"
)

// GetCategories retrieves all categories from the database
func GetCategories() ([]Category, error) {
	// SQL query to select all categories
	query := "SELECT id, name FROM categories"

	// Execute the query
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetCategories: %v", err)
	}
	defer rows.Close()

	var categories []Category

	// Iterate over the result set
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, fmt.Errorf("GetCategories: %v", err)
		}
		categories = append(categories, category)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCategories: %v", err)
	}

	return categories, nil
}

// AddCategories inserts multiple categories and returns the count of newly inserted categories.
func AddCategories(categories []Category) (int64, error) {
	if len(categories) == 0 {
		return 0, nil // No categories to insert, so return early.
	}

	// Base query string
	query := "INSERT INTO categories (name) VALUES "

	// Prepare value placeholders and arguments slice
	values := []string{}
	args := []interface{}{}

	// Generate placeholders for each value
	for i, category := range categories {
		values = append(values, fmt.Sprintf("($%d)", i+1))
		args = append(args, category.Name)
	}

	// Join all value placeholders into the query
	query += strings.Join(values, ", ")

	// Add conflict handling to avoid duplicates
	query += " ON CONFLICT (name) DO NOTHING RETURNING id;"

	// Execute the query and return the count of inserted rows
	rows, err := DB.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("AddCategories: failed to insert categories: %v", err)
	}
	defer rows.Close()

	// Count the number of rows successfully inserted
	var count int64
	for rows.Next() {
		count++
	}

	return count, nil
}
