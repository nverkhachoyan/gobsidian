---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/26
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Orcus"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 153
---
# [Orcus](3-Mechanics\CLI\bestiary\npc/orcus-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 153*  

## Orcus

Orcus is the Demon Prince of Undeath, known as the Blood Lord. He takes some pleasure in the sufferings of the living, but far prefers the company and service of the undead. His desire is to see all life quenched and the multiverse transformed into a vast necropolis populated solely by undead creatures under his command.

Orcus rewards those who spread death in his name by granting them a small portion of his power. The least of these become ghouls and zombies who serve in his legions, while his favored servants are the cultists and necromancers who murder the living and then manipulate the dead, emulating their dread master.

Orcus is a bestial creature of corruption with a diseased, decaying look. He has the lower torso of a goat, and a humanoid upper body with a corpulent belly swollen with rot. Great bat wings sprout from his shoulders, and his head is like the skull of a goat, the flesh nearly rotted from it. In one hand, he wields the legendary [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md), which is described in "chapter 7" of the "Dungeon Master's Guide".

## Orcus's Lair

Orcus makes his lair in the fortress city of Naratyr, which is on Thanatos, the layer of the Abyss that he rules. Surrounded by a moat fed by the River Styx, Naratyr is an eerily quiet and cold city, its streets often empty for hours at a time. The central castle of bone has interior walls of flesh and carpets made of woven hair. The city contains wandering undead, many of which are engaged in continuous battles with one another.

