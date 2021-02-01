package main

import (
	"postgresql/db"
)

var schema = `

CREATE TABLE "teamlead" (
	"id" serial,
	"full_name" VARCHAR(255),
	CONSTRAINT "teamlead_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- name: create-director-table
CREATE TABLE "director" (
	"id" serial,
	"full_name" VARCHAR(255),
	CONSTRAINT "director_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- name: create-task-table
CREATE TABLE "task" (
	"id" serial,
	"title" VARCHAR(255),
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
	CONSTRAINT "programmer_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("t_id") REFERENCES "teamlead"("id");
ALTER TABLE "task" ADD CONSTRAINT "task_fk1" FOREIGN KEY ("p_id") REFERENCES "programmer"("id");


`

func main() {

	//schema ni mazaga migrate qiladi
	// db.RUN().MustExec(schema)

	tx := db.RUN().MustBegin()

	// tx.MustExec("INSERT INTO director (id, full_name) VALUES ($1, $2)", 1, "Sardor")
	// tx.MustExec("INSERT INTO teamlead (id, full_name) VALUES ($1, $2)", 1, "Doston")
	// tx.MustExec("INSERT INTO programmer (id, full_name) VALUES ($1, $2)", 1, "Navruz")

	// tx.MustExec("INSERT INTO task (id, title) VALUES ($1, $2)", 1, "Web Ilova yaratish")

	tx.Commit()

}
