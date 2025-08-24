# 📌 Project Plan: Personal Task Manager (CLI App in Go)

A command-line application to manage tasks (like TODOs), with file-based persistence, concurrency features, and modular design.

---

## 🎯 Goals

* Strengthen understanding of Go basics (structs, interfaces, slices, maps).
* Practice concurrency (goroutines, channels, worker pools).
* Work with file I/O (JSON, CSV, or custom format).
* Explore Go modules, packages, error handling, and testing.
* Implement a CLI with commands and subcommands.
* Gain experience with dependency management and project organization.

---

## 🏗️ Features & Milestones

### **Phase 1: Core Task Management**

* [ ] Define a `Task` struct (`ID`, `Title`, `Description`, `Status`, `CreatedAt`, `DueDate`).
* [ ] Store tasks in memory (slice/map).
* [ ] Implement CRUD operations:

  * `add` a task
  * `list` tasks
  * `update` task status
  * `delete` a task

### **Phase 2: File Persistence**

* [ ] Save tasks to a JSON file on disk.
* [ ] Load tasks from file at startup.
* [ ] Implement `export` (CSV) and `import` commands.

### **Phase 3: CLI Interaction**

* [ ] Implement a text-based menu system or command-line flags (`task add "Buy milk"`).
* [ ] Support filtering tasks by `status`, `due date`, etc.
* [ ] Pretty-print tasks in tabular format.

### **Phase 4: Concurrency**

* [ ] Use goroutines + channels for:

  * Running background reminders for due tasks.
  * Parallel file save/load (simulate worker pool).
* [ ] Add a "bulk import" that processes multiple files concurrently.

### **Phase 5: Enhancements**

* [ ] Add task categories/tags.
* [ ] Add priority levels (low/medium/high).
* [ ] Support undo/redo (stack-based approach).
* [ ] Optional encryption for saved files (practice with `crypto`).

## 🧪 Testing Plan

* Unit tests for:

  * Adding/updating/deleting tasks
  * File read/write
  * Import/export
* Integration tests:

  * CLI command flow
  * Concurrency tests (timeouts, reminders)
