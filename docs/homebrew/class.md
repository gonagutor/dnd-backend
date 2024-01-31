# How to create a custom class

## Defining the class

First you must define a top level key inside a json file called `class`
This key must be an array so that you can define class variants

### Class name

A class must have a `name` parameter. This will be the display name of the class.  
For an easier translation this object should be a map with the key being a [ISO-639](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) two letter code and the value being the name in that language

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      }
    }
  ]
}
```

### (Optional) Source

You can define a source parameter.  
It will default to `homebrew` if it is not set.  
Setting this parameter will allow you to include full rulesets in a campaign by just including anything with the same source.  
This can also be used as a filter parameter on other places by adding `#name_in_lowercase`

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB"
    }
  ]
}
```

### (Optional) Page

You can also define a page for easy reference if this is from an actual rules expansion and not just homebrew

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46
    }
  ]
}
```

### (Optional) SRD

If this is from the SRD you can mark this as true. Used for filtering by using `#srd`. Will default to false

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true
    }
  ]
}
```

### Hit dice

The number and sides of the dice used to calculate hit points

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      }
    }
  ]
}
```

### Proficiencies

#### Saving throws proficiencies

Must be an array inside a proficiencies object containing the three letter shorthand for main stats

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"]
      }
    }
  ]
}
```

#### Skills proficiencies

Must be an object that may contain a `choose` and a `forced` key.  
Forced will automatically be added as proficiencies, choose will let them choose using `count` as amount any from the array `from`.  
Both `from` and `forced` array items must be the full english name in [Snake case](https://es.wikipedia.org/wiki/Snake_case)

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        }
      }
    }
  ]
}
```

#### Weapons proficiencies

Must be a string array.  
Use a `#` and then the name of the type of weapon in [Snake case](https://es.wikipedia.org/wiki/Snake_case)  
Custom weapon kinds will work. Can be combined with a source `#` like `#srd` or other filters to limit the choices

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"]
      }
    }
  ]
}
```

#### Armor proficiencies

Must be a string array.  
Use a `#` and then the name of the type of armor in [Snake case](https://es.wikipedia.org/wiki/Snake_case)  
Custom armor kinds will work. Can be combined with a source `#` like `#srd` or other filters to limit the choices

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      }
    }
  ]
}
```

### Equipment

#### Additional from background

Whether or not it will allow to get additional starting equipment from background

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true
      }
    }
  ]
}
```

#### Starting equipment

Must be a key in the `equipment` object called `starting`  
Must be an array containing another array.  
If this second array contains just one key it will be automatically given to the player.  
If it contains multiple items the player will be prompted to choose one of them

Here is introduced a new `#`, the `#id?<item id>` this can be used to reference a specific item
How many items will be given is defined by adding an `x<item count>` to the end like so: `#id?<item id> x2`
Although using `#martial x2` is supported and the user will be prompted twice to choose 2 martial weapons we recommend just using multiple arrays for multiple choices of the same type

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"], // d135c12e-4f2e-11ee-be56-0242ac120002 should be great axe's id
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"], // 9bff66a2-4f30-11ee-be56-0242ac120002 should be explorer's pack id
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"] // 6d0ce432-4f30-11ee-be56-0242ac120002 should be javelin's id
        ],
      }
    }
  ]
}
```

#### Alternate gold

Must be a key in the `equipment` object called `alternateGold`  
Must be an object containing `count` (how many dice will be rolled), `sides` (how many sides does the die have), `multiplier` (the resulting amount will be multiplied by this valud ) and `additional` (the resulting amount after multiplication will be added this)

This amount of gold may be given to purchase their own equipment instead of the starting equipment

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 0,
          "sides": 8,
          "multiplier": 10,
          "additional": 0,
        }
      }
    }
  ]
}
```

### Multiclassing

#### Requirements

Stats requirements must be defined in a `requirements` object inside the `multiclassing` object  
Attributes must be defined by using the english main stat shorthand in uppercase.

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 0,
          "sides": 8,
          "multiplier": 10,
          "additional": 0,
        }
      },
      "multiclassing": {
        "requirements": {
          "STR": 13
        }
      }
    }
  ]
}
```

#### Proficiencies gained

Proficiencies gained must be defined in a `proficiencies` object inside the `multiclassing` object  
This object follows the previous proficiencies object

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 0,
          "sides": 8,
          "multiplier": 10,
          "additional": 0,
        }
      },
      "multiclassing": {
        "requirements": {
          "STR": 13
        },
        "proficiencies": {
          "weapons": ["#simple", "#martial"],
          "armor": ["#shield #phb"]
        }
      }
    }
  ]
}
```

### Class specifics

#### Counters

Counters must have:

