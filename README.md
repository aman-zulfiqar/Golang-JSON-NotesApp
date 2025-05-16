📝 JSON-NotesApp-Golang
JSON-NotesApp-Golang is a lightweight and efficient command-line application written in Go for managing personal notes. This tool allows users to add, list, and delete notes, storing them in a structured JSON file. It utilizes the sirupsen/logrus library for structured and leveled logging, enhancing the application's maintainability and debuggability.

📋 Table of Contents
Features

Prerequisites

Usage

Project Structure

Logging

License

Acknowledgments

✨ Features
Add Notes: Prompt-based input for note title and content.

List Notes: Display all saved notes with timestamps.

Delete Notes: Remove notes by specifying the title.

Persistent Storage: Notes are stored in a notes.json file.

Structured Logging: Utilizes logrus for informative logging.

🛠️ Prerequisites
Go installed on your system (version 1.16 or later).

🚀 Usage
Run the application using the following command:

bash
Copy
Edit
go run main.go
Upon execution, you'll be presented with a menu:

markdown
Copy
Edit
Notes Application
1. Add Note
2. List Notes
3. Delete Note
4. Exit
Choose an option:
Follow the on-screen prompts to interact with the application.

📁 Project Structure
main.go: Main application file containing all functionalities.

notes.json: JSON file used for storing notes persistently.

📊 Logging
The application uses the sirupsen/logrus library for structured logging. This provides leveled logging (Info, Warning, Error) and includes timestamps for better traceability.

🙌 Acknowledgments
Inspired by the need for a simple and efficient way to manage notes directly from the command line.
