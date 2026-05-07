Piggy is a personal finance management application.

It is a web application that allows users to track their income and expenses, set budgets, and monitor their savings goals.

This app does the following:
1. Allows users to create an account
2. Allows users to login
3. Allows users to logout
4. Allows users to add income
5. Allows users to add expenses
6. Allows users to view their income and expenses
7. Allows users to view their savings goals
8. Allows users to add savings goals
9. Allows users to view their budgets
10. Allows users to add budgets
11. Allows users to view their savings goals
12. Allows users to add savings goals

docker exec -it backend-db-1 psql -U piggy -d piggydb 


The Code Explained Line-by-Line
SQL

WHERE (CAST(@type_filter AS TEXT) = '' OR type = @type_filter)

1. @type_filter

This is a Named Parameter. When you run sqlc generate, sqlc will look at this and create a Go function that asks for a variable called typeFilter (a string).
2. CAST(@type_filter AS TEXT) = ''

This checks if the string coming from your frontend is empty.

    If it's empty: This whole part becomes TRUE. In SQL, if the WHERE clause is TRUE, it returns every row. This is how we handle the "Show All" case.
