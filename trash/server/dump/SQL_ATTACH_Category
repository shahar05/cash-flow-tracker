Your approach using `UNION` can work, but you don't need to overcomplicate it. The `UNION` operator is used when you want to combine results from two or more `SELECT` statements, which makes sense here since you're querying from both `transactions` and `category_links`. However, it's worth noting that `UNION` removes duplicates, and since you're using `LIMIT 1`, this shouldn't be a major issue.

### Suggestion:

The idea is to check both `transactions` and `category_links` tables for a matching `category_id` based on either `name` or `international_branch_id`. Here's how you can structure your query:

```sql
(
    SELECT category_id 
    FROM transactions 
    WHERE name = $1 OR international_branch_id = $2 
    LIMIT 1
)
UNION
(
    SELECT category_id 
    FROM category_links 
    WHERE international_branch_id = $2 
    LIMIT 1
)
LIMIT 1;
```

### Explanation:
- The first `SELECT` looks for a `category_id` in the `transactions` table based on `name` or `international_branch_id`.
- The second `SELECT` looks for a `category_id` in the `category_links` table where `international_branch_id` matches.
- The `LIMIT 1` after the `UNION` ensures that you get only one result, even if both queries return rows.

This query efficiently finds the first matching `category_id` from either the `transactions` or `category_links` table, based on `name` or `international_branch_id`.

### Alternative Option:
You could also simplify the logic by prioritizing `transactions` and falling back on `category_links` only if there is no result from `transactions`:

```sql
WITH first_match AS (
    SELECT category_id 
    FROM transactions 
    WHERE name = $1 OR international_branch_id = $2 
    LIMIT 1
)
SELECT category_id
FROM first_match
UNION ALL
SELECT category_id 
FROM category_links 
WHERE international_branch_id = $2
AND NOT EXISTS (SELECT 1 FROM first_match)
LIMIT 1;
```

This way, the query first tries to find a match in `transactions`, and if it doesn’t, it checks `category_links`.