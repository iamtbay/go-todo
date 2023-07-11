package database

import (
	"context"
	"errors"
	"fmt"
)

func (t *Todo) CreateNewTodo(todo *Todo) error {
	fmt.Println("New todo = ", todo)
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
			INSERT INTO 
				"todos" 
					("userId", "title", "body", "questTime") 
			VALUES ($1, $2, $3, $4);
			`
	_, err := dbConn.DB.ExecContext(ctx, query, todo.UserId, todo.Title, todo.Body, todo.QuestTime)
	if err != nil {
		return err
	}
	return nil
}

// GET TODOS
func (t *Todo) GetTodos(id string) ([]Todo, error) {
	var AllTodos []Todo
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
				SELECT * FROM "todos" WHERE "userId" = $1
			`
	rows, err := dbConn.DB.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println("error while querying", err)
		return nil, err
	}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.Id,
			&todo.UserId,
			&todo.Title,
			&todo.Body,
			&todo.IsCompleted,
			&todo.QuestTime,
			&todo.CreatedAt,
		)
		if err != nil {
			fmt.Println("error while scaning", err)
			return nil, err
		}

		AllTodos = append(AllTodos, todo)
	}
	return AllTodos, nil

}
//GET SINGLE TODO
func (t *Todo) GetSingleTodo(id, userId string) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM todos WHERE id = $1 AND "userId" = $2`
	row := dbConn.DB.QueryRowContext(ctx, query, id, userId)
	var newTodo Todo
	err := row.Scan(
		&newTodo.Id,
		&newTodo.UserId,
		&newTodo.Title,
		&newTodo.Body,
		&newTodo.IsCompleted,
		&newTodo.QuestTime,
		&newTodo.CreatedAt,
	)
	if err != nil {
		fmt.Println("Error while scaning", err)
		return nil, err
	}
	return &newTodo, nil
}
func (t *Todo) MarkAsCompleted(id, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT "isCompleted" FROM todos WHERE id=$1 AND "userId"=$2`
	var isCompleted bool
	row := dbConn.DB.QueryRowContext(ctx, query, id, userId)
	err := row.Scan(&isCompleted)
	if err != nil {
		return err
	}
	fmt.Println(isCompleted)
	query = ` UPDATE todos 
				SET "isCompleted"=$1
				WHERE
				id=$2 AND "userId"=$3
				`
	rows, err := dbConn.DB.ExecContext(ctx, query, !isCompleted, id, userId)
	if err != nil {
		return err
	}
	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows < 1 {
		return errors.New("unvalid credentials")
	}
	return nil
}
func (t *Todo) UpdateATodo(todo *Todo, id, userId string) (*Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
			UPDATE todos
			SET title=$1, body=$2, "questTime"=$3
			WHERE
			id=$4 AND "userId" = $5
			`
	_, err := dbConn.DB.ExecContext(ctx, query, todo.Title, todo.Body, todo.QuestTime, id, userId)
	if err != nil {
		return nil, err
	}
	return nil, nil

}

//DELETE TODO !
func (t *Todo) DeleteATodo(id, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
			DELETE FROM todos WHERE id = $1 AND "userId" = $2
			`

	rows, err := dbConn.DB.ExecContext(ctx, query, id, userId)
	if err != nil {
		return err
	}
	affected, err := rows.RowsAffected()
	if affected < 1 {
		return errors.New("unvalid credentials")
	}
	if err != nil {
		fmt.Println(err)
		return errors.New("errors while affecting")
	}
	return nil

}
