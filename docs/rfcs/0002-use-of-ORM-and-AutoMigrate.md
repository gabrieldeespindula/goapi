# RFC: Use of ORM (GORM) and AutoMigrate for Database Management in the goapi Project

## Context

In the `goapi` project, we need an efficient, simple, and reliable way to manage the PostgreSQL database, including the creation and updating of tables as the application's data model evolves.

## Considered Options

1. **Raw SQL with manual migrations**
   - Writing SQL scripts manually to create and modify the database schema.
   - Using tools like `golang-migrate` to manage migration versions.
   - Offers greater control and transparency but requires writing and maintaining SQL scripts.

2. **Using an ORM (GORM) with automatic migrations (AutoMigrate)**
   - Mapping Go structs to database tables.
   - Using the `AutoMigrate` function to automatically create or modify tables based on the structs.
   - Reduces complexity and speeds up initial development with less direct control over the generated SQL.

3. **Using ORM with manually written migrations (e.g., Gormigrate)**
   - Combining ORM usage with more controlled migration files.
   - Still requires writing and managing migration files.

## Decision

We decided to use **GORM ORM** with the **AutoMigrate** method to manage the database schema.

### Justifications

- The project is in an early phase and requires agility in development.
- AutoMigrate reduces the need to write SQL manually, accelerating schema changes.
- It simplifies maintaining the codebase and the data model in Go.
- If needed in the future, we can introduce additional tools for versioned migrations (like Gormigrate or golang-migrate).

## Current Implementation

- The `db` package manages the database connection and exports the `*gorm.DB` object.
- At the application's entry point (`main.go`), after connecting to the database, we call `db.DB.AutoMigrate(&models.Record{})` to ensure tables are up to date with the `Record` model.
- The application uses GORM for all database operations, keeping the code simple and idiomatic in Go.

## Possible Next Steps

- Evaluate the adoption of versioned migrations for better control in production environments.
- Monitor AutoMigrate's behavior to avoid undesired changes in critical environments.
- Implement backup and rollback strategies to mitigate risks.
