package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/salaleser/vitalina/util"
)

func GetTTSCount() int {
	sql := "SELECT COUNT(*) FROM tts"

	rows := query(sql)

	if rows == nil {
		fmt.Fprintf(os.Stderr, "tts execute query: %v\n", errors.New("rows == nil"))
		return -1
	}

	var c int
	for rows.Next() {
		err := rows.Scan(&c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan1 tts count failed: %v\n", err)
			return -1
		}
	}

	return c
}

func GetTTS(tts string, language string) (string, int) {
	sql := "SELECT filename, counter FROM tts WHERE text = '" + tts + "' AND language = '" + language + "'"

	rows := query(sql)

	var filename string
	var counter int
	for rows.Next() {
		err := rows.Scan(&filename, &counter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan tts failed: %v\n", err)
		}
	}

	return filename, counter
}

func UpdateTTS(tts string, filename string, counter int, language string, updated string) {
	sql := fmt.Sprintf("INSERT INTO tts (text, filename, counter, language, updated) VALUES ('%s','%s',%d,'%s','%s') ON CONFLICT (text) DO UPDATE SET (counter, updated) = (%d,'%s')",
		tts, filename, counter, language, updated, counter, updated)

	query(sql)
}

func query(sql string) pgx.Rows {
	port, _ := strconv.ParseInt(util.Config["db-port"], 10, 16)
	config, err := pgx.ParseConfig("")
	config.Host = util.Config["db-host"]
	config.Port = uint16(port)
	config.User = util.Config["db-user"]
	config.Password = util.Config["db-password"]
	config.Database = "vacdbot"
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		return nil
	}

	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		return nil
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query1 failed: %v\n", err)
		return nil
	}
	defer rows.Close()

	return rows
}
