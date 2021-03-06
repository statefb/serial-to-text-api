// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "serial-csv-converter",
    "version": "1.0"
  },
  "host": "localhost:3000",
  "paths": {
    "/data": {
      "get": {
        "tags": [
          "Common"
        ],
        "summary": "get current collected data.",
        "operationId": "get-data",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CollectedData"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/reset": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "reset status.",
        "operationId": "put-reset",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/send": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "send file to specified location.",
        "operationId": "put-send",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CollectedData"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/sendall": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "send all collected data.",
        "operationId": "put-sendall",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/start": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "start collecting data",
        "operationId": "put-start",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/KeyData"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "tags": [
          "Common"
        ],
        "summary": "get current status.",
        "operationId": "get-status",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "enum": [
                "waiting",
                "collecting",
                "sending",
                "error"
              ]
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/stop": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "stop collecting data.",
        "operationId": "put-stop",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CollectedData": {
      "type": "object",
      "title": "CollectedData",
      "required": [
        "timestamp",
        "value"
      ],
      "properties": {
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "Error.v1": {
      "description": "A standard error object.",
      "type": "object",
      "title": "Error",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      },
      "x-tags": [
        "Common"
      ]
    },
    "KeyData": {
      "type": "object",
      "title": "KeyData",
      "required": [
        "lotId",
        "bagId"
      ],
      "properties": {
        "bagId": {
          "type": "string"
        },
        "lotId": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "serial-csv-converter",
    "version": "1.0"
  },
  "host": "localhost:3000",
  "paths": {
    "/data": {
      "get": {
        "tags": [
          "Common"
        ],
        "summary": "get current collected data.",
        "operationId": "get-data",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CollectedData"
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/reset": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "reset status.",
        "operationId": "put-reset",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/send": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "send file to specified location.",
        "operationId": "put-send",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/CollectedData"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/sendall": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "send all collected data.",
        "operationId": "put-sendall",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/start": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "start collecting data",
        "operationId": "put-start",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/KeyData"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "tags": [
          "Common"
        ],
        "summary": "get current status.",
        "operationId": "get-status",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "enum": [
                "waiting",
                "collecting",
                "sending",
                "error"
              ]
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    },
    "/stop": {
      "put": {
        "tags": [
          "Common"
        ],
        "summary": "stop collecting data.",
        "operationId": "put-stop",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error.v1"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CollectedData": {
      "type": "object",
      "title": "CollectedData",
      "required": [
        "timestamp",
        "value"
      ],
      "properties": {
        "timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "Error.v1": {
      "description": "A standard error object.",
      "type": "object",
      "title": "Error",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      },
      "x-tags": [
        "Common"
      ]
    },
    "KeyData": {
      "type": "object",
      "title": "KeyData",
      "required": [
        "lotId",
        "bagId"
      ],
      "properties": {
        "bagId": {
          "type": "string"
        },
        "lotId": {
          "type": "string"
        }
      }
    }
  }
}`))
}
