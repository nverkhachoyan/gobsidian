---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/14
- monster/environment/underdark
- monster/size/large
- monster/type/aberration/mind-flayer
aliases: ["Elder Brain"]
NoteIcon: monster
BestiaryType: aberration (mind flayer)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 120, Volo's Guide to Monsters p. 173
---
# [Elder Brain](3-Mechanics\CLI\bestiary\aberration/elder-brain-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 120, Volo's Guide to Monsters p. 173*  

```statblock
"name": "Elder Brain (MPMM)"
"size": "Large"
"type": "aberration"
"subtype": "mind flayer"
"alignment": "Typically  Lawful Evil"
"ac": !!int "10"
"hp": !!int "210"
"hit_dice": "20d10 + 100"
"stats":
- !!int "15"
- !!int "10"
- !!int "20"
- !!int "21"
- !!int "19"
- !!int "24"
"speed": "5 ft., swim 10 ft."
"saves":
  "Charisma": !!int "12"
  "Wisdom": !!int "9"
  "Intelligence": !!int "10"
"skillsaves":
  "Intimidation": !!int "12"
  "Deception": !!int "12"
  "Insight": !!int "14"
  "Arcana": !!int "10"
  "Persuasion": !!int "12"
"senses": "blindsight 120 ft., passive Perception 14"
"languages": "understands Common, Deep Speech, and Undercommon but can't speak, telepathy\
  \ 5 miles"
"cr": "14"
"traits":
- "desc": "The elder brain casts one of the following spells, requiring no spell components\
    \ and using Intelligence as the spellcasting ability (spell save DC 18):\n\nAt\
    \ will: [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\
    \n1/day each: [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)\n\n3/day:\
    \ [modify memory](/3-Mechanics/CLI/spells/modify-memory.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "The elder brain is aware of creatures within 5 miles of it that have an\
    \ Intelligence score of 4 or higher. It knows the distance and direction to each\
    \ creature, as well as each one's Intelligence score, but can't sense anything\
    \ else about it. A creature protected by a [mind blank](/3-Mechanics/CLI/spells/mind-blank.md)\
    \ spell, a [nondetection](/3-Mechanics/CLI/spells/nondetection.md) spell, or similar\
    \ magic can't be perceived in this manner."
  "name": "Creature Sense"
- "desc": "If the elder brain fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The elder brain has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The elder brain can use its telepathy to initiate and maintain telepathic\
    \ conversations with up to ten creatures at a time. The elder brain can let those\
    \ creatures telepathically hear each other while connected in this way."
  "name": "Telepathic Hub"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 30 ft., one target.\
    \ Hit: dice:4d8 + 2|text(20) (4d8 + 2) bludgeoning damage. If the target\
    \ is a Huge or smaller creature, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 15) and takes dice:1d8 + 5|text(9) (1d8 + 5) psychic damage at\
    \ the start of each of its turns until the grapple ends. The elder brain can have\
    \ up to four targets [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ at a time."
  "name": "Tentacle"
- "desc": "Creatures of the elder brain's choice within 60 feet of it must succeed\
    \ on a DC 18 Intelligence saving throw or take dice:5d10 + 5|text(32) (5d10\
    \ + 5) psychic damage and be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. A target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Mind Blast (Recharge 5-6)"
"bonus_actions":
- "desc": "The elder brain targets one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ creature it senses with its Creature Sense trait and establishes a psychic link\
    \ with the target. Until the link ends, the elder brain can perceive everything\
    \ the target senses. The target becomes aware that something is linked to its\
    \ mind once it is no longer [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ and the elder brain can terminate the link at any time (no action required).\
    \ The target can use an action on its turn to attempt to break the link, doing\
    \ so with a successful DC 18 Charisma saving throw. On a successful save, the\
    \ target takes dice:3d6|text(10) (3d6) psychic damage. The link also ends\
    \ if the target and the elder brain are more than 5 miles apart. The elder brain\
    \ can form psychic links with up to ten creatures at a time."
  "name": "Psychic Link"
- "desc": "The elder brain targets a creature with which it has a psychic link. The\
    \ elder brain gains insight into the target's emotional state and foremost thoughts\
    \ (including worries, loves, and hates)."
  "name": "Sense Thoughts"
"legendary_actions":
- "desc": "The elder brain targets one creature within 120 feet of it with which it\
    \ has a psychic link. The elder brain breaks the creature's [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ on a spell it has cast. The creature also takes dice:1d4|text(2) (1d4) psychic\
    \ damage per level of the spell."
  "name": "Break Concentration"
- "desc": "The elder brain targets one creature within 120 feet of it with which it\
    \ has a psychic link. The target and enemies of the elder brain within 30 feet\
    \ of target take dice:3d6|text(10) (3d6) psychic damage."
  "name": "Psychic Pulse"
- "desc": "The elder brain targets one creature within 120 feet of it with which it\
    \ has a psychic link. The elder brain ends the link, causing the creature to have\
    \ disadvantage on all ability checks, attack rolls, and saving throws until the\
    \ end of the creature's next turn."
  "name": "Sever Psychic Link"
- "desc": "The elder brain makes one Tentacle attack."
  "name": "Tentacle (Costs 2 Actions)"
"source":
- "MPMM"
- "VGM"
- "PaBTSO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Elder%20Brain.webp"
```
^statblock

```encounter-table
name: Elder Brain
creatures:
 - 1: Elder Brain
```

## Environment

underdark