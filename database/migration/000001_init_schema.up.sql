CREATE TABLE "students" (
  "student_id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "middle_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "gender" char(1) NOT NULL,
  "date_of_birth" date NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "email" varchar NOT NULL,
  "year_of_enroll" integer NOT NULL,
  "country" varchar NOT NULL,
  "major" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "courses" (
  "course_id" bigserial PRIMARY KEY,
  "course_name" varchar NOT NULL,
  "instructor" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "enrollments" (
  "enrollment_id" bigserial PRIMARY KEY,
  "enrollment_date" date NOT NULL,
  "student_id" integer NOT NULL,
  "course_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "enrollments" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("student_id");

ALTER TABLE "enrollments" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("course_id");