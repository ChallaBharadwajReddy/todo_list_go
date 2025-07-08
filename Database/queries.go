package database

// import(
// 	"github.com/gin-gonic/gin"
// )

func Get_list() ([]Todo, error) {
	rows, err := Db.Query("SELECT id, todo , deleted FROM todos WHERE deleted=FALSE")
	var todos []Todo
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Todo, &todo.Deleted); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func Insert_todo(body Todo) error {
	_, err := Db.Exec("INSERT INTO todos (id, todo) VALUES ($1, $2)", body.Id, body.Todo)

	return err
}

func Edit_todo(body Todo) (int, error) {
	res, err := Db.Exec("UPDATE todos SET todo=$1 WHERE id=$2", body.Todo, body.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()

	return int(rowsAffected), err
}

func Delete_todo(id string) (int, error) {
	res, err := Db.Exec("UPDATE todos SET deleted=TRUE WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()

	return int(rowsAffected), err
}
