// This file was autogenerated. DO NOT EDIT.

package core

import (
	"encoding/json"
	"strings"
)

const (
	clear  = "\x1b[0m"
	bright = "\x1b[1m"
	gray   = "\x1b[90m"
	yellow = "\x1b[33m"
)

// Command represents a Tile38 command.
type Command struct {
	Name       string     `json:"-"`
	Summary    string     `json:"summary"`
	Complexity string     `json:"complexity"`
	Arguments  []Argument `json:"arguments"`
	Since      string     `json:"since"`
	Group      string     `json:"group"`
	DevOnly    bool       `json:"dev"`
}

// String returns a string representation of the command.
func (c Command) String() string {
	var s = c.Name
	for _, arg := range c.Arguments {
		s += " " + arg.String()
	}
	return s
}

// TermOutput returns a string representation of the command suitable for displaying in a terminal.
func (c Command) TermOutput(indent string) string {
	line := c.String()
	var line1 string
	if strings.HasPrefix(line, c.Name) {
		line1 = bright + c.Name + clear + gray + line[len(c.Name):] + clear
	} else {
		line1 = bright + strings.Replace(c.String(), " ", " "+clear+gray, 1) + clear
	}
	line2 := yellow + "summary: " + clear + c.Summary
	//line3 := yellow + "since: " + clear + c.Since
	return indent + line1 + "\n" + indent + line2 + "\n" //+ indent + line3 + "\n"
}

// EnumArg represents a enum arguments.
type EnumArg struct {
	Name      string     `json:"name"`
	Arguments []Argument `json:"arguments"`
}

// String returns a string representation of an EnumArg.
func (a EnumArg) String() string {
	var s = a.Name
	for _, arg := range a.Arguments {
		s += " " + arg.String()
	}
	return s
}

// Argument represents a command argument.
type Argument struct {
	Command  string      `json:"command"`
	NameAny  interface{} `json:"name"`
	TypeAny  interface{} `json:"type"`
	Optional bool        `json:"optional"`
	Multiple bool        `json:"multiple"`
	Variadic bool        `json:"variadic"`
	Enum     []string    `json:"enum"`
	EnumArgs []EnumArg   `json:"enumargs"`
}

// String returns a string representation of an Argument.
func (a Argument) String() string {
	var s string
	if a.Command != "" {
		s += " " + a.Command
	}
	if len(a.EnumArgs) > 0 {
		eargs := ""
		for _, arg := range a.EnumArgs {
			v := arg.String()
			if strings.Contains(v, " ") {
				v = "(" + v + ")"
			}
			eargs += v + "|"
		}
		if len(eargs) > 0 {
			eargs = eargs[:len(eargs)-1]
		}
		s += " " + eargs
	} else if len(a.Enum) > 0 {
		s += " " + strings.Join(a.Enum, "|")
	} else {
		names, _ := a.NameTypes()
		subs := ""
		for _, name := range names {
			subs += " " + name
		}
		subs = strings.TrimSpace(subs)
		s += " " + subs
		if a.Variadic {
			s += " [" + subs + " ...]"
		}
		if a.Multiple {
			s += " ..."
		}
	}
	s = strings.TrimSpace(s)
	if a.Optional {
		s = "[" + s + "]"
	}
	return s
}

func parseAnyStringArray(any interface{}) []string {
	if str, ok := any.(string); ok {
		return []string{str}
	} else if any, ok := any.([]interface{}); ok {
		arr := []string{}
		for _, any := range any {
			if str, ok := any.(string); ok {
				arr = append(arr, str)
			}
		}
		return arr
	}
	return []string{}
}

// NameTypes returns the types and names of an argument as separate arrays.
func (a Argument) NameTypes() (names, types []string) {
	names = parseAnyStringArray(a.NameAny)
	types = parseAnyStringArray(a.TypeAny)
	if len(types) > len(names) {
		types = types[:len(names)]
	} else {
		for len(types) < len(names) {
			types = append(types, "")
		}
	}
	return
}