- An `id` which will be used as a reference later
- A `name` which will be shown to the user. This is translatable the same way the class is
- A `progression` array which must contain 20 entries (one per level) and will be used to determine how many does a player have per level. -1 for infinite
- A `recoveredOn` which can be either `long_rest` or `short_rest` which will define when the user gets their points back

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 0,
          "sides": 8,
          "multiplier": 10,
          "additional": 0,
        }
      },
      "multiclassing": {
        "requirements": {
          "STR": 13
        },
        "proficiencies": {
          "weapons": ["#simple", "#martial"],
          "armor": ["#shield #phb"]
        }
      },
      "specifics": {
        "counters": [{
          "id": "rage",
          "name": {
            "es": "Rábia",
            "en": "Rage"
          },
          "progression": [2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, -1],
          "recoveredOn": "long_rest"
        }]
      }
    }
  ]
}
```

#### Bonus

Counters must have:

- An `id` which will be used as a reference later
- A `name` which will be shown to the user. This is translatable the same way the class is
- A `progression` array which must contain 20 entries (one per level) and will be used to determine how many does a player have per level. -1 for infinite
- A `recoveredOn` which can be either `long_rest` or `short_rest` which will define when the user gets their points back

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 2,
          "sides": 4,
          "multiplier": 10,
          "additional": 0,
        }
      },
      "multiclassing": {
        "requirements": {
          "STR": 13
        },
        "proficiencies": {
          "weapons": ["#simple", "#martial"],
          "armor": ["#shield #phb"]
        }
      },
      "specifics": {
        "counters": [{
          "id": "rage",
          "name": {
            "es": "Rabia",
            "en": "Rage"
          },
          "progression": [2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, -1],
          "recoveredOn": "long_rest"
        }],
        "bonuses": [{
          "type": "attack",
          "dependsOn": "%rage",
          "progression": [2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4],
        }]
      },
    }
  ]
}
```

### Class features

Counters must have:

- An `id` which will be used as a reference later
- A `name` which will be shown to the user. This is translatable the same way the class is
- A `progression` array which must contain 20 entries (one per level) and will be used to determine how many does a player have per level. -1 for infinite
- A `recoveredOn` which can be either `long_rest` or `short_rest` which will define when the user gets their points back

```json
{
  "class": [
    {
      "name": {
        "es": "Bárbaro",
        "en": "Barbarian"
      },
      "source": "PHB",
      "page": 46,
      "srd": true,
      "hitDice": {
        "count": 1,
        "faces": 12
      },
      "proficiencies": {
        "savingThrows": ["CON", "STR"],
        "skills": {
          "choose": {
            "count": 2,
            "from": [
              "animal_handling",
              "athletics",
              "intimidation",
              "nature",
              "perception",
              "survival"
            ]
          },
          "forced": []
        },
        "weapons": ["#simple", "#martial"],
        "armor": ["#shield #phb", "#light", "#medium"]
      },
      "equipment": {
        "additionalFromBackground": true,
        "starting": [
          ["#id?d135c12e-4f2e-11ee-be56-0242ac120002", "#martial #melee"],
          ["#handaxe x2", "#simple"]
          ["#id?9bff66a2-4f30-11ee-be56-0242ac120002"],
          ["#id?6d0ce432-4f30-11ee-be56-0242ac120002 x4"]
        ],
        "goldAlternative": {
          "count": 2,
          "sides": 4,
          "multiplier": 10,
          "additional": 0,
        }
      },
      "multiclassing": {
        "requirements": {
          "STR": 13
        },
        "proficiencies": {
          "weapons": ["#simple", "#martial"],
          "armor": ["#shield #phb"]
        }
      },
      "specifics": {
        "counters": [{
          "id": "rage",
          "name": {
            "es": "Rabia",
            "en": "Rage"
          },
          "progression": [2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, -1],
          "recoveredOn": "long_rest"
        }],
        "effects": [{
          "id": "rage",
          "name": {
            "es": "Rage",
            "en": "Rabia"
          },
          "bonus": [2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4],
          "resistance": ["bludgeoning", "piercing", "slashing"],
          "duration": 1000,
          "dependsOn": ["!#heavy", "!unconscious", "attacked|attacked_other"],
          "advantage": {
            "stats": ["STR"],
            "savingThrows": ["STR"]
          },
        }]
      },
      "features": [
        [{
          "name": {
            "en": "Rage",
            "es": "Rabia"
          },
          "logic": {
            "extra": {
              "distance": "self",
              "dependsOn": "%rage",
              "onSuccess": {
                "applies": "action",
                "onEnd": {
                  "after": "%rage",
                  "applies": ["exhaustion"]
                }
              }
            }
          },
          "description": {
            "en": "Starting when you choose this path at 3rd level, you can go into a frenzy when you rage. If you do so, for the duration of your rage you can make a single melee weapon attack as a bonus action on each of your turns after this one. When your rage ends, you suffer one level of {@condition exhaustion}.",
            "es": "A partir del momento en el que escoges esta senda, a nivel 3, puedes abandonarte al frenesí cuando te enfurezcas. Si eliges hacer esto, mientras estés enfurecido podrás hacer un ataque con arma cuerpo a cuerpo como acción adicional durante cada uno de tus turnos (empezando por el siguiente a haberte enfurecido). A cambio, cuando tu furia termine sufrirás un nivel de {@condition exhaustion}"
          }
        },{
          "name": {
            "en": "Frenzy",
            "es": "Frenesí"
          },
          "logic": {
            "distance": "self",
            "dependsOn": "%rage",
            "onSuccess": {
              "applies": "rage",
              "onEnd": {
                "after": "%rage",
                "applies": ["exhaustion"]
              }
            }
          },
          "description": {
            "en": "Starting when you choose this path at 3rd level, you can go into a frenzy when you rage. If you do so, for the duration of your rage you can make a single melee weapon attack as a bonus action on each of your turns after this one. When your rage ends, you suffer one level of {@condition exhaustion}.",
            "es": "A partir del momento en el que escoges esta senda, a nivel 3, puedes abandonarte al frenesí cuando te enfurezcas. Si eliges hacer esto, mientras estés enfurecido podrás hacer un ataque con arma cuerpo a cuerpo como acción adicional durante cada uno de tus turnos (empezando por el siguiente a haberte enfurecido). A cambio, cuando tu furia termine sufrirás un nivel de {@condition exhaustion}"
          }
        }],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
        [],
      ]
    }
  ]
}
```

