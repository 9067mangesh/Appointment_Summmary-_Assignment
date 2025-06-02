package sender

import (
    "database/sql"
    "fmt"
)

func GenerateDoctorSummary(db *sql.DB, date string) error {
    query := `
        SELECT d.DoctorStaffID, d.Name, COUNT(a.AppointmentID)
        FROM DoctorStaff d
        JOIN Appointment a ON d.DoctorStaffID = a.DoctorStaffID
        WHERE DATE(a.StartTime) = ? AND a.Status = 'S'
        GROUP BY d.DoctorStaffID, d.Name;
    `
    rows, err := db.Query(query, date)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var doctorID int
        var doctorName string
        var count int

        if err := rows.Scan(&doctorID, &doctorName, &count); err != nil {
            return err
        }

        summary := fmt.Sprintf("Hi Dr. %s, you have %d appointments scheduled today.", doctorName, count)
        _, err := db.Exec("INSERT INTO DoctorSummary (DoctorStaffID, Summary) VALUES (?, ?)", doctorID, summary)
        if err != nil {
            return err
        }
    }

    return nil
}

func GenerateCenterSummary(db *sql.DB, date string) error {
    query := `
        SELECT c.CenterID, c.CenterName, COUNT(DISTINCT a.DoctorStaffID), COUNT(a.AppointmentID)
        FROM Center c
        JOIN Appointment a ON c.CenterID = a.CenterID
        WHERE DATE(a.StartTime) = ? AND a.Status = 'S'
        GROUP BY c.CenterID, c.CenterName;
    `
    rows, err := db.Query(query, date)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var centerID int
        var centerName string
        var doctorCount, apptCount int

        if err := rows.Scan(&centerID, &centerName, &doctorCount, &apptCount); err != nil {
            return err
        }

        summary := fmt.Sprintf("%s has %d appointments across %d doctors for today.", centerName, apptCount, doctorCount)
        _, err := db.Exec("INSERT INTO CenterSummary (CenterID, Summary) VALUES (?, ?)", centerID, summary)
        if err != nil {
            return err
        }
    }

    return nil
}
