package timefmt

import (
	"bytes"
)

// https://man7.org/linux/man-pages/man3/strftime.3.html
// https://golang.org/src/time/format.go
type dict map[byte]string

var txdict = dict{
	'a': "Mon",
	'A': "Monday",
	'b': "Jan",
	'B': "January",
	'c': "%a %b %e %H:%M:%S %Y",
	'C': "",   // not support
	'd': "02", // day, 01-31
	'D': "%m/%d/%y",
	'e': "", // not support
	'E': "", // not support
	'F': "%Y-%m-%d",
	'G': "", // not support
	'g': "", // not support
	'h': "%b",
	'H': "15", // need test
	'I': "03",
	'j': "002",
	'k': "15", // need test
	'l': "3",  // need test
	'm': "01",
	'M': "04",
	'n': "\n",
	'O': "", // not support
	'p': "PM",
	'P': "pm",
	'r': "%I:%M:%S %p",
	'R': "%H:%M",
	's': "", // not support
	'S': "05",
	't': "\t",
	'T': "%H:%M:%S",
	'u': "", // not support
	'U': "", // not support
	'V': "", // not support
	'w': "", // not support
	'W': "", // not support
	'x': "%m/%d/%y",
	'X': "%H:%M:%S",
	'y': "2006",
	'Y': "06",
	'z': "-0700",
	'Z': "", // not
	'%': "%",
	//	'-': "\\-",
}

var tab []string

func maketable() []string {
	if tab != nil {
		return tab
	}
	tab = make([]string, 256)
	var latedict = map[byte]string{}
	for k, v := range txdict {
		if v == "" {
			continue
		} else if v[0] == '%' {
			latedict[k] = v
		} else {
			tab[k] = v
		}
	}
	for k, v := range latedict {
		s := translateFmt(v)
		tab[k] = s
	}
	return tab
}

func translateFmt(strfmt string) string {
	var out bytes.Buffer
	for i := 0; i < len(strfmt); i++ {
		if t := strfmt[i]; t == '%' {
			if i+1 == len(strfmt) {
				break
			}
			i += 1
			c := strfmt[i]
			if s := tab[c]; s != "" {
				out.WriteString(s)
			}
		} else {
			out.WriteByte(t)
		}
	}
	return out.String()
}

func FromPosix(strfmt string) string {
	maketable()
	return translateFmt(strfmt)
}
