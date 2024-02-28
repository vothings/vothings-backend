CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar,
  "email" varchar UNIQUE,
  "password" varchar,
  "role" varchar,
  "created_at" timestamp
);

CREATE TABLE "question_multiple_choice" (
  "id" uuid PRIMARY KEY,
  "question" varchar,
  "users_id" uuid
);

CREATE TABLE "multiple_choice_candidate" (
  "id" uuid PRIMARY KEY,
  "choice" varchar,
  "question_id" uuid
);

CREATE TABLE "multiple_choice_result" (
  "id" uuid PRIMARY KEY,
  "candidate_id" uuid,
  "users_id" uuid
);

CREATE TABLE "question_word_cloud" (
  "id" uuid PRIMARY KEY,
  "question" varchar,
  "users_id" uuid
);

CREATE TABLE "word_cloud_result" (
  "id" uuid PRIMARY KEY,
  "answer" varchar,
  "users_id" uuid,
  "question_id" uuid
);

CREATE TABLE "question_quiz" (
  "id" uuid PRIMARY KEY,
  "question" varchar,
  "users_id" uuid
);

CREATE TABLE "question_answer_quiz" (
  "id" uuid PRIMARY KEY,
  "choice" varchar,
  "question_id" uuid
);

CREATE TABLE "quiz_result" (
  "id" uuid PRIMARY KEY,
  "answer" varchar,
  "question_answer_id" uuid,
  "users_id" uuid
);

CREATE TABLE "question_rating" (
  "id" uuid PRIMARY KEY,
  "question" varchar,
  "users_id" uuid
);

CREATE TABLE "rating_result" (
  "id" uuid PRIMARY KEY,
  "answer" varchar,
  "question_rating_id" uuid,
  "users_id" uuid
);

CREATE TABLE "open_text_question" (
  "id" uuid PRIMARY KEY,
  "question" varchar,
  "users_id" uuid
);

CREATE TABLE "open_text_result" (
  "id" uuid PRIMARY KEY,
  "answer" varchar,
  "question_id" uuid,
  "users_id" uuid
);

ALTER TABLE "multiple_choice_result" ADD FOREIGN KEY ("candidate_id") REFERENCES "multiple_choice_candidate" ("id");

ALTER TABLE "multiple_choice_result" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "question_multiple_choice" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "multiple_choice_candidate" ADD FOREIGN KEY ("question_id") REFERENCES "question_multiple_choice" ("id");

ALTER TABLE "question_word_cloud" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "word_cloud_result" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "word_cloud_result" ADD FOREIGN KEY ("question_id") REFERENCES "question_word_cloud" ("id");

ALTER TABLE "question_quiz" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "question_answer_quiz" ADD FOREIGN KEY ("question_id") REFERENCES "question_quiz" ("id");

ALTER TABLE "quiz_result" ADD FOREIGN KEY ("question_answer_id") REFERENCES "question_answer_quiz" ("id");

ALTER TABLE "question_rating" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "rating_result" ADD FOREIGN KEY ("question_rating_id") REFERENCES "question_rating" ("id");

ALTER TABLE "rating_result" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "open_text_question" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");

ALTER TABLE "open_text_result" ADD FOREIGN KEY ("question_id") REFERENCES "open_text_question" ("id");

ALTER TABLE "open_text_result" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("id");
