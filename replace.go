package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fasttemplate"
)

func fire(c *cli.Context) error {
	// load configs
	godotenv.Overload(c.StringSlice("file")...)

	args := c.StringSlice("env")

	for _, item := range args {
		splited := strings.Split(item, "=")
		if len(splited) == 2 {
			key := strings.TrimSpace(splited[0])
			value := strings.TrimSpace(splited[1])
			os.Setenv(key, value)
		}
	}

	source := c.String("source")

	template, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("read source file %s failed, %v", source, err)
	}

	startTag := c.String("start")
	endTag := c.String("end")

	t, err := fasttemplate.NewTemplate(string(template), startTag, endTag)

	if err != nil {
		return err
	}

	writer := os.Stdout
	output := c.String("output")
	if output != "-" {
		writer, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return fmt.Errorf("error while open target file, %v", err)
		}
	}
	defer writer.Close()

	tmp := func(w io.Writer, tag string) (int, error) {
		k := strings.TrimSpace(tag)
		return w.Write([]byte(os.Getenv(k)))
	}

	t.ExecuteFunc(writer, tmp)

	return nil
}