## Defining subclasses

A `subclass` key should be created at the same level as `class`
A `subclassName` key will name the subclass group  
It should be an array of objects as follows

- A `name` following the translation format previously described or a simple string
- A `shortname` following the translation format previously described or a simple string
- A `source` that will be used for filtering like previous occurences
- A `page` that is just used for a quick reference
- A `srd` that is also used for filtering
- A `isSubclassChoice` that is used to determine if that is the start of a new subclass. This name and description will be used as a description of the subclass and to determine when to start the path
- A `features` array that contains another 20 arrays, one per level, which may contain 0 or more objects like
  - A `name` following the translation format previously described or a simple string
  - A `description` following the translation format previously described or a simple string. You can separate paragraphs using arrays
  - A `logic` object that determines how this feature will be handled. This will be explained later
  - A optional `choices` entry, which lets players decide one of a variant

```json
{
  ...
  "subclassTitle": {
    "es": "Senda primordial",
    "en": "Primal path"
  },
  "subclass": [{
    "name": {
      "en": "Path of the Berserker",
      "es": "Senda del Berserker"
    },
    "shortname": "Berserker",
    "source": "PHB",
    "page": 49,
    "srd": true,
    "features": [
      [],
      [],
      [{
        "name": {
          "en": "Path of the Berserker",
          "es": "Senda del Berserker"
        },
        "isSubclassChoice": true,
        "description": {
          "en": "For some barbarians, rage is a means to an end - that end being violence. The Path of the Berserker is a path of untrammeled fury, slick with blood. As you enter the berserker's rage, you thrill in the chaos of battle, heedless of your own health or well-being.",
          "es": "Para ciertos bárbaros la furia es el medio para alcanzar un fin: la violencia. La Senda del Berserker se asienta en una rabia sin cortapisas, cubierta de sangre. Al dejarte llevar por la furia del berserker experimentarás el entusiasmo que proporciona el caos del combate, sin reparar en tu propia seguridad o bienestar"
        }
      }, {
        "name": {
          "en": "Frenzy",
          "es": "Frenesí"
        },
        "logic": {
          "distance": "self",
          "dependsOn": "rage",
          "onSuccess": {
            "applies": "action",
            "onEnd": {
              "after": "rage",
              "applies": ["exhaustion"]
            }
          }
        },
        "description": {
          "en": "Starting when you choose this path at 3rd level, you can go into a frenzy when you rage. If you do so, for the duration of your rage you can make a single melee weapon attack as a bonus action on each of your turns after this one. When your rage ends, you suffer one level of {@condition exhaustion}.",
          "es": "A partir del momento en el que escoges esta senda, a nivel 3, puedes abandonarte al frenesí cuando te enfurezcas. Si eliges hacer esto, mientras estés enfurecido podrás hacer un ataque con arma cuerpo a cuerpo como acción adicional durante cada uno de tus turnos (empezando por el siguiente a haberte enfurecido). A cambio, cuando tu furia termine sufrirás un nivel de {@condition exhaustion}"
        }
      }],
      [],
      [],
      [{
        "name": {
          "en": "Mindless rage",
          "es": "Furia irracional"
        },
        "logic": {
          "passive": {
            "prevents": ["charmed", "frightened"]
          }
        },
        "description": {
          "en": "Beginning at 6th level, you can't be {@condition charmed} or {@condition frightened} while raging. If you are {@condition charmed} or {@condition frightened} when you enter your rage, the effect is suspended for the duration of the rage.",
          "es": "A partir de nivel 6, no puedes ser {@condition charmed} ni {@condition frightened} mientras estés enfurecido. Si estabas {@condition charmed} o {@condition frightened} antes de enfurecerte, dejarás de estarlo temporalmente, mientras dure la furia"
        }
      }],
      [],
      [],
      [],
      [{
        "name": {
          "en": "Intimidating Presence",
          "es": "Presencia Intimidante"
        },
        "logic": {
          "action": {
            "distance": 30,
            "savingThrow": {
              "base": 8,
              "bonifier": ["PROF", "CHA"],
              "onFailure": {
                "prevents": ["reuse"]
              },
              "onSuccess": {
                "frightened": {
                  "duration": 1,
                  "extendsWith": "action",
                  "condition": {
                    "distance": 60
                  }
                }
              }
            }
          }
        },
        "description": {
          "en": [
				    "Beginning at 10th level, you can use your action to frighten someone with your menacing presence. When you do so, choose one creature that you can see within 30 feet of you. If the creature can see or hear you, it must succeed on a Wisdom saving throw (DC equal to 8 + your proficiency bonus + your Charisma modifier) or be {@condition frightened} of you until the end of your next turn. On subsequent turns, you can use your action to extend the duration of this effect on the {@condition frightened} creature until the end of your next turn. This effect ends if the creature ends its turn out of line of sight or more than 60 feet away from you.",
				    "If the creature succeeds on its saving throw, you can't use this feature on that creature again for 24 hours."
			    ],
          "es": [
            "A partir de nivel 10, puedes usar tu acción para asustar a alguien con tu mera presencia. Para hacer esto, elige a una criatura que puedas ver a 30 pies o menos de ti. Si dicha criatura puede verte u oírte, deberá superar una tirada de salvación de Sabiduría (CD 8 + tu bonificador por competencia + tu bonificador por Carisma) o estará {@condition frightened} de ti hasta el final de tu siguiente turno. En cada uno de los turnos subsiguientes, podrás utilizar tu acción para alargar la duración de este efecto sobre la criatura {@condition frightened} hasta el final de tu siguiente turno. Este efecto termina si la criatura termina su turno fuera de tu línea de visión o a 60 pies o más de ti", "Si tiene éxito en su tirada de salvación, no podrás volver a usar este rasgo sobre esa criatura hasta que pasen 24 horas"
          ]
        }
      }],
      [],
      [],
      [],
      [{
        "name": {
          "en": "Retaliation",
          "es": "Represalia"
        },
        "logic": {
          "passive": {
            "distance": 5,
            "attack": "melee",
            "uses": "reaction"
          }
        },
        "description": {
          "en": "Starting at 14th level, when you take damage from a creature that is within 5 feet of you, you can use your reaction to make a melee weapon attack against that creature.",
          "es": "A partir de nivel 14, cuando recibas daño de una criatura que se encuentre a 5 pies o menos de ti, podrás utilizar tu reacción para hacer un ataque con arma cuerpo a cuerp o contra ella"
        }
      }],
      [],
      [],
      [],
      [],
      [],
      []
    ]
  }],
}
```

