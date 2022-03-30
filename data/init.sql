CREATE DATABASE dbsystem;
USE dbsystem;

DROP TABLE IF EXISTS employees, roles, timesheet, departments, requests;

CREATE TABLE employees (
                       id INT NOT NULL AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       birthday DATE NOT NULL,
                       roleId INT NOT NULL,
                       departmentId INT NOT NULL,
                       index (id),
                       PRIMARY KEY (id)
);

CREATE TABLE roles (
                       id INT NOT NULL AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       roleId INT NOT NULL,
                       PRIMARY KEY (id),
                       index (id)
);

CREATE TABLE timesheet(
                            id INT NOT NULL AUTO_INCREMENT,
                            employeeId INT NOT NULL,
                            year varchar(255) NOT NULL,
                            month varchar(255) NOT NULL,
                            day varchar(255) NOT NULL,
                            checkIn VARCHAR(255),
                            checkOut VARCHAR(255),
                            PRIMARY KEY (id),
                            FOREIGN KEY (employeeId) REFERENCES employees(id),
                            index (employeeId)
);

CREATE TABLE departments(
                           id INT NOT NULL AUTO_INCREMENT,
                           name VARCHAR(255) NOT NULL,
                           managerId INT NOT NULL,
                           PRIMARY KEY (id),
                           foreign key (managerId) references employees(id),
                           index (name)
);

CREATE TABLE requests(
                        id INT NOT NULL AUTO_INCREMENT,
                        typeRequest VARCHAR(255) default 'leaving request',
                        employeeId INT NOT NULL,
                        startDate DATE NOT NULL,
                        endDate DATE NOT NULL,
                        reason VARCHAR(255) NOT NULL,
                        statusRequest VARCHAR(255) NOT NULL default 'pending',
                        PRIMARY KEY (id),
                        FOREIGN KEY (employeeId) REFERENCES employees(id),
                        index (employeeId)
);

INSERT INTO employees (name, email, password, birthday, roleId, departmentId)
VALUES
    ('nguyen thanh hai', 'hai@zinza.com.vn', '123456', '1981-05-17', 1, 1),

    ('le quang hung', 'hung@zinza.com.vn', '123456', '1985-05-17', 2, 1),
    ('phan hoai son', 'son.ph1@zinza.com.vn', '123456', '1995-05-17', 3, 1),
    ('nguyen van huy', 'huy.nv@zinza.com.vn', '123456', '2000-05-17', 3, 1),


    ('nguyen xuan cuong', 'cuong@zinza.com.vn', '123456', '1991-05-17', 2, 2),
    ('nguyen van thi', 'thi.nv@zinza.com.vn', '123456', '1990-05-17', 3, 2),

    ('le minh quan', 'quan@zinza.com.vn', '123456', '1992-05-17', 2, 3),
    ('nguyen van hieu', 'hieu.nv@zinza.com.vn', '123456', '2000-05-17', 3, 3),

    ('nguyen thi phuong thao', 'thao.ntp@zinza.com.vn', '123456', '1993-01-01', 2, 4),
    ('nguyen hong doan', 'doan.nh@zinza.com.vn', '123456', '1999-05-17', 3, 4);



INSERT INTO roles (name, roleId)
VALUES ('admin', 1), ('divison lead', 2), ('employee', 3);

INSERT INTO departments (name, managerId)
VALUES ('Divison 1', 2), ('Divison 2', 5), ('Divison 3', 7), ('Support', 9);


INSERT INTO timesheet (employeeId, year, month, day, checkIn, checkOut)
VALUES (1, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 17:30:00"),
    (2, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 17:30:00"),
    (3, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 17:30:00"),
    (4, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 17:30:00"),
    (5, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 18:30:00"),
    (6, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 18:30:00"),
    (7, "2022", "3", "24", "2022-03-24 08:30:00", "2022-03-24 18:30:00"),
    (1, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (2, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (3, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (4, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (5, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (6, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (7, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (8, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (9, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00"),
    (10, "2022", "3", "25", "2022-03-25 08:30:00", "2022-03-25 18:30:00");




INSERT INTO requests (employeeId, typeRequest, startDate, endDate, reason, statusRequest)
VALUE (1, "leaving request", "2022-01-01", "2022-01-03", "I want to go to Japan", "accepted"),
    (2, "leaving request", "2022-01-01", "2022-01-03", "I want to go to Japan", "accepted"),
    (3, "leaving request", "2022-01-01", "2022-01-03", "I want to go to Japan", "accepted"),
    (4, "leaving request", "2022-01-01", "2022-01-03", "I want to go to Japan", "accepted");


