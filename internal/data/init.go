package data

import (
	"Shiro/internal/config"
	"Shiro/internal/tools"
	"database/sql"
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
)

func SetupDB(app *iris.Application) *sql.DB {
	// 连接到MySQL服务器
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Application.Database.User, config.Application.Database.Password, config.Application.Database.Server, config.Application.Database.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		app.Logger().Fatal("Failed to setup data: %s", err.Error())
		return nil
	}

	// 选择数据库
	SelectDatabaseSQL := fmt.Sprintf("USE `%s`", config.Application.Database.Dbname)
	if _, err := db.Exec(SelectDatabaseSQL); err != nil {
		app.Logger().Fatal("failed to select data: %w", err)
	}

	if !tools.IsFileExists("install.lock") {
		initializeTable(app, db)
	}

	app.Logger().Info("The data table has been successfully initialized")
	return db
}

func initializeTable(app *iris.Application, db *sql.DB) {
	// 创建 users 表
	usersSQL := `CREATE TABLE IF NOT EXISTS users (
        username VARCHAR(255) PRIMARY KEY,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(usersSQL); err != nil {
		app.Logger().Fatal("error creating users table: %w", err)
	}

	// 创建 categories 表
	categoriesSQL := `CREATE TABLE IF NOT EXISTS categories (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(categoriesSQL); err != nil {
		app.Logger().Fatal("error creating categories table: %w", err)
	}

	// 创建 posts 表
	postsSQL := `CREATE TABLE IF NOT EXISTS posts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        author_id VARCHAR(255) NOT NULL,
        category_id INT NOT NULL,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        status VARCHAR(20) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (author_id) REFERENCES users(username),
        FOREIGN KEY (category_id) REFERENCES categories(id)
    );`
	if _, err := db.Exec(postsSQL); err != nil {
		app.Logger().Fatal("error creating posts table: %w", err)
	}

	// 创建 comments 表
	commentsSQL := `CREATE TABLE IF NOT EXISTS comments (
        id INT AUTO_INCREMENT PRIMARY KEY,
        post_id INT NOT NULL,
        author VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        status VARCHAR(20) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (post_id) REFERENCES posts(id)
    );`
	if _, err := db.Exec(commentsSQL); err != nil {
		app.Logger().Fatal("error creating comments table: %w", err)
	}

	// 创建 tags 表
	tagsSQL := `CREATE TABLE IF NOT EXISTS tags (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(tagsSQL); err != nil {
		app.Logger().Fatal("error creating tags table: %w", err)
	}

	// 创建 posts_tags 关联表
	postsTagsSQL := `CREATE TABLE IF NOT EXISTS posts_tags (
        post_id INT NOT NULL,
        tag_id INT NOT NULL,
        PRIMARY KEY (post_id, tag_id),
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
    );`
	if _, err := db.Exec(postsTagsSQL); err != nil {
		app.Logger().Fatal("error creating posts_tags table: %w", err)
	}

	// 创建 files 表
	filesSQL := `CREATE TABLE IF NOT EXISTS files (
        id INT AUTO_INCREMENT PRIMARY KEY,
        post_id INT,
        file_name VARCHAR(255) NOT NULL,
        file_path VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE SET NULL
    );`
	if _, err := db.Exec(filesSQL); err != nil {
		app.Logger().Fatal("error creating files table: %w", err)
	}

	// 填入默认的内容
	insertDefaultUser(app, db)
	insertDefaultCategory(app, db)
	insertDefaultPost(app, db)

	// 创建名为install.lock的空文件
	file, err := os.Create("install.lock")
	if err != nil {
		// 如果创建文件时发生错误，使用log包打印错误并退出
		app.Logger().Fatal(err)
	}
	// 关闭文件句柄
	defer file.Close()
}

// 插入一个默认的user
func insertDefaultUser(app *iris.Application, db *sql.DB) {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	result, err := db.Exec(query, "Shiro", "Shiroloveyou", "Shiro@example.com")
	if err != nil {
		app.Logger().Fatal("error executing insert user query: %w", err)
		return
	}
	rowAffected, err := result.RowsAffected()
	if err != nil || rowAffected < 1 {
		app.Logger().Fatal("error getting last insert for user: %w", err)
		return
	}

	app.Logger().Info("Inserted default user successfully.")
}

// 插入一个默认的category
func insertDefaultCategory(app *iris.Application, db *sql.DB) int64 {
	query := `INSERT INTO categories (name, description) VALUES (?, ?)`
	result, err := db.Exec(query, "default_category", "Default category description")
	if err != nil {
		app.Logger().Fatal("error executing insert category query: %w", err)
		return 0
	}
	categoryID, err := result.LastInsertId()
	if err != nil {
		app.Logger().Fatal("error getting last insert ID for category: %w", err)
		return 0
	}

	app.Logger().Info("Inserted default category successfully with ID:", categoryID)
	return categoryID
}

// 插入一个默认的post
func insertDefaultPost(app *iris.Application, db *sql.DB) {
	// 定义插入语句
	// 假设 author_id 和 category_id 为 1, 根据实际情况调整
	query := `INSERT INTO posts (author_id, category_id, title, content, status) VALUES (?, ?, ?, ?, ?)`
	// 使用预定义的值执行插入
	_, err := db.Exec(query, "Shiro", 1, "Hello World", "Default content of the post", "draft")
	if err != nil {
		app.Logger().Fatal("error executing insert query: %w", err)
	}

	app.Logger().Info("Inserted default post successfully")
}
