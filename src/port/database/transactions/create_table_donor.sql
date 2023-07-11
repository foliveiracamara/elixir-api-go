CREATE TABLE tab_donor (
    donor_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf CHAR(11) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    blood_type VARCHAR(255),
    organ_donor BOOL,
    person_picture VARCHAR(255),
    donation_date DATE,
    birth_date DATE
);