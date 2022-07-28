package iniparser

import (
	"log"
	"os"
	"strings"
)

func ParseMultipleConfigs(config_files ...string) {
	for _, cf := range config_files {
		ParseConfig(cf)
	}
}

func ParseConfig(config_file string) {
	if config_file == "" {
		config_file = "config.ini"
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("\033[34m[INFO]\033[0m ")
	log.Default().Println("Config File:", config_file)
	log.SetPrefix("")
	bytes_read, err := os.ReadFile(config_file)
	if err != nil {
		log.Fatalf("[Error] %v\n", err.Error())
	}
	file_content := normalize_lines_list(strings.Split(strings.TrimSpace(string(bytes_read)), "\n"))
	for _, elem := range file_content {
		keyvalue_tuple := strings.SplitN(elem, "=", 2)
		os.Setenv(keyvalue_tuple[0], keyvalue_tuple[1])
	}
}

const (
	state_not_string = iota
	state_string     = iota
)

var current_state = state_not_string

const (
	ss_semicolon = iota
	ss_hashtag   = iota
)

var (
	ss = []string{ // Special Sequences
		":semi-colon:",
		":hashtag:",
	}
	sscp = []string{ // Special Sequences Counterpart
		";",
		"#",
	}
)

func normalize_line(iline string) (oline string) {
	current_state = state_not_string
	var assign_index = strings.Index(iline, "=") + 1

	if strings.ContainsAny(iline, "=") {
		for i := assign_index; i < len(iline); i++ {
			if current_state == state_string {
				if iline[i] == '"' || iline[i] == '\'' {
					current_state = state_not_string
					oline += string(iline[i])
					continue
				}
				oline += string(iline[i])
				continue
			}
			if current_state == state_not_string {
				if iline[i] == '"' || iline[i] == '\'' {
					current_state = state_string
					oline += string(iline[i])
					continue
				} else if iline[i] == '#' || iline[i] == ';' {
					break
				}
				oline += string(iline[i])
				continue
			}
		}
		oline = strings.TrimSpace(iline[:assign_index-1]) + "=" + strings.TrimSpace(oline)
	}
	return // Implicitly returns `oline`
}

func normalize_lines_list(ilist []string) (olist []string) {
	for _, elem := range ilist {
		if !strings.ContainsAny(elem, "=") && (!strings.HasPrefix(elem, "[") && !strings.HasSuffix(elem, "]")) {
			continue
		} else if strings.HasPrefix(elem, "[") && strings.HasSuffix(elem, "]") {
			continue
		}
		elem = normalize_line(elem)
		if elem != "" && !strings.HasPrefix(elem, ";") && !strings.HasPrefix(elem, "#") {
			olist = append(
				olist,
				strings.ReplaceAll(
					strings.ReplaceAll(
						elem, ss[ss_semicolon], sscp[ss_semicolon],
					),
					ss[ss_hashtag], sscp[ss_hashtag],
				),
			)
		}
	}
	return // Implicitly returns `olist`
}
