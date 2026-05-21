CREATE TABLE lesson (
    id         SERIAL PRIMARY KEY,
    teacher_id UUID NOT NULL,
    group_id   INTEGER NOT NULL,
    subject_id INTEGER NOT NULL,
    room       INTEGER NOT NULL,
    date       TIMESTAMP NOT NULL
);

CREATE TABLE work (
    lesson_id  INTEGER NOT NULL,
    student_id UUID NOT NULL,
    data       JSONB,
    FOREIGN KEY (lesson_id) REFERENCES lesson(id) ON DELETE CASCADE
);