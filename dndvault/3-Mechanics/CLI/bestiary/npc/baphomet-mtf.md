---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/23
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Baphomet"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 143
---
# [Baphomet](3-Mechanics\CLI\bestiary\npc/baphomet-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 143*  

## Baphomet

Civilization is weakness and savagery is strength in the credo of Baphomet, the Horned King and the Prince of Beasts. He rules over minotaurs and others with savage hearts. He is worshiped by those who want to break the confines of civility and unleash their bestial natures, for Baphomet envisions a world without restraint, where creatures live out their most savage desires.

Cults devoted to Baphomet use mazes and complex knots as their emblems, creating secret places to indulge themselves, including labyrinths of the sort their master favors. Bloodstained crowns and weapons of iron and brass decorate their profane altars.

Over time, Baphomet's cultists become tainted by his influence, gaining bloodshot eyes and coarse, thickening hair. Small horns eventually sprout from the forehead. In time, a devoted cultist might transform entirely into a minotaur—considered the greatest gift of the Prince of Beasts.

Baphomet appears as a great, black-furred minotaur, 20 feet tall with six iron horns. An infernal light burns in his red eyes. Although he is filled with bestial blood lust, there lies within him a cruel and cunning intellect devoted to subverting all of civilization.

Baphomet wields a great glaive called Heartcleaver. He sometimes casts this deadly weapon aside so that he can charge his enemies and gore them with his horns, trampling them into the earth and rending them with his teeth like a beast.

## Baphomet's Lair

Baphomet's lair is his palace, the Lyktion, which is on the layer of the Abyss called the Endless Maze. Nestled within the twisting passages of the plane-wide labyrinth, the Lyktion is immaculately maintained and surrounded by a moat constructed in the fashion of a three-dimensional maze. The palace itself is a towering structure whose interior is as labyrinthine as the plane on which it resides, populated by [minotaurs](/3-Mechanics/CLI/bestiary/monstrosity/minotaur.md), [goristros](/3-Mechanics/CLI/bestiary/fiend/goristro.md), and [quasits](/3-Mechanics/CLI/bestiary/fiend/quasit.md).

```statblock
"name": "Baphomet (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "275"
"hit_dice": "19d12 + 152"
"stats":
- !!int "30"
- !!int "14"
- !!int "26"
- !!int "18"
- !!int "24"
- !!int "16"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "14"
  "Constitution": !!int "15"
"skillsaves":
  "Intimidation": !!int "17"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Baphomet's spellcasting ability is Charisma (spell save DC 18). He can\
    \ innately cast the following spells, requiring no material components:\n\nAt\
    \ will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day:\
    \ [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [dominate beast](/3-Mechanics/CLI/spells/dominate-beast.md), [hunter's mark](/3-Mechanics/CLI/spells/hunters-mark.md),\
    \ [maze](/3-Mechanics/CLI/spells/maze.md), [wall of stone](/3-Mechanics/CLI/spells/wall-of-stone.md)"
  "name": "Innate Spellcasting"
- "desc": "If Baphomet moves at least 10 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:3d10|text(16)\
    \ (3d10) piercing damage. If the target is a creature, it must succeed on a\
    \ DC 25 Strength saving throw or be pushed up to 10 feet away and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
- "desc": "Baphomet can perfectly recall any path he has traveled, and he is immune\
    \ to the [maze](/3-Mechanics/CLI/spells/maze.md) spell."
  "name": "Labyrinthine Recall"
- "desc": "If Baphomet fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Baphomet has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Baphomet's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "At the start of his turn, Baphomet can gain advantage on all melee weapon\
    \ attack rolls during that turn, but attack rolls against him have advantage until\
    \ the start of his next turn."
  "name": "Reckless"
"actions":
- "desc": "Baphomet makes three attacks: one with Heartcleaver, one with his bite,\
    \ and one with his gore attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 10|text(21) (2d10 + 10) slashing damage."
  "name": "Heartcleaver"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 10|text(19) (2d8 + 10) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d6 + 10|text(17) (2d6 + 10) piercing damage."
  "name": "Gore"
- "desc": "Each creature of Baphomet's choice within 120 feet of him and aware of\
    \ him must succeed on a DC 18 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. These later saves have disadvantage if Baphomet\
    \ is within line of sight of the creature.\n\nIf a creature succeeds on any of\
    \ these saves or the effect ends on it, the creature is immune to Baphomet's Frightful\
    \ Presence for the next 24 hours."
  "name": "Frightful Presence"
"legendary_actions":
- "desc": "Baphomet makes a melee attack with Heartcleaver."
  "name": "Heartcleaver Attack"
- "desc": "Baphomet moves up to his speed, then makes a gore attack."
  "name": "Charge (Costs 2 Actions)"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), Baphomet can take a lair\
    \ action to cause one of the following magical effects; he can't use the same\
    \ effect two rounds in a row:"
  "name": ""
- "desc": "- Baphomet seals one doorway or other entryway within the lair. The opening\
    \ must be unoccupied. It is filled with solid stone for 1 minute or until Baphomet\
    \ creates this effect again.  \n- Baphomet chooses a room within the lair that\
    \ is no larger in any dimension than 100 feet. Until the next initiative count\
    \ 20, gravity is reversed within that room. Any creatures or objects in the room\
    \ when this happens fall in the direction of the new pull of gravity, unless they\
    \ have some means of remaining aloft. Baphomet can ignore the gravity reversal\
    \ if he's in the room, although he likes to use this action to land on a ceiling\
    \ to attack targets flying near it.  \n- Baphomet casts [mirage arcane](/3-Mechanics/CLI/spells/mirage-arcane.md),\
    \ affecting a room within the lair that is no larger in any dimension than 100\
    \ feet. The effect ends on the next initiative count 20.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Baphomet's lair is warped by his magic, creating\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- Plant life within 1 mile of the lair grows thick and forms walls of trees,\
    \ hedges, and other flora in the form of small mazes.  \n- Beasts within 1 mile\
    \ of the lair become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ and disoriented, as though constantly under threat of being hunted, and might\
    \ lash out or panic even when no visible threat is nearby.  \n- If a humanoid\
    \ spends at least 1 hour within 1 mile of the lair, that creature must succeed\
    \ on a DC 18 Wisdom saving throw or descend into a madness determined by the Madness\
    \ of Baphomet table. A creature that succeeds on this saving throw can't be affected\
    \ by this regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Baphomet dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Baphomet's lair or within line of sight of the\
    \ demon lord, roll on the Madness of Baphomet table to determine the nature of\
    \ the madness, which is a character flaw that lasts until cured. See the \"Dungeon\
    \ Master's Guide\" for more on madness.\n\nMadness of Baphomet\n\ndice: [](baphomet-mtf.md#^madness-of-baphomet)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"My anger consumes me. I can't be reasoned with when my rage has been\
    \ stoked.\" |\n| 21–40 | \"I degenerate into beastly behavior, seeming more like\
    \ a wild animal than a thinking being.\" |\n| 41–60 | \"The world is my hunting\
    \ ground. Others are my prey.\" |\n| 61–80 | \"Hate comes easily to me and explodes\
    \ into rage.\" |\n| 81–00 | \"I see those who oppose me not as people, but as\
    \ beasts meant to be preyed upon.\" |\n^madness-of-baphomet"
  "name": "Madness of Baphomet"
"source":
- "MTF"
- "OotA"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Baphomet.webp"
```
^statblock

```encounter-table
name: Baphomet
creatures:
 - 1: Baphomet
```