### Logic

Many optional fields are allowed here so a skeleton for each will be provided

- The top object may be one of `action`, `extra`, `action+extra`, `passive` or it may directly contain logic that is applied on self
  - An `action`, `action+extra` and `extra` may contain
    - A `distance` which may be on feet or `self`
    - A optional `target`. If omitted will default to 1
    - A `savingThrow` which specifies the saving throw that the targeted creature(s) have to suceed on. Most parameters inside are optional
      - `base` which determines the base which is added uppon. This can also be `roll` for forcing to roll for it
      - `bonifier` which determines which values will be added to the resulting value
      - `onFailure` which determines what will happen if this doesn't hit
        - `prevents` it must be an array of what prevents. It may be `reuse`
      - `onSuccess` which determines what will happen if this hits
        - `<an effect>` which will be applied
          - `duration` how long it will last
          - `extendsWith` if the effect can be continued this can have an `action`, an `extra`, an `action+extra` or it may consume a counter by the `%<counter id>` format
          - `condition` that will determine if the effect ends by a creature action or other
        - `damage`

```json
{
  "logic": {
    "action": {
      "distance": 30,
      "savingThrow": {
        "base": 8,
        "bonifier": ["PROF", "CHA"],
        "onFailure": {
          "prevents": ["reuse"]
        },
        "onSuccess": {
          "frightened": {
            "duration": 1,
            "extendsWith": "action",
            "condition": {
              "distance": 60
            }
          }
        }
      }
    }
  }
}
```