// Commands is a map of all of the commands.
var Commands = func() map[string]Command {
	var commands map[string]Command
	if err := json.Unmarshal([]byte(commandsJSON), &commands); err != nil {
		panic(err.Error())
	}
	for name, command := range commands {
		command.Name = strings.ToUpper(name)
		commands[name] = command
	}
	return commands
}()

var commandsJSON = `{
  "SET":{
    "summary": "Sets the value of an id",
    "complexity": "O(1)",
    "arguments": [
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      },
      {
        "command": "FIELD",
        "name": ["name", "value"],
        "type": ["string", "double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "EX",
        "name": ["seconds", "value"],
        "type": ["string", "double"],
        "optional": true,
        "multiple": false
      },
	  {
		"name": "value",
        "enumargs": [
          {
            "name": "OBJECT",
            "arguments":[
              {
                "name": "geojson",
                "type": "geojson"    
              }
            ]
          },
          {
            "name": "POINT",
            "arguments":[
              {
                "name": "lat",
                "type": "double"    
              },
              {
                "name": "lon",
                "type": "double"    
              },
              {
                "name": "z",
                "type": "double",
                "optional": true
              }
            ]
          },
          {
            "name": "BOUNDS",
            "arguments":[
              {
                "name": "minlat",
                "type": "double"    
              },
              {
                "name": "minlon",
                "type": "double"    
              },
              {
                "name": "maxlat",
                "type": "double"    
              },
              {
                "name": "maxlon",
                "type": "double"    
              }
            ]
          },
          {
            "name": "HASH",
            "arguments":[
              {
                "name": "geohash",
                "type": "geohash"    
              }
            ]
          },
          {
            "name": "STRING",
            "arguments":[
              {
                "name": "value",
                "type": "string"    
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "EXPIRE": {
    "summary": "Set a timeout on an id",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      },
      {
        "name": "seconds",
        "type": "double"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "TTL": {
    "summary": "Get a timeout on an id",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "PERSIST": {
    "summary": "Remove the existing timeout on an id",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "FSET": {
    "summary": "Set the value for a single field of an id",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      },
      {
        "name": "field",
        "type": "string"
      },
      {
        "name": "value",
        "type": "double"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "BOUNDS": {
    "summary": "Get the combined bounds of all the objects in a key",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      }
    ],
    "since": "1.3.0",
    "group": "keys"
  },
  "GET": {
    "summary": "Get the object of an id",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      },
      {
        "command": "WITHFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "OBJECT"
          },
          {
            "name": "POINT"
          },
          {
            "name": "BOUNDS"
          },
          {
            "name": "HASH",
            "arguments": [
              {
                "name": "geohash",
                "type": "geohash"
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "DEL": {
    "summary": "Delete an id from a key",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "name": "id",
        "type": "string"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "DROP": {
    "summary": "Remove a key from the database",
    "complexity": "O(1)",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "KEYS": {
    "summary": "Finds all keys matching the given pattern",
    "complexity": "O(N) where N is the number of keys in the database",
    "arguments":[
      {
        "name": "pattern",
        "type": "pattern"
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "STATS": {
    "summary": "Show stats for one or more keys",
    "complexity": "O(N) where N is the number of keys being requested",
    "arguments":[
      {
        "name": "key",
        "type": "string",
        "variadic": true
      }
    ],
    "since": "1.0.0",
    "group": "keys"
  },
  "SEARCH": {
    "summary": "Search for string values in a key",
    "complexity": "O(N) where N is the number of values in the key",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "CURSOR",
        "name": "start",
        "type": "integer",
        "optional": true
      },
      {
        "command": "LIMIT",
        "name": "count",
        "type": "integer",
        "optional": true
      },
      {
        "command": "MATCH",
        "name": "pattern",
        "type": "pattern",
        "optional": true
      },
      {
        "name": "order",
        "optional": true,
        "enumargs": [
          {
            "name": "ASC"
          },
          {
            "name": "DESC"
          }
		]
	  },
      {
        "command": "WHERE",
        "name": ["field","min","max"],
        "type": ["string","double","double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "NOFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "COUNT"
          },
          {
            "name": "IDS"
          }
        ]
      }
    ],
    "since": "1.4.2",
    "group": "search"
  },
  "SCAN": {
    "summary": "Incrementally iterate though a key",
    "complexity": "O(N) where N is the number of ids in the key",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "CURSOR",
        "name": "start",
        "type": "integer",
        "optional": true
      },
      {
        "command": "LIMIT",
        "name": "count",
        "type": "integer",
        "optional": true
      },
      {
        "command": "MATCH",
        "name": "pattern",
        "type": "pattern",
        "optional": true
      },
      {
        "name": "order",
        "optional": true,
        "enumargs": [
          {
            "name": "ASC"
          },
          {
            "name": "DESC"
          }
		]
	  },
      {
        "command": "WHERE",
        "name": ["field","min","max"],
        "type": ["string","double","double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "NOFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "COUNT"
          },
          {
            "name": "IDS"
          },
          {
            "name": "OBJECTS"
          },
          {
            "name": "POINTS"
          },
          {
            "name": "BOUNDS"
          },
          {
            "name": "HASHES",
            "arguments": [
              {
                "name": "precision",
                "type": "integer"
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "search"
  },
  "NEARBY": {
    "summary": "Searches for ids that are nearby a point",
    "complexity": "O(log(N)) where N is the number of ids in the area",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "CURSOR",
        "name": "start",
        "type": "integer",
        "optional": true
      },
      {
        "command": "LIMIT",
        "name": "count",
        "type": "integer",
        "optional": true
      },
      {
        "command": "SPARSE",
        "name": "spread",
        "type": "integer",
        "optional": true
      },
      {
        "command": "MATCH",
        "name": "pattern",
        "type": "pattern",
        "optional": true
      },
      {
        "command": "WHERE",
        "name": ["field","min","max"],
        "type": ["string","double","double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "NOFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "FENCE",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "DETECT",
        "name": ["what"],
        "type": ["string"],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "COUNT"
          },
          {
            "name": "IDS"
          },
          {
            "name": "OBJECTS"
          },
          {
            "name": "POINTS"
          },
          {
            "name": "BOUNDS"
          },
          {
            "name": "HASHES",
            "arguments": [
              {
                "name": "precision",
                "type": "integer"
              }
            ]
          }
        ]
      },
      {
        "name": "area",
        "enumargs": [
          {
            "name": "POINT",
            "arguments": [
              {
                "name": "lat",
                "type": "double"
              },
              {
                "name": "lon",
                "type": "double"
              },
              {
                "name": "meters",
                "type": "double"
              }
            ]
          },
          {
            "name": "ROAM",
            "arguments":[
              {
                "name": "key",
                "type": "string"    
              },
              {
                "name": "pattern",
                "type": "pattern"    
              },
              {
                "name": "meters",
                "type": "double"    
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "search"
  },
  "WITHIN": {
    "summary": "Searches for ids that are nearby a point",
    "complexity": "O(log(N)) where N is the number of ids in the area",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "CURSOR",
        "name": "start",
        "type": "integer",
        "optional": true
      },
      {
        "command": "LIMIT",
        "name": "count",
        "type": "integer",
        "optional": true
      },
      {
        "command": "SPARSE",
        "name": "spread",
        "type": "integer",
        "optional": true
      },
      {
        "command": "MATCH",
        "name": "pattern",
        "type": "pattern",
        "optional": true
      },
      {
        "command": "WHERE",
        "name": ["field","min","max"],
        "type": ["string","double","double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "NOFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "FENCE",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "DETECT",
        "name": ["what"],
        "type": ["string"],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "COUNT"
          },
          {
            "name": "IDS"
          },
          {
            "name": "OBJECTS"
          },
          {
            "name": "POINTS"
          },
          {
            "name": "BOUNDS"
          },
          {
            "name": "HASHES",
            "arguments": [
              {
                "name": "precision",
                "type": "integer"
              }
            ]
          }
        ]
      },
      {
        "name": "area",
        "enumargs": [
          {
            "name": "GET",
            "arguments": [
              {
                "name": "key",
                "type": "string"
              },
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "BOUNDS",
            "arguments":[
              {
                "name": "minlat",
                "type": "double"    
              },
              {
                "name": "minlon",
                "type": "double"    
              },
              {
                "name": "maxlat",
                "type": "double"    
              },
              {
                "name": "maxlon",
                "type": "double"    
              }
            ]
          },
          {
            "name": "OBJECT",
            "arguments":[
              {
                "name": "geojson",
                "type": "geojson"    
              }
            ]
          },
          {
            "name": "TILE",
            "arguments":[
              {
                "name": "x",
                "type": "double"    
              },
              {
                "name": "y",
                "type": "double"    
              },
              {
                "name": "z",
                "type": "double"    
              }
            ]
          },
          {
            "name": "QUADKEY",
            "arguments":[
              {
                "name": "quadkey",
                "type": "string"    
              }
            ]
          },
          {
            "name": "HASH",
            "arguments": [
              {
                "name": "geohash",
                "type": "geohash"
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "search"
  },
  "INTERSECTS": {
    "summary": "Searches for ids that are nearby a point",
    "complexity": "O(log(N)) where N is the number of ids in the area",
    "arguments":[
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "CURSOR",
        "name": "start",
        "type": "integer",
        "optional": true
      },
      {
        "command": "LIMIT",
        "name": "count",
        "type": "integer",
        "optional": true
      },
      {
        "command": "SPARSE",
        "name": "spread",
        "type": "integer",
        "optional": true
      },
      {
        "command": "MATCH",
        "name": "pattern",
        "type": "pattern",
        "optional": true
      },
      {
        "command": "WHERE",
        "name": ["field","min","max"],
        "type": ["string","double","double"],
        "optional": true,
        "multiple": true
      },
      {
        "command": "NOFIELDS",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "FENCE",
        "name": [],
        "type": [],
        "optional": true
      },
      {
        "command": "DETECT",
        "name": ["what"],
        "type": ["string"],
        "optional": true
      },
      {
        "name": "type",
        "optional": true,
        "enumargs": [
          {
            "name": "COUNT"
          },
          {
            "name": "IDS"
          },
          {
            "name": "OBJECTS"
          },
          {
            "name": "POINTS"
          },
          {
            "name": "BOUNDS"
          },
          {
            "name": "HASHES",
            "arguments": [
              {
                "name": "precision",
                "type": "integer"
              }
            ]
          }
        ]
      },
      {
        "name": "area",
        "enumargs": [
          {
            "name": "GET",
            "arguments": [
              {
                "name": "key",
                "type": "string"
              },
              {
                "name": "id",
                "type": "string"
              }
            ]
          },
          {
            "name": "BOUNDS",
            "arguments":[
              {
                "name": "minlat",
                "type": "double"    
              },
              {
                "name": "minlon",
                "type": "double"    
              },
              {
                "name": "maxlat",
                "type": "double"    
              },
              {
                "name": "maxlon",
                "type": "double"    
              }
            ]
          },
          {
            "name": "OBJECT",
            "arguments":[
              {
                "name": "geojson",
                "type": "geojson"    
              }
            ]
          },
          {
            "name": "TILE",
            "arguments":[
              {
                "name": "x",
                "type": "double"    
              },
              {
                "name": "y",
                "type": "double"    
              },
              {
                "name": "z",
                "type": "double"    
              }
            ]
          },
          {
            "name": "QUADKEY",
            "arguments":[
              {
                "name": "quadkey",
                "type": "string"    
              }
            ]
          },
          {
            "name": "HASH",
            "arguments": [
              {
                "name": "geohash",
                "type": "geohash"
              }
            ]
          }
        ]
      }
    ],
    "since": "1.0.0",
    "group": "search"
  },
  "CONFIG GET": {
    "summary": "Get the value of a configuration parameter",
    "arguments":[
      {
        "name": "parameter",
        "type": "string"    
      }
    ],
    "group": "server"
  },
  "CONFIG SET": {
    "summary": "Set a configuration parameter to the given value",
    "arguments":[
      {
        "name": "parameter",
        "type": "string"    
      },
      {
        "name": "value",
        "type": "string",
        "optional": true
      }
    ],
    "group": "server"
  },
  "CONFIG REWRITE": {
    "summary": "Rewrite the configuration file with the in memory configuration",
    "arguments":[],
    "group": "server"
  },
  "SERVER": {
    "summary":"Show server stats and details",
    "complexity": "O(1)",
    "arguments": [],
    "since": "1.0.0",
    "group": "server"
  },
  "GC": {
    "summary":"Forces a garbage collection",
    "complexity": "O(1)",
    "arguments": [],
    "since": "1.0.0",
    "group": "server"
  },
  "READONLY": {
    "summary": "Turns on or off readonly mode",
    "complexity": "O(1)",
    "arguments": [
      {
        "enum": ["yes","no"]
      }
    ],
    "since": "1.0.0",
    "group": "server"
  },
  "FLUSHDB": {
    "summary":"Removes all keys",
    "complexity": "O(1)",
    "arguments": [],
    "since": "1.0.0",
    "group": "server"
  },
  "FOLLOW": {
    "summary": "Follows a leader host",
    "complexity": "O(1)",
    "arguments": [
      {
        "name": "host",
        "type": "string"
      },
      {
        "name": "port",
        "type": "integer"
      }
    ],
    "since": "1.0.0",
    "group": "replication"
  },
  "AOF": {
    "summary": "Downloads the AOF starting from pos and keeps the connection alive",
    "complexity": "O(1)",
    "arguments": [
      {
        "name": "pos",
        "type": "integer"
      }
    ],
    "since": "1.0.0",
    "group": "replication"
  },
  "AOFMD5": {
    "summary": "Performs a checksum on a portion of the aof",
    "complexity": "O(1)",
    "arguments": [
      {
        "name": "pos",
        "type": "integer"
      },
      {
        "name": "size",
        "type": "integer"
      }
    ],
    "since": "1.0.0",
    "group": "replication"
  },
  "AOFSHRINK": {
    "summary": "Shrinks the aof in the background",
    "group": "replication"
  },
  "PING": {
    "summary": "Ping the server",
    "group": "connection"
  },
  "QUIT": {
    "summary": "Close the connection",
    "group": "connection"
  },
  "AUTH": {
    "summary": "Authenticate to the server",
    "arguments": [
      {
        "name": "password",
        "type": "string"
      }
    ],
    "group": "connection"
  },
  "OUTPUT": {
    "summary": "Gets or sets the output format for the current connection.",
    "arguments": [
      {
        "name": "format",
        "optional": true,
        "enumargs": [
          {
            "name": "json"
          },
          {
            "name": "resp"
          }
        ]
      }
    ],
    "group": "connection"
  },
  "SETHOOK": {
    "summary": "Creates a webhook which points to geofenced search",
    "arguments": [
      {
        "name": "name",
        "type": "string"
      },
      {
        "name": "endpoint",
        "type": "string"
      },
      {
        "enum": ["NEARBY", "WITHIN", "INTERSECTS"]
      },
      {
        "name": "key",
        "type": "string"
      },
      {
        "command": "FENCE",
        "name": [],
        "type": []
      },
      {
        "command": "DETECT",
        "name": ["what"],
        "type": ["string"],
        "optional": true
      },
      {
        "name": "param",
        "type": "string",
        "variadic": true
      }

    ],
    "group": "webhook"
  },
  "DELHOOK": {
    "summary": "Removes a webhook",
    "arguments": [
      {
        "name": "name",
        "type": "string"
      }
    ],
    "group": "webhook"
  },
  "HOOKS": {
    "summary": "Finds all hooks matching a pattern",
    "arguments":[
      {
        "name": "pattern",
        "type": "pattern"
      }
    ],
    "group": "webhook"
  }
}`
