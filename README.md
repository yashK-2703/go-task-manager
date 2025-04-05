# 📘 Go Task Manager

A simple command-line task manager built in Go.  
Supports adding, listing, marking tasks as done, and deleting them — with JSON-based storage.

> 🚀 GitHub Repo: [yashK-2703/go-task-manager](https://github.com/yashK-2703/go-task-manager)

---

## 📂 Project Structure

```
go-task-manager/
├── main.go             # CLI entry point
├── task/
│   ├── task.go         # Task logic (add, delete, mark done)
│   └── task_test.go    # Unit tests for task package
├── data/
│   └── tasks.json      # Task storage file (auto-created)
└── README.md
```

---

## 🔧 Usage

### ⬇️ Build the CLI

```bash
go build -o taskmanager
```

### ➕ Add a Task

```bash
./taskmanager -add="Buy groceries"
```

### 📋 List All Tasks

```bash
./taskmanager -list
```

### ✅ Mark Task as Done

```bash
./taskmanager -done=2
```

### ❌ Delete a Task

```bash
./taskmanager -delete=2
```

### 📁 Use a Custom File

```bash
./taskmanager -file="mytasks.json" -add="Read Go docs"
```

---

## 🧪 Run Tests

From the root directory:

```bash
go test ./...
```

Or test specific function:

```bash
go test -v -run TestAddTask ./task
```

---

## 📦 Todo / Improvements

- [ ] Add due dates or priority
- [ ] Add search/filter CLI options
- [ ] Migrate to a database (optional)
- [ ] Export to CSV or markdown

---

## 🧑‍💻 Author

Made with ❤️ by [@yashK-2703](https://github.com/yashK-2703)

---

## 📄 License

MIT License

