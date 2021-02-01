-- name: create-teamlead-table
CREATE TABLE "teamlead" (
	"id" serial,
	"full_name" VARCHAR(255),
	"created_at" DATETIME,
	CONSTRAINT "teamlead_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- name: create-director-table
CREATE TABLE "director" (
	"id" serial,
	"full_name" VARCHAR(255),
	"created_at" DATETIME,
	CONSTRAINT "director_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- name: create-task-table
CREATE TABLE "task" (
	"id" serial,
	"title" VARCHAR(255),
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL,
	"t_id" integer NOT NULL,
	"p_id" integer NOT NULL,
	"end_task" BOOLEAN NOT NULL,
	"finish_task" BOOLEAN NOT NULL,
	"check_task" BOOLEAN NOT NULL,
	CONSTRAINT "task_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


-- name: create-programmer-table
CREATE TABLE "programmer" (
	"id" serial NOT NULL,
	"full_name" VARCHAR(255),
	"created_at" DATETIME,
	CONSTRAINT "programmer_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("t_id") REFERENCES "teamlead"("id");
ALTER TABLE "task" ADD CONSTRAINT "task_fk1" FOREIGN KEY ("p_id") REFERENCES "programmer"("id");

