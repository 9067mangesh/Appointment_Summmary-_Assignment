# Appointment_Summmary-_Assignment
🩺 Appointment Summary - Assignment
📘 Project Overview
Appointment Summary Assignment is a simple backend project developed using Go (Golang) and SQLite3, designed to manage and summarize appointments between patients and doctors. It allows for efficient creation, tracking, and querying of appointments with associated patient and doctor details.

🛠️ Tech Stack
Language: Go (Golang)

Database: SQLite3

🚀 Features
Add and retrieve Patient, Doctor, and Appointment records

View appointment summaries with doctor-patient information

Store and update medical notes and prescriptions

RESTful API endpoints for CRUD operations

Lightweight and portable SQLite3 database

📂 Project Structure
plaintext
Copy
Edit
appointment_summary/
├── main.go
├── db/
│   └── database.go        # DB connection and setup
├── models/
│   ├── patient.go
│   ├── doctor.go
│   └── appointment.go
├── handlers/
│   ├── patient_handler.go
│   ├── doctor_handler.go
│   └── appointment_handler.go
├── utils/
│   └── response.go        # Helper functions
├── go.mod
└── README.md
📦 Setup & Run
Prerequisites
Go 1.18+
SQLite3
