# New DB

This is the first step in run main. 

```go
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to db")
		return err
	}
```

## database connection

A database connection is returned to Run()

`db` contains a pointer to a `sql.DB` struct.

and the db package has a `NewDatabase()` function that returns a `*sql.DB` and an `error`

```go
func NewDatabase() (*Database, error) {
	ConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	fmt.Printf("connection string: %s\n", ConnectionString)
	// comments-api  | connection string:
	// host=db port=5432 user= dbname= password=postgres sslmode=disable
	db, err := sqlx.Connect("postgres", ConnectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to db : %s", err)
	}

	return &Database{Client: db}, nil

}
```

# New Service

the second step in run main is to create a new service

```go
	cmtService := comment.NewService(db)
```

it has passed to it `db` which is a pointer to a `sql.DB` struct

the comment service has an interface
    
```go
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}
```

so that anytyhing that implements the interface can be passed to it

# Get service record

the third step in run main is to get a record from the service, using both the service record and the db instance that was passed to it

```go
	cm, err := cmtService.GetComment(
		context.Background(),
		"602d7576-bc5c-11ee-8e50-00155dd54b61",
	)
```

## service and db get record

Both the db and the comment services have a `GetComment()` method that accept a context and a string and return a comment and an error

the comment service has 

```go
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("returning comment from service")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return cmt, ErrFetchingComment
	}

	return cmt, nil
}
```

and it then callse the `GetComment()` method on the store that was passed to it as `s.Store.GetComment()`

```go
func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching the comment by uuid: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
```

# Conclusion

The `Run()` function in the main package is the entry point to the application. It is responsible for creating a new database connection, creating a new comment service, and getting a comment record from the database.

this pattern is used so that the implementation of the database can be changed without affecting the rest of the application, so long as the new database implementation implements the `Store` interface that is used by the comment service and has the same signature for the `GetComment()` method.

Also, the database implementation is kept separate from the comment service, so that the comment service can be tested without having to connect to a database.

A mock database implementation can be created that implements the `Store` interface and used to test the comment service.