// Example get-active-window reads the _NET_ACTIVE_WINDOW property of the root
// window and uses the result (a window id) to get the name of the window.
package main

import (
	"database/sql"
	"fmt"
	"github.com/erasche/gologme"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Golog struct {
	Db *sql.DB
}

func (t *Golog) logToDb(uid int, windowlogs []gologme.WindowLogs, keylogs []gologme.KeyLogs, wll int) {
	tx, err := t.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	wl_stmt, err := tx.Prepare("insert into windowLogs (uid, time, name) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	kl_stmt, err := tx.Prepare("insert into keyLogs (uid, time, count) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer wl_stmt.Close()
	defer kl_stmt.Close()

	log.Printf("%d window logs %d key logs from [%d]\n", wll, len(keylogs), uid)

	for _, w := range keylogs {
		_, err = kl_stmt.Exec(uid, w.Time.Unix(), w.Count)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, w := range windowlogs {
		_, err = wl_stmt.Exec(uid, w.Time.Unix(), w.Name)
		if err != nil {
			log.Fatal(err)
		}
		// Dunno why windowLogs comes through two too big, so whatever.
		if i >= wll-1 {
			break
		}
	}

	tx.Commit()
}

func (t *Golog) ensureAuth(user string, key string) (int, error) {
	// Pretty assuredly not safe from timing attacks.
	stmt, err := t.Db.Prepare("select id from users where username = ? and api_key = ?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	// Maybe helps against timing attacks? Hmm.
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)

	var uid int
	err = stmt.QueryRow(user, key).Scan(&uid)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

func (t *Golog) Log(args gologme.RpcArgs, result *int) error {
	uid, err := t.ensureAuth(args.User, args.ApiKey)
	if err != nil {
		log.Fatal(err)
		*result = 1
		return nil
	} else {
		log.Printf("%s authenticated successfully as uid %d\n", args.User, uid)
	}

	t.logToDb(
		uid,
		args.Windows,
		args.KeyLogs,
		args.WindowLogsLength,
	)

	*result = 0
	return nil
}

func (t *Golog) setupDb(db *sql.DB) {
	dbCreation := `
	create table if not exists users (
		id integer not null primary key autoincrement,
		username text,
		api_key text
	);

	create table if not exists windowLogs (
		id integer not null primary key autoincrement,
		uid integer,
		time integer,
		name text,
		foreign key (uid) references users(id)
	);

	create table if not exists keyLogs (
		id integer not null primary key autoincrement,
		uid integer,
		time integer,
		count integer,
		foreign key (uid) references users(id)
	);

	create table if not exists notes (
		id integer not null primary key autoincrement,
		uid integer,
		time integer,
		type integer,
		contents text,
		foreign key (uid) references users(id)
	);
	`
	// notes.type is either 0 for daily blog or 1 for point-in-time note

	_, err := db.Exec(dbCreation)
	if err != nil {
		log.Fatal(err)
	}
	t.Db = db
}

func main() {
	db, err := sql.Open("sqlite3", "file.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	golog := new(Golog)
	golog.setupDb(db)
	rpc.Register(golog)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	fmt.Println("Listening...")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}