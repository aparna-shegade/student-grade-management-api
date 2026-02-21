CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), username VARCHAR(50) UNIQUE NOT NULL, password_hash VARCHAR(255) NOT NULL, role VARCHAR(20) NOT NULL CHECK (role IN ('student', 'teacher', 'admin')), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
CREATE TABLE courses (id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), name VARCHAR(100) NOT NULL, teacher_id UUID REFERENCES users(id), credits FLOAT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
CREATE TABLE grades (id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), student_id UUID REFERENCES users(id), course_id UUID REFERENCES courses(id), grade FLOAT CHECK (grade >= 0 AND grade <= 100), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, UNIQUE(student_id, course_id));
