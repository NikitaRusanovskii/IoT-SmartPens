
CREATE TABLE IF NOT EXISTS student_groups (
    group_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS subjects (
    subject_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS students (
    student_id UUID PRIMARY KEY,
    fname TEXT NOT NULL,
    mname TEXT,
    lname TEXT NOT NULL,
    group_id INTEGER NOT NULL REFERENCES student_groups(group_id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS teachers (
    teacher_id UUID PRIMARY KEY,
    fname TEXT NOT NULL,
    mname TEXT,
    lname TEXT NOT NULL,
    subject_id INTEGER NOT NULL REFERENCES subjects(subject_id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS lessons (
    lesson_id SERIAL PRIMARY KEY,
    teacher_id UUID NOT NULL REFERENCES teachers(teacher_id) ON DELETE CASCADE,
    date TIMESTAMP NOT NULL,
    group_id INTEGER NOT NULL REFERENCES student_groups(group_id) ON DELETE CASCADE,
    subject_id INTEGER NOT NULL REFERENCES subjects(subject_id) ON DELETE CASCADE,
    room INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS works (
    work_id SERIAL PRIMARY KEY,
    lesson_id INTEGER NOT NULL REFERENCES lessons(lesson_id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES students(student_id) ON DELETE CASCADE,
    data JSONB NOT NULL,
    mark INTEGER DEFAULT 0 CHECK (mark = 0 OR (mark >= 2 AND mark <= 5))
);

CREATE INDEX IF NOT EXISTS idx_students_group_id ON students(group_id);
CREATE INDEX IF NOT EXISTS idx_teachers_subject_id ON teachers(subject_id);
CREATE INDEX IF NOT EXISTS idx_lessons_teacher_id ON lessons(teacher_id);
CREATE INDEX IF NOT EXISTS idx_lessons_group_id ON lessons(group_id);
CREATE INDEX IF NOT EXISTS idx_lessons_subject_id ON lessons(subject_id);
CREATE INDEX IF NOT EXISTS idx_lessons_date ON lessons(date);
CREATE INDEX IF NOT EXISTS idx_works_lesson_id ON works(lesson_id);
CREATE INDEX IF NOT EXISTS idx_works_student_id ON works(student_id);
CREATE INDEX IF NOT EXISTS idx_student_groups_name ON student_groups(name);
CREATE INDEX IF NOT EXISTS idx_subjects_name ON subjects(name);
CREATE INDEX IF NOT EXISTS idx_works_data_gin ON works USING gin (data);
