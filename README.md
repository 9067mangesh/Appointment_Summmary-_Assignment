# Appointment_Summmary-_Assignment

🩺 Appointment Summary - Assignment
📘 Project Overview
Appointment Summary Assignment is a demo healthcare-focused project designed to track and display appointment summaries between patients and doctors. This project aims to provide a structured and user-friendly way to manage and view appointment details, supporting better medical record management.

🚀 Features
Create and manage Patients, Doctors, and Appointments

View detailed Appointment Summaries including time, date, patient & doctor info

Record and update medical notes and prescriptions

Seamless integration between different objects (Doctor, Patient, Appointment)

Data validation and error handling for reliable performance

🛠️ Technologies Used
Update this section based on your actual tech stack.

Salesforce Platform (Lightning Experience)

Lightning Web Components (LWC)

Apex (for server-side logic)

Experience Cloud (for patient portal) – if applicable

SOQL & SOSL

Flows and Process Builder

Custom Objects and Fields

🧩 Data Model
This project uses the following custom Salesforce objects:

Patient__c

Name

Age

Contact Info

Doctor__c

Name

Specialization

Availability

Appointment__c

Patient (Lookup to Patient__c)

Doctor (Lookup to Doctor__c)

Appointment Date & Time

Notes

Prescription__c

Associated with Appointment

Medication details

Dosage

🔄 Flows & Automation
Book Appointment Flow – used to create an appointment

Follow-up Reminder Flow – sends a reminder after a set duration

Auto-create Prescription Record – after appointment completion