```statblock
"name": "Orcus (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor; 20 with the [Wand of Orcus](/3-Mechanics/CLI/items/wand-of-orcus.md)"
"hp": !!int "405"
"hit_dice": "30d12 + 210"
"stats":
- !!int "27"
- !!int "14"
- !!int "25"
- !!int "20"
- !!int "20"
- !!int "25"
"speed": "40 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "13"
  "Constitution": !!int "15"
"skillsaves":
  "Perception": !!int "13"
  "Arcana": !!int "13"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "necrotic; poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Orcus's spellcasting ability is Charisma (spell save DC 23, dice: d20+15\
    \ (+15) to hit with spell attacks). He can innately cast the following spells,\
    \ requiring no material components:\n\nAt will: [chill touch](/3-Mechanics/CLI/spells/chill-touch.md)\
    \ (17th level), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day:\
    \ [time stop](/3-Mechanics/CLI/spells/time-stop.md)\n\n3/day each: [create\
    \ undead](/3-Mechanics/CLI/spells/create-undead.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md)"
  "name": "Innate Spellcasting"
- "desc": "The wand has 7 charges, and any of its properties that require a saving\
    \ throw have a save DC of 18. While holding it, Orcus can use an action to cast\
    \ [animate dead](/3-Mechanics/CLI/spells/animate-dead.md), [blight](/3-Mechanics/CLI/spells/blight.md),\
    \ or [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md). Alternatively,\
    \ he can expend 1 or more of the wand's charges to cast one of the following spells\
    \ from it: [circle of death](/3-Mechanics/CLI/spells/circle-of-death.md) (1 charge),\
    \ [finger of death](/3-Mechanics/CLI/spells/finger-of-death.md) (1 charge), or\
    \ [power word kill](/3-Mechanics/CLI/spells/power-word-kill.md) (2 charges). The\
    \ wand regains dice: 1d4 + 3|avg|noform (1d4 + 3) charges daily at dawn.\n\
    \nWhile holding the wand, Orcus can use an action to conjure undead creatures\
    \ whose combined average hit points don't exceed 500. These undead magically rise\
    \ up from the ground or otherwise form in unoccupied spaces within 300 feet of\
    \ Orcus and obey his commands until they are destroyed or until he dismisses them\
    \ as an action. Once this property of the wand is used, the property can't be\
    \ used again until the next dawn."
  "name": "Wand of Orcus"
- "desc": "If Orcus fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Orcus has advantage on saving throws against spells and other magical effects."
  "name": "Magic Resistance"
- "desc": "Orcus's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "When Orcus casts [animate dead](/3-Mechanics/CLI/spells/animate-dead.md)\
    \ or [create undead](/3-Mechanics/CLI/spells/create-undead.md), he chooses the\
    \ level at which the spell is cast, and the creatures created by the spells remain\
    \ under his control indefinitely. Additionally, he can cast [create undead](/3-Mechanics/CLI/spells/create-undead.md)\
    \ even when it isn't night."
  "name": "Master of Undeath"
"actions":
- "desc": "Orcus makes two Wand of Orcus attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+19 (+19) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) bludgeoning damage plus dice:2d12|text(13)\
    \ (2d12) necrotic damage."
  "name": "Wand of Orcus"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) piercing damage plus dice:2d8|text(9)\
    \ (2d8) poison damage."
  "name": "Tail"
"legendary_actions":
- "desc": "Orcus makes one tail attack."
  "name": "Tail"
- "desc": "Orcus casts [chill touch](/3-Mechanics/CLI/spells/chill-touch.md) (17th\
    \ level)."
  "name": "A Taste of Undeath"
- "desc": "Orcus chooses a point on the ground that he can see within 100 feet of\
    \ him. A cylinder of swirling necrotic energy 60 feet tall and with a 10-foot\
    \ radius rises from that point and lasts until the end of Orcus's next turn. Creatures\
    \ in that area have vulnerability to necrotic damage."
  "name": "Creeping Death (Costs 2 Actions)"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Orcus can take a lair\
    \ action to cause one of the following effects; he can't use the same effect two\
    \ rounds in a row:"
  "name": ""
- "desc": "- Orcus's voice booms throughout the lair. His utterance causes one creature\
    \ of his choice to be subjected to [power word kill](/3-Mechanics/CLI/spells/power-word-kill.md).\
    \ Orcus needn't see the creature, but he must be aware that the individual is\
    \ in the lair.  \n- Orcus causes up to six corpses within the lair to rise as\
    \ [skeletons](/3-Mechanics/CLI/bestiary/undead/skeleton.md), [zombies](/3-Mechanics/CLI/bestiary/undead/zombie.md),\
    \ or [ghouls](/3-Mechanics/CLI/bestiary/undead/ghoul.md). These undead obey his\
    \ telepathic commands, which can reach anywhere in the lair.  \n- Orcus causes\
    \ skeletal arms to rise from an area on the ground in a 20-foot square that he\
    \ can see. They last until the next initiative count 20. Each creature in that\
    \ area when the arms appear must succeed on a DC 23 Strength saving throw or be\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) until the arms\
    \ disappear or until Orcus releases their grasp (no action required).  "
  "name": ""
"regional_effects":
- "desc": "The region containing Orcus's lair is warped by his magic, creating one\
    \ or more of the following effects:"
  "name": ""
- "desc": "- Dead beasts periodically animate as undead mockeries of their former\
    \ selves. Skeletal and zombie versions of local wildlife are commonly seen in\
    \ the area.  \n- The air becomes filled with the stench of rotting flesh, and\
    \ buzzing flies grow thick within the region, even when there is no carrion to\
    \ be found.  \n- If a humanoid spends at least 1 hour within 1 mile of the lair,\
    \ that creature must succeed on a DC 23 Wisdom saving throw or descend into a\
    \ madness determined by the Madness of Orcus table. A creature that succeeds on\
    \ this saving throw can't be affected by this regional effect again for 24 hours.\
    \  "
  "name": ""
- "desc": "If Orcus dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Orcus's lair or within line of sight of the demon\
    \ lord, roll on the Madness of Orcus table to determine the nature of the madness,\
    \ which is a character flaw that lasts until cured. See the \"Dungeon Master's\
    \ Guide\" for more on madness.\n\nMadness of Orcus\n\ndice: [](orcus-mtf.md#^madness-of-orcus)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"I often become withdrawn and moody, dwelling on the insufferable state\
    \ of life.\" |\n| 21–40 | \"I am compelled to make the weak suffer.\" |\n| 41–\
    60 | \"I have no compunction against tampering with the dead in my search to better\
    \ understand death.\" |\n| 61–80 | \"I want to achieve the everlasting existence\
    \ of undeath.\" |\n| 81–00 | \"I am awash in the awareness of life's futility.\"\
    \ |\n^madness-of-orcus"
  "name": "Madness of Orcus"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Orcus.webp"
```
^statblock

```encounter-table
name: Orcus
creatures:
 - 1: Orcus
